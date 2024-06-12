package globals

type ContextKey string

const (
	UserContextKey       = ContextKey("user")
	SessionContextKey    = ContextKey("session")
	AuthorisedContextKey = ContextKey("authorised")
)
