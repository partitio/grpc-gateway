# Atlas Patch for protoc-gen-swagger

atlas-patch is build on top of original protoc-gen-swagger and is intended to conform [atlas-app-toolkit REST API Sepcification](https://github.com/infobloxopen/atlas-app-toolkit#rest-api-syntax-specification).

Patch includes following changes:

 * Fixed method comments extraction

 * Rendering of messages that have a primitive type (STRING, INT, BOOLEAN)
   does not occur if message is used only as a field (not an rpc Request or Response),
   hence recursive message definitions and complex-structured messages can be presented
   as plain string query parameters.

 * Introduced new `atlas_patch` flag. If this flag is enabled `--swagger_out="atlas_patch=true:."`
   following changes are made to a swagger spec:

   * All responses are assigned to an appropriate response code:
     GET - 200/OK, POST - 201/CREATED, PUT - 202/UPDATED, DELETE - 204/DELETED.

   * Recursive references are broken up. Such references occur while using protoc-gen-gorm plugin
     with many-to-many/one-to-many relations.

   * Collection operators from atlas-app-toolkit are provided with documentation and correct
     names.

   * atlas.rpc.identifier in path is treated correctly and not distributed among path and
     query parameters, also id.payload_id is replaced with id in path.

   * Unused references elimination.
   
   * Exclude all operations tagged as "private" see example below
       ```
       rpc Update (UpdateNetworkRequest) returns (UpdateNetworkResponse) {
           option (google.api.http) = {
             put: "/network/{payload.id.resource_id}"
             body: "payload"
       
             additional_bindings {
               patch: "/network/{payload.id.resource_id}",
               body:  "payload"
             }
       
           };
           option (grpc.gateway.protoc_gen_swagger.options.openapiv2_operation) = {
             tags: "private"
           };
         }
       ```
   * Introduced new `with_private` flag if set generate service.private.swagger.json
   with all operation (including tagged as "private")
   
   * Provide couple annotations for replacing values in swagger schema you need specify flag ```atlas_patch=true``` and ```with_custom_annotations=true```
     - ```@example``` annotation can be used for replacing default example with custom one
      support few value types includes float64, string, map[string]interface{}, []map[string]interface{} []float64, []string
       - ```@example 5.0```
       - ```@example "Internal error"```
       - ```@example {"Location": "Tacoma"}```
       - ```@example ["First", "Second"]```
       - ```@example [1, 5, 44]```
       - ```@example [{"Location": "Tacoma"}, {"Group": "Engineering"}]```
     
     - ```@title``` annotation can be used for replacing default title with custom one
       - ```@title "StringCondition"```
      
    If you example too long to be presented on one line you could use multiple lines annotation
    ```
       @example <<<EOF
       {
           "Location": "Tacoma"
       }
    ```
    
    or
    
   ```
       @example <<<EOF
       {
           "Location": "Tacoma"
       }
       EOF
    ```
        
    In first case all what presented after line ```@example <<<EOF``` will be rendered as example,
    if you want to manually set boundaries please use ```EOF``` as a closing line