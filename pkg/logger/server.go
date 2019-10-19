/*
Copyright 2019 kubeflow.org.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package logger

import (
	"bytes"
	"github.com/go-logr/logr"
	"io/ioutil"
	"net/http"
	"net/url"
)

type loggerHandler struct {
	log       logr.Logger
	svcPort   string
	logUrl    *url.URL
	transport http.RoundTripper
}

func New(log logr.Logger, svcPort string, logUrl *url.URL) http.Handler {
	return &loggerHandler{
		log:     log,
		svcPort: svcPort,
		logUrl:  logUrl,
	}
}

func (eh *loggerHandler) post(url *url.URL, body []byte, contentType string) ([]byte, error) {
	eh.log.Info("Calling server", "url", url.String(), "contentType", contentType)
	response, err := http.Post(url.String(), contentType, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if err = response.Body.Close(); err != nil {
		return nil, err
	}
	return b, nil
}

func (eh *loggerHandler) callService(r *http.Request) ([]byte, error) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	b, err = eh.post(&url.URL{
		Scheme: "http",
		Host:   "0.0.0.0:" + eh.svcPort,
		Path:   r.URL.Path,
	}, b, r.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (eh *loggerHandler) logPayload(b []byte, contentType string) error {
	b, err := eh.post(eh.logUrl, b, contentType)
	if err != nil {
		return err
	}
	return nil
}

func (eh *loggerHandler) logRequest(r *http.Request) error {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return eh.logPayload(b, r.Header.Get("Content-Type"))
}

// call svc and add send request/responses to logUrl
func (eh *loggerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// log Request
	err := eh.logRequest(r)
	if err != nil {
		eh.log.Error(err, "Failed to log request")
	}

	// Predict
	b, err := eh.callService(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// log response
	err = eh.logPayload(b, r.Header.Get("Content-Type"))
	if err != nil {
		eh.log.Error(err, "Failed to log response")
	}

	// Write final response
	_, err = w.Write(b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
