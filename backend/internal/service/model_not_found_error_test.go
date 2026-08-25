package service

import (
	"net/http"
	"testing"
)

func TestIsUpstreamModelUnavailableError(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		body       []byte
		want       bool
	}{
		{
			name:       "404 model not found message",
			statusCode: http.StatusNotFound,
			body:       []byte(`{"error":{"message":"model not found"}}`),
			want:       true,
		},
		{
			name:       "404 model_not_found code",
			statusCode: http.StatusNotFound,
			body:       []byte(`{"error":{"code":"model_not_found","message":"The requested model was not found"}}`),
			want:       true,
		},
		{
			name:       "404 unknown model message",
			statusCode: http.StatusNotFound,
			body:       []byte(`{"error":{"message":"unknown model gpt-5.4"}}`),
			want:       true,
		},
		{
			name:       "404 endpoint not found is not model specific",
			statusCode: http.StatusNotFound,
			body:       []byte(`{"error":{"message":"endpoint not found"}}`),
			want:       false,
		},
		{
			name:       "404 arbitrary body is not model specific",
			statusCode: http.StatusNotFound,
			body:       []byte(`404 page not found`),
			want:       false,
		},
		{
			name:       "non 404 does not match",
			statusCode: http.StatusBadRequest,
			body:       []byte(`{"error":{"message":"model not found"}}`),
			want:       false,
		},
		{
			name:       "400 ChatGPT account does not support model",
			statusCode: http.StatusBadRequest,
			body:       []byte(`{"error":{"message":"The 'gpt-5.6-sol' model is not supported when using Codex with a ChatGPT account.","type":"invalid_request_error"}}`),
			want:       true,
		},
		{
			name:       "400 unsupported parameter is not model capability",
			statusCode: http.StatusBadRequest,
			body:       []byte(`{"error":{"message":"Unsupported parameter: temperature","type":"invalid_request_error"}}`),
			want:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUpstreamModelUnavailableError(tt.statusCode, tt.body); got != tt.want {
				t.Fatalf("isUpstreamModelUnavailableError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAntigravityModelNotFoundKeepsBare404Fallback(t *testing.T) {
	if !isModelNotFoundError(http.StatusNotFound, []byte(`endpoint not found`)) {
		t.Fatal("antigravity model-not-found helper should keep bare 404 fallback")
	}
}
