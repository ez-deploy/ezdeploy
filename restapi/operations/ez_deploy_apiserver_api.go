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

	"github.com/ez-deploy/ezdeploy/models"
	"github.com/ez-deploy/ezdeploy/restapi/operations/identity"
	"github.com/ez-deploy/ezdeploy/restapi/operations/pod"
	"github.com/ez-deploy/ezdeploy/restapi/operations/project"
	"github.com/ez-deploy/ezdeploy/restapi/operations/r_b_a_c"
	"github.com/ez-deploy/ezdeploy/restapi/operations/service"
)

// NewEzDeployApiserverAPI creates a new EzDeployApiserver instance
func NewEzDeployApiserverAPI(spec *loads.Document) *EzDeployApiserverAPI {
	return &EzDeployApiserverAPI{
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

		PodCheckPodTicketHandler: pod.CheckPodTicketHandlerFunc(func(params pod.CheckPodTicketParams) middleware.Responder {
			return middleware.NotImplemented("operation pod.CheckPodTicket has not yet been implemented")
		}),
		PodCreatePodTicketHandler: pod.CreatePodTicketHandlerFunc(func(params pod.CreatePodTicketParams, principal *models.AuthInfo) middleware.Responder {
			return middleware.NotImplemented("operation pod.CreatePodTicket has not yet been implemented")
		}),
		ProjectCreateProjectHandler: project.CreateProjectHandlerFunc(func(params project.CreateProjectParams, principal *models.AuthInfo) middleware.Responder {
			return middleware.NotImplemented("operation project.CreateProject has not yet been implemented")
		}),
		ServiceCreateServiceHandler: service.CreateServiceHandlerFunc(func(params service.CreateServiceParams, principal *models.AuthInfo) middleware.Responder {
			return middleware.NotImplemented("operation service.CreateService has not yet been implemented")
		}),
		ServiceCreateServiceVersionHandler: service.CreateServiceVersionHandlerFunc(func(params service.CreateServiceVersionParams, principal *models.AuthInfo) middleware.Responder {
			return middleware.NotImplemented("operation service.CreateServiceVersion has not yet been implemented")
		}),
		IdentityCreateUserHandler: identity.CreateUserHandlerFunc(func(params identity.CreateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation identity.CreateUser has not yet been implemented")
		}),
		ServiceDeleteServiceHandler: service.DeleteServiceHandlerFunc(func(params service.DeleteServiceParams, principal *models.AuthInfo) middleware.Responder {
			return middleware.NotImplemented("operation service.DeleteService has not yet been implemented")
		}),
		ProjectGetProjectHandler: project.GetProjectHandlerFunc(func(params project.GetProjectParams, principal *models.AuthInfo) middleware.Responder {
			return middleware.NotImplemented("operation project.GetProject has not yet been implemented")
		}),
		RbacGetProjectRBACHandler: r_b_a_c.GetProjectRBACHandlerFunc(func(params r_b_a_c.GetProjectRBACParams, principal *models.AuthInfo) middleware.Responder {
			return middleware.NotImplemented("operation r_b_a_c.GetProjectRBAC has not yet been implemented")
		}),
		ServiceGetServiceVersionHandler: service.GetServiceVersionHandlerFunc(func(params service.GetServiceVersionParams, principal *models.AuthInfo) middleware.Responder {
			return middleware.NotImplemented("operation service.GetServiceVersion has not yet been implemented")
		}),
		IdentityGetUserHandler: identity.GetUserHandlerFunc(func(params identity.GetUserParams) middleware.Responder {
			return middleware.NotImplemented("operation identity.GetUser has not yet been implemented")
		}),
		RbacGetUserRBACHandler: r_b_a_c.GetUserRBACHandlerFunc(func(params r_b_a_c.GetUserRBACParams, principal *models.AuthInfo) middleware.Responder {
			return middleware.NotImplemented("operation r_b_a_c.GetUserRBAC has not yet been implemented")
		}),
		ServiceListServiceHandler: service.ListServiceHandlerFunc(func(params service.ListServiceParams, principal *models.AuthInfo) middleware.Responder {
			return middleware.NotImplemented("operation service.ListService has not yet been implemented")
		}),
		ServiceListServicePodHandler: service.ListServicePodHandlerFunc(func(params service.ListServicePodParams, principal *models.AuthInfo) middleware.Responder {
			return middleware.NotImplemented("operation service.ListServicePod has not yet been implemented")
		}),
		ServiceListServiceVersionHandler: service.ListServiceVersionHandlerFunc(func(params service.ListServiceVersionParams, principal *models.AuthInfo) middleware.Responder {
			return middleware.NotImplemented("operation service.ListServiceVersion has not yet been implemented")
		}),
		IdentityLoginHandler: identity.LoginHandlerFunc(func(params identity.LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation identity.Login has not yet been implemented")
		}),
		IdentityLogoutHandler: identity.LogoutHandlerFunc(func(params identity.LogoutParams, principal *models.AuthInfo) middleware.Responder {
			return middleware.NotImplemented("operation identity.Logout has not yet been implemented")
		}),
		ServiceUpdateServiceDescriptionHandler: service.UpdateServiceDescriptionHandlerFunc(func(params service.UpdateServiceDescriptionParams, principal *models.AuthInfo) middleware.Responder {
			return middleware.NotImplemented("operation service.UpdateServiceDescription has not yet been implemented")
		}),
		ServiceUpdateServiceVersionHandler: service.UpdateServiceVersionHandlerFunc(func(params service.UpdateServiceVersionParams, principal *models.AuthInfo) middleware.Responder {
			return middleware.NotImplemented("operation service.UpdateServiceVersion has not yet been implemented")
		}),
		IdentityWhoamiHandler: identity.WhoamiHandlerFunc(func(params identity.WhoamiParams, principal *models.AuthInfo) middleware.Responder {
			return middleware.NotImplemented("operation identity.Whoami has not yet been implemented")
		}),

		// Applies when the "Cookie" header is set
		KeyAuth: func(token string) (*models.AuthInfo, error) {
			return nil, errors.NotImplemented("api key auth (key) Cookie from header param [Cookie] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*EzDeployApiserverAPI apiserver */
type EzDeployApiserverAPI struct {
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

	// KeyAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Cookie provided in the header
	KeyAuth func(string) (*models.AuthInfo, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// PodCheckPodTicketHandler sets the operation handler for the check pod ticket operation
	PodCheckPodTicketHandler pod.CheckPodTicketHandler
	// PodCreatePodTicketHandler sets the operation handler for the create pod ticket operation
	PodCreatePodTicketHandler pod.CreatePodTicketHandler
	// ProjectCreateProjectHandler sets the operation handler for the create project operation
	ProjectCreateProjectHandler project.CreateProjectHandler
	// ServiceCreateServiceHandler sets the operation handler for the create service operation
	ServiceCreateServiceHandler service.CreateServiceHandler
	// ServiceCreateServiceVersionHandler sets the operation handler for the create service version operation
	ServiceCreateServiceVersionHandler service.CreateServiceVersionHandler
	// IdentityCreateUserHandler sets the operation handler for the create user operation
	IdentityCreateUserHandler identity.CreateUserHandler
	// ServiceDeleteServiceHandler sets the operation handler for the delete service operation
	ServiceDeleteServiceHandler service.DeleteServiceHandler
	// ProjectGetProjectHandler sets the operation handler for the get project operation
	ProjectGetProjectHandler project.GetProjectHandler
	// RbacGetProjectRBACHandler sets the operation handler for the get project r b a c operation
	RbacGetProjectRBACHandler r_b_a_c.GetProjectRBACHandler
	// ServiceGetServiceVersionHandler sets the operation handler for the get service version operation
	ServiceGetServiceVersionHandler service.GetServiceVersionHandler
	// IdentityGetUserHandler sets the operation handler for the get user operation
	IdentityGetUserHandler identity.GetUserHandler
	// RbacGetUserRBACHandler sets the operation handler for the get user r b a c operation
	RbacGetUserRBACHandler r_b_a_c.GetUserRBACHandler
	// ServiceListServiceHandler sets the operation handler for the list service operation
	ServiceListServiceHandler service.ListServiceHandler
	// ServiceListServicePodHandler sets the operation handler for the list service pod operation
	ServiceListServicePodHandler service.ListServicePodHandler
	// ServiceListServiceVersionHandler sets the operation handler for the list service version operation
	ServiceListServiceVersionHandler service.ListServiceVersionHandler
	// IdentityLoginHandler sets the operation handler for the login operation
	IdentityLoginHandler identity.LoginHandler
	// IdentityLogoutHandler sets the operation handler for the logout operation
	IdentityLogoutHandler identity.LogoutHandler
	// ServiceUpdateServiceDescriptionHandler sets the operation handler for the update service description operation
	ServiceUpdateServiceDescriptionHandler service.UpdateServiceDescriptionHandler
	// ServiceUpdateServiceVersionHandler sets the operation handler for the update service version operation
	ServiceUpdateServiceVersionHandler service.UpdateServiceVersionHandler
	// IdentityWhoamiHandler sets the operation handler for the whoami operation
	IdentityWhoamiHandler identity.WhoamiHandler

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
func (o *EzDeployApiserverAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *EzDeployApiserverAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *EzDeployApiserverAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *EzDeployApiserverAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *EzDeployApiserverAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *EzDeployApiserverAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *EzDeployApiserverAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *EzDeployApiserverAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *EzDeployApiserverAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the EzDeployApiserverAPI
func (o *EzDeployApiserverAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.KeyAuth == nil {
		unregistered = append(unregistered, "CookieAuth")
	}

	if o.PodCheckPodTicketHandler == nil {
		unregistered = append(unregistered, "pod.CheckPodTicketHandler")
	}
	if o.PodCreatePodTicketHandler == nil {
		unregistered = append(unregistered, "pod.CreatePodTicketHandler")
	}
	if o.ProjectCreateProjectHandler == nil {
		unregistered = append(unregistered, "project.CreateProjectHandler")
	}
	if o.ServiceCreateServiceHandler == nil {
		unregistered = append(unregistered, "service.CreateServiceHandler")
	}
	if o.ServiceCreateServiceVersionHandler == nil {
		unregistered = append(unregistered, "service.CreateServiceVersionHandler")
	}
	if o.IdentityCreateUserHandler == nil {
		unregistered = append(unregistered, "identity.CreateUserHandler")
	}
	if o.ServiceDeleteServiceHandler == nil {
		unregistered = append(unregistered, "service.DeleteServiceHandler")
	}
	if o.ProjectGetProjectHandler == nil {
		unregistered = append(unregistered, "project.GetProjectHandler")
	}
	if o.RbacGetProjectRBACHandler == nil {
		unregistered = append(unregistered, "r_b_a_c.GetProjectRBACHandler")
	}
	if o.ServiceGetServiceVersionHandler == nil {
		unregistered = append(unregistered, "service.GetServiceVersionHandler")
	}
	if o.IdentityGetUserHandler == nil {
		unregistered = append(unregistered, "identity.GetUserHandler")
	}
	if o.RbacGetUserRBACHandler == nil {
		unregistered = append(unregistered, "r_b_a_c.GetUserRBACHandler")
	}
	if o.ServiceListServiceHandler == nil {
		unregistered = append(unregistered, "service.ListServiceHandler")
	}
	if o.ServiceListServicePodHandler == nil {
		unregistered = append(unregistered, "service.ListServicePodHandler")
	}
	if o.ServiceListServiceVersionHandler == nil {
		unregistered = append(unregistered, "service.ListServiceVersionHandler")
	}
	if o.IdentityLoginHandler == nil {
		unregistered = append(unregistered, "identity.LoginHandler")
	}
	if o.IdentityLogoutHandler == nil {
		unregistered = append(unregistered, "identity.LogoutHandler")
	}
	if o.ServiceUpdateServiceDescriptionHandler == nil {
		unregistered = append(unregistered, "service.UpdateServiceDescriptionHandler")
	}
	if o.ServiceUpdateServiceVersionHandler == nil {
		unregistered = append(unregistered, "service.UpdateServiceVersionHandler")
	}
	if o.IdentityWhoamiHandler == nil {
		unregistered = append(unregistered, "identity.WhoamiHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *EzDeployApiserverAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *EzDeployApiserverAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "key":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.KeyAuth(token)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *EzDeployApiserverAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *EzDeployApiserverAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
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
func (o *EzDeployApiserverAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
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
func (o *EzDeployApiserverAPI) HandlerFor(method, path string) (http.Handler, bool) {
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

// Context returns the middleware context for the ez deploy apiserver API
func (o *EzDeployApiserverAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *EzDeployApiserverAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/visit/pod/ticket/check"] = pod.NewCheckPodTicket(o.context, o.PodCheckPodTicketHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/visit/pod/ticket/create"] = pod.NewCreatePodTicket(o.context, o.PodCreatePodTicketHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/project/create"] = project.NewCreateProject(o.context, o.ProjectCreateProjectHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/service/create"] = service.NewCreateService(o.context, o.ServiceCreateServiceHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/service/version/create"] = service.NewCreateServiceVersion(o.context, o.ServiceCreateServiceVersionHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/user/create"] = identity.NewCreateUser(o.context, o.IdentityCreateUserHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/service/delete"] = service.NewDeleteService(o.context, o.ServiceDeleteServiceHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/project/get"] = project.NewGetProject(o.context, o.ProjectGetProjectHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/rbac/project/get"] = r_b_a_c.NewGetProjectRBAC(o.context, o.RbacGetProjectRBACHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/service/version/get"] = service.NewGetServiceVersion(o.context, o.ServiceGetServiceVersionHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/get"] = identity.NewGetUser(o.context, o.IdentityGetUserHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/rbac/user/get"] = r_b_a_c.NewGetUserRBAC(o.context, o.RbacGetUserRBACHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/service/list"] = service.NewListService(o.context, o.ServiceListServiceHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/service/pod/list"] = service.NewListServicePod(o.context, o.ServiceListServicePodHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/service/version/list"] = service.NewListServiceVersion(o.context, o.ServiceListServiceVersionHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/user/login"] = identity.NewLogin(o.context, o.IdentityLoginHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/logout"] = identity.NewLogout(o.context, o.IdentityLogoutHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/service/update/desc"] = service.NewUpdateServiceDescription(o.context, o.ServiceUpdateServiceDescriptionHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/service/update/deploy"] = service.NewUpdateServiceVersion(o.context, o.ServiceUpdateServiceVersionHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/whoami"] = identity.NewWhoami(o.context, o.IdentityWhoamiHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *EzDeployApiserverAPI) Serve(builder middleware.Builder) http.Handler {
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
func (o *EzDeployApiserverAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *EzDeployApiserverAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *EzDeployApiserverAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *EzDeployApiserverAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
