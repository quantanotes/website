package handlers

import (
	"context"
	"net/http"
	"quanta/internal/globals"
	"quanta/internal/model"
	"strings"

	"github.com/go-chi/chi/v5"
)

type signinRequest struct {
	Email    string
	Password string
}

var signin = jsonRequestWithWriter(func(ctx context.Context, w http.ResponseWriter, req signinRequest) error {
	if req.Email == "" || req.Password == "" {
		return globals.ErrBadRequest("Email and password are required. Please check your input and try again.")
	}

	req.Email = strings.ToLower(req.Email)

	sess, verified, err := model.VerifyUserAndCreateSession(ctx, req.Email, req.Password)
	if err != nil {
		return err
	} else if !verified {
		return globals.ErrUnauthorised("The email or password you entered is incorrect. Please check your credentials and try again.")
	}

	setSessionCookie(w, sess)
	w.Write([]byte("OK"))
	return nil
})

type signupRequest struct {
	Email         string `validate:"required,email"`
	Username      string `validate:"required,min=3,max=24,alphanumerichyphenunderscore"`
	FullName      string `validate:"required,min=1,max=255"`
	PreferredName string `validate:"required,min=1,max=255"`
	Password      string `validate:"required,password"`
}

func (r *signupRequest) toUser() model.User {
	return model.User{
		Email:         r.Email,
		Username:      r.Username,
		FullName:      r.FullName,
		PreferredName: r.PreferredName,
		Password:      &r.Password,
	}
}

var signup = jsonRequest(func(ctx context.Context, req signupRequest) error {
	if err := globals.Validator.Struct(req); err != nil {
		return globals.ErrBadRequest("")
	}
	return model.SendEmailVerification(ctx, req.toUser(), "/")
})

func verify(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	redirect := r.URL.Query().Get("redirect")
	if redirect == "" {
		redirect = "/"
	}

	user, ok, err := model.GetUserEmailVerificationCode(r.Context(), code)
	if err != nil {
		writeError(w, err)
		return
	} else if !ok {
		writeError(w, globals.ErrBadRequest("Your verification link does not exist, or has expired. Please check your verification link."))
		return
	}

	user.Email = strings.ToLower(user.Email)
	if user, err = model.CreateUser(r.Context(), user); err != nil {
		writeError(w, err)
		return
	}

	sess, _ := model.CreateSession(r.Context(), user.ID)
	setSessionCookie(w, sess)
	http.Redirect(w, r, redirect, http.StatusSeeOther)
}

func signout(w http.ResponseWriter, r *http.Request) {
	sess, ok := r.Context().Value(globals.SessionContextKey).(string)
	if !ok {
		// No session to expire
		return
	}
	if err := model.DeleteSession(r.Context(), sess); err != nil {
		writeError(w, err)
		return
	}
	setSessionCookie(w, "")
	http.Redirect(w, r, "/", http.StatusFound)
}

var details = protectedJSONResponse(func(ctx context.Context, userID string) (model.User, error) {
	user, err := model.GetUser(ctx, userID)
	user.Password = nil
	user.CustomerID = nil
	return user, err
})

func setSessionCookie(w http.ResponseWriter, sess string) {
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
