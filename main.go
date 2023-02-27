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

// Define the stucts and the variables that we will use
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

// processing the function that handles request
func processOperation(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var requestBody input
	_ = json.NewDecoder(r.Body).Decode(&requestBody)
	var requestResponse output

	if requestBody.Operation == "deduplicate" || requestBody.Operation == "getPairs" {

		id := uuid.New()
		if requestBody.Operation == "deduplicate" {
			requestResponse = output{
				ID:        string(id.String()),
				Operation: requestBody.Operation,
				Data:      methods.Deduplicate(requestBody.Data),
			}
			requestHistory = append(requestHistory, requestResponse)

			json.NewEncoder(w).Encode(requestResponse)
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

	} else {
		fmt.Fprintf(w, "Invalid operation!")
		return
	}

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/upload", processOperation).Methods("POST")

	fmt.Println("Starting server at port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
