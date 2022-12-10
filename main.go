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

	writer.Header().Set("Access-Control-Allow-Origin", "*")
	
	question := request.URL.Query().Get("q")
	if question == "" {
		writer.Write([]byte("pls ask your question"))
		return
	}

	answer, err := gptRequest(question)
	if err != nil {
		responseError(writer, err)
		return
	}

	fmt.Printf("question: %s, answer: %s\n", question, answer)

	writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	writer.Write([]byte(answer))
}

func responseError(writer http.ResponseWriter, err any) {
	errInfo, _ := json.Marshal(err)
	writer.Write(errInfo)
}
