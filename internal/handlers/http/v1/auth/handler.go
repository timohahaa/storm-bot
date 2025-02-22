package auth

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/timohahaa/gw/internal/errors"
	"github.com/timohahaa/gw/internal/modules/customer"
	"github.com/timohahaa/gw/internal/modules/user"
	"github.com/timohahaa/gw/internal/utils/render"
	"github.com/timohahaa/gw/pkg/validate"
)

type mod struct {
	customer *customer.Module
	user     *user.Module
}

type handler struct {
	mod                    mod
	customerAuthCookieName string
	userAuthCookieName     string
	cookieDomain           string
	cookieSecure           bool
}

//	@Summary	Create customer
//	@Tags		Auth Customer
//	@Param		CreateParams	body		customer.CreateParams	true	"Create Params"
//	@Success	200				{object}	customer.Model			"Customer Model"
//	@Failure	default			{object}	errors.HTTPError		"Error"
//	@Router		/api/v1/auth/customers/register [post]
func (h *handler) createCustomer(w http.ResponseWriter, r *http.Request) {
	var (
		form customer.CreateParams
	)

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		render.Error(w, r, err)
		return
	}

	if invParams := validate.Struct(&form); len(invParams) != 0 {
		render.Error(w, r, invParams)
		return
	}

	customer, err := h.mod.customer.Create(r.Context(), form)
	if err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	render.JSON(w, customer)
}

//	@Summary	Confirm customer
//	@Tags		Auth Customer
//	@Param		user_id	path		string				true	"User ID"
//	@Param		code	query		string				true	"Confirmation Code"
//	@Success	200		{object}	customer.Model		"Customer Model"
//	@Failure	default	{object}	errors.HTTPError	"Error"
//	@Router		/api/v1/auth/customers/{user_id}/confirm [post]
func (h *handler) confirmCustomer(w http.ResponseWriter, r *http.Request) {
	var (
		userId, err = uuid.Parse(chi.URLParam(r, "user_id"))
		code        = r.URL.Query().Get("code")
	)

	if err != nil {
		render.Error(w, r, err)
		return
	}

	customer, err := h.mod.customer.ConfirmCode(r.Context(), userId, code)
	if err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	render.JSON(w, customer)
}

//	@Summary	Send customer code
//	@Tags		Auth Customer
//	@Param		user_id	path		string				true	"User ID"
//	@Success	200		{object}	string				"Customer Model"
//	@Failure	default	{object}	errors.HTTPError	"Error"
//	@Router		/api/v1/auth/customers/{user_id}/send-code [post]
func (h *handler) sendCodeCustomer(w http.ResponseWriter, r *http.Request) {
	var (
		userId, err = uuid.Parse(chi.URLParam(r, "user_id"))
	)

	if err != nil {
		render.Error(w, r, err)
		return
	}

	if err := h.mod.customer.SendCode(r.Context(), userId); err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

//	@Summary	Login customer
//	@Tags		Auth Customer
//	@Param		LoginForm	body		customer.LoginForm	true	"Login Form"
//	@Success	200			{object}	customer.Model		"Customer Model"
//	@Failure	default		{object}	errors.HTTPError	"Error"
//	@Router		/api/v1/auth/customers/login [post]
func (h *handler) loginCustomer(w http.ResponseWriter, r *http.Request) {
	var (
		form customer.LoginForm
	)

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		render.Error(w, r, err)
		return
	}

	if invParams := validate.Struct(&form); len(invParams) != 0 {
		render.Error(w, r, invParams)
		return
	}

	loginResult, err := h.mod.customer.Login(r.Context(), form)
	if err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	// SET COOKIE loginResult.Token
	h.setTokenCookie(w, loginResult.Token, h.customerAuthCookieName, form.RememberMe)

	render.JSON(w, loginResult.User)
}

//	@Summary	Create user
//	@Tags		Auth User
//	@Param		CreateParams	body		user.CreateParams	true	"Create Params"
//	@Success	200				{object}	user.Model			"User Model"
//	@Failure	default			{object}	errors.HTTPError	"Error"
//	@Router		/api/v1/auth/users/register [post]
func (h *handler) createUser(w http.ResponseWriter, r *http.Request) {
	var (
		form user.CreateParams
	)

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		render.Error(w, r, err)
		return
	}

	if invParams := validate.Struct(&form); len(invParams) != 0 {
		render.Error(w, r, invParams)
		return
	}

	user, err := h.mod.user.Create(r.Context(), form)
	if err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	render.JSON(w, user)
}

//	@Summary	Confirm user
//	@Tags		Auth User
//	@Param		user_id	path		string				true	"User ID"
//	@Param		code	query		string				true	"Confirmation Code"
//	@Success	200		{object}	user.Model			"User Model"
//	@Failure	default	{object}	errors.HTTPError	"Error"
//	@Router		/api/v1/auth/users/{user_id}/confirm [post]
func (h *handler) confirmUser(w http.ResponseWriter, r *http.Request) {
	var (
		userId, err = uuid.Parse(chi.URLParam(r, "user_id"))
		code        = r.URL.Query().Get("code")
	)

	if err != nil {
		render.Error(w, r, err)
		return
	}

	user, err := h.mod.user.ConfirmCode(r.Context(), userId, code)
	if err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	render.JSON(w, user)
}

//	@Summary	Send user code
//	@Tags		Auth User
//	@Param		user_id	path		string				true	"User ID"
//	@Success	200		{object}	customer.Model		"Customer Model"
//	@Failure	default	{object}	errors.HTTPError	"Error"
//	@Router		/api/v1/auth/users/{user_id}/send-code [post]
func (h *handler) sendCodeUser(w http.ResponseWriter, r *http.Request) {
	var (
		userId, err = uuid.Parse(chi.URLParam(r, "user_id"))
	)

	if err != nil {
		render.Error(w, r, err)
		return
	}

	if err := h.mod.user.SendCode(r.Context(), userId); err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

//	@Summary	Login user
//	@Tags		Auth User
//	@Param		LoginForm	body		user.LoginForm		true	"Login Form"
//	@Success	200			{object}	user.Model			"User Model"
//	@Failure	default		{object}	errors.HTTPError	"Error"
//	@Router		/api/v1/auth/users/login [post]
func (h *handler) loginUser(w http.ResponseWriter, r *http.Request) {
	var (
		form user.LoginForm
	)

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		render.Error(w, r, err)
		return
	}

	if invParams := validate.Struct(&form); len(invParams) != 0 {
		render.Error(w, r, invParams)
		return
	}

	loginResult, err := h.mod.user.Login(r.Context(), form)
	if err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	// SET COOKIE loginResult.Token
	h.setTokenCookie(w, loginResult.Token, h.userAuthCookieName, form.RememberMe)

	render.JSON(w, loginResult.User)
}
