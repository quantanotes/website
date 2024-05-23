package handlers

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"quanta/internal/globals"
)

func okResponse(w http.ResponseWriter) {
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

func authorisedResponse(w http.ResponseWriter) {
	w.Write([]byte(globals.MsgAuthorised))
}

func jsonResponse(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		errorInternalResponse(w, err)
	}
}

func jsonRequest(w http.ResponseWriter, r *http.Request, v any) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		errorInternalResponse(w, err)
		return err
	}
	return nil
}

func errorResponse(w http.ResponseWriter, msg string, status int) {
	http.Error(w, msg, status)
}

func errorBadRequestResponse(w http.ResponseWriter) {
	http.Error(w, globals.MsgBadRequest, http.StatusBadRequest)
}

func errorUnauthorisedResponse(w http.ResponseWriter) {
	http.Error(w, globals.MsgUnauthorised, http.StatusUnauthorized)
}

func errorInternalResponse(w http.ResponseWriter, err error) {
	slog.Error(err.Error())
	http.Error(w, globals.MsgError, http.StatusInternalServerError)
}

func pubJSONResHandler[T any](h func(context.Context) (T, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := h(r.Context())
		if err != nil {
			errorInternalResponse(w, err)
			return
		}

		jsonResponse(w, res)
	}
}

func pubJSONReqJSONResHandler[T any, U any](h func(context.Context, T) (U, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req T
		if err := jsonRequest(w, r, &req); err != nil {
			return
		}

		res, err := h(r.Context(), req)
		if err != nil {
			errorInternalResponse(w, err)
			return
		}

		jsonResponse(w, res)
	}
}

func protJSONResHandler[T any](h func(context.Context, string) (T, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(string)
		res, err := h(r.Context(), user)
		if err != nil {
			errorInternalResponse(w, err)
			return
		}

		jsonResponse(w, res)
	}
}

func protJSONReqJSONResHandler[T any, U any](h func(context.Context, string, T) (U, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req T
		if err := jsonRequest(w, r, &req); err != nil {
			return
		}

		user := r.Context().Value("user").(string)
		res, err := h(r.Context(), user, req)
		if err != nil {
			errorInternalResponse(w, err)
			return
		}

		jsonResponse(w, res)
	}
}

func protJSONReqOKResHandler[T any](h func(context.Context, string, T) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req T
		if err := jsonRequest(w, r, &req); err != nil {
			return
		}

		user := r.Context().Value("user").(string)
		if err := h(r.Context(), user, req); err != nil {
			errorInternalResponse(w, err)
			return
		}

		okResponse(w)
	}
}
