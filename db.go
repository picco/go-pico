package pico

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// SQLDatabase type
type SQLDatabase struct {
	Core
}

// UseSQLDatabase func
func (service *SQLDatabase) UseSQLDatabase(dialect string, conn string) *gorm.DB {
	log.Println("DB: initializing")

	db, err := gorm.Open(dialect, conn)
	if err != nil {
		log.Fatal("DB: Failed to connect")
	} else {
		log.Println("DB: Connected successfully")
	}

	return db
}

// GORMResponse func
func (service *SQLDatabase) GORMResponse(res http.ResponseWriter, result *gorm.DB) {
	if result.RecordNotFound() {
		service.APIErrorNotFoundResponse(res)
	} else {
		service.APIResponse(res, &JSONResponse{
			Data: result.Value,
		})
	}
}
