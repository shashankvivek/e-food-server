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

	"e-food/restapi/operations/guest"
	"e-food/restapi/operations/menu"
	"e-food/restapi/operations/products"
	"e-food/restapi/operations/user"
)

// NewEFoodAPI creates a new EFood instance
func NewEFoodAPI(spec *loads.Document) *EFoodAPI {
	return &EFoodAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		GuestAddItemHandler: guest.AddItemHandlerFunc(func(params guest.AddItemParams) middleware.Responder {
			return middleware.NotImplemented("operation guest.AddItem has not yet been implemented")
		}),
		GuestAddSessionHandler: guest.AddSessionHandlerFunc(func(params guest.AddSessionParams) middleware.Responder {
			return middleware.NotImplemented("operation guest.AddSession has not yet been implemented")
		}),
		UserAddToCartHandler: user.AddToCartHandlerFunc(func(params user.AddToCartParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.AddToCart has not yet been implemented")
		}),
		MenuCategoryListHandler: menu.CategoryListHandlerFunc(func(params menu.CategoryListParams) middleware.Responder {
			return middleware.NotImplemented("operation menu.CategoryList has not yet been implemented")
		}),
		UserGetCartHandler: user.GetCartHandlerFunc(func(params user.GetCartParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.GetCart has not yet been implemented")
		}),
		ProductsGetFromSubCategoryHandler: products.GetFromSubCategoryHandlerFunc(func(params products.GetFromSubCategoryParams) middleware.Responder {
			return middleware.NotImplemented("operation products.GetFromSubCategory has not yet been implemented")
		}),
		GuestGetItemsHandler: guest.GetItemsHandlerFunc(func(params guest.GetItemsParams) middleware.Responder {
			return middleware.NotImplemented("operation guest.GetItems has not yet been implemented")
		}),
		UserLoginHandler: user.LoginHandlerFunc(func(params user.LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation user.Login has not yet been implemented")
		}),
		UserRemoveFromCartHandler: user.RemoveFromCartHandlerFunc(func(params user.RemoveFromCartParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.RemoveFromCart has not yet been implemented")
		}),
		GuestRemoveItemHandler: guest.RemoveItemHandlerFunc(func(params guest.RemoveItemParams) middleware.Responder {
			return middleware.NotImplemented("operation guest.RemoveItem has not yet been implemented")
		}),
		UserCheckoutHandler: user.CheckoutHandlerFunc(func(params user.CheckoutParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.Checkout has not yet been implemented")
		}),

		// Applies when the "Authorization" header is set
		BearerAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (Bearer) Authorization from header param [Authorization] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*EFoodAPI the e food API */
type EFoodAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

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

	// BearerAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	BearerAuth func(string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// GuestAddItemHandler sets the operation handler for the add item operation
	GuestAddItemHandler guest.AddItemHandler
	// GuestAddSessionHandler sets the operation handler for the add session operation
	GuestAddSessionHandler guest.AddSessionHandler
	// UserAddToCartHandler sets the operation handler for the add to cart operation
	UserAddToCartHandler user.AddToCartHandler
	// MenuCategoryListHandler sets the operation handler for the category list operation
	MenuCategoryListHandler menu.CategoryListHandler
	// UserGetCartHandler sets the operation handler for the get cart operation
	UserGetCartHandler user.GetCartHandler
	// ProductsGetFromSubCategoryHandler sets the operation handler for the get from sub category operation
	ProductsGetFromSubCategoryHandler products.GetFromSubCategoryHandler
	// GuestGetItemsHandler sets the operation handler for the get items operation
	GuestGetItemsHandler guest.GetItemsHandler
	// UserLoginHandler sets the operation handler for the login operation
	UserLoginHandler user.LoginHandler
	// UserRemoveFromCartHandler sets the operation handler for the remove from cart operation
	UserRemoveFromCartHandler user.RemoveFromCartHandler
	// GuestRemoveItemHandler sets the operation handler for the remove item operation
	GuestRemoveItemHandler guest.RemoveItemHandler
	// UserCheckoutHandler sets the operation handler for the checkout operation
	UserCheckoutHandler user.CheckoutHandler
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

// SetDefaultProduces sets the default produces media type
func (o *EFoodAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *EFoodAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *EFoodAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *EFoodAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *EFoodAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *EFoodAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *EFoodAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the EFoodAPI
func (o *EFoodAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.BearerAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.GuestAddItemHandler == nil {
		unregistered = append(unregistered, "guest.AddItemHandler")
	}
	if o.GuestAddSessionHandler == nil {
		unregistered = append(unregistered, "guest.AddSessionHandler")
	}
	if o.UserAddToCartHandler == nil {
		unregistered = append(unregistered, "user.AddToCartHandler")
	}
	if o.MenuCategoryListHandler == nil {
		unregistered = append(unregistered, "menu.CategoryListHandler")
	}
	if o.UserGetCartHandler == nil {
		unregistered = append(unregistered, "user.GetCartHandler")
	}
	if o.ProductsGetFromSubCategoryHandler == nil {
		unregistered = append(unregistered, "products.GetFromSubCategoryHandler")
	}
	if o.GuestGetItemsHandler == nil {
		unregistered = append(unregistered, "guest.GetItemsHandler")
	}
	if o.UserLoginHandler == nil {
		unregistered = append(unregistered, "user.LoginHandler")
	}
	if o.UserRemoveFromCartHandler == nil {
		unregistered = append(unregistered, "user.RemoveFromCartHandler")
	}
	if o.GuestRemoveItemHandler == nil {
		unregistered = append(unregistered, "guest.RemoveItemHandler")
	}
	if o.UserCheckoutHandler == nil {
		unregistered = append(unregistered, "user.CheckoutHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *EFoodAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *EFoodAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "Bearer":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.BearerAuth)

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *EFoodAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *EFoodAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
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
func (o *EFoodAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
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
func (o *EFoodAPI) HandlerFor(method, path string) (http.Handler, bool) {
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

// Context returns the middleware context for the e food API
func (o *EFoodAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *EFoodAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/guest/cart"] = guest.NewAddItem(o.context, o.GuestAddItemHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/sessionInfo"] = guest.NewAddSession(o.context, o.GuestAddSessionHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/user/cart"] = user.NewAddToCart(o.context, o.UserAddToCartHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/categories"] = menu.NewCategoryList(o.context, o.MenuCategoryListHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/cart"] = user.NewGetCart(o.context, o.UserGetCartHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/productListBySubCategory/{id}"] = products.NewGetFromSubCategory(o.context, o.ProductsGetFromSubCategoryHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/guest/cart"] = guest.NewGetItems(o.context, o.GuestGetItemsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/login"] = user.NewLogin(o.context, o.UserLoginHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/user/cart"] = user.NewRemoveFromCart(o.context, o.UserRemoveFromCartHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/guest/cart"] = guest.NewRemoveItem(o.context, o.GuestRemoveItemHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/checkoutCart"] = user.NewCheckout(o.context, o.UserCheckoutHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *EFoodAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *EFoodAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *EFoodAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *EFoodAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *EFoodAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
