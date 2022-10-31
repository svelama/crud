package handlers

import (
	"github.com/svelama/go/http/data"
	"log"
	"net/http"
)

type Products struct{
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products{
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request){

	if r.Method == http.MethodGet {
		getProducts(w, r)
		return
	}

	// Return 
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func getProducts(w http.ResponseWriter, r *http.Request){

	lp := data.GetProducts()
	if err := lp.ToJson(w); err != nil {
		http.Error(w, "failed to fetch products", http.StatusInternalServerError)
	}
}