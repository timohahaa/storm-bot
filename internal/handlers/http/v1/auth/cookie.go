package auth

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const (
	ExpiredRememberMe = 7 * 24 * time.Hour
)

func (h *handler) setTokenCookie(w http.ResponseWriter, token, cookieName string, rememberMe bool) {
	cookie := http.Cookie{
		Name:     cookieName,
		Value:    token,
		Domain:   "." + h.cookieDomain,
		Secure:   h.cookieSecure,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	cookieRememberMe := http.Cookie{
		Name:     fmt.Sprintf("%s-remember-me", cookieName),
		Value:    strconv.FormatBool(rememberMe),
		Domain:   "." + h.cookieDomain,
		Secure:   h.cookieSecure,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	if rememberMe {
		cookie.Expires = time.Now().Add(ExpiredRememberMe)
		cookieRememberMe.Expires = time.Now().Add(ExpiredRememberMe)
	}

	http.SetCookie(w, &cookieRememberMe)
	http.SetCookie(w, &cookie)
}
