package middleware

import (
	"context"
	"net/http"
	"quanta/internal/globals"
	"quanta/internal/model"
)

func Private(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess, user, err := getSessionAndUser(r)
		if err != nil {
			writeError(w, err)
		} else {
			ctx := storeSessionAndUser(r, sess, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

func Restricted(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess, user, err := getSessionAndUser(r)
		if err != nil {
			next.ServeHTTP(w, r)
		} else {
			ctx := storeSessionAndUser(r, sess, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

func getSessionAndUser(r *http.Request) (string, string, error) {
	sess, err := r.Cookie("session")
	if err == http.ErrNoCookie {
		return "", "", globals.ErrUnauthorised("Session expired or not found. Please log in again to continue.")
	} else if err != nil {
		return "", "", err
	}

	userID, ok, err := model.GetSession(r.Context(), sess.Value)
	if err != nil {
		return "", "", err
	} else if !ok {
		return "", "", globals.ErrUnauthorised("Session expired or not found. Please log in again to continue.")
	}

	return sess.Value, userID, nil
}

func storeSessionAndUser(r *http.Request, sess, userID string) context.Context {
	ctx := context.WithValue(r.Context(), globals.UserContextKey, userID)
	ctx = context.WithValue(ctx, globals.SessionContextKey, sess)
	ctx = context.WithValue(ctx, globals.AuthorisedContextKey, true)
	return ctx
}
