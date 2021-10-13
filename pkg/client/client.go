package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gojek/heimdall/v7"
	"github.com/gojek/heimdall/v7/httpclient"

	"github.com/nurislam03/template/pkg/logging"
)

func Client() *httpclient.Client {
	// First set a backoff mechanism. Constant backoff increases the backoff at a constant rate
	backoffInterval := 2 * time.Millisecond
	// Define a maximum jitter interval. It must be more than 1*time.Millisecond
	maximumJitterInterval := 5 * time.Millisecond

	backoff := heimdall.NewConstantBackoff(backoffInterval, maximumJitterInterval)

	// Create a new retry mechanism with the backoff
	retrier := heimdall.NewRetrier(backoff)

	timeout := 20 * time.Second
	// Create a new client, sets the retry mechanism, and the number of times you would like to retry
	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(timeout),
		httpclient.WithRetrier(retrier),
		httpclient.WithRetryCount(4),
	)
	return client

}

func SendRequest(method string, url string, data interface{}, headers http.Header) (map[string]interface{}, *http.Response, error) {
	jsonBytes, jsonErr := json.Marshal(data)
	if jsonErr != nil {
		logging.Info("json Marshal error ", jsonErr)
		return nil, nil, jsonErr
	}
	var resp *http.Response
	var reqError error
	if method == http.MethodGet {
		resp, reqError = Client().Get(url, headers)
	} else if method == http.MethodPost {
		resp, reqError = Client().Post(url, bytes.NewReader(jsonBytes), headers)
	} else if method == http.MethodPatch {
		resp, reqError = Client().Patch(url, bytes.NewReader(jsonBytes), headers)
	} else {
		logging.Info("Does not support the http method ", method)
		return nil, nil, errors.New("method does not implement")
	}

	if reqError != nil {
		cErr := errors.New("server")
		logging.Info("Server request failed ", reqError)
		return nil, nil, cErr
	}
	return ResponseBuilder(resp)
}

func ResponseBuilder(resp *http.Response) (map[string]interface{}, *http.Response, error) {
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logging.Info("Service post response body error ", err)
		return nil, nil, err
	}

	logging.Info("Response body ", string(body))
	// Convert JSON to Map data type
	var respData map[string]interface{}

	errs := json.Unmarshal(body, &respData)
	if errs != nil {
		logging.Info("service json unmarshal error ", errs)
		return nil, nil, errs
	}
	if resp.StatusCode >= 500 && resp.StatusCode <= 503 {
		cErr := errors.New("server")
		logging.Info("500 error ", respData)
		return nil, nil, cErr
	}

	if resp.StatusCode >= 400 && resp.StatusCode <= 499 {
		cErr := errors.New(string(body))
		logging.Info("4xx error ", respData)
		return respData, resp, cErr
	}
	return respData, resp, nil
}
