package handler

import (
	"StanCodingChallenge/app/model"
	"encoding/json"
	"net/http"
)

func PostItems(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	request := model.Request{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		respondError(w, http.StatusBadRequest, "Could not decode request: JSON parsing failed")
		return
	}

	if len(request.Items) <= 0 {
		respondError(w, http.StatusBadRequest, "Could not decode request: JSON parsing failed")
		return
	}

	respondJSON(w, http.StatusOK, getResponse(request))
}

func getResponse(values model.Request) []model.ResponseItem {
	var response []model.ResponseItem
	for _, value := range values.Items {
		if value.DRM && value.EpisodeCount > 0 {
			resp := model.ResponseItem{
				Image: value.Image.ShowImage,
				Title: value.Title,
				Slug:  value.Slug,
			}
			response = append(response, resp)
		}
	}
	return response
}
