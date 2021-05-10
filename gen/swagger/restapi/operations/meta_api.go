// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/yetialex/meta_test/gen/swagger/restapi/operations/core"
	"github.com/yetialex/meta_test/gen/swagger/restapi/operations/iba"
	"github.com/yetialex/meta_test/gen/swagger/restapi/operations/signals"
	"github.com/yetialex/meta_test/gen/swagger/restapi/operations/swagger"
)

// NewMetaAPI creates a new Meta instance
func NewMetaAPI(spec *loads.Document) *MetaAPI {
	return &MetaAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		CoreCreateDirectoryHandler: core.CreateDirectoryHandlerFunc(func(params core.CreateDirectoryParams) middleware.Responder {
			return middleware.NotImplemented("operation core.CreateDirectory has not yet been implemented")
		}),
		IbaCreateIBAGateHandler: iba.CreateIBAGateHandlerFunc(func(params iba.CreateIBAGateParams) middleware.Responder {
			return middleware.NotImplemented("operation iba.CreateIBAGate has not yet been implemented")
		}),
		CoreCreateNodeHandler: core.CreateNodeHandlerFunc(func(params core.CreateNodeParams) middleware.Responder {
			return middleware.NotImplemented("operation core.CreateNode has not yet been implemented")
		}),
		CoreCreateSignalHandler: core.CreateSignalHandlerFunc(func(params core.CreateSignalParams) middleware.Responder {
			return middleware.NotImplemented("operation core.CreateSignal has not yet been implemented")
		}),
		CoreDeleteDirectoryHandler: core.DeleteDirectoryHandlerFunc(func(params core.DeleteDirectoryParams) middleware.Responder {
			return middleware.NotImplemented("operation core.DeleteDirectory has not yet been implemented")
		}),
		CoreDeleteNodeHandler: core.DeleteNodeHandlerFunc(func(params core.DeleteNodeParams) middleware.Responder {
			return middleware.NotImplemented("operation core.DeleteNode has not yet been implemented")
		}),
		CoreDeleteSignalHandler: core.DeleteSignalHandlerFunc(func(params core.DeleteSignalParams) middleware.Responder {
			return middleware.NotImplemented("operation core.DeleteSignal has not yet been implemented")
		}),
		CoreGetDirectoriesHandler: core.GetDirectoriesHandlerFunc(func(params core.GetDirectoriesParams) middleware.Responder {
			return middleware.NotImplemented("operation core.GetDirectories has not yet been implemented")
		}),
		CoreGetDirectoryHandler: core.GetDirectoryHandlerFunc(func(params core.GetDirectoryParams) middleware.Responder {
			return middleware.NotImplemented("operation core.GetDirectory has not yet been implemented")
		}),
		IbaGetIBAGateByNameHandler: iba.GetIBAGateByNameHandlerFunc(func(params iba.GetIBAGateByNameParams) middleware.Responder {
			return middleware.NotImplemented("operation iba.GetIBAGateByName has not yet been implemented")
		}),
		IbaGetIBAGateMntHandler: iba.GetIBAGateMntHandlerFunc(func(params iba.GetIBAGateMntParams) middleware.Responder {
			return middleware.NotImplemented("operation iba.GetIBAGateMnt has not yet been implemented")
		}),
		IbaGetIBAGateMntsHandler: iba.GetIBAGateMntsHandlerFunc(func(params iba.GetIBAGateMntsParams) middleware.Responder {
			return middleware.NotImplemented("operation iba.GetIBAGateMnts has not yet been implemented")
		}),
		IbaGetIBAGatesHandler: iba.GetIBAGatesHandlerFunc(func(params iba.GetIBAGatesParams) middleware.Responder {
			return middleware.NotImplemented("operation iba.GetIBAGates has not yet been implemented")
		}),
		IbaGetIBAMappingByGateAndMountHandler: iba.GetIBAMappingByGateAndMountHandlerFunc(func(params iba.GetIBAMappingByGateAndMountParams) middleware.Responder {
			return middleware.NotImplemented("operation iba.GetIBAMappingByGateAndMount has not yet been implemented")
		}),
		IbaGetIBAMappingByServerIDHandler: iba.GetIBAMappingByServerIDHandlerFunc(func(params iba.GetIBAMappingByServerIDParams) middleware.Responder {
			return middleware.NotImplemented("operation iba.GetIBAMappingByServerID has not yet been implemented")
		}),
		IbaGetIBAMappingByTopicNameHandler: iba.GetIBAMappingByTopicNameHandlerFunc(func(params iba.GetIBAMappingByTopicNameParams) middleware.Responder {
			return middleware.NotImplemented("operation iba.GetIBAMappingByTopicName has not yet been implemented")
		}),
		IbaGetIBAServerHandler: iba.GetIBAServerHandlerFunc(func(params iba.GetIBAServerParams) middleware.Responder {
			return middleware.NotImplemented("operation iba.GetIBAServer has not yet been implemented")
		}),
		IbaGetIBAServersHandler: iba.GetIBAServersHandlerFunc(func(params iba.GetIBAServersParams) middleware.Responder {
			return middleware.NotImplemented("operation iba.GetIBAServers has not yet been implemented")
		}),
		CoreGetNodeHandler: core.GetNodeHandlerFunc(func(params core.GetNodeParams) middleware.Responder {
			return middleware.NotImplemented("operation core.GetNode has not yet been implemented")
		}),
		CoreGetNodeChildrenHandler: core.GetNodeChildrenHandlerFunc(func(params core.GetNodeChildrenParams) middleware.Responder {
			return middleware.NotImplemented("operation core.GetNodeChildren has not yet been implemented")
		}),
		CoreGetNodesHandler: core.GetNodesHandlerFunc(func(params core.GetNodesParams) middleware.Responder {
			return middleware.NotImplemented("operation core.GetNodes has not yet been implemented")
		}),
		CoreGetNodesTreeWithSignalsHandler: core.GetNodesTreeWithSignalsHandlerFunc(func(params core.GetNodesTreeWithSignalsParams) middleware.Responder {
			return middleware.NotImplemented("operation core.GetNodesTreeWithSignals has not yet been implemented")
		}),
		CoreGetRootNodeHandler: core.GetRootNodeHandlerFunc(func(params core.GetRootNodeParams) middleware.Responder {
			return middleware.NotImplemented("operation core.GetRootNode has not yet been implemented")
		}),
		CoreGetSignalHandler: core.GetSignalHandlerFunc(func(params core.GetSignalParams) middleware.Responder {
			return middleware.NotImplemented("operation core.GetSignal has not yet been implemented")
		}),
		CoreGetSignalsHandler: core.GetSignalsHandlerFunc(func(params core.GetSignalsParams) middleware.Responder {
			return middleware.NotImplemented("operation core.GetSignals has not yet been implemented")
		}),
		SignalsGetSignalsClassesHandler: signals.GetSignalsClassesHandlerFunc(func(params signals.GetSignalsClassesParams) middleware.Responder {
			return middleware.NotImplemented("operation signals.GetSignalsClasses has not yet been implemented")
		}),
		SignalsGetSignalsValueTypesHandler: signals.GetSignalsValueTypesHandlerFunc(func(params signals.GetSignalsValueTypesParams) middleware.Responder {
			return middleware.NotImplemented("operation signals.GetSignalsValueTypes has not yet been implemented")
		}),
		SwaggerGetSwaggerJSONHandler: swagger.GetSwaggerJSONHandlerFunc(func(params swagger.GetSwaggerJSONParams) middleware.Responder {
			return middleware.NotImplemented("operation swagger.GetSwaggerJSON has not yet been implemented")
		}),
		IbaRegisterIBAServerHandler: iba.RegisterIBAServerHandlerFunc(func(params iba.RegisterIBAServerParams) middleware.Responder {
			return middleware.NotImplemented("operation iba.RegisterIBAServer has not yet been implemented")
		}),
		IbaRegisterSignalHandler: iba.RegisterSignalHandlerFunc(func(params iba.RegisterSignalParams) middleware.Responder {
			return middleware.NotImplemented("operation iba.RegisterSignal has not yet been implemented")
		}),
		CoreUpdateDirectoryHandler: core.UpdateDirectoryHandlerFunc(func(params core.UpdateDirectoryParams) middleware.Responder {
			return middleware.NotImplemented("operation core.UpdateDirectory has not yet been implemented")
		}),
		IbaUpdateIBAGateMntHandler: iba.UpdateIBAGateMntHandlerFunc(func(params iba.UpdateIBAGateMntParams) middleware.Responder {
			return middleware.NotImplemented("operation iba.UpdateIBAGateMnt has not yet been implemented")
		}),
		CoreUpdateNodeHandler: core.UpdateNodeHandlerFunc(func(params core.UpdateNodeParams) middleware.Responder {
			return middleware.NotImplemented("operation core.UpdateNode has not yet been implemented")
		}),
		CoreUpdateSignalHandler: core.UpdateSignalHandlerFunc(func(params core.UpdateSignalParams) middleware.Responder {
			return middleware.NotImplemented("operation core.UpdateSignal has not yet been implemented")
		}),
	}
}

