package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/example/helpers"
)

func request(url string, method string, body string) (string, error) {
	provider := common.DefaultConfigProvider()

	// Prepare HTTP request
	var reqBody *bytes.Buffer
	if body != "" {
		reqBody = bytes.NewBuffer([]byte(body))
	} else {
		reqBody = bytes.NewBuffer(nil)
	}

	req, err := http.NewRequest(method, url, reqBody)
	helpers.FatalIfError(err)

	// Set required headers
	req.Header.Set("date", time.Now().UTC().Format(http.TimeFormat))
	req.Header.Set("content-type", "application/json")

	// Sign the request
	signer := common.DefaultRequestSigner(provider)
	signer.Sign(req)

	client := http.Client{}
	resp, err := client.Do(req)
	helpers.FatalIfError(err)
	defer resp.Body.Close()

	// Read the response body
	respBody, err := ioutil.ReadAll(resp.Body)
	helpers.FatalIfError(err)

	return string(respBody), nil
}
