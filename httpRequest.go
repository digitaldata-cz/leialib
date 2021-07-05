package leialib

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HttpRequest
// url - target URL
// method - GET,POST,...
// vReq - request body or nil
// headers - HTTP headers
// rootCApool - CA certificate pool or nil
func HttpsRequest(url string, method string, vReq interface{}, headers *http.Header, rootCApool *x509.CertPool, vResp interface{}) (err error) {
	reqBody, err := json.Marshal(vReq)
	if err != nil {
		return err
	}
	client := http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:            rootCApool,
				InsecureSkipVerify: rootCApool == nil,
			},
		},
	}
	req, err := http.NewRequest(method, url, bytes.NewReader(reqBody))
	if err != nil {
		return err
	}
	if headers != nil {
		req.Header = *headers
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	resp.Body.Close()
	err = json.Unmarshal(respBody, vResp)
	if err != nil {
		return err
	}
	return nil
}
