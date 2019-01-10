// Package generator provides an abstract interface to code generators.
package generator

import (
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/partitio/grpc-gateway/protoc-gen-micro-gateway/descriptor"
)

// Generator is an abstraction of code generators.
type Generator interface {
	// Generate generates output files from input .proto files.
	Generate(targets []*descriptor.File) ([]*plugin.CodeGeneratorResponse_File, error)
}
