package pico

import (
	"encoding/gob"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

// SessionInterface type
type SessionInterface interface {
	Get(*http.Request, string) (*sessions.Session, error)
}

// Session type
type Session struct {
	name  string
	store *sessions.FilesystemStore
}

// NewSession func
func NewSession(name string) SessionInterface {
	log.Println("Session: initializing")
	gob.Register(map[string]interface{}{})

	service := &Session{
		name:  name,
		store: sessions.NewFilesystemStore("", []byte(name)),
	}

	return service
}

// Get func
func (service *Session) Get(r *http.Request, name string) (*sessions.Session, error) {
	return service.store.Get(r, name)
}
