package pico

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// Core type
type Core struct{}

// JSONResponse type
type JSONResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

// UseConfig func
func (core *Core) UseConfig() ConfigInterface {
	return NewConfig()
}

// UseHTTPServer func
func (core *Core) UseHTTPServer(addr string) HTTPServerInterface {
	return NewHTTPServer(addr)
}

// URLParam func
func (core *Core) URLParam(req *http.Request, name string) string {
	return chi.URLParam(req, name)
}

// URLParamInt func
func (core *Core) URLParamInt(req *http.Request, name string) int {
	result, _ := strconv.Atoi(core.URLParam(req, name))
	return result
}

// APIResponse func
func (core *Core) APIResponse(res http.ResponseWriter, data interface{}) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(data)
}

// APIErrorResponse func
func (core *Core) APIErrorResponse(res http.ResponseWriter, message string) {
	core.APIResponse(res, &JSONResponse{
		Error: message,
	})
}

// APIErrorNotFoundResponse func
func (core *Core) APIErrorNotFoundResponse(res http.ResponseWriter) {
	core.APIErrorResponse(res, "Not found")
}
