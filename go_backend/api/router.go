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

		// ./auth.go
		r.Route("/api/v1/auth", func(r chi.Router) {
			r.Get("/refresh_token", refreshTokenHandler)
			r.Get("/logout", logoutHandler)
		})

		// ./config.go
		r.Route("/api/v1/config", func(r chi.Router) {
			r.Get("/get", GetConfigHandler)
			r.Patch("/update", UpdateConfigHandler)
		})

		// ./bangumi.go
		r.Route("/api/v1/bangumi", func(r chi.Router) {
			r.Get("/get/all", getAllBangumiHandler)
			/*
				r.Get("/get/{bangumi_id}", GetConfigHandler)
				r.Patch("/update/{bangumi_id}", GetConfigHandler)
				r.Delete("/delete/{bangumi_id}", GetConfigHandler)
				r.Delete("/delete/many", GetConfigHandler)
				r.Delete("/disable/{bangumi_id}", GetConfigHandler)
				r.Delete("/disable/many", GetConfigHandler)
				r.Get("/enable/{bangumi_id}", GetConfigHandler)
				r.Get("/refresh/poster/all", GetConfigHandler)
				r.Get("/reset/all", GetConfigHandler)
			*/
		})

	})

	// Public routes
	r.Group(func(r chi.Router) {
		// ./auth.go
		r.Post("/api/v1/auth/login", loginHandler)

		r.Get("/", templater)
		r.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("./dist/assets"))))
		r.Handle("/images/*", http.StripPrefix("/images/", http.FileServer(http.Dir("./dist/images"))))
	})

	return r
}
