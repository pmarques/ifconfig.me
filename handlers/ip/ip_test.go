package ip

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestHandlerDefault(t *testing.T) {
	req := httptest.NewRequest("GET", "http:example.com/ip", nil)
	w := httptest.NewRecorder()
	Handler(w, req)

	if 200 != w.Code {
		t.Error("Expected HTTP status code 200, got [", w.Code, "]")
	}

	// 192.0.2.0/24 is "TEST-NET" and is forced @ httptest.go
	if `{"ip":"192.0.2.1"}` != w.Body.String() {
		t.Error(`Expected '{"ip":"192.0.2.1"}', got [`, w.Body.String(), "]")
	}
}

func TestHandlerXML(t *testing.T) {
	req := httptest.NewRequest("GET", "http:example.com/ip?f=xml", nil)
	w := httptest.NewRecorder()
	Handler(w, req)

	if 200 != w.Code {
		t.Error("Expected HTTP status code 200, got [", w.Code, "]")
	}

	var expectedOutput = `<?xml version="1.0" encoding="UTF-8"?>
  <response>
      <ip>192.0.2.1</ip>
  </response>`
	if expectedOutput != w.Body.String() {
		t.Error(`Expected "`, expectedOutput, `, got "`, w.Body.String())
	}
}

func TestHandlerYaml(t *testing.T) {
	req := httptest.NewRequest("GET", "http:example.com/ip?f=yaml", nil)
	w := httptest.NewRecorder()
	Handler(w, req)

	if 501 != w.Code {
		t.Error("Expected HTTP status code 501, got [", w.Code, "]")
	}

	if `Encoding response to [yaml] is not implemented
` != w.Body.String() {
		t.Error(`Expected "Encoding response to [yaml] is not implemented", got `, w.Body.String())
	}
}

func TestHandlerIPParseError(t *testing.T) {
	req := httptest.NewRequest("GET", "http:example.com/ip", nil)
	// monkey patch to throw error when parsing address
	req.RemoteAddr = "123"
	w := httptest.NewRecorder()
	Handler(w, req)

	if 500 != w.Code {
		t.Error("Expected HTTP status code 500, got [", w.Code, "]")
	}

	if `Error retrieving client IP
` != w.Body.String() {
		t.Error(`Expected "Error retrieving client IP
", got [`, w.Body.String(), "]")
	}
}

func TestHandlerXForwardedFor(t *testing.T) {
	req := httptest.NewRequest("GET", "http:example.com/ip", nil)
	req.Header["X-Forwarded-For"] = []string{"1.1.1.1, 2.2.2.2"}
	w := httptest.NewRecorder()
	Handler(w, req)

	if 200 != w.Code {
		t.Error("Expected HTTP status code 200, got [", w.Code, "]")
	}

	if `{"ip":"1.1.1.1"}` != w.Body.String() {
		t.Error(`Expected '{"ip":"1.1.1.1"}' got `, w.Body.String())
	}
}

func FuzzHandler(f *testing.F) {
	f.Add("1.1.1.1")
	f.Add("2.2.2.2")
	f.Add("1.1.1.1, 2.2.2.2")
	f.Add("1.1.1.1, 2.2.2.2, 3.3.3.3")
	f.Fuzz(func(t *testing.T, xForwardFor string) {
		req := httptest.NewRequest("GET", "http:example.com/ip", nil)
		req.Header["X-Forwarded-For"] = []string{xForwardFor}
		w := httptest.NewRecorder()
		Handler(w, req)

		if 200 != w.Code {
			t.Error("Expected HTTP status code 200, got [", w.Code, "]")
		}

		if !json.Valid([]byte(w.Body.Bytes())) {
			t.Error("Invalid JSON ", w.Body.String())
		}
	})
}
