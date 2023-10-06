package api

import (
	"net/http"

	"github.com/flosch/pongo2/v6"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

// Faster to compile first at startup
var index = pongo2.Must(pongo2.FromFile("./dist/index.html"))

func templater(w http.ResponseWriter, r *http.Request) {
	err := index.ExecuteWriter(pongo2.Context{"query": r.FormValue("query")}, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func jwtVerifier(jwtInstance *jwtauth.JWTAuth) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return jwtauth.Verify(jwtInstance, jwtFromCookie)(next)
	}
}

func jwtFromCookie(r *http.Request) string {
	cookie, err := r.Cookie("token")
	if err != nil {
		return ""
	}
	return cookie.Value
}

func Router() http.Handler {
	r := chi.NewRouter()

	// Secured routes
	r.Group(func(r chi.Router) {
		// jwt middlewares
		r.Use(jwtVerifier(jwtInstance))
		r.Use(verifyAccessToken)

		// Routes
		r.Get("/api/v1/auth/refresh_token", refreshTokenHandler)
		r.Get("/api/v1/auth/logout", logoutHandler)

	})

	// Public routes
	r.Group(func(r chi.Router) {
		r.Post("/api/v1/auth/login", loginHandler)
		r.Get("/", templater)
		r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("./dist/assets"))))
		r.Handle("/images/*", http.StripPrefix("/images/", http.FileServer(http.Dir("./dist/images"))))
	})

	return r
}
