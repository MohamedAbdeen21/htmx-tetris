package main

import (
	"bytes"
	"context"
	"embed"
	"log"
	"net/http"
	"net/url"
	"strings"
	"tetris/cmd/routes" // Update this import to match your project structure

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

//go:embed style/*
var content embed.FS

// responseWriter is a custom implementation of http.ResponseWriter
type responseWriter struct {
	headers http.Header
	body    bytes.Buffer
	status  int
}

func (rw *responseWriter) Header() http.Header {
	return rw.headers
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	return rw.body.Write(b)
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.status = statusCode
}

// handleRequest is the Lambda handler function
func handleRequest(ctx context.Context, request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	var response events.LambdaFunctionURLResponse
	rw := &responseWriter{headers: http.Header{}}
	log.Printf("Request Headers = %v", request.Headers)

	switch {
	case strings.HasPrefix(request.RawPath, "/style/"):
		// Serve CSS content
		serveStaticContent(rw, request.RawPath)
	case request.RawPath == "/" && request.RequestContext.HTTP.Method == "GET":
		// Handle root route
		routes.Root(rw, nil)
	case request.RawPath == "/tick" && request.RequestContext.HTTP.Method == "POST":
		// Handle tick route
		routes.Tick(rw, request.Headers["action"])
	case request.RawPath == "/restart" && request.RequestContext.HTTP.Method == "GET":
		// Handle restart route
		routes.Restart(rw, nil)
	default:
		// Handle unknown routes
		http.NotFound(rw, nil)
	}

	response = events.LambdaFunctionURLResponse{
		Body:       rw.body.String(),
		StatusCode: rw.status,
		Headers:    map[string]string{"Content-Type": "text/html"},
	}

	for key, values := range rw.headers {
		if len(values) > 0 {
			response.Headers[key] = values[0]
		}
	}

	return response, nil
}

func serveStaticContent(rw http.ResponseWriter, path string) {
	fs := http.FileServer(http.FS(content))
	fs.ServeHTTP(rw, &http.Request{URL: &url.URL{Path: path}})
}

func main() {
	log.Print("Starting Lambda function")
	lambda.Start(handleRequest)
}
