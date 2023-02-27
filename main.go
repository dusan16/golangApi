package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	methods "go-task/lib"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Define the stucts and the variables that we will use for storing data with project
// Note: I haven't implemented persistent storeage as it wasn't stated in the task, so everything is stored in requestHistory slice

type input struct {
	Operation string  `json:"operation"`
	Data      []int64 `json:"data"`
}

type output struct {
	ID        string  `json:"id"`
	Operation string  `json:"operation"`
	Data      []int64 `json:"data"`
}

var requestHistory []output

// processing the function that handles POST request
func processOperation(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var requestBody input
	_ = json.NewDecoder(r.Body).Decode(&requestBody) // pass the value of the request body
	var requestResponse output

	//first we check if the valid operation has been passed
	if requestBody.Operation == "deduplicate" || requestBody.Operation == "getPairs" {

		id := uuid.New() // create new UUID value of id usnig google/uuid library function
		if requestBody.Operation == "deduplicate" {
			requestResponse = output{
				ID:        string(id.String()),
				Operation: requestBody.Operation,
				Data:      methods.Deduplicate(requestBody.Data),
			}
			requestHistory = append(requestHistory, requestResponse) // add the most recent request to the history

			json.NewEncoder(w).Encode(requestResponse) // send the response to client
			return

		} else {
			_, numbersWithPairs := methods.GetPairs(requestBody.Data)
			requestResponse = output{
				ID:        string(id.String()),
				Operation: requestBody.Operation,
				Data:      numbersWithPairs,
			}
			requestHistory = append(requestHistory, requestResponse)

			json.NewEncoder(w).Encode(requestResponse)
			return

		}

	} else { // if it isn't a valid request send back the notification
		fmt.Fprintf(w, "Invalid operation!")
		return
	}

}

func main() {
	r := mux.NewRouter() // create new instance of the mux router

	// create POST /upload endpoint
	r.HandleFunc("/upload", processOperation).Methods("POST")

	fmt.Println("Starting server at port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
