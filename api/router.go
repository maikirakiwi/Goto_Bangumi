package api

import (
	"net/http"
	"strings"
	"time"

	"github.com/arl/statsviz"
	"github.com/flosch/pongo2/v6"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
)

// Faster to compile first at startup
var index = pongo2.Must(pongo2.FromFile("./webui/index.html"))

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

func Router(dev bool) http.Handler {
	r := chi.NewRouter()

	// Dev mode middlewares
	if dev {
		r.Use(hlog.NewHandler(log.Logger))
		r.Use(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
			hlog.FromRequest(r).Info().Msgf("%s -> %d %s %s", strings.Replace(r.RemoteAddr, "127.0.0.1", "", 1), status, r.Method, r.URL)
		}))

		// Statsviz
		srv, _ := statsviz.NewServer()
		r.Get("/debug/statsviz/ws", srv.Ws())
		r.Get("/debug/statsviz", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/debug/statsviz/", http.StatusMovedPermanently)
		})
		r.Handle("/debug/statsviz/*", srv.Index())
		log.Info().Msg("Statsviz available at /debug/statsviz")
	}

	// Secured routes
	r.Group(func(r chi.Router) {
		// jwt middlewares, not loaded in dev mode
		if !dev {
			r.Use(jwtVerifier(jwtInstance))
			r.Use(verifyAccessToken)
		}

		// ./auth.go
		r.Route("/api/v1/auth", func(r chi.Router) {
			r.Get("/refresh_token", refreshTokenHandler)
			r.Get("/logout", logoutHandler)
			r.Post("/update", updateUserHandler)
		})

		// ./config.go
		r.Route("/api/v1/config", func(r chi.Router) {
			r.Get("/get", GetConfigHandler)
			r.Patch("/update", UpdateConfigHandler)
		})

		// ./bangumi.go
		r.Route("/api/v1/bangumi", func(r chi.Router) {
			r.Get("/get/all", getAllBangumiHandler)
			r.Get("/get/{bangumi_id}", getBangumiHandler)
			r.Patch("/update/{bangumi_id}", updateBangumiHandler)
			/*
				r.Delete("/delete/{bangumi_id}", GetConfigHandler)
				r.Delete("/delete/many", GetConfigHandler)
				r.Delete("/disable/{bangumi_id}", GetConfigHandler)
				r.Delete("/disable/many", GetConfigHandler)
				r.Get("/enable/{bangumi_id}", GetConfigHandler)
				r.Get("/refresh/poster/all", GetConfigHandler)
				r.Get("/reset/all", GetConfigHandler)
			*/
		})

		// ./program.go
		r.Get("/api/v1/status", ProgramStatusHandler)
		r.Get("/api/v1/log", LogOutputHandler)
		r.Get("/api/v1/log/clear", LogClearHandler)

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
