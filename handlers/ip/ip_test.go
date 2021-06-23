package ip

import (
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

	if `Encoding responso to [yaml] is not implemented
` != w.Body.String() {
		t.Error(`Expected "Encoding responso to [yaml] is not implemented", got `, w.Body.String())
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

	if `Error parsing remote address [123]
` != w.Body.String() {
		t.Error(`Expected "Error parsing remote address [123]", got [`, w.Body.String(), "]")
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
