package route

import (
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
)

var tokenAuth *jwtauth.JWTAuth

type Route struct {
	Router *chi.Mux
}

func initTokenAuth() {
	tokenAuth = jwtauth.New("HS256", []byte(os.Getenv("SECRET_JWT")), nil)
}

func New() *Route {
	initTokenAuth()
	r := &Route{
		Router: chi.NewRouter(),
	}
	r.Router.Use(middleware.Logger)
	r.Router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.initRoutes()
	return r
}

func (ro *Route) initRoutes() {

	ro.Router.Group(func(r chi.Router) {
		// r.Use(jwtauth.Verifier(tokenAuth))
		// r.Use(jwtauth.Authenticator)
		r.Get("/api/search/{term}-{from}-{maxResults}", ro.CreateSearchHandler())
	})

	ro.Router.Group(func(r chi.Router) {
		r.Post("/api/token", ro.getToken())
	})

}
