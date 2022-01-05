/*Challenge : Key-value database
* Write an HTTP server that functions as a key/value database.
* - Users will make a POST request to /db create/assign a value
* - User will make a GET request to /db/<key> to get a value
* - Use a map[string]interface{} to hold values
* - Limit Access the map with sync.Mutex
 */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

// this is our fake database
var fakeKeyValueDatabase = struct {
	sync.RWMutex
	db map[string]interface{}
} {
	db: make(map[string]interface{}),
} ;

// func to interface with database to read value from key
func readValueFromDatabase(key string) (interface{}, bool) {
	fakeKeyValueDatabase.RLock();
	elem, ok := fakeKeyValueDatabase.db[key];
	fakeKeyValueDatabase.RUnlock();
	return elem, ok;
}

// func to inteface with database to write value
func writeValueToDatabase(key string, value interface{}) {
	fakeKeyValueDatabase.Lock();
	fakeKeyValueDatabase.db[key] = value;
	fakeKeyValueDatabase.Unlock();
}

// helloHandler returns a greeting
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Gophers!\n")
}

// dbGetRequest is a request to get a value from database
type dbGetRequest struct {
	Key string `json:"key"`
	Value interface{} `json:"value"`
}

// dbGetResponse is a response to get a value from database
type dbGetResponse struct {
	Key string `json:"key"`
	Value interface{} `json:"value"`
}

// dbGetRequest is a request to create/assign a value from database
type dbPostRequest struct {
	Key string `json:"key"`
	Value interface{} `json:"value"`
}

// dbGetResponse is a response to create/assign a value from database
type dbPostResponse struct {
	Key string `json:"key"`
	Value interface{} `json:"value"`
}

// mathHandler does a math operation
func dbGetHandler(w http.ResponseWriter, r *http.Request) {
	// Decode request
	defer r.Body.Close();

	urlPathElements := strings.Split(strings.Trim(r.URL.Path,"/"),"/");
	if len(urlPathElements) > 2 {
		http.Error(w, fmt.Sprintf("invalid URL Path %s contains %q\n", r.URL.Path, urlPathElements), http.StatusBadRequest);
		return;
	}

	reqData := dbGetRequest {
		Key : urlPathElements[1],
	}

	value, ok := readValueFromDatabase(reqData.Key);
	if !ok {
		http.Error(w, fmt.Sprintf("Key <%s> not found in database ", reqData.Key), http.StatusNotFound);
		return;
	}

	// process operation to get a value
	respData := &dbGetResponse{
		Key: reqData.Key,
		Value : value,
	};

	// Encode and return result
	w.Header().Set("Content-Type", "application/json");

	respEncoder := json.NewEncoder(w);
	if err := respEncoder.Encode(respData); err != nil {
		// Cannot return error to client here
		log.Printf("Cannot encode %v - Error : %s", respData, err);
	}

}

// mathHandler does a math operation
func dbPostHandler(w http.ResponseWriter, r *http.Request) {
	// Decode request
	defer r.Body.Close();
	reqDecoder := json.NewDecoder(r.Body);
	reqData := &dbPostRequest{}

	if err := reqDecoder.Decode(reqData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest);
		return;
	}

	// return an error if they key is empty (value can be empty to reset/delete value)
	if reqData.Key == "" {
		http.Error(w, "invalid JSON Payload", http.StatusBadRequest);
		return;
	}

	// Process operation from request data and write to database
	writeValueToDatabase(reqData.Key, reqData.Value);

	// write response with content from database
	value, _ := readValueFromDatabase(reqData.Key);
	respData := &dbPostResponse{
		Key: reqData.Key,
		Value: value,
	};

	// Encode and return result
	w.Header().Set("Content-Type", "application/json");
	w.WriteHeader(http.StatusCreated); // confirm a new entity has been created

	respEncoder := json.NewEncoder(w);
	if err := respEncoder.Encode(respData); err != nil {
		// Cannot return error to client here
		log.Printf("Cannot encode %v - Error : %s", respData, err);
	}

}

func main() {
	// register http handler
	http.HandleFunc("/", helloHandler);
	http.HandleFunc("/db", dbPostHandler);
	http.HandleFunc("/db/", dbGetHandler);


	// start http server

	fmt.Println("Starting http server on port 8080");
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err);
	}

}
