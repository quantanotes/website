package middleware

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"quanta/internal/globals"
	"quanta/internal/model"
)

func Private(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess, user, err := getSessionAndUserVerbose(w, r)
		if err != nil {
			return
		}
		ctx := storeSessionAndUser(r, sess, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Restricted(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess, user, err := getSessionAndUser(w, r)
		if err == nil {
			ctx := storeSessionAndUser(r, sess, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func getSessionAndUserVerbose(w http.ResponseWriter, r *http.Request) (string, string, error) {
	sess, err := r.Cookie("session")
	if err == http.ErrNoCookie {
		http.Error(w, globals.MsgUnauthorised, http.StatusUnauthorized)
		return "", "", err
	} else if err != nil {
		slog.Error(err.Error())
		http.Error(w, globals.MsgError, http.StatusInternalServerError)
		return "", "", err
	}

	userId, ok, err := model.GetSession(r.Context(), sess.Value)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, globals.MsgError, http.StatusInternalServerError)
		return "", "", err
	} else if !ok {
		http.Error(w, globals.MsgUnauthorised, http.StatusUnauthorized)
		return "", "", errors.New("session not found")
	}

	return sess.Value, userId, nil
}

func getSessionAndUser(w http.ResponseWriter, r *http.Request) (string, string, error) {
	sess, err := r.Cookie("session")
	if err == http.ErrNoCookie {
		return "", "", err
	} else if err != nil {
		return "", "", err
	}

	user, ok, err := model.GetSession(r.Context(), sess.Value)
	if err != nil {
		return "", "", err
	} else if !ok {
		return "", "", errors.New("session not found")
	}

	return sess.Value, user, nil
}

func storeSessionAndUser(r *http.Request, sess, user string) context.Context {
	ctx := context.WithValue(r.Context(), "user", user)
	ctx = context.WithValue(ctx, "session", sess)
	ctx = context.WithValue(ctx, "authorised", true)
	return ctx
}