/*MetaAPI Сервис мета информации */
type MetaAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// CoreCreateDirectoryHandler sets the operation handler for the create directory operation
	CoreCreateDirectoryHandler core.CreateDirectoryHandler
	// IbaCreateIBAGateHandler sets the operation handler for the create i b a gate operation
	IbaCreateIBAGateHandler iba.CreateIBAGateHandler
	// CoreCreateNodeHandler sets the operation handler for the create node operation
	CoreCreateNodeHandler core.CreateNodeHandler
	// CoreCreateSignalHandler sets the operation handler for the create signal operation
	CoreCreateSignalHandler core.CreateSignalHandler
	// CoreDeleteDirectoryHandler sets the operation handler for the delete directory operation
	CoreDeleteDirectoryHandler core.DeleteDirectoryHandler
	// CoreDeleteNodeHandler sets the operation handler for the delete node operation
	CoreDeleteNodeHandler core.DeleteNodeHandler
	// CoreDeleteSignalHandler sets the operation handler for the delete signal operation
	CoreDeleteSignalHandler core.DeleteSignalHandler
	// CoreGetDirectoriesHandler sets the operation handler for the get directories operation
	CoreGetDirectoriesHandler core.GetDirectoriesHandler
	// CoreGetDirectoryHandler sets the operation handler for the get directory operation
	CoreGetDirectoryHandler core.GetDirectoryHandler
	// IbaGetIBAGateByNameHandler sets the operation handler for the get i b a gate by name operation
	IbaGetIBAGateByNameHandler iba.GetIBAGateByNameHandler
	// IbaGetIBAGateMntHandler sets the operation handler for the get i b a gate mnt operation
	IbaGetIBAGateMntHandler iba.GetIBAGateMntHandler
	// IbaGetIBAGateMntsHandler sets the operation handler for the get i b a gate mnts operation
	IbaGetIBAGateMntsHandler iba.GetIBAGateMntsHandler
	// IbaGetIBAGatesHandler sets the operation handler for the get i b a gates operation
	IbaGetIBAGatesHandler iba.GetIBAGatesHandler
	// IbaGetIBAMappingByGateAndMountHandler sets the operation handler for the get i b a mapping by gate and mount operation
	IbaGetIBAMappingByGateAndMountHandler iba.GetIBAMappingByGateAndMountHandler
	// IbaGetIBAMappingByServerIDHandler sets the operation handler for the get i b a mapping by server ID operation
	IbaGetIBAMappingByServerIDHandler iba.GetIBAMappingByServerIDHandler
	// IbaGetIBAMappingByTopicNameHandler sets the operation handler for the get i b a mapping by topic name operation
	IbaGetIBAMappingByTopicNameHandler iba.GetIBAMappingByTopicNameHandler
	// IbaGetIBAServerHandler sets the operation handler for the get i b a server operation
	IbaGetIBAServerHandler iba.GetIBAServerHandler
	// IbaGetIBAServersHandler sets the operation handler for the get i b a servers operation
	IbaGetIBAServersHandler iba.GetIBAServersHandler
	// CoreGetNodeHandler sets the operation handler for the get node operation
	CoreGetNodeHandler core.GetNodeHandler
	// CoreGetNodeChildrenHandler sets the operation handler for the get node children operation
	CoreGetNodeChildrenHandler core.GetNodeChildrenHandler
	// CoreGetNodesHandler sets the operation handler for the get nodes operation
	CoreGetNodesHandler core.GetNodesHandler
	// CoreGetNodesTreeWithSignalsHandler sets the operation handler for the get nodes tree with signals operation
	CoreGetNodesTreeWithSignalsHandler core.GetNodesTreeWithSignalsHandler
	// CoreGetRootNodeHandler sets the operation handler for the get root node operation
	CoreGetRootNodeHandler core.GetRootNodeHandler
	// CoreGetSignalHandler sets the operation handler for the get signal operation
	CoreGetSignalHandler core.GetSignalHandler
	// CoreGetSignalsHandler sets the operation handler for the get signals operation
	CoreGetSignalsHandler core.GetSignalsHandler
	// SignalsGetSignalsClassesHandler sets the operation handler for the get signals classes operation
	SignalsGetSignalsClassesHandler signals.GetSignalsClassesHandler
	// SignalsGetSignalsValueTypesHandler sets the operation handler for the get signals value types operation
	SignalsGetSignalsValueTypesHandler signals.GetSignalsValueTypesHandler
	// SwaggerGetSwaggerJSONHandler sets the operation handler for the get swagger JSON operation
	SwaggerGetSwaggerJSONHandler swagger.GetSwaggerJSONHandler
	// IbaRegisterIBAServerHandler sets the operation handler for the register i b a server operation
	IbaRegisterIBAServerHandler iba.RegisterIBAServerHandler
	// IbaRegisterSignalHandler sets the operation handler for the register signal operation
	IbaRegisterSignalHandler iba.RegisterSignalHandler
	// CoreUpdateDirectoryHandler sets the operation handler for the update directory operation
	CoreUpdateDirectoryHandler core.UpdateDirectoryHandler
	// IbaUpdateIBAGateMntHandler sets the operation handler for the update i b a gate mnt operation
	IbaUpdateIBAGateMntHandler iba.UpdateIBAGateMntHandler
	// CoreUpdateNodeHandler sets the operation handler for the update node operation
	CoreUpdateNodeHandler core.UpdateNodeHandler
	// CoreUpdateSignalHandler sets the operation handler for the update signal operation
	CoreUpdateSignalHandler core.UpdateSignalHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *MetaAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *MetaAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *MetaAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *MetaAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *MetaAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *MetaAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *MetaAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *MetaAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *MetaAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the MetaAPI
