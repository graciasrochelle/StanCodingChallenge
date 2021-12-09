package handler_test

import (
	handler "StanCodingChallenge/app/handlers"
	"StanCodingChallenge/app/model"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostItem(t *testing.T) {
	testCases := map[string]struct {
		request    model.Request
		response   string
		statusCode int
	}{
		"success": {
			request: model.Request{
				Payload: []model.RequestItem{
					{
						Title:        "16 Kids and Counting",
						DRM:          true,
						Slug:         "show/16kidsandcounting",
						EpisodeCount: 3,
						Image: model.Image{
							ShowImage: "http://catchup.ninemsn.com.au/img/jump-in/shows/16KidsandCounting1280.jpg",
						},
					},
					{
						Title:        "The Taste (Le Goût)",
						DRM:          true,
						Slug:         "show/thetaste",
						EpisodeCount: 0,
						Image: model.Image{
							ShowImage: "http://catchup.ninemsn.com.au/img/jump-in/shows/16KidsandCounting1280.jpg",
						},
					},
					{
						Title:        "Thunderbirds",
						DRM:          true,
						EpisodeCount: 1,
						Slug:         "show/seapatrol",
						Image: model.Image{
							ShowImage: "http://catchup.ninemsn.com.au/img/jump-in/shows/16KidsandCounting1280.jpg",
						},
					},
					{
						Title:        "Scooby-Doo! Mystery Incorporated",
						DRM:          false,
						EpisodeCount: 1,
						Slug:         "show/scoobydoomysteryincorporated",
						Image: model.Image{
							ShowImage: "http://catchup.ninemsn.com.au/img/jump-in/shows/16KidsandCounting1280.jpg",
						},
					},
				},
			},
			response:   "{\"response\":[{\"title\":\"16 Kids and Counting\",\"slug\":\"show/16kidsandcounting\",\"image\":\"http://catchup.ninemsn.com.au/img/jump-in/shows/16KidsandCounting1280.jpg\"},{\"title\":\"Thunderbirds\",\"slug\":\"show/seapatrol\",\"image\":\"http://catchup.ninemsn.com.au/img/jump-in/shows/16KidsandCounting1280.jpg\"}]}",
			statusCode: 200,
		},
		"bad request": {
			request:    model.Request{},
			response:   "{\"error\":\"Could not decode request: JSON parsing failed\"}",
			statusCode: 400,
		},
	}
	for tc, tp := range testCases {
		t.Logf("Test Case: <%s>", tc)
		w := httptest.NewRecorder()
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(tp.request)
		req, _ := http.NewRequest("POST", "/item", b)
		handler.PostItems(w, req)
		assert.Equal(t, tp.statusCode, w.Code)
		assert.Equal(t, tp.response, w.Body.String())
	}
}
