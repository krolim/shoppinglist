package main

import (
	// 	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/krolim/shoppinglist/dbmanager"
	"github.com/krolim/shoppinglist/messageparser"
	"log"
	"net/http"
	"strconv"
	// 	_ "github.com/lib/pq"
)

type Hello struct{}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

type ProductsJson struct {
	Products []*dbmanager.Product
}

type OrdersJson struct {
	Orders []*dbmanager.Order
}

type NewProductOrederedJson struct {
	ProductId int
	Amount    int
}

func initHttpHandlers() {
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		switch r.Method {
		case "DELETE":
			if len(r.Form["id"]) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Missing id")

			} else {
				id, err := strconv.ParseInt(r.Form["id"][0], 0, 64)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintf(w, "Invalid id %v", err)
				} else {
					// dbmanager.RemoveProduct(int(id))
					fmt.Fprintf(w, "Deleted: %v", id)
				}
			}
		default:
			err, products := dbmanager.FetchAllProducts()
			if err != nil {
				log.Fatal(err)
				panic("")
			}
			var response ProductsJson
			response.Products = products
			enc := json.NewEncoder(w)
			enc.Encode(response)
		}

	})

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		switch r.Method {
		case "GET":
			enc := json.NewEncoder(w)
			if len(r.Form["id"]) == 0 {
				// load all
				orders := dbmanager.LoadAllOrders()
				response := OrdersJson{Orders: orders}
				enc.Encode(response)
			} else {
				id, err := strconv.ParseInt(r.Form["id"][0], 0, 64)
				if err == nil {
					order := dbmanager.LoadOrder(int(id))
					enc.Encode(order)
				} else {
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintf(w, "Invalid id %v", err)
				}
			}
		case "PUT":
			var orderedProduct NewProductOrederedJson
			dec := json.NewDecoder(r.Body)
			err := dec.Decode(&orderedProduct)
			if err != nil {
				fmt.Printf("Error: %v", err)
			} else {
				// fmt.Println(orderedProduct.ProductId)
				dbmanager.AddToOrder(orderedProduct.ProductId, orderedProduct.Amount)
			}
			// fmt.Println(orderedProduct)
			//
		default:
			err, products := dbmanager.FetchAllProducts()
			if err != nil {
				log.Fatal(err)
				panic("")
			}
			var response ProductsJson
			response.Products = products
			enc := json.NewEncoder(w)
			enc.Encode(response)
		}

	})
}

func startWebServer() {
	err := http.ListenAndServe("localhost:4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	messageparser.ParseMsg("едно butter два килограма peaches")
	// initHttpHandlers()
	// startWebServer()
}
