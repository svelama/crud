package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/svelama/go/http/handlers"
)

type Person struct{
	name string
	age int
}

func NewPerson(name string, age int) Person{
	return Person{name, age}
}

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	products := handlers.NewProducts(l)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", products.GetProducts)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", products.AddProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", products.UpdateProducts)

	delRouter := sm.Methods(http.MethodDelete).Subrouter()
	delRouter.HandleFunc("/products/{id:[0-9]+}", products.DeleteProducts)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	go func(){
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	sig := <- sigChan
	l.Println("Received terminate, graceful shutdown", sig)
	
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}