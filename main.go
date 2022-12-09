package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	openaiToken = ""
	serverPort  = 8999
)

func main() {
	if err := http.ListenAndServe(fmt.Sprintf(":%d", serverPort), &GPT3Handler{}); err != nil {
		panic(err)
	}
}

type GPT3Handler struct {
}

func (G GPT3Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			responseError(writer, err)
		}
	}()

	question := request.URL.Query().Get("q")
	if question == "" {
		writer.Write([]byte("pls ask your question"))
		return
	}

	resp, err := gptRequest(question)
	if err != nil {
		responseError(writer, err)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(resp))
}

func responseError(writer http.ResponseWriter, err any) {
	errInfo, _ := json.Marshal(err)
	writer.Write(errInfo)
}
