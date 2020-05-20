package pico

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// HTTPServerInterface type
type HTTPServerInterface interface {
	UseRouter(*chi.Mux) HTTPServerInterface
	Router() *chi.Mux
	Server() *http.Server
	Start()
	Stop()
}

// HTTPServer type
type HTTPServer struct {
	addr   string
	router *chi.Mux
	server *http.Server
}

// NewHTTPServer func
func NewHTTPServer(addr string) HTTPServerInterface {
	log.Println("HTTP: initializing")

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Welcome to Pico!"))
	})

	server := &HTTPServer{
		addr:   addr,
		router: router,
		server: &http.Server{
			Addr:    addr,
			Handler: router,
		},
	}

	return server
}

// Server func
func (server *HTTPServer) Server() *http.Server {
	return server.server
}

// UseRouter func
func (server *HTTPServer) UseRouter(router *chi.Mux) HTTPServerInterface {
	// TODO
	return server
}

// Router func
func (server *HTTPServer) Router() *chi.Mux {
	return server.router
}

// Start func
func (server *HTTPServer) Start() {
	go func() {
		log.Printf("HTTP: server staring at %v\n", server.addr)
		if err := server.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	//select {}
}

// Stop func
func (server *HTTPServer) Stop() {
	log.Println("HTTP: server stopping")
	server.server.Shutdown(context.Background())
}
