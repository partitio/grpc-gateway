package runtime

import (
	"context"
	"encoding/base64"
	"fmt"
	"net"
	"net/http"
	"net/textproto"
	"strconv"
	"strings"
	"time"

	"github.com/micro/go-micro/v2/metadata"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

// MetadataHeaderPrefix is the http prefix that represents custom gmetadata
// parameters to or from a gRPC call.
const MetadataHeaderPrefix = "Grpc-Metadata-"

// MetadataPrefix is prepended to permanent HTTP header keys (as specified
// by the IANA) when added to the gRPC context.
// Removed "grpcgateway-"
const MetadataPrefix = ""

// MetadataTrailerPrefix is prepended to gRPC gmetadata as it is converted to
// HTTP headers in a response handled by micro-gateway
const MetadataTrailerPrefix = "Grpc-Trailer-"

const metadataGrpcTimeout = "Grpc-Timeout"
const metadataHeaderBinarySuffix = "-Bin"

const xForwardedFor = "X-Forwarded-For"
const xForwardedHost = "X-Forwarded-Host"

var (
	// DefaultContextTimeout is used for gRPC call context.WithTimeout whenever a Grpc-Timeout inbound
	// header isn't present. If the value is 0 the sent `context` will not have a timeout.
	DefaultContextTimeout = 0 * time.Second
)

func decodeBinHeader(v string) ([]byte, error) {
	if len(v)%4 == 0 {
		// Input was padded, or padding was not necessary.
		return base64.StdEncoding.DecodeString(v)
	}
	return base64.RawStdEncoding.DecodeString(v)
}

/*
AnnotateContext adds context information such as gmetadata from the request.

At a minimum, the RemoteAddr is included in the fashion of "X-Forwarded-For",
except that the forwarded destination is not another HTTP service but rather
a gRPC service.
*/
func AnnotateContext(ctx context.Context, mux *ServeMux, req *http.Request) (context.Context, error) {
	md := metadata.Metadata{}
	timeout := DefaultContextTimeout
	if tm := req.Header.Get(metadataGrpcTimeout); tm != "" {
		var err error
		timeout, err = timeoutDecode(tm)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid grpc-timeout: %s", tm)
		}
	}

	for key, vals := range req.Header {
		for _, val := range vals {
			key = textproto.CanonicalMIMEHeaderKey(key)
			if h, ok := mux.incomingHeaderMatcher(key); ok {
				// Handles "-bin" gmetadata in grpc, since grpc will do another base64
				// encode before sending to server, we need to decode it first.
				if strings.HasSuffix(key, metadataHeaderBinarySuffix) {
					b, err := decodeBinHeader(val)
					if err != nil {
						return nil, status.Errorf(codes.InvalidArgument, "invalid binary header %s: %s", key, err)
					}

					val = string(b)
				}
				md = addToMetadata(md, h, val)
			}
		}
	}
	if host := req.Header.Get(xForwardedHost); host != "" {
		md = addToMetadata(md, xForwardedHost, host)
	} else if req.Host != "" {
		md = addToMetadata(md, xForwardedHost, req.Host)
	}

	if addr := req.RemoteAddr; addr != "" {
		if remoteIP, _, err := net.SplitHostPort(addr); err == nil {
			if fwd := req.Header.Get(xForwardedFor); fwd == "" {
				md = addToMetadata(md, xForwardedFor, remoteIP)
			} else {
				md = addToMetadata(md, xForwardedFor, fmt.Sprintf("%s, %s", fwd, remoteIP))
			}
		} else {
			grpclog.Infof("invalid remote addr: %s", addr)
		}
	}

	if timeout != 0 {
		ctx, _ = context.WithTimeout(ctx, timeout)
	}
	if len(md) == 0 {
		return ctx, nil
	}
	for _, mda := range mux.metadataAnnotators {
		md = mergeMetadata(md, mda(ctx, req))
	}
	return metadata.NewContext(ctx, md), nil
}

//// ServerMetadata consists of gmetadata sent from gRPC server.
//type ServerMetadata struct {
//	HeaderMD  gmetadata.MD
//	TrailerMD gmetadata.MD
//}

type serverMetadataKey struct{}

func timeoutDecode(s string) (time.Duration, error) {
	size := len(s)
	if size < 2 {
		return 0, fmt.Errorf("timeout string is too short: %q", s)
	}
	d, ok := timeoutUnitToDuration(s[size-1])
	if !ok {
		return 0, fmt.Errorf("timeout unit is not recognized: %q", s)
	}
	t, err := strconv.ParseInt(s[:size-1], 10, 64)
	if err != nil {
		return 0, err
	}
	return d * time.Duration(t), nil
}

func timeoutUnitToDuration(u uint8) (d time.Duration, ok bool) {
	switch u {
	case 'H':
		return time.Hour, true
	case 'M':
		return time.Minute, true
	case 'S':
		return time.Second, true
	case 'm':
		return time.Millisecond, true
	case 'u':
		return time.Microsecond, true
	case 'n':
		return time.Nanosecond, true
	default:
	}
	return
}

// isPermanentHTTPHeader checks whether hdr belongs to the list of
// permenant request headers maintained by IANA.
// http://www.iana.org/assignments/message-headers/message-headers.xml
func isPermanentHTTPHeader(hdr string) bool {
	switch hdr {
	case
		"Accept",
		"Accept-Charset",
		"Accept-Language",
		"Accept-Ranges",
		"Authorization",
		"Cache-Control",
		"Content-Type",
		"Cookie",
		"Date",
		"Expect",
		"From",
		"Host",
		"If-Match",
		"If-Modified-Since",
		"If-None-Match",
		"If-Schedule-Tag-Match",
		"If-Unmodified-Since",
		"Max-Forwards",
		"Origin",
		"Pragma",
		"Referer",
		"User-Agent",
		"Via",
		"Warning":
		return true
	}
	return false
}

func addToMetadata(metadata metadata.Metadata, key string, val string) metadata.Metadata {
	if v, ok := metadata[key]; ok {
		metadata[key] = v + "," + val
		return metadata
	}
	metadata[key] = val
	return metadata
}

func mergeMetadata(m1 metadata.Metadata, m2 metadata.Metadata) metadata.Metadata {
	for k1, v1 := range m1 {
		if v, ok := m2[k1]; ok {
			m2[k1] = v + "," + v1
		} else {
			m2[k1] = v1
		}
	}
	return m2
}
