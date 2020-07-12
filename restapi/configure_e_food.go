// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"e-food/clients"
	"e-food/handlers"
	"e-food/pkg/utils"
	"e-food/restapi/operations"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/google/uuid"
	"github.com/rs/cors"
	"net/http"
)

//go:generate swagger generate server --target ..\..\e-food-server --name EFood --spec ..\swagger.yaml

func configureFlags(api *operations.EFoodAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.EFoodAPI) http.Handler {
	// configure the api here

	api.ServeError = errors.ServeError

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	clientBuilder := clients.NewClientBuilder()

	dbClient := clientBuilder.BuildSqlClient()

	api.BearerAuth = utils.ValidateHeader

	api.UserLoginHandler = handlers.NewUserLoginHandler(dbClient)

	api.MenuCategoryListHandler = handlers.NewMenuCategoryHandler(dbClient)

	api.ProductsGetFromSubCategoryHandler = handlers.NewProductsFromSubCategoryHandler(dbClient)

	api.GuestGetItemsHandler = handlers.NewCartGetItemsHandler(dbClient)

	api.GuestAddItemHandler = handlers.NewCartAddItemHandler(dbClient)

	api.GuestRemoveItemHandler = handlers.NewCartRemoveItemHandler(dbClient)

	api.GuestAddSessionHandler = handlers.NewGuestAddSessionHandler(dbClient)

	api.UserAddToCartHandler = handlers.NewUserAddToCartHandler(dbClient)

	api.UserGetCartHandler = handlers.NewUserGetCartHandler(dbClient)

	api.UserRemoveFromCartHandler = handlers.NewUserRemoveFromCartHandler(dbClient)

	api.UserCheckoutHandler = handlers.NewCartCheckoutHandler(dbClient)

	api.UserRegisterHandler = handlers.NewUserRegisterHandler(dbClient)

	api.AdminGenerateCouponHandler = handlers.NewAdminGenerateCouponHandler(dbClient)

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

// The middleware configuration is for the handlers executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	corsHandler := cors.New(cors.Options{
		Debug:            false,
		AllowedHeaders:   []string{"*"}, // TODO: config accordingly
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{},
		AllowCredentials: true,
		MaxAge:           1000,
	})
	h := corsHandler.Handler(handler)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var _, err = r.Cookie("guest_session")
		// cookie not set
		if err != nil {
			cookie := &http.Cookie{
				Name:     "guest_session",
				Value:    uuid.New().String(),
				SameSite: http.SameSiteDefaultMode,
				MaxAge:   260000}
			http.SetCookie(w, cookie)
		}
		h.ServeHTTP(w, r)
	})
}
