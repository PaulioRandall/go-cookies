package uhttp

import (
	"fmt"
	"log"
	"net/http"
)

// LogRequest logs the details of a request such as the URL.
func LogRequest(req *http.Request) {
	log.Printf("(%s) %s\n", req.Method, req.URL.String())
}

// RelURL creates the absolute path of a requests URL without any fragment.
func RelURL(req *http.Request) string {
	r := req.URL.Path
	if req.URL.RawQuery != "" {
		r += "?" + req.URL.RawQuery
	}
	return r
}

// CorsHeaders represents a set of CORS headers for a response.
type CorsHeaders struct {
	Origin  string
	Headers string
	Methods string
}

// NewOpenCors returns a CorsHeaders struct with all values set as '*'.
func NewOpenCors() CorsHeaders {
	return CorsHeaders{"*", "*", "*"}
}

// UseCors sets within the response 'res' the CORS headers 'h'.
func UseCors(res *http.ResponseWriter, h *CorsHeaders) {
	(*res).Header().Set("Access-Control-Allow-Origin", h.Origin)
	(*res).Header().Set("Access-Control-Allow-Headers", h.Headers)
	(*res).Header().Set("Access-Control-Allow-Methods", h.Methods)
}

// UseUTF8Json sets the 'Content-Type' header of response 'res' as JSON. If the
// 'extension' is not empty it will inserted into the content type.
//
// E.g.
// Without extension: 	'application/json; charset=utf-8'
// With extension:		 	'application/openapi+json; charset=utf-8'
func UseUTF8Json(res *http.ResponseWriter, extension string) {
	var ct string
	if extension == "" {
		ct = "application/json; charset=utf-8"
	} else {
		ct = fmt.Sprintf("application/%s+json; charset=utf-8", extension)
	}
	(*res).Header().Set("Content-Type", ct)
}
