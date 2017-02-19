package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ApiResponse struct {
	Data DataResponse `json:"data"`
}

type DataResponse struct {
	Images ImagesList `json:"images"`
}

type ImagesList struct {
	Original      Image `json:"original"`
	OriginalStill Image `json:"original_still"`
}

type Image struct {
	Url string `json:"url"`
}

func imageContent(phrase string, image_type string) []byte {
	apiResponse := getGiphyResponse(phrase)

	var image_body *http.Response
	var err error

	if image_type == "gif" {
		image_body, err = http.Get(apiResponse.Data.Images.Original.Url)
	} else {
		image_body, err = http.Get(apiResponse.Data.Images.OriginalStill.Url)
	}

	bytes, err := ioutil.ReadAll(image_body.Body)
	checkError(err)

	return bytes
}

func getGiphyResponse(word string) *ApiResponse {
	wordUrl := "http://api.giphy.com/v1/gifs/translate?s=" + url.QueryEscape(word) + "&api_key=dc6zaTOxFJmzC"
	fmt.Println(wordUrl)

	resp, err := http.Get(wordUrl)
	checkError(err)

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	var api_response = new(ApiResponse)
	json.Unmarshal(body, &api_response)
	return api_response
}
