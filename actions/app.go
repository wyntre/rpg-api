package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo-pop/v2/pop/popmw"
	"github.com/gobuffalo/envy"
	forcessl "github.com/gobuffalo/mw-forcessl"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
	"github.com/unrolled/secure"

	"github.com/dgrijalva/jwt-go"
	contenttype "github.com/gobuffalo/mw-contenttype"
	tokenauth "github.com/gobuffalo/mw-tokenauth"
	"github.com/gobuffalo/x/sessions"
	"github.com/rs/cors"
	"github.com/wyntre/rpg_api/models"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
//
// Routing, middleware, groups, etc... are declared TOP -> DOWN.
// This means if you add a middleware to `app` *after* declaring a
// group, that group will NOT have that new middleware. The same
// is true of resource declarations as well.
//
// It also means that routes are checked in the order they are declared.
// `ServeFiles` is a CATCH-ALL route, so it should always be
// placed last in the route declarations, as it will prevent routes
// declared after it to never be called.
func App() *buffalo.App {
	if app == nil {
		corsHandler := cors.New(cors.Options{
			AllowedOrigins: []string{"http://localhost*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
			AllowedHeaders: []string{"Authorization", "Content-Type"},
		})

		app = buffalo.New(buffalo.Options{
			Env:          ENV,
			SessionStore: sessions.Null{},
			PreWares: []buffalo.PreWare{
				corsHandler.Handler,
			},
			SessionName: "_rpg_api_session",
		})

		// Automatically redirect to SSL
		app.Use(forceSSL())

		if ENV == "test" {
			envy.Set("JWT_PUBLIC_KEY", "test_keys/test_key_public.pem")
			envy.Set("JWT_PRIVATE_KEY", "test_keys/test_key_private.pem")
		}

		// Setup JWT
		TokenAuthentication := tokenauth.New(tokenauth.Options{
			SignMethod: jwt.SigningMethodRS256,
		})
		app.Use(TokenAuthentication)

		// Log request parameters (filters apply).
		app.Use(paramlogger.ParameterLogger)

		// Set the request content type to JSON
		app.Use(contenttype.Set("application/json"))

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.Connection)
		// Remove to disable this.
		app.Use(popmw.Transaction(models.DB))
		app.Use(checkToken)

		app.GET("/", HomeHandler)

		//define API version
		v1 := app.Group("/v1")

		//Routes for Auth
		auth := v1.Group("/auth")
		auth.POST("/", AuthCreate)
		auth.DELETE("/", AuthDestroy)
		// auth.Middleware.Skip(Authorize, AuthCreate)
		auth.Middleware.Skip(TokenAuthentication, AuthCreate)
		auth.Middleware.Skip(checkToken, AuthCreate)

		//Routes for User registration
		users := v1.Group("/users")
		users.POST("/", UsersCreate)
		users.Middleware.Skip(TokenAuthentication, UsersCreate)
		users.Middleware.Skip(checkToken, UsersCreate)

		// Routers for Characters
		characters := v1.Group("/characters")
		characters.GET("/", CharactersList)
		characters.POST("/new", CharactersCreate)
		characters.GET("/{id}", CharactersShow)
		characters.PUT("/{id}", CharactersUpdate)
		characters.DELETE("/{id}", CharactersDestroy)

		campaigns := v1.Group("/campaigns")
		campaignsResource := CampaignsResource{}
		campaigns.GET("/", campaignsResource.List)
		campaigns.POST("/new", campaignsResource.Create)
		campaigns.GET("/show/{id}", campaignsResource.Show)
		campaigns.PUT("/{id}", campaignsResource.Update)
		campaigns.DELETE("/{id}", campaignsResource.Destroy)
		// campaigns.GET("/{id}/quests", campaignsResource.Quests)

		quests := v1.Group("/quests")
		questsResource := QuestsResource{}
		quests.GET("/{id}", questsResource.List)
		quests.POST("/new", questsResource.Create)
		quests.GET("/show/{id}", questsResource.Show)
		quests.PUT("/{id}", questsResource.Update)
		quests.DELETE("/{id}", questsResource.Destroy)

		maps := v1.Group("/maps")
		mapsResource := MapsResource{}
		maps.GET("/{id}", mapsResource.List)
		maps.POST("/new", mapsResource.Create)
		maps.GET("/show/{id}", mapsResource.Show)
		maps.PUT("/{id}", mapsResource.Update)
		maps.DELETE("/{id}", mapsResource.Destroy)

		levels := v1.Group("/levels")
		levelsResource := LevelsResource{}
		levels.GET("/{id}", levelsResource.List)
		levels.POST("/new", levelsResource.Create)
		levels.GET("/show/{id}", levelsResource.Show)
		levels.PUT("/{id}", levelsResource.Update)
		levels.DELETE("/{id}", levelsResource.Destroy)

		tileCategories := v1.Group("/tile_categories")
		tileCategoriesResource := TileCategoriesResource{}
		tileCategories.GET("/", tileCategoriesResource.List)
		tileCategories.POST("/new", tileCategoriesResource.Create)
		tileCategories.GET("/show/{id}", tileCategoriesResource.Show)
		tileCategories.PUT("/{id}", tileCategoriesResource.Update)
		tileCategories.DELETE("/{id}", tileCategoriesResource.Destroy)
		tileCategories.GET("/{id}/tile_types", tileCategoriesResource.Types)

		tileTypes := v1.Group("/tile_types")
		tileTypesResource := TileTypesResource{}
		tileTypes.POST("/new", tileTypesResource.Create)
		tileTypes.GET("/{id}", tileTypesResource.Show)
		tileTypes.PUT("/{id}", tileTypesResource.Update)
		tileTypes.DELETE("/{id}", tileTypesResource.Destroy)
	}

	return app
}

// forceSSL will return a middleware that will redirect an incoming request
// if it is not HTTPS. "http://example.com" => "https://example.com".
// This middleware does **not** enable SSL. for your application. To do that
// we recommend using a proxy: https://gobuffalo.io/en/docs/proxy
// for more information: https://github.com/unrolled/secure/
func forceSSL() buffalo.MiddlewareFunc {
	return forcessl.Middleware(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}
