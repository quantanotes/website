package handlers

import (
	"net/http"
	"quanta/internal/model"
	"quanta/internal/single"
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

	sess, verified, err := model.VerifyUserAndCreateSession(r.Context(), req.Email, req.Password)
	if err != nil {
		errorInternalResponse(w, err)
		return
	} else if !verified {
		errorUnauthorisedResponse(w)
		return
	}

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
	}); err != nil {
		errorInternalResponse(w, err)
		return
	}

	okResponse(w)
}

func verify(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Code string
	}

	if err := jsonRequest(w, r, &req); err != nil {
		return
	}

	user, ok, err := model.GetVerificationUser(r.Context(), req.Code)
	if err != nil {
		errorInternalResponse(w, err)
		return
	} else if !ok {
		errorResponse(w, "Verification code does not exist or has expired.", http.StatusBadRequest)
		return
	}

	if _, err := model.CreateUser(r.Context(), user); err != nil {
		errorInternalResponse(w, err)
		return
	}

	okResponse(w)
}

func signout(w http.ResponseWriter, r *http.Request) {
	sess := r.Context().Value("session").(string)

	if err := model.DeleteSession(r.Context(), sess); err != nil {
		errorInternalResponse(w, err)
		return
	}

	cookie := http.Cookie{
		Name:     "session",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)
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
