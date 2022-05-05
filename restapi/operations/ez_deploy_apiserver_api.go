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
	"github.com/ez-deploy/ezdeploy/restapi/operations/project"
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

		ProjectCreateProjectHandler: project.CreateProjectHandlerFunc(func(params project.CreateProjectParams, principal *models.AuthInfo) middleware.Responder {
			return middleware.NotImplemented("operation project.CreateProject has not yet been implemented")
		}),
		IdentityCreateUserHandler: identity.CreateUserHandlerFunc(func(params identity.CreateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation identity.CreateUser has not yet been implemented")
		}),
		ProjectListProjectHandler: project.ListProjectHandlerFunc(func(params project.ListProjectParams, principal *models.AuthInfo) middleware.Responder {
			return middleware.NotImplemented("operation project.ListProject has not yet been implemented")
		}),
		IdentityLoginHandler: identity.LoginHandlerFunc(func(params identity.LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation identity.Login has not yet been implemented")
		}),
		IdentityLogoutHandler: identity.LogoutHandlerFunc(func(params identity.LogoutParams, principal *models.AuthInfo) middleware.Responder {
			return middleware.NotImplemented("operation identity.Logout has not yet been implemented")
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

	// ProjectCreateProjectHandler sets the operation handler for the create project operation
	ProjectCreateProjectHandler project.CreateProjectHandler
	// IdentityCreateUserHandler sets the operation handler for the create user operation
	IdentityCreateUserHandler identity.CreateUserHandler
	// ProjectListProjectHandler sets the operation handler for the list project operation
	ProjectListProjectHandler project.ListProjectHandler
	// IdentityLoginHandler sets the operation handler for the login operation
	IdentityLoginHandler identity.LoginHandler
	// IdentityLogoutHandler sets the operation handler for the logout operation
	IdentityLogoutHandler identity.LogoutHandler
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

	if o.ProjectCreateProjectHandler == nil {
		unregistered = append(unregistered, "project.CreateProjectHandler")
	}
	if o.IdentityCreateUserHandler == nil {
		unregistered = append(unregistered, "identity.CreateUserHandler")
	}
	if o.ProjectListProjectHandler == nil {
		unregistered = append(unregistered, "project.ListProjectHandler")
	}
	if o.IdentityLoginHandler == nil {
		unregistered = append(unregistered, "identity.LoginHandler")
	}
	if o.IdentityLogoutHandler == nil {
		unregistered = append(unregistered, "identity.LogoutHandler")
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

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/project/create"] = project.NewCreateProject(o.context, o.ProjectCreateProjectHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/user/create"] = identity.NewCreateUser(o.context, o.IdentityCreateUserHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/project/list"] = project.NewListProject(o.context, o.ProjectListProjectHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/user/login"] = identity.NewLogin(o.context, o.IdentityLoginHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/logout"] = identity.NewLogout(o.context, o.IdentityLogoutHandler)
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
