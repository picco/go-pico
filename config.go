package pico

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// ConfigInterface type
type ConfigInterface interface {
	Import(string) string
	Get(string) string
	Set(string, string) string
	SetInt(string, int) string
}

// Config type
type Config struct {
	config map[string]string
}

// NewConfig func
func NewConfig() ConfigInterface {
	log.Println("Config: initializing")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Config: error loading .env")
	}

	return &Config{
		config: make(map[string]string),
	}
}

// Import func
func (config *Config) Import(name string) string {
	config.config[name] = os.Getenv(name)
	if len(config.config[name]) > 0 {
		log.Printf("Config: setting %v = %v", name, config.config[name])
	} else {
		log.Fatalf("Config: missing environment variable: %v", name)
	}

	return config.config[name]
}

// Get func
func (config *Config) Get(name string) string {
	value, ok := config.config[name]
	if !ok {
		return config.Import(name)
	}

	return value
}

// Set func
func (config *Config) Set(name string, value string) string {
	config.config[name] = value
	return value
}

// SetInt func
func (config *Config) SetInt(name string, value int) string {
	return config.Set(name, strconv.Itoa(value))
}
