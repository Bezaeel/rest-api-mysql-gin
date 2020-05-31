package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"

	"github.com/bezaeel/rest-api-mysql-gin/pkg/contact"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kelseyhightower/envconfig"

)

func main() {

	var cfg Config

	// Configuration
	err := envconfig.Process("", &cfg)
	if err != nil {
		panic(fmt.Errorf("error loading configuration: %s", err.Error()))
	}

	db, err := gorm.Open("mysql", cfg.toURL())
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer db.Close()

	db.AutoMigrate(contact.Contact{})

	// init routes
	mux := NewBaseMux(
		TextHandler(http.StatusOK, "application/json", `"ready"`).ServeHTTP,
	)

	statusRepo := contact.NewRepo(db)

	routes, err := initRoutes(
		statusRepo,
	)
	if err != nil {
		log.Fatal("failure initialising routes", err)
	}

	mux.Handle("/", routes)
}

func NewBaseMux(ready http.HandlerFunc) *http.ServeMux {
	mux := http.NewServeMux()

	// /alive always responds with 200 OK
	mux.HandleFunc("/alive", TextHandler(http.StatusOK, "application/json", `"OK"`))

	// /ready is a custom handler, `httputil.Ready` can be used for this
	mux.Handle("/health", ready)

	// pprof allows remote profiling
	// more info: https://golang.org/pkg/net/http/pprof/
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	return mux
}


// TextHandler returns a HandlerFunc that writes a constant Content-Type and string as a response.
func TextHandler(status int, contentType, response string) http.HandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) {
		if contentType != "" {
			wr.Header().Set("Content-Type", contentType)
		}

		wr.WriteHeader(status)
		fmt.Fprint(wr, response)
	}
}
