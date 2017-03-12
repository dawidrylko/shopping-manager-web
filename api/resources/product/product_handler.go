package product

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"

	"../../server"
)

// Handler method show all endpoint in product package
func Handler() {
	server.Router.HandleFunc("/product", create).Methods("POST")

	server.Router.HandleFunc("/product/{id}", update).Methods("PUT")

	server.Router.HandleFunc("/product", getAll).Methods("GET")
	server.Router.HandleFunc("/product/{id}", get).Methods("GET")

	server.Router.HandleFunc("/product/{id}", remove).Methods("DELETE")
}

func create(responseWriter http.ResponseWriter, request *http.Request) {
	product := Product{}

	json.NewDecoder(request.Body).Decode(&product)
	product.ID = bson.NewObjectId()

	server.GetSession().DB("shopping-manager").C("products").Insert(product)

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(201)

	productJSON, _ := json.Marshal(product)
	fmt.Fprintf(responseWriter, "%s", productJSON)
}

func update(responseWriter http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	if !bson.IsObjectIdHex(params["id"]) {
		responseWriter.WriteHeader(404)

		return
	}

	oid := bson.ObjectIdHex(params["id"])

	product := Product{}
	product.ID = oid
	json.NewDecoder(request.Body).Decode(&product)

	if error := server.GetSession().DB("shopping-manager").C("products").UpdateId(oid, product); error != nil {
		responseWriter.WriteHeader(404)

		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(200)

	productJSON, _ := json.Marshal(product)
	fmt.Fprintf(responseWriter, "%s", productJSON)
}

func getAll(responseWriter http.ResponseWriter, request *http.Request) {
	products := Products{}

	if error := server.GetSession().DB("shopping-manager").C("products").Find(nil).All(&products); error != nil {
		responseWriter.WriteHeader(404)

		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(200)

	productsJSON, _ := json.Marshal(products)
	fmt.Fprintf(responseWriter, "%s", productsJSON)
}

func get(responseWriter http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	oid := bson.ObjectIdHex(params["id"])
	product := Product{}

	if error := server.GetSession().DB("shopping-manager").C("products").FindId(oid).One(&product); error != nil {
		responseWriter.WriteHeader(404)

		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(200)

	productJSON, _ := json.Marshal(product)
	fmt.Fprintf(responseWriter, "%s", productJSON)
}

func remove(responseWriter http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	if !bson.IsObjectIdHex(params["id"]) {
		responseWriter.WriteHeader(404)

		return
	}

	oid := bson.ObjectIdHex(params["id"])

	if error := server.GetSession().DB("shopping-manager").C("products").RemoveId(oid); error != nil {
		responseWriter.WriteHeader(404)

		return
	}

	responseWriter.WriteHeader(200)
}