package client

import (
	"bytes"
	"com/github/graylog-code/config/template"
	errorHandler "com/github/graylog-code/error"
	"com/github/graylog-code/graylog/api"
	"com/github/graylog-code/util"
	"encoding/json"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"time"
)
type loggingTransport struct {}

func (lrt loggingTransport) RoundTrip(request *http.Request) (*http.Response,error) {
	var roundTripId = uuid.NewString()
	var logger = log.WithField("roundTripId",roundTripId)

	var requestBody []byte
	var reqErr error
	if request.Body != nil {
		requestBody, reqErr = io.ReadAll(request.Body)
		if reqErr != nil {
			return nil, reqErr
		}
	}
	request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
	logger.Trace("==================================")
	logger.Tracef("Sending request to: %s%s", request.Host,request.URL.Path)
	logger.Trace("Request header(s): ")
	for key, element := range request.Header {
		logger.Tracef("%s: %s", key, element)
	}
	logger.Trace("Request body: ", string(requestBody))
	logger.Trace("==================================")
	resp, err  := http.DefaultTransport.RoundTrip(request)
	if err != nil {
		logger.Error(resp)
		logger.Errorf(err.Error())
		return nil,err
	}

	var respErr error
	var responseBody []byte
	if resp.Body != nil {
		responseBody, respErr = io.ReadAll(resp.Body)
		if respErr != nil {
			return nil, respErr
		}
	}
	logger.Trace("==================================")
	logger.Trace("Response header(s): ")
	for key, element := range resp.Header {
		logger.Tracef("%s: %s", key, element)
	}
	logger.Trace("Response body:\n ", string(responseBody))
	logger.Trace("==================================")
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(responseBody))
	return resp, err
}

var client = http.Client{
	Transport: loggingTransport{},
	Timeout:   30 * time.Second,
}

func DeleteFromGraylog(config template.GraylogConfigTemplate,endpoint string) int {
	log.Debugf("Deleting resource from graylog: %s", endpoint)
	req, reqErr := http.NewRequest("DELETE",
		getBaseAddress(config)+endpoint, nil)
	errorHandler.Check(reqErr)
	addHeaders(config, req)
	resp, respErr := client.Do(req)
	errorHandler.Check(respErr)

	return resp.StatusCode
}

func PostToGraylog(config template.GraylogConfigTemplate,
	data interface{}, endpoint string,
    response interface{}) int {
	var reader io.Reader
	var dataType = reflect.TypeOf(data).Name()
	//This is needed as /api/system/whitelist/check needs json data with quotes intact.
	//So we can't use normal Unmarshalling, but the json.RawMessage type
	if dataType == "RawMessage" {
		rawJsonMessage := data.(json.RawMessage)
		rawJsonMessageAsString := string(rawJsonMessage)
		log.Debugf("Sending data to Graylog: %s", rawJsonMessageAsString)
		reader = bytes.NewBuffer(rawJsonMessage)
	} else {
		toPost, marshalError := json.Marshal(data)
		errorHandler.Check(marshalError)
		toPostAsString := string(toPost)
		log.Debugf("Sending data to Graylog: %s", toPostAsString)
		reader = strings.NewReader(toPostAsString)
	}
	req, reqErr := http.NewRequest("POST", getBaseAddress(config)+endpoint, reader)
	errorHandler.Check(reqErr)
	addHeaders(config, req)
	resp, respErr := client.Do(req)
	errorHandler.Check(respErr)
	defer util.CloseResource(resp.Body)
	body, readBodyError := io.ReadAll(resp.Body)
	errorHandler.Check(readBodyError)
	log.Debugf("Response received: %s", string(body))
	if response != nil {
		unmarshalError := json.Unmarshal(body, &response)
		errorHandler.Check(unmarshalError)
	}

	return resp.StatusCode
}

func getBaseAddress(config template.GraylogConfigTemplate) string {
	return config.GraylogServer + config.ApiPrefix
}

func GetFromGraylog(config template.GraylogConfigTemplate,endpoint string,data interface{}) {
	req, reqErr := http.NewRequest("GET",
		getBaseAddress(config)+endpoint,nil)
	errorHandler.Check(reqErr)
	addHeaders(config,req)
	resp, respErr := client.Do(req)
	errorHandler.Check(respErr)
	defer util.CloseResource(resp.Body)
	body, readBodyError := io.ReadAll(resp.Body)
	errorHandler.Check(readBodyError)
	unmarshalErr := json.Unmarshal(body,&data)
	errorHandler.Check(unmarshalErr)
}

func PutToGraylog(config template.GraylogConfigTemplate,
	data interface{},
    endpoint string) (int, *api.ResourceCreationResponse) {
	var reader io.Reader
	if data != nil {
		toPost, marshalError := json.Marshal(data)
		errorHandler.Check(marshalError)
		log.Debugf("Sending data to Graylog: %s", string(toPost))
		reader = strings.NewReader(string(toPost))
	}
	req, reqErr := http.NewRequest("PUT",
		getBaseAddress(config)+endpoint, reader)
	errorHandler.Check(reqErr)
	addHeaders(config,req)
	resp, respErr := client.Do(req)
	errorHandler.Check(respErr)
	defer util.CloseResource(resp.Body)
	body, readBodyError := io.ReadAll(resp.Body)
	errorHandler.Check(readBodyError)
	if IsNoContent(resp.StatusCode) {
		return resp.StatusCode, nil
	}
	log.Debugf("Response received: %s",string(body))
	var resourceCreationResponse api.ResourceCreationResponse
	unmarshalError := json.Unmarshal(body, &resourceCreationResponse)
	errorHandler.Check(unmarshalError)

	return resp.StatusCode,&resourceCreationResponse
}

func addHeaders(config template.GraylogConfigTemplate, req *http.Request) {
	req.SetBasicAuth(config.AccessToken, "token")
	req.Header.Add("X-Requested-By", "cli")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
}

func ResponseIsSuccess(responseCode int) bool {
	return inBetween(responseCode,200,208)
}

func ResponseIsFailure(responseCode int) bool {
	return !ResponseIsSuccess(responseCode)
}

func inBetween(i, min, max int) bool {
	if (i >= min) && (i <= max) {
		return true
	} else {
		return false
	}
}

func IsNoContent(respCode int) bool {
	return respCode == 204
}