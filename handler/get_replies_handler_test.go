package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ameykpatil/chatbot/client"
	"github.com/ameykpatil/chatbot/service"
	"github.com/ameykpatil/chatbot/storage"
)

func TestGetRepliesHandler_MissingMessage(t *testing.T) {

	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/replies", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("bot_id", "123")
	req.URL.RawQuery = q.Encode()

	// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response
	rr := httptest.NewRecorder()
	handler := NewGetReplyHandler(service.ReplyService{})

	// Handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	// Check the response body is what we expect
	expected := `invalid query parameters`
	got := strings.TrimSuffix(rr.Body.String(), "\n")
	if got != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			got, expected)
	}
}

func TestGetRepliesHandler_MissingBotID(t *testing.T) {

	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/replies", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("message", "hello")
	req.URL.RawQuery = q.Encode()

	// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response
	rr := httptest.NewRecorder()
	handler := NewGetReplyHandler(service.ReplyService{})

	// Handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	// Check the response body is what we expect
	expected := `invalid query parameters`
	got := strings.TrimSuffix(rr.Body.String(), "\n")
	if got != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			got, expected)
	}
}

func TestGetRepliesHandler_InternalError(t *testing.T) {

	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/replies", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("bot_id", "123")
	q.Add("message", "hello")
	req.URL.RawQuery = q.Encode()

	// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response
	rr := httptest.NewRecorder()

	// Mock handler dependencies
	ic := client.NewIntentClient(&http.Client{}, "", "")
	rs := service.NewReplyService(storage.ReplyReader{}, *ic)
	handler := NewGetReplyHandler(*rs)

	// Handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
