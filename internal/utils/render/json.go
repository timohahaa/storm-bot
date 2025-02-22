package render

import (
	"encoding/json"
	"net/http"
)

type container struct {
	Meta interface{} `json:"meta,omitempty"`
	Data interface{} `json:"data"`
}

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

func JSON(w http.ResponseWriter, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	return enc.Encode(&container{
		Data: v,
	})
}

func JSONWithMeta(w http.ResponseWriter, data, meta interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	return enc.Encode(&container{
		Data: data,
		Meta: meta,
	})
}
