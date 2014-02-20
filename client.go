package prevoty

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type PrevotyClient struct {
	ApiKey string
	Base   string
}

func NewPrevotyClient(apiKey string) PrevotyClient {
	return PrevotyClient{
		ApiKey: apiKey,
		Base:   "http://api.prevoty.com/1",
	}
}

func (self *PrevotyClient) Verify() (bool, error) {
	verifyUrl := fmt.Sprintf("%s/key/verify?api_key=%s", self.Base, self.ApiKey)
	response, err := http.Get(verifyUrl)
	if err == nil {
		defer response.Body.Close()
		switch response.StatusCode {
		case 200:
			return true, nil
		case 400:
			return false, &BadInputParameter{}
		case 403:
			return false, &BadAPIKey{}
		case 500:
			return false, &InternalError{}
		}
	}
	return false, err
}

func (self *PrevotyClient) Info() (*KeyInformation, error) {
	infoUrl := fmt.Sprintf("%s/key/info?api_key=%s", self.Base, self.ApiKey)
	response, err := http.Get(infoUrl)
	information := new(KeyInformation)
	if err == nil {
		defer response.Body.Close()
		switch response.StatusCode {
		case 200:
			body, ioErr := ioutil.ReadAll(response.Body)
			if ioErr != nil {
				return information, ioErr
			}
			decodingErr := json.Unmarshal(body, &information)
			if decodingErr != nil {
				return information, decodingErr
			}
			return information, nil
		case 400:
			return information, &BadInputParameter{}
		case 403:
			return information, &BadAPIKey{}
		case 500:
			return information, &InternalError{}
		}
	}
	return information, err
}

func (self *PrevotyClient) VerifyConfigurationKey(configurationKey string) (bool, error) {
	verifyUrl := fmt.Sprintf("%s/rule/verify?api_key=%s&rule_key=%s", self.Base, self.ApiKey, configurationKey)
	response, err := http.Get(verifyUrl)
	if err == nil {
		defer response.Body.Close()
		switch response.StatusCode {
		case 200:
			return true, nil
		case 400:
			return false, &BadInputParameter{}
		case 403:
			return false, &BadAPIKey{}
		case 500:
			return false, &InternalError{}
		}
	}
	return false, err
}

func (self *PrevotyClient) Filter(input string, configurationKey string) (*TrustedContentResult, error) {
	filterUrl := fmt.Sprintf("%s/xss/filter", self.Base)
	response, err := http.PostForm(filterUrl, url.Values{"api_key": {self.ApiKey}, "rule_key": {configurationKey}, "input": {input}})
	result := new(TrustedContentResult)
	if err == nil {
		defer response.Body.Close()
		switch response.StatusCode {
		case 200:
			body, ioErr := ioutil.ReadAll(response.Body)
			if ioErr != nil {
				return result, ioErr
			}
			decodingErr := json.Unmarshal(body, &result)
			if decodingErr != nil {
				return result, decodingErr
			}
			return result, nil
		case 400:
			return result, &BadInputParameter{}
		case 403:
			return result, &BadAPIKey{}
		case 413:
			return result, &RequestTooLarge{}
		case 500:
			return result, &InternalError{}
		case 507:
			return result, &AccountQuotaExceeded{}
		}
	}
	return result, err
}
