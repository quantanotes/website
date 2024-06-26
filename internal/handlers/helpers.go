package handlers

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"quanta/internal/globals"
	"quanta/internal/middleware"
)

func readJSON(w http.ResponseWriter, r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		writeError(w, err)
		return err
	}
	return nil
}

func writeOK(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

func writeJSON(w http.ResponseWriter, v any) error {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		writeError(w, err)
		return err
	}
	return nil
}

func writeError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "text/plain")
	switch err := err.(type) {
	case globals.Error:
		w.WriteHeader(err.Status)
		w.Write([]byte(err.Message))
	default:
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something bad happened on our end. Please try again later."))
		slog.Error(err.Error())
	}
}

type jsonRequestFunc[Req any] func(context.Context, Req) error

func jsonRequestHandler[Req any](h jsonRequestFunc[Req]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Req
		if err := readJSON(w, r, &req); err != nil {
			writeError(w, err)
			return
		}
		if err := h(r.Context(), req); err != nil {
			writeError(w, err)
		}
		writeOK(w)
	}
}

type jsonResponseFunc[Res any] func(context.Context) (Res, error)

func jsonResponseHandler[Res any](h jsonResponseFunc[Res]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := h(r.Context())
		if err != nil {
			writeError(w, err)
		}
		writeJSON(w, res)
	}
}

type jsonRequestJSONResponseFunc[Req any, Res any] func(context.Context, Req) (Res, error)

func jsonRequestJSONResponseHandler[Req any, Res any](h jsonRequestJSONResponseFunc[Req, Res]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Req
		if err := readJSON(w, r, &req); err != nil {
			return
		}
		res, err := h(r.Context(), req)
		if err != nil {
			writeError(w, err)
			return
		}
		writeJSON(w, res)
	}
}

type jsonRequestWithWriterFunc[Req any] func(context.Context, http.ResponseWriter, Req) error

func jsonRequestWithWriterHandler[Req any](h jsonRequestWithWriterFunc[Req]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Req
		if err := readJSON(w, r, &req); err != nil {
			return
		}
		if err := h(r.Context(), w, req); err != nil {
			writeError(w, err)
			return
		}
	}
}

type protectedJSONRequestWithWriterFunc[Req any] func(context.Context, http.ResponseWriter, string, Req) error

func protectedJSONRequestWithWriterHandler[Req any](h protectedJSONRequestWithWriterFunc[Req]) http.HandlerFunc {
	hn := func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(globals.UserContextKey).(string)
		var req Req
		if err := readJSON(w, r, &req); err != nil {
			return
		}
		if err := h(r.Context(), w, userID, req); err != nil {
			writeError(w, err)
		}
		return
	}
	return middleware.Private(http.HandlerFunc(hn)).ServeHTTP
}

type protectedJSONRequestFunc[Req any] func(context.Context, string, Req) error

func protectedJSONRequestHandler[Req any](h protectedJSONRequestFunc[Req]) http.HandlerFunc {
	hn := func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(globals.UserContextKey).(string)
		var req Req
		if err := readJSON(w, r, &req); err != nil {
			return
		}
		if err := h(r.Context(), userID, req); err != nil {
			writeError(w, err)
			return
		}
		writeOK(w)
	}
	return middleware.Private(http.HandlerFunc(hn)).ServeHTTP
}

type protectedJSONResponseFunc[Res any] func(context.Context, string) (Res, error)

func protectedJSONResponseHandler[Res any](h protectedJSONResponseFunc[Res]) http.HandlerFunc {
	hn := func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(globals.UserContextKey).(string)
		res, err := h(r.Context(), userID)
		if err != nil {
			writeError(w, err)
			return
		}
		writeJSON(w, res)
	}
	return middleware.Private(http.HandlerFunc(hn)).ServeHTTP
}

type protectedJSONResponseWithRequestFunc[Res any] func(context.Context, *http.Request, string) (Res, error)

func protectedJSONResponseWithRequestHandler[Res any](h protectedJSONResponseWithRequestFunc[Res]) http.HandlerFunc {
	hn := func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(globals.UserContextKey).(string)
		res, err := h(r.Context(), r, userID)
		if err != nil {
			writeError(w, err)
			return
		}
		writeJSON(w, res)
	}
	return middleware.Private(http.HandlerFunc(hn)).ServeHTTP
}

type protectedJSONRequestJSONResponseFunc[Req any, Res any] func(context.Context, string, Req) (Res, error)

func protectedJSONRequestJSONResponseHandler[Req any, Res any](h protectedJSONRequestJSONResponseFunc[Req, Res]) http.HandlerFunc {
	hn := func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(globals.UserContextKey).(string)
		var req Req
		if err := readJSON(w, r, &req); err != nil {
			return
		}
		res, err := h(r.Context(), userID, req)
		if err != nil {
			writeError(w, err)
			return
		}
		writeJSON(w, res)
	}
	return middleware.Private(http.HandlerFunc(hn)).ServeHTTP
}
