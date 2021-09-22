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

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// getAuthenticatedUrl returns ready to use URL to call Marvel API
func getAuthenticatedUrl(url string) string {
	now := time.Now() // current local time
	ts := now.Unix()

	baseUrl := config.Config("MARVEL_BASE_URL")
	publicKey := config.Config("MARVEL_PUBLIC_KEY")
	privateKey := config.Config("MARVEL_PRIVATE_KEY")

	hash := getMD5Hash(fmt.Sprintf("%d%s%s", ts, privateKey, publicKey))

	authUrl := fmt.Sprintf("%s/%s?ts=%d&apikey=%s&hash=%s", baseUrl, url, ts, publicKey, hash)
	fmt.Println(authUrl)

	return authUrl
}

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
