package main

import (
// 	"database/sql"
	"log"
	"fmt"
	"strconv"
	"net/http"
	"github.com/krolim/shoppinglist/dbmanager"
// 	_ "github.com/lib/pq"
)

type Hello struct{}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello!")
}


func main() {
	// p := &dbmanager.Product{ Type_id: 1, 
	// 						Title: "Bread 1", 
	// 						Description: "Wholegrained bread 1", 
	// 						Images: "{}"}
	// err := dbmanager.AddProduct(p)
	// err := dbmanager.RemoveProduct(1)
	err,products := dbmanager.FetchAllProducts()
	if (err != nil) {
		log.Fatal(err)
		panic("")
	}
	for i := 0; i < len(products); i++ {
		fmt.Printf("row %v: %+v\n", i, products[i])	
	}

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		switch r.Method {
			case "DELETE":
				if len(r.Form["id"]) == 0  {
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintf(w, "Missing id")

				} else {
					id,err := strconv.ParseInt(r.Form["id"][0], 0, 64)
					if (err != nil) {
						w.WriteHeader(http.StatusBadRequest)
						fmt.Fprintf(w, "Invalid id %v", err)
					} else {
						dbmanager.RemoveProduct(int(id))
						fmt.Fprintf(w, "Deleted: %v", r.Form["id"][0])
					}
				}
			default:
				err,products := dbmanager.FetchAllProducts()
				if (err != nil) {
					log.Fatal(err)
					panic("")
				}
				for i := 0; i < len(products); i++ {
					fmt.Fprintf(w, "row %v: %+v\n", i, products[i])	
				}
							// fmt.Fprintf(w, "Welcome to the bar %v", r.URL.Path)	
		}		
					
	})
	// var h Hello
	err = http.ListenAndServe("localhost:4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}