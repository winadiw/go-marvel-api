package external

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/winadiw/go-marvel-api/config"
)

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

// MarvelGetCharactersById returns character by ID
func MarvelGetCharactersById(ID string) error {
	response, err := http.Get(getAuthenticatedUrl("/v1/public/characters/" + ID))

	if err != nil {
		fmt.Print(err.Error())
		return err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

	return nil

}
