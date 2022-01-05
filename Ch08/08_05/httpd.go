// HTTP server example
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// helloHandler returns a greeting
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Gophers!")
}

// MathRequest is a request of math operation
type MathRequest struct {
	Op string `json:"op"`
	Left float64 `json:"left"`
	Right float64 `json:"right"`
}

// MathResponse is a response to MathRequest
type MathResponse struct {
	Error string `json:"error"`
	Result float64 `json:"result"`
}

// mathHandler does a math operation
func mathHandler(w http.ResponseWriter, r *http.Request) {
	// Decode request
	defer r.Body.Close();
	reqDecoder := json.NewDecoder(r.Body);
	reqData := &MathRequest{}

	if err := reqDecoder.Decode(reqData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest);
		return;
	}

	// Process operation from request data and build response data
	respData := &MathResponse{};
	switch reqData.Op {
	case "+":
		respData.Result = reqData.Left + reqData.Right;
	case "-":
		respData.Result = reqData.Left - reqData.Right;
	case "*":
		respData.Result = reqData.Left * reqData.Right;
	case "/":
		if reqData.Right == 0.0 {
			respData.Error = "division by 0"
		} else {
			respData.Result = reqData.Left / reqData.Right;
		}
	default:
		respData.Error = fmt.Sprintf("'%s' is not a valid operation", reqData.Op);
	}

	// Encode and return result
	w.Header().Set("Content-Type", "application/json");
	if respData.Error != "" {
		w.WriteHeader(http.StatusBadRequest);
	}

	respEncoder := json.NewEncoder(w);
	if err := respEncoder.Encode(respData); err != nil {
		// Cannot return error to client here
		log.Printf("Cannot encode %v - Error : %s", respData, err);
	}


}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/math", mathHandler);

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
