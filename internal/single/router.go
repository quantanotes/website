package single

import "github.com/go-chi/chi/v5"

var (
	Router chi.Router
)

func init() {
	Router = chi.NewRouter()
}
