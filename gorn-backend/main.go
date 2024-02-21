package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/drpulpo/gorn/model"
	"github.com/drpulpo/gorn/repository"
	"github.com/gorilla/mux"
)

func getAllOrderHandler(w http.ResponseWriter, r *http.Request) {

	// get all the users in the db
	orders, err := repository.GetAllOrders()

	if err != nil {
		log.Fatalf("Unable to get all order. %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(orders)
}

func userHandler(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	log.Print(path)
	vars := mux.Vars(req)

	id := vars["id"]

	i, err := strconv.Atoi(id)

	log.Print(id)
	if err != nil {
		fmt.Printf("unable to convert id:%v to int", id)
	}
	user, err := repository.GetUserData(i)

	log.Print(user)

	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}

func orderHandler(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	log.Print(path)
	vars := mux.Vars(req)

	id := vars["id"]

	i, err := strconv.Atoi(id)

	log.Print(id)
	if err != nil {
		fmt.Printf("Incorrect Id:%v to int", id)
	}
	order, err := repository.GetOrderData(i)

	log.Print(order)

	if err != nil {
		log.Fatalf("Error at getting order. %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(order)
}

func addUserHandler(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	log.Printf("Requested path %v", path)

	decoder := json.NewDecoder(req.Body)

	var user model.User
	// decode the data and assign it to model.User type
	err := decoder.Decode(&user)

	if err != nil {
		log.Fatalf("Error at body %v", err)
	}

	repository.InsertRecord(user)
	log.Printf("User Inserted: %+v", user)

}

func addOrderHandler(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	log.Printf("path is %v", path)

	decoder := json.NewDecoder(req.Body)

	var order model.Order

	err := decoder.Decode(&order)

	if err != nil {
		log.Fatalf("Error at body %v", err)
	}

	repository.InsertOrder(order)

	log.Printf("Order Inserted:  %+v", order)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

}

func deleteOrderHandler(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	log.Print(path)
	vars := mux.Vars(req)

	id := vars["id"]

	i, err := strconv.Atoi(id)

	log.Print(id)
	if err != nil {
		fmt.Printf("Incorrect id:%v to int", id)
	}
	order, err := repository.DeleteOrder(i)

	log.Print(order)

	if err != nil {
		log.Fatalf("Error getting Order. %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(order)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GoRN API root!")
}

func main() {
	r := mux.NewRouter()

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := repository.MigrateDb()

	if err != nil {
		log.Fatalf("Error at db migration %v", err)
	}
	// Start Server
	go func() {
		log.Println("API Server Starting :8080")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	//TODO . Will userHandler() with brace. work instead of userHandler
	r.HandleFunc("/users/{id:[0-9]+}", userHandler)
	r.HandleFunc("/addUser", addUserHandler)
	r.HandleFunc("/orders", getAllOrderHandler)
	r.HandleFunc("/orders/{id:[0-9]+}", orderHandler)
	r.HandleFunc("/addOrder", addOrderHandler)
	r.HandleFunc("/deleteOrder/{id:[0-9]+}", deleteOrderHandler)
	r.HandleFunc("/", rootHandler)

	// Graceful Shutdown
	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}
