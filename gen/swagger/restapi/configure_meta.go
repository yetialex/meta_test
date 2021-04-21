// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/yetialex/meta_test/gen/swagger/restapi/operations"
	"github.com/yetialex/meta_test/gen/swagger/restapi/operations/core"
	"github.com/yetialex/meta_test/gen/swagger/restapi/operations/iba"
	"github.com/yetialex/meta_test/gen/swagger/restapi/operations/signals"
	"github.com/yetialex/meta_test/gen/swagger/restapi/operations/swagger"
)

//go:generate swagger generate server --target ../../swagger --name Meta --spec ../../../swagger-test/swagger.json

func configureFlags(api *operations.MetaAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.MetaAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.CoreCreateCoreDirectoryHandler == nil {
		api.CoreCreateCoreDirectoryHandler = core.CreateCoreDirectoryHandlerFunc(func(params core.CreateCoreDirectoryParams) middleware.Responder {
			return middleware.NotImplemented("operation core.CreateCoreDirectory has not yet been implemented")
		})
	}
	if api.CoreCreateCoreNodeHandler == nil {
		api.CoreCreateCoreNodeHandler = core.CreateCoreNodeHandlerFunc(func(params core.CreateCoreNodeParams) middleware.Responder {
			return middleware.NotImplemented("operation core.CreateCoreNode has not yet been implemented")
		})
	}
	if api.CoreCreateCoreSignalHandler == nil {
		api.CoreCreateCoreSignalHandler = core.CreateCoreSignalHandlerFunc(func(params core.CreateCoreSignalParams) middleware.Responder {
			return middleware.NotImplemented("operation core.CreateCoreSignal has not yet been implemented")
		})
	}
	if api.CoreDeleteDirectoryHandler == nil {
		api.CoreDeleteDirectoryHandler = core.DeleteDirectoryHandlerFunc(func(params core.DeleteDirectoryParams) middleware.Responder {
			return middleware.NotImplemented("operation core.DeleteDirectory has not yet been implemented")
		})
	}
	if api.CoreDeleteNodeHandler == nil {
		api.CoreDeleteNodeHandler = core.DeleteNodeHandlerFunc(func(params core.DeleteNodeParams) middleware.Responder {
			return middleware.NotImplemented("operation core.DeleteNode has not yet been implemented")
		})
	}
	if api.CoreDeleteSignalHandler == nil {
		api.CoreDeleteSignalHandler = core.DeleteSignalHandlerFunc(func(params core.DeleteSignalParams) middleware.Responder {
			return middleware.NotImplemented("operation core.DeleteSignal has not yet been implemented")
		})
	}
	if api.CoreGetCoreDirectoriesHandler == nil {
		api.CoreGetCoreDirectoriesHandler = core.GetCoreDirectoriesHandlerFunc(func(params core.GetCoreDirectoriesParams) middleware.Responder {
			return middleware.NotImplemented("operation core.GetCoreDirectories has not yet been implemented")
		})
	}
	if api.CoreGetCoreNodesHandler == nil {
		api.CoreGetCoreNodesHandler = core.GetCoreNodesHandlerFunc(func(params core.GetCoreNodesParams) middleware.Responder {
			return middleware.NotImplemented("operation core.GetCoreNodes has not yet been implemented")
		})
	}
	if api.CoreGetCoreSignalsHandler == nil {
		api.CoreGetCoreSignalsHandler = core.GetCoreSignalsHandlerFunc(func(params core.GetCoreSignalsParams) middleware.Responder {
			return middleware.NotImplemented("operation core.GetCoreSignals has not yet been implemented")
		})
	}
	if api.CoreGetDirectoryHandler == nil {
		api.CoreGetDirectoryHandler = core.GetDirectoryHandlerFunc(func(params core.GetDirectoryParams) middleware.Responder {
			return middleware.NotImplemented("operation core.GetDirectory has not yet been implemented")
		})
	}
	if api.IbaGetIBAGateByNameHandler == nil {
		api.IbaGetIBAGateByNameHandler = iba.GetIBAGateByNameHandlerFunc(func(params iba.GetIBAGateByNameParams) middleware.Responder {
			return middleware.NotImplemented("operation iba.GetIBAGateByName has not yet been implemented")
		})
	}
	if api.IbaGetIBAGateMntHandler == nil {
		api.IbaGetIBAGateMntHandler = iba.GetIBAGateMntHandlerFunc(func(params iba.GetIBAGateMntParams) middleware.Responder {
			return middleware.NotImplemented("operation iba.GetIBAGateMnt has not yet been implemented")
		})
	}
	if api.IbaGetIBAGateMntsHandler == nil {
		api.IbaGetIBAGateMntsHandler = iba.GetIBAGateMntsHandlerFunc(func(params iba.GetIBAGateMntsParams) middleware.Responder {
			return middleware.NotImplemented("operation iba.GetIBAGateMnts has not yet been implemented")
		})
	}
	if api.IbaGetIBAGatesHandler == nil {
		api.IbaGetIBAGatesHandler = iba.GetIBAGatesHandlerFunc(func(params iba.GetIBAGatesParams) middleware.Responder {
			return middleware.NotImplemented("operation iba.GetIBAGates has not yet been implemented")
		})
	}
	if api.IbaGetIBAMappingByGateAndMountHandler == nil {
		api.IbaGetIBAMappingByGateAndMountHandler = iba.GetIBAMappingByGateAndMountHandlerFunc(func(params iba.GetIBAMappingByGateAndMountParams) middleware.Responder {
			return middleware.NotImplemented("operation iba.GetIBAMappingByGateAndMount has not yet been implemented")
		})
	}
	if api.IbaGetIBAMappingByServerIDHandler == nil {
		api.IbaGetIBAMappingByServerIDHandler = iba.GetIBAMappingByServerIDHandlerFunc(func(params iba.GetIBAMappingByServerIDParams) middleware.Responder {
			return middleware.NotImplemented("operation iba.GetIBAMappingByServerID has not yet been implemented")
		})
	}
	if api.IbaGetIBAMappingByTopicNameHandler == nil {
		api.IbaGetIBAMappingByTopicNameHandler = iba.GetIBAMappingByTopicNameHandlerFunc(func(params iba.GetIBAMappingByTopicNameParams) middleware.Responder {
			return middleware.NotImplemented("operation iba.GetIBAMappingByTopicName has not yet been implemented")
		})
	}
	if api.IbaGetIBAServerHandler == nil {
		api.IbaGetIBAServerHandler = iba.GetIBAServerHandlerFunc(func(params iba.GetIBAServerParams) middleware.Responder {
			return middleware.NotImplemented("operation iba.GetIBAServer has not yet been implemented")
		})
	}
	if api.IbaGetIBAServersHandler == nil {
		api.IbaGetIBAServersHandler = iba.GetIBAServersHandlerFunc(func(params iba.GetIBAServersParams) middleware.Responder {
			return middleware.NotImplemented("operation iba.GetIBAServers has not yet been implemented")
		})
	}
	if api.CoreGetNodeHandler == nil {
		api.CoreGetNodeHandler = core.GetNodeHandlerFunc(func(params core.GetNodeParams) middleware.Responder {
			return middleware.NotImplemented("operation core.GetNode has not yet been implemented")
		})
	}
	if api.CoreGetNodeChildrenHandler == nil {
		api.CoreGetNodeChildrenHandler = core.GetNodeChildrenHandlerFunc(func(params core.GetNodeChildrenParams) middleware.Responder {
			return middleware.NotImplemented("operation core.GetNodeChildren has not yet been implemented")
		})
	}
	if api.CoreGetNodesTreeWithSignalsHandler == nil {
		api.CoreGetNodesTreeWithSignalsHandler = core.GetNodesTreeWithSignalsHandlerFunc(func(params core.GetNodesTreeWithSignalsParams) middleware.Responder {
			return middleware.NotImplemented("operation core.GetNodesTreeWithSignals has not yet been implemented")
		})
	}
	if api.CoreGetRootNodeHandler == nil {
		api.CoreGetRootNodeHandler = core.GetRootNodeHandlerFunc(func(params core.GetRootNodeParams) middleware.Responder {
			return middleware.NotImplemented("operation core.GetRootNode has not yet been implemented")
		})
	}
	if api.CoreGetSignalHandler == nil {
		api.CoreGetSignalHandler = core.GetSignalHandlerFunc(func(params core.GetSignalParams) middleware.Responder {
			return middleware.NotImplemented("operation core.GetSignal has not yet been implemented")
		})
	}
	if api.SignalsGetSignalsClassesHandler == nil {
		api.SignalsGetSignalsClassesHandler = signals.GetSignalsClassesHandlerFunc(func(params signals.GetSignalsClassesParams) middleware.Responder {
			return middleware.NotImplemented("operation signals.GetSignalsClasses has not yet been implemented")
		})
	}
	if api.SignalsGetSignalsValueTypesHandler == nil {
		api.SignalsGetSignalsValueTypesHandler = signals.GetSignalsValueTypesHandlerFunc(func(params signals.GetSignalsValueTypesParams) middleware.Responder {
			return middleware.NotImplemented("operation signals.GetSignalsValueTypes has not yet been implemented")
		})
	}
	if api.SwaggerGetSwaggerJSONHandler == nil {
		api.SwaggerGetSwaggerJSONHandler = swagger.GetSwaggerJSONHandlerFunc(func(params swagger.GetSwaggerJSONParams) middleware.Responder {
			return middleware.NotImplemented("operation swagger.GetSwaggerJSON has not yet been implemented")
		})
	}
	if api.IbaRegisterIBAGateHandler == nil {
		api.IbaRegisterIBAGateHandler = iba.RegisterIBAGateHandlerFunc(func(params iba.RegisterIBAGateParams) middleware.Responder {
			return middleware.NotImplemented("operation iba.RegisterIBAGate has not yet been implemented")
		})
	}
	if api.IbaRegisterIBAGateMntHandler == nil {
		api.IbaRegisterIBAGateMntHandler = iba.RegisterIBAGateMntHandlerFunc(func(params iba.RegisterIBAGateMntParams) middleware.Responder {
			return middleware.NotImplemented("operation iba.RegisterIBAGateMnt has not yet been implemented")
		})
	}
	if api.IbaRegisterIBAServerHandler == nil {
		api.IbaRegisterIBAServerHandler = iba.RegisterIBAServerHandlerFunc(func(params iba.RegisterIBAServerParams) middleware.Responder {
			return middleware.NotImplemented("operation iba.RegisterIBAServer has not yet been implemented")
		})
	}
	if api.IbaRegisterSignalHandler == nil {
		api.IbaRegisterSignalHandler = iba.RegisterSignalHandlerFunc(func(params iba.RegisterSignalParams) middleware.Responder {
			return middleware.NotImplemented("operation iba.RegisterSignal has not yet been implemented")
		})
	}
	if api.CoreUpdateDirectoryHandler == nil {
		api.CoreUpdateDirectoryHandler = core.UpdateDirectoryHandlerFunc(func(params core.UpdateDirectoryParams) middleware.Responder {
			return middleware.NotImplemented("operation core.UpdateDirectory has not yet been implemented")
		})
	}
	if api.IbaUpdateIBAGateMetadataHandler == nil {
		api.IbaUpdateIBAGateMetadataHandler = iba.UpdateIBAGateMetadataHandlerFunc(func(params iba.UpdateIBAGateMetadataParams) middleware.Responder {
			return middleware.NotImplemented("operation iba.UpdateIBAGateMetadata has not yet been implemented")
		})
	}
	if api.IbaUpdateIBAGateMntHandler == nil {
		api.IbaUpdateIBAGateMntHandler = iba.UpdateIBAGateMntHandlerFunc(func(params iba.UpdateIBAGateMntParams) middleware.Responder {
			return middleware.NotImplemented("operation iba.UpdateIBAGateMnt has not yet been implemented")
		})
	}
	if api.CoreUpdateNodeHandler == nil {
		api.CoreUpdateNodeHandler = core.UpdateNodeHandlerFunc(func(params core.UpdateNodeParams) middleware.Responder {
			return middleware.NotImplemented("operation core.UpdateNode has not yet been implemented")
		})
	}
	if api.CoreUpdateSignalHandler == nil {
		api.CoreUpdateSignalHandler = core.UpdateSignalHandlerFunc(func(params core.UpdateSignalParams) middleware.Responder {
			return middleware.NotImplemented("operation core.UpdateSignal has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
