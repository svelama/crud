package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/svelama/go/http/data"
)

type Products struct{
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products{
	return &Products{l}
}

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request){

	p.l.Println("Get method called")
	lp := data.GetProducts()
	if err := lp.ToJson(w); err != nil {
		http.Error(w, "failed to fetch products", http.StatusInternalServerError)
	}
}

func (p *Products) AddProducts(w http.ResponseWriter, r *http.Request){

	p.l.Println("Post method called")
	raw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to add products", http.StatusBadRequest)
		return
	}

    product := &data.Product{}
	err = json.Unmarshal(raw, product)
	if err != nil {
		http.Error(w, "failed to unmarshal json", http.StatusBadRequest)
		return
	}
	
	data.AddProducts(product)
}

func (p *Products) UpdateProducts(w http.ResponseWriter, r *http.Request){
	
	vars := mux.Vars(r)
	id := vars["id"]
	
	p.l.Println("Put method called, id: ", id)

	// Get Products object from Json
	product := &data.Product{}
	err := product.FromJson(r.Body)
	if err != nil {
		http.Error(w, "failed to unmarshal json", http.StatusBadRequest)
		return
	}

	err = data.UpdateProducts(product, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (p *Products) DeleteProducts(w http.ResponseWriter, r *http.Request){
	p.l.Println("Delete method called")
	vars := mux.Vars(r)
	id := vars["id"]

	if !data.RemoveProducts(id) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Product does not exist"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Product successfully deleted"))
}