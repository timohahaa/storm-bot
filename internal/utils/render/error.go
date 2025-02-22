package render

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/timohahaa/gw/internal/errors"
	"github.com/timohahaa/gw/pkg/validate"
)

func Error(w http.ResponseWriter, r *http.Request, err error) error {
	switch {
	case err == io.EOF:
		err = errors.Get(errors.IO)
	case strings.Contains(err.Error(), "invalid UUID"):
		err = errors.Get(errors.InvalidUUID)
	}

	switch err := err.(type) {
	case validate.InvalidParamsErr:
		return writeError(w, r, errors.HTTPError{
			Code:          errors.Validation,
			Status:        http.StatusBadRequest,
			Message:       "invalid params found",
			InvalidParams: err,
		})
	case *errors.AppError:
		return writeError(w, r, err.HTTPError())
	case *json.SyntaxError:
		return writeError(w, r, errors.Get(errors.JsonSyntax).WithDetail(err.Error()).HTTPError())
	}

	log.Errorf("[UNRECOGNIZED ERROR]: %T, %+v (%s: %s)\n", err, err, r.Method, r.URL.String())
	return writeError(w, r, errors.Get(errors.Internal).HTTPError())
}

func writeError(w http.ResponseWriter, r *http.Request, httpErr errors.HTTPError) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpErr.Status)
	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	return enc.Encode(map[string]interface{}{
		"error": httpErr,
	})
}
