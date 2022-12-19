package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type aocFacade struct {
	rootUrl       string
	sessionCookie string
	client        *http.Client
}

func NewAocHttpClient() (aocFacade, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
		return aocFacade{}, err
	}

	return aocFacade{
		rootUrl:       "https://adventofcode.com",
		sessionCookie: os.Getenv("SESSION_COOKIE"),
		client:        &http.Client{},
	}, nil
}

func (facade *aocFacade) doRequestWithCookie(req *http.Request) (string, error) {
	// Declare HTTP Method and Url

	// Set cookie
	req.Header.Set("Cookie", fmt.Sprintf("session=%s;", facade.sessionCookie))
	resp, err := facade.client.Do(req)
	if err != nil {
		log.Fatal("An error occurred, please try again")
		return "", err
	}
	defer resp.Body.Close()
	// Read response
	data, err := ioutil.ReadAll(resp.Body)

	// error handle
	if err != nil {
		fmt.Printf("error = %s \n", err)
	}

	return string(data[:]), nil
}

func (facade *aocFacade) GetDayInput(year int, day int) (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%d/day/%d/input", facade.rootUrl, year, day), nil)
	if err != nil {
		return "", err
	}
	return facade.doRequestWithCookie(req)
}

func (facade *aocFacade) SubmitDay(year int, day int, part string, answer string) (string, error) {
	form := url.Values{}
	form.Add("level", part)
	form.Add("answer", answer)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%d/day/%d/answer", facade.rootUrl, year, day), strings.NewReader(form.Encode()))

	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return facade.doRequestWithCookie(req)
}
