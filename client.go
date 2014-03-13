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

func (self *PrevotyClient) FilterContent(input string, configurationKey string) (*TrustedContentResult, error) {
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

func (self *PrevotyClient) GenerateTimedToken(action string, userIdentifier string, ttl string) (*TrustedTokenGenerationResult, error) {
	generateUrl := fmt.Sprintf("%s/token/timed/generate?api_key=%s&action=%s&user_identifier=%s&ttl=%s", self.Base, self.ApiKey, action, userIdentifier, ttl)
	response, err := http.Get(generateUrl)
	result := new(TrustedTokenGenerationResult)
	if err == nil {
		defer response.Body.Close()
		switch response.StatusCode {
		case 200:
			body, ioErr := ioutil.ReadAll(response.Body)
			if ioErr != nil {
				return result, nil
			}
			decodingErr := json.Unmarshal(body, &result)
			if decodingErr != nil {
				return result, nil
			}
		case 400:
			return result, &BadInputParameter{}
		case 403:
			return result, &BadAPIKey{}
		case 500:
			return result, &InternalError{}
		}
	}
	return result, err
}

func (self *PrevotyClient) ValidateTimedToken(action string, userIdentifier string, token string) (*TrustedTokenValidationResult, error) {
	validateUrl := fmt.Sprintf("%s/token/timed/validate?api_key=%s&action=%s&user_identifier=%s&token=%s", self.Base, self.ApiKey, action, userIdentifier, token)
	response, err := http.Get(validateUrl)
	result := new(TrustedTokenValidationResult)
	if err == nil {
		defer response.Body.Close()
		switch response.StatusCode {
		case 200:
			body, ioErr := ioutil.ReadAll(response.Body)
			if ioErr != nil {
				return result, nil
			}
			decodingErr := json.Unmarshal(body, &result)
			if decodingErr != nil {
				return result, nil
			}
		case 400:
			return result, &BadInputParameter{}
		case 403:
			return result, &BadAPIKey{}
		case 500:
			return result, &InternalError{}
		}
	}
	return result, err
}

func (self *PrevotyClient) AnalyzeQuery(query string, configKey string) (*TrustedQueryResult, error) {
	parseUrl := fmt.Sprintf("%s/query/parse", self.Base)
	response, err := http.PostForm(parseUrl, url.Values{"api_key": {self.ApiKey}, "query": {query}, "config_key": {configKey}})
	result := new(TrustedQueryResult)
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
		case 500:
			return result, &InternalError{}
		}
	}
	return result, err
}
