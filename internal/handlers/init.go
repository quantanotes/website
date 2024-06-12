package handlers

import (
	"quanta/internal/middleware"

	"github.com/go-chi/chi/v5"
	mw "github.com/go-chi/chi/v5/middleware"
)

var R = chi.NewRouter()

func init() {
	r := R

	r.Handle("/static/*", static())
	r.Handle("/assets/*", assets())

	r.Group(func(r chi.Router) {
		r.Use(mw.RequestID)
		r.Use(mw.RealIP)
		r.Use(mw.Logger)
		r.Use(mw.Recoverer)

		r.Get("/notes", restrictedPage("notes-public", "notes", nil))
		r.Get("/io", restrictedPage("io-public", "io", nil))
		r.Get("/heisenberg", restrictedPage("heisenberg-public", "heisenberg", nil))

		r.Get("/", page("index", nil))
		r.Get("/space", page("space", nil))
		r.Get("/space/*", page("space", nil))

		r.Get("/signin", page("signin", getQuery("redirect")))
		r.Get("/signup", page("signup", getQuery("redirect")))
		// TODO: render a page in case of error
		r.Get("/verify/{code}", verify)

		r.Post("/api/node/root/{category}", nodeRoot)
		r.Post("/api/node/children", nodeChildren)
		r.Post("/api/node/create", nodeCreate)
		r.Post("/api/node/update", nodeUpdate)
		r.Post("/api/node/delete", nodeDelete)
		r.Post("/api/node/move", nodeMove)

		r.With(middleware.Private).Post("/api/io/chat", ioChat)

		r.With(middleware.Private).Post("/api/auth/signout", signout)
		r.Post("/api/auth/signin", signin)
		r.Post("/api/auth/signup", signup)
		r.Post("/api/auth/details", details)
	})
}
