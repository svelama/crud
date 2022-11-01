package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/svelama/go/http/data"
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

	if r.Method == http.MethodPost {
		addProducts(w, r)
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

func addProducts(w http.ResponseWriter, r *http.Request){

	raw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to add products", http.StatusBadRequest)
		return
	}

    p := &data.Product{}
	err = json.Unmarshal(raw, p)
	if err != nil {
		http.Error(w, "failed to unmarshal json", http.StatusBadRequest)
		return
	}
	
	data.AddProducts(p)
}