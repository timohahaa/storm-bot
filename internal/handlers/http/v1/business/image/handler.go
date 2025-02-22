package image

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/timohahaa/gw/internal/errors"
	"github.com/timohahaa/gw/internal/modules/image"
	"github.com/timohahaa/gw/internal/utils/render"
)

const (
	MB = 1 << 20
)

type handler struct {
	mod *image.Module
}

//	@Summary	Upload new image for business
//	@Tags		Business Images
//	@Param		business_id	path		string				true	"Business ID"
//	@Param		file		formData	file				true	"Image"
//	@Param		type		query		string				true	"Image Type (stamp, card, reward)"
//	@Success	200			{object}	image.JSONModel		"Response"
//	@Failure	default		{object}	errors.HTTPError	"Error"
//	@Router		/api/v1/businesses/{business_id}/images/upload [post]
func (h *handler) upload(w http.ResponseWriter, r *http.Request) {
	var (
		businessId, err = uuid.Parse(chi.URLParam(r, "business_id"))
		imageType       = r.URL.Query().Get("type")
	)

	if err != nil {
		render.Error(w, r, err)
		return
	}

	if err := r.ParseMultipartForm(10 * MB); err != nil {
		render.Error(w, r, err)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		render.Error(w, r, err)
		return
	}

	fileSize, fileType, contentType, err := getFileMeta(header)
	if err != nil {
		render.Error(w, r, err)
		return
	}

	model, err := h.mod.Create(r.Context(), image.CreateParams{
		BusinessID:  businessId,
		FileSize:    fileSize,
		FileType:    fileType,
		ContentType: contentType,
		Type:        image.ImageType(imageType),
	}, file)

	if err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	render.JSON(w, model.ToJSON())
}

//	@Summary	Download specific business image
//	@Tags		Business Images
//	@Param		business_id	path		string				true	"Business ID"
//	@Param		image_id	path		string				true	"Image ID"
//	@Success	200			{object}	nil					"Response"
//	@Failure	default		{object}	errors.HTTPError	"Error"
//	@Router		/api/v1/businesses/{business_id}/images/{image_id}/ [get]
func (h *handler) download(w http.ResponseWriter, r *http.Request) {
	var (
		imageId, err = uuid.Parse(chi.URLParam(r, "image_id"))
		businessId   uuid.UUID
	)

	if err != nil {
		render.Error(w, r, err)
		return
	}

	businessId, err = uuid.Parse(chi.URLParam(r, "business_id"))
	if err != nil {
		render.Error(w, r, err)
		return
	}

	url, err := h.mod.GetDownloadLink(r.Context(), imageId, businessId)
	if err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}

//	@Summary	Delete specific business image
//	@Tags		Business Images
//	@Param		business_id	path		string				true	"Business ID"
//	@Param		image_id	path		string				true	"Image ID"
//	@Success	204			{object}	nil					"Response"
//	@Failure	default		{object}	errors.HTTPError	"Error"
//	@Router		/api/v1/businesses/{business_id}/images/{image_id}/ [delete]
func (h *handler) delete(w http.ResponseWriter, r *http.Request) {
	var (
		imageId, err = uuid.Parse(chi.URLParam(r, "image_id"))
		businessId   uuid.UUID
	)

	if err != nil {
		render.Error(w, r, err)
		return
	}

	businessId, err = uuid.Parse(chi.URLParam(r, "business_id"))
	if err != nil {
		render.Error(w, r, err)
		return
	}

	if err := h.mod.Delete(r.Context(), imageId, businessId); err != nil {
		render.Error(w, r, errors.DB(err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

//	@Summary	List business images
//	@Tags		Business Images
//	@Param		business_id	path		string				true	"Business ID"
//	@Param		type		query		string				true	"Image Type (stamp, card, reward)"
//	@Success	200			{object}	[]image.JSONModel	"Response"
//	@Failure	default		{object}	errors.HTTPError	"Error"
//	@Router		/api/v1/businesses/{business_id}/images/ [get]
func (h *handler) list(w http.ResponseWriter, r *http.Request) {
	var (
		businessId, err = uuid.Parse(chi.URLParam(r, "business_id"))
		imageType       = r.URL.Query().Get("type")
	)

	if err != nil {
		render.Error(w, r, err)
		return
	}

	images, err := h.mod.List(r.Context(), businessId, image.ImageType(imageType))
	if err != nil {
		render.Error(w, r, err)
		return
	}

	var imagesJson = make([]image.JSONModel, 0, len(images))
	for _, i := range images {
		imagesJson = append(imagesJson, i.ToJSON())
	}

	render.JSON(w, imagesJson)
}
