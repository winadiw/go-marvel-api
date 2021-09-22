package external

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/winadiw/go-marvel-api/config"
	"github.com/winadiw/go-marvel-api/utils"
)

type MarvelErrorResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}
type BaseMarvelResponse struct {
	Code            int    `json:"code"`
	Status          string `json:"status"`
	Copyright       string `json:"copyright"`
	AttributionText string `json:"attributionText"`
	AttributionHTML string `json:"attributionHTML"`
	Etag            string `json:"etag"`
}

type BaseMarvelResponseData struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Total  int `json:"total"`
	Count  int `json:"count"`
}

type MarvelGetCharacterByIdResponse struct {
	BaseMarvelResponse
	Data struct {
		BaseMarvelResponseData
		Results []struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Modified    string `json:"modified"`
		} `json:"results"`
	} `json:"data"`
}

var baseUrl = config.Config("MARVEL_BASE_URL")
var publicKey = config.Config("MARVEL_PUBLIC_KEY")
var privateKey = config.Config("MARVEL_PRIVATE_KEY")

// getMD5Hash returns md5 of a given string
func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// getAuthenticatedUrl returns ready to use URL to call Marvel API
func getAuthenticatedUrl(url string) string {
	ts := time.Now().Unix()
	hash := getMD5Hash(fmt.Sprintf("%d%s%s", ts, privateKey, publicKey))
	authUrl := fmt.Sprintf("%s/%s?ts=%d&apikey=%s&hash=%s", baseUrl, url, ts, publicKey, hash)
	return authUrl
}

// makeGetRequest handles all repetitive code for get request
func makeGetRequest(url string) ([]byte, map[string]interface{}) {
	response, err := http.Get(getAuthenticatedUrl(url))

	if err != nil {
		fmt.Print(err.Error())
		return nil, utils.ResponseError(http.StatusFailedDependency, "Network Error", err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return nil, utils.ResponseError(http.StatusUnprocessableEntity, "Unable to read response data", err)
	}

	if response.StatusCode != http.StatusOK {
		var errorResponse MarvelErrorResponse
		err = json.Unmarshal(responseData, &errorResponse)
		if err != nil {
			return nil, utils.ResponseError(http.StatusUnprocessableEntity, "Unable to parse error response from marvel", err)
		}

		return nil, utils.ResponseError(errorResponse.Code, errorResponse.Status, nil)
	}

	return responseData, nil
}

// MarvelGetCharacterById returns character by ID
func MarvelGetCharacterById(ID string) (MarvelGetCharacterByIdResponse, map[string]interface{}) {

	response, errRequest := makeGetRequest("/v1/public/characters/" + ID)

	if errRequest != nil {
		return MarvelGetCharacterByIdResponse{}, errRequest
	}

	var responseObject MarvelGetCharacterByIdResponse
	errMarshall := json.Unmarshal(response, &responseObject)

	if errMarshall != nil {
		log.Fatal(errMarshall)
		return MarvelGetCharacterByIdResponse{}, utils.ResponseError(http.StatusUnprocessableEntity, "Unable to parse response", errMarshall)
	}

	return responseObject, nil

}
