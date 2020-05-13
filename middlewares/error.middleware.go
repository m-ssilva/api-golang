package middlewares

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// Message struct
type Message struct {
	Message string `json:"message"`
}

// ErrorModel struct
type ErrorModel struct {
	Name   string  `json:"name"`
	Status int     `json:"status"`
	Body   Message `json:"body"`
}

func getErrorSchemas() []ErrorModel {
	schemas := make([]ErrorModel, 30)
	raw, _ := ioutil.ReadFile("./helpers/errors-schema.json")
	json.Unmarshal(raw, &schemas)
	return schemas
}

func getErrorMessage(errorString string, errorModels ...ErrorModel) ErrorModel {
	for _, schema := range errorModels {
		if errorString == schema.Name {
			errorModel := ErrorModel{schema.Name, schema.Status, schema.Body}
			return errorModel
		}
	}
	errorModel := ErrorModel{errorModels[0].Name, errorModels[0].Status, errorModels[0].Body}
	return errorModel
}

// RootHandler is the default handler
type RootHandler func(http.ResponseWriter, *http.Request) error

func (fn RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	errorModels := getErrorSchemas()

	err := fn(w, r)
	if err == nil {
		return
	}

	responseMessage := getErrorMessage(err.Error(), errorModels...)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseMessage.Status)
	io.WriteString(w, `{ "message": "`+responseMessage.Body.Message+`"}`)
}