func (o *MetaAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.CoreCreateDirectoryHandler == nil {
		unregistered = append(unregistered, "core.CreateDirectoryHandler")
	}
	if o.IbaCreateIBAGateHandler == nil {
		unregistered = append(unregistered, "iba.CreateIBAGateHandler")
	}
	if o.CoreCreateNodeHandler == nil {
		unregistered = append(unregistered, "core.CreateNodeHandler")
	}
	if o.CoreCreateSignalHandler == nil {
		unregistered = append(unregistered, "core.CreateSignalHandler")
	}
	if o.CoreDeleteDirectoryHandler == nil {
		unregistered = append(unregistered, "core.DeleteDirectoryHandler")
	}
	if o.CoreDeleteNodeHandler == nil {
		unregistered = append(unregistered, "core.DeleteNodeHandler")
	}
	if o.CoreDeleteSignalHandler == nil {
		unregistered = append(unregistered, "core.DeleteSignalHandler")
	}
	if o.CoreGetDirectoriesHandler == nil {
		unregistered = append(unregistered, "core.GetDirectoriesHandler")
	}
	if o.CoreGetDirectoryHandler == nil {
		unregistered = append(unregistered, "core.GetDirectoryHandler")
	}
	if o.IbaGetIBAGateByNameHandler == nil {
		unregistered = append(unregistered, "iba.GetIBAGateByNameHandler")
	}
	if o.IbaGetIBAGateMntHandler == nil {
		unregistered = append(unregistered, "iba.GetIBAGateMntHandler")
	}
	if o.IbaGetIBAGateMntsHandler == nil {
		unregistered = append(unregistered, "iba.GetIBAGateMntsHandler")
	}
	if o.IbaGetIBAGatesHandler == nil {
		unregistered = append(unregistered, "iba.GetIBAGatesHandler")
	}
	if o.IbaGetIBAMappingByGateAndMountHandler == nil {
		unregistered = append(unregistered, "iba.GetIBAMappingByGateAndMountHandler")
	}
	if o.IbaGetIBAMappingByServerIDHandler == nil {
		unregistered = append(unregistered, "iba.GetIBAMappingByServerIDHandler")
	}
	if o.IbaGetIBAMappingByTopicNameHandler == nil {
		unregistered = append(unregistered, "iba.GetIBAMappingByTopicNameHandler")
	}
	if o.IbaGetIBAServerHandler == nil {
		unregistered = append(unregistered, "iba.GetIBAServerHandler")
	}
	if o.IbaGetIBAServersHandler == nil {
		unregistered = append(unregistered, "iba.GetIBAServersHandler")
	}
	if o.CoreGetNodeHandler == nil {
		unregistered = append(unregistered, "core.GetNodeHandler")
	}
	if o.CoreGetNodeChildrenHandler == nil {
		unregistered = append(unregistered, "core.GetNodeChildrenHandler")
	}
	if o.CoreGetNodesHandler == nil {
		unregistered = append(unregistered, "core.GetNodesHandler")
	}
	if o.CoreGetNodesTreeWithSignalsHandler == nil {
		unregistered = append(unregistered, "core.GetNodesTreeWithSignalsHandler")
	}
	if o.CoreGetRootNodeHandler == nil {
		unregistered = append(unregistered, "core.GetRootNodeHandler")
	}
	if o.CoreGetSignalHandler == nil {
		unregistered = append(unregistered, "core.GetSignalHandler")
	}
	if o.CoreGetSignalsHandler == nil {
		unregistered = append(unregistered, "core.GetSignalsHandler")
	}
	if o.SignalsGetSignalsClassesHandler == nil {
		unregistered = append(unregistered, "signals.GetSignalsClassesHandler")
	}
	if o.SignalsGetSignalsValueTypesHandler == nil {
		unregistered = append(unregistered, "signals.GetSignalsValueTypesHandler")
	}
	if o.SwaggerGetSwaggerJSONHandler == nil {
		unregistered = append(unregistered, "swagger.GetSwaggerJSONHandler")
	}
	if o.IbaRegisterIBAServerHandler == nil {
		unregistered = append(unregistered, "iba.RegisterIBAServerHandler")
	}
	if o.IbaRegisterSignalHandler == nil {
		unregistered = append(unregistered, "iba.RegisterSignalHandler")
	}
	if o.CoreUpdateDirectoryHandler == nil {
		unregistered = append(unregistered, "core.UpdateDirectoryHandler")
	}
	if o.IbaUpdateIBAGateMntHandler == nil {
		unregistered = append(unregistered, "iba.UpdateIBAGateMntHandler")
	}
	if o.CoreUpdateNodeHandler == nil {
		unregistered = append(unregistered, "core.UpdateNodeHandler")
	}
	if o.CoreUpdateSignalHandler == nil {
		unregistered = append(unregistered, "core.UpdateSignalHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *MetaAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *MetaAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *MetaAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *MetaAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *MetaAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *MetaAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the meta API
func (o *MetaAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *MetaAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/core/directories"] = core.NewCreateDirectory(o.context, o.CoreCreateDirectoryHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/ibas/gates/${gate_name}"] = iba.NewCreateIBAGate(o.context, o.IbaCreateIBAGateHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/core/nodes"] = core.NewCreateNode(o.context, o.CoreCreateNodeHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/core/signals"] = core.NewCreateSignal(o.context, o.CoreCreateSignalHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/core/directories/{directory_id}"] = core.NewDeleteDirectory(o.context, o.CoreDeleteDirectoryHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/core/nodes/{node_id}"] = core.NewDeleteNode(o.context, o.CoreDeleteNodeHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/core/signals/{signal_id}"] = core.NewDeleteSignal(o.context, o.CoreDeleteSignalHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/core/directories"] = core.NewGetDirectories(o.context, o.CoreGetDirectoriesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/core/directories/{directory_id}"] = core.NewGetDirectory(o.context, o.CoreGetDirectoryHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/ibas/gates/${gate_name}"] = iba.NewGetIBAGateByName(o.context, o.IbaGetIBAGateByNameHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/ibas/gates/${gate_name}/mnts/${mnt}"] = iba.NewGetIBAGateMnt(o.context, o.IbaGetIBAGateMntHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/ibas/gates/${gate_name}/mnts"] = iba.NewGetIBAGateMnts(o.context, o.IbaGetIBAGateMntsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/ibas/gates"] = iba.NewGetIBAGates(o.context, o.IbaGetIBAGatesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/ibas/gates/${gate_name}/mnts/${mnt}/signals/mapping"] = iba.NewGetIBAMappingByGateAndMount(o.context, o.IbaGetIBAMappingByGateAndMountHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/ibas/servers/${iba_server_id}/signals/mapping"] = iba.NewGetIBAMappingByServerID(o.context, o.IbaGetIBAMappingByServerIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/ibas/topics/${topic_name}/mapping"] = iba.NewGetIBAMappingByTopicName(o.context, o.IbaGetIBAMappingByTopicNameHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/ibas/servers/${iba_server_id}"] = iba.NewGetIBAServer(o.context, o.IbaGetIBAServerHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/ibas/servers"] = iba.NewGetIBAServers(o.context, o.IbaGetIBAServersHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/core/nodes/{node_id}"] = core.NewGetNode(o.context, o.CoreGetNodeHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/core/nodes/{node_id}/children"] = core.NewGetNodeChildren(o.context, o.CoreGetNodeChildrenHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/core/nodes"] = core.NewGetNodes(o.context, o.CoreGetNodesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/core/nodes/tree"] = core.NewGetNodesTreeWithSignals(o.context, o.CoreGetNodesTreeWithSignalsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/core/root_node"] = core.NewGetRootNode(o.context, o.CoreGetRootNodeHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/core/signals/{signal_id}"] = core.NewGetSignal(o.context, o.CoreGetSignalHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/core/signals"] = core.NewGetSignals(o.context, o.CoreGetSignalsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/core/signals/classes"] = signals.NewGetSignalsClasses(o.context, o.SignalsGetSignalsClassesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/core/signals/value_types"] = signals.NewGetSignalsValueTypes(o.context, o.SignalsGetSignalsValueTypesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/swagger.json"] = swagger.NewGetSwaggerJSON(o.context, o.SwaggerGetSwaggerJSONHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/ibas/servers"] = iba.NewRegisterIBAServer(o.context, o.IbaRegisterIBAServerHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/ibas/servers/${iba_server_id}/signals"] = iba.NewRegisterSignal(o.context, o.IbaRegisterSignalHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/core/directories/{directory_id}"] = core.NewUpdateDirectory(o.context, o.CoreUpdateDirectoryHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/ibas/gates/${gate_name}/mnts/${mnt}"] = iba.NewUpdateIBAGateMnt(o.context, o.IbaUpdateIBAGateMntHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/core/nodes/{node_id}"] = core.NewUpdateNode(o.context, o.CoreUpdateNodeHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/core/signals/{signal_id}"] = core.NewUpdateSignal(o.context, o.CoreUpdateSignalHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *MetaAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *MetaAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *MetaAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *MetaAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *MetaAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
