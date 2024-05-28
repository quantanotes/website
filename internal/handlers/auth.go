package handlers

import (
	"net/http"
	"quanta/internal/model"
	"quanta/internal/single"
	"strings"

	"github.com/go-chi/chi/v5"
)

func signin(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string
		Password string
	}
	if err := jsonRequest(w, r, &req); err != nil {
		return
	}

	if req.Email == "" || req.Password == "" {
		errorBadRequestResponse(w)
		return
	}

	req.Email = strings.ToLower(req.Email)

	sess, verified, err := model.VerifyUserAndCreateSession(r.Context(), req.Email, req.Password)
	if err != nil {
		errorInternalResponse(w, err)
		return
	} else if !verified {
		errorUnauthorisedResponse(w)
		return
	}

	setSessionCookie(sess, w)
	authorisedResponse(w)
}

func signup(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email         string `validate:"required,email"`
		Username      string `validate:"required,min=3,max=24,alphanumerichyphenunderscore"`
		FullName      string `validate:"required,min=1,max=255"`
		PreferredName string `validate:"required,min=1,max=255"`
		Password      string `validate:"required,password"`
	}

	redirect := r.URL.Query().Get("redirect")

	if err := jsonRequest(w, r, &req); err != nil {
		return
	}

	if err := single.Validator.Struct(req); err != nil {
		errorBadRequestResponse(w)
		return
	}

	if err := model.SendVerificationEmail(r.Context(), model.User{
		Email:         req.Email,
		Username:      req.Username,
		FullName:      req.FullName,
		PreferredName: req.PreferredName,
		Password:      &req.Password,
	}, redirect); err != nil {
		errorInternalResponse(w, err)
		return
	}

	okResponse(w)
}

func verify(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	redirect := r.URL.Query().Get("redirect")
	if redirect == "" {
		redirect = "/"
	}

	user, ok, err := model.GetVerificationUser(r.Context(), code)
	if err != nil {
		errorInternalResponse(w, err)
		return
	} else if !ok {
		errorResponse(w, "Verification code does not exist or has expired.", http.StatusBadRequest)
		return
	}

	user.Email = strings.ToLower(user.Email)
	if user, err = model.CreateUser(r.Context(), user); err != nil {
		errorInternalResponse(w, err)
		return
	}

	sess, _ := model.CreateSession(r.Context(), user.Id)
	setSessionCookie(sess, w)
	http.Redirect(w, r, redirect, http.StatusSeeOther)
}

func signout(w http.ResponseWriter, r *http.Request) {
	sess := r.Context().Value("session").(string)
	if err := model.DeleteSession(r.Context(), sess); err != nil {
		errorInternalResponse(w, err)
		return
	}
	setSessionCookie("", w)
	http.Redirect(w, r, "/", http.StatusFound)
}

func details(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(string)
	details, err := model.GetUser(r.Context(), user)
	if err != nil {
		errorInternalResponse(w, err)
		return
	}
	details.Password = nil
	jsonResponse(w, details)
}

func setSessionCookie(sess string, w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:     "session",
		Value:    sess,
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 30,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
}
