package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("HomeHandler")
	w.Write([]byte("HomeHandler"))
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("ProductsHandler")
	w.Write([]byte("ProductsHandler"))
}

func articlesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("ArticlesHandler")
	w.Write([]byte("ArticlesHandler"))
}

func m1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware1")
		next.ServeHTTP(w, r)
	})
}

func m2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware2")
		next.ServeHTTP(w, r)
	})
}

func m3(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware3")
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)

	r.Use(m3)

	rp := r.PathPrefix("/products").Subrouter()
	rp.HandleFunc("", productsHandler)
	rp.Use(m1)

	ra := r.PathPrefix("/articles").Subrouter()
	ra.HandleFunc("", articlesHandler)
	ra.Use(m2)

	server := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
