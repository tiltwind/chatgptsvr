package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type GPT3Request struct {
	Model       string      `json:"model"`
	Prompt      string      `json:"prompt"`
	MaxTokens   int         `json:"max_tokens"`
	Temperature int         `json:"temperature"`
	TopP        int         `json:"top_p"`
	N           int         `json:"n"`
	Stream      bool        `json:"stream"`
	Logprobs    interface{} `json:"logprobs"`
	Stop        string      `json:"stop"`
}

type GPT3Choice struct {
	Text string `json:"text"`
}
type GPT3Response struct {
	Choices []*GPT3Choice `json:"choices"`
}

func gptRequest(question string) (string, error) {
	req := GPT3Request{
		Model:       "text-davinci-003",
		Prompt:      question,
		MaxTokens:   4000,
		Temperature: 0,
		TopP:        1,
		N:           1,
		Stream:      false,
		Logprobs:    nil,
	}
	body, _ := json.Marshal(req)
	httpReq, _ := http.NewRequest(http.MethodPost, "https://api.openai.com/v1/completions", bytes.NewReader(body))
	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", openaiToken))
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := HttpClient.Do(httpReq)
	if err != nil {
		return "", err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	gpt3Reps := &GPT3Response{}
	json.Unmarshal(respBody, gpt3Reps)

	if len(gpt3Reps.Choices) == 0 {
		return "No Response", nil
	}

	return gpt3Reps.Choices[0].Text, nil
}

const (
	DefaultMaxIdleConns        = 32
	DefaultMaxIdleConnsPerHost = 8
	DefaultMaxConnsPerHost     = 64
	DefaultIdleConnTimeout     = time.Second * 8

	DefaultRequestTimeout = time.Second * 120
)

// nolint:exhaustivestruct // ignore this
var HttpClient = &http.Client{
	Transport: &http.Transport{
		MaxIdleConns:        DefaultMaxIdleConns,
		MaxIdleConnsPerHost: DefaultMaxIdleConnsPerHost,
		MaxConnsPerHost:     DefaultMaxConnsPerHost,
		IdleConnTimeout:     DefaultIdleConnTimeout,
	},
	Timeout: DefaultRequestTimeout,
}
