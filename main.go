package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	url2 "net/url"
	"os"
	"strings"
)

var baseUrl = "https://mcare.siriusxm.ca"
var siriusConfigJson = &SiriusConfig{}

var authToken string

var cookieJar, _ = cookiejar.New(nil)
var client = &http.Client{
	Jar: cookieJar,
}

func getAppConfig() (SiriusConfig, error) {
	req, err := http.NewRequest("GET", baseUrl+"/authService/100000002/appconfig", nil)
	returnConfig := SiriusConfig{}
	if err != nil {
		return returnConfig, errors.New("couldn't fetch appconfig")
	}
	AddSecretHeaders(req)

	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, _ := io.ReadAll(response.Body)
	if err := json.Unmarshal(body, &returnConfig); err != nil {
		return returnConfig, errors.New("couldn't unmarshal json")
	}

	return returnConfig, nil
}

func AddSecretHeaders(req *http.Request) {
	req.Header.Add("X-Kony-App-Key", os.Getenv("X-Kony-App-Key"))
	req.Header.Add("X-Kony-App-Secret", os.Getenv("X-Kony-App-Secret"))
	req.Header.Add("X-Kony-Platform-Type", "android")
}

func getAuthToken() string {
	req, err := http.NewRequest("POST", baseUrl+"/authService/100000002/login", nil)
	if err != nil {
		panic("Couldn't perform login request")
	}

	AddSecretHeaders(req)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		log.Fatalln("Couldn't log in")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	loginTokens := &LoginTokens{}
	_ = json.Unmarshal(body, &loginTokens)
	return loginTokens.ClaimsToken.Value
}

func renewSubscription(radioId string) error {
	refreshUrl := siriusConfigJson.ServicesMeta.USESBRefresh.URL + "/RefreshDevice"

	data := url2.Values{}
	data.Add("events", "[]")
	data.Add("AuthName", "")
	data.Add("AuthPwd", "")
	data.Add("deviceId", radioId)
	data.Add("vin", "")
	data.Add("phone", "")
	data.Add("note", "1")
	data.Add("provisionType", "activate")
	data.Add("provisionPackageName", "")
	data.Add("provisionDate", "")
	data.Add("provisionPriority", "2")
	data.Add("appVersion", "2.1.0")
	data.Add("dmCode", "")
	data.Add("flow_name", "Enter Radio ID")
	data.Add("lat", "41.8265505")
	data.Add("lng", "-88.0056906")
	data.Add("vehicle_active_flag", "")

	// make the request
	req, err := http.NewRequest("POST", refreshUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return errors.New("couldn't create request")
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	AddSecretHeaders(req)

	resp, err := client.Do(req)
	if err != nil {
		return errors.New("couldn't perform request")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("couldn't read response body")
	}
	renewResp := &RenewResponse{}

	if err := json.Unmarshal(body, &renewResp); err != nil {
		return errors.New("couldn't unmarshal json")
	}

	if renewResp.Errors[0].ResultCode != "SUCCESS" {
		log.Println(renewResp)
		return errors.New("couldn't renew subscription")
	}

	return nil
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading .env file")
	}

}
func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide a radio id as the first argument. Tune to channel 0 to get the radio id.")
		os.Exit(1)
	}

	// get the device id from the command line
	radioId := os.Args[1]

	var err error
	*siriusConfigJson, err = getAppConfig()

	if err != nil {
		log.Fatalln("Failed to get the config")
	}

	authToken = getAuthToken()

	err = renewSubscription(radioId)
	if err != nil {
		fmt.Println("❌ Couldn't renew, check radio id")
	} else {
		fmt.Println("✅ Successfully Renewed! Turn off your radio/car, wait a few minutes, and turn it back on.")
	}

}
