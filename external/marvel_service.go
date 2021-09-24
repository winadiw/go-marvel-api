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

var (
	MarvelService marvelServiceInterface = &marvelService{}
)

type marvelService struct{}

type marvelServiceInterface interface {
	MarvelGetCharacterById(ID string) (MarvelGetCharactersResponse, *utils.ResponseErrorData)
	MarvelGetCharacters(limit, offset int) (MarvelGetCharactersResponse, *utils.ResponseErrorData)
}

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

type MarvelCharacterData struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Modified    string `json:"modified"`
}

type MarvelGetCharactersResponse struct {
	BaseMarvelResponse
	Data struct {
		BaseMarvelResponseData
		Results []MarvelCharacterData `json:"results"`
	} `json:"data"`
}

// IdList flattens id from character list
func (m MarvelGetCharactersResponse) IdList() []int {
	var list []int
	for _, user := range m.Data.Results {
		list = append(list, user.ID)
	}
	return list
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
func makeGetRequest(url string) ([]byte, *utils.ResponseErrorData) {
	fmt.Println("call url: " + url)
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err)
		return nil, utils.ResponseError(http.StatusFailedDependency, "Network Error", nil)
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
func (m *marvelService) MarvelGetCharacterById(ID string) (MarvelGetCharactersResponse, *utils.ResponseErrorData) {

	response, errRequest := makeGetRequest(getAuthenticatedUrl("v1/public/characters/" + ID))

	if errRequest != nil {
		return MarvelGetCharactersResponse{}, errRequest
	}

	var responseObject MarvelGetCharactersResponse
	errMarshall := json.Unmarshal(response, &responseObject)

	if errMarshall != nil {
		log.Fatal(errMarshall)
		return MarvelGetCharactersResponse{}, utils.ResponseError(http.StatusUnprocessableEntity, "Unable to parse response", errMarshall)
	}

	return responseObject, nil

}

// MarvelGetCharacters returns characters list
func (m *marvelService) MarvelGetCharacters(limit, offset int) (MarvelGetCharactersResponse, *utils.ResponseErrorData) {

	response, errRequest := makeGetRequest(getAuthenticatedUrl("v1/public/characters") + fmt.Sprintf("&limit=%d&offset=%d", limit, offset))

	if errRequest != nil {
		return MarvelGetCharactersResponse{}, errRequest
	}

	var responseObject MarvelGetCharactersResponse
	errMarshall := json.Unmarshal(response, &responseObject)

	if errMarshall != nil {
		log.Fatal(errMarshall)
		return MarvelGetCharactersResponse{}, utils.ResponseError(http.StatusUnprocessableEntity, "Unable to parse response", errMarshall)
	}

	return responseObject, nil

}
