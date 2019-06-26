package main

import (
	"encoding/json"
	"log"
	"net/http"

	. "./config"
	"gopkg.in/mgo.v2/bson"

	. "./dao"

	. "./models"
	"github.com/gorilla/mux"
)

var config = Config{}
var dao = BooksDAO{}

func AllBooksEndPoint(w http.ResponseWriter, r *http.Request) {
	books, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, books)
}

func FindBookEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	book, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Book ID")
		return
	}
	respondWithJson(w, http.StatusOK, book)
}

func CreateBookEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	book.ID = bson.NewObjectId()
	if err := dao.Insert(book); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, book)
}

func UpdateBookEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	if err := dao.Update(book); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteBookEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	if err := dao.Delete(book); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books", AllBooksEndPoint).Methods("GET")
	r.HandleFunc("/books", CreateBookEndPoint).Methods("POST")
	r.HandleFunc("/books", UpdateBookEndPoint).Methods("PUT")
	r.HandleFunc("/books", DeleteBookEndPoint).Methods("DELETE")
	r.HandleFunc("/books/{id}", FindBookEndPoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
