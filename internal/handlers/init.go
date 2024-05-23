package handlers

import (
	"quanta/internal/middleware"
	"quanta/internal/single"

	"github.com/go-chi/chi/v5"
	mw "github.com/go-chi/chi/v5/middleware"
)

func init() {
	r := single.Router

	r.Handle("/static/*", static())
	r.Handle("/assets/*", assets())

	r.Group(func(r chi.Router) {
		r.Use(mw.RequestID)
		r.Use(mw.RealIP)
		r.Use(mw.Logger)
		r.Use(mw.Recoverer)

		r.Get("/", page("index", nil))
		r.Get("/notes", restrictedPage("notes-public", "notes", nil))
		r.Get("/io", restrictedPage("io-public", "io", nil))
		r.Get("/space", page("space", nil))
		r.Get("/space/*", page("space", nil))
		r.Get("/heisenberg", restrictedPage("heisenberg-public", "heisenberg", nil))

		r.Get("/subscribe", page("subscribe", nil))

		r.Get("/signin", page("signin", nil))
		r.Get("/signup", page("signup", nil))
		r.Get("/verify", page("verify", nil))
		r.Get("/signin/*", page("signin", getQuery("redirect")))
		r.Get("/signup/*", page("signup", getQuery("redirect")))
		r.Get("/verify/*", page("verify", getQuery("redirect")))

		r.Post("/api/space/algorithm", spaceAlgorithm)
		r.Post("/api/space/children", spaceChildren)

		r.Group(func(r chi.Router) {
			r.Use(middleware.Private)

			r.Post("/api/notes/roots", notesRoots)
			r.Post("/api/notes/create", notesCreate)
			r.Post("/api/notes/update", notesUpdate)
			r.Post("/api/notes/delete", notesDelete)
			r.Post("/api/notes/move", notesMove)
			r.Post("/api/notes/children", notesChildren)
			r.Post("/api/notes/publish", notesPublish)
			r.Post("/api/notes/unpublish", notesUnpublish)

			r.Post("/api/maxwell/roots", maxwellRoots)
			r.Post("/api/maxwell/children", maxwellChildren)
			// r.Post("/api/maxwell/create", maxwellCreate)
			// r.Post("/api/maxwell/delete", maxwellUpdate)
			r.Post("/api/maxwell/chat", maxwellChat)

			r.Post("/api/permissions/grant", grant)
			r.Post("/api/permissions/links/retrieve", retrieveLinks)
			r.Post("/api/permissions/links/create", createLink)
			r.Post("/api/permissions/links/delete", deleteLink)
		})

		r.Post("/api/auth/signin", signin)
		r.Post("/api/auth/signup", signup)
		r.Post("/api/auth/verify", verify)
		r.Group(func(r chi.Router) {
			r.Use(middleware.Private)

			r.Post("/api/auth/signout", signout)
			r.Post("/api/auth/details", details)
		})
	})
}
