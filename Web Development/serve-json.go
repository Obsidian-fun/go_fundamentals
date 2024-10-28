/***
sample on serving json files over the network,
related files : data.json , script.js


***/

package main

import (
    "encoding/json"
    "net/http"
    "log"
    "os"
)

type Data struct {
    // Define the structure of your JSON data
    // Example:
    ID   int    `json:"id"`
    Name string `json:"name"`
}

// Handler to serve the JSON
func jsonHandler(w http.ResponseWriter, r *http.Request) {
    // Open the large JSON file
    file, err := os.Open("data.json") // Assuming your large JSON file is named 'data.json'
    if err != nil {
        http.Error(w, "Unable to open file", http.StatusInternalServerError)
        return
    }
    defer file.Close()

    // Decode the JSON file
    var data []Data
    decoder := json.NewDecoder(file)
    if err := decoder.Decode(&data); err != nil {
        http.Error(w, "Unable to decode JSON", http.StatusInternalServerError)
        return
    }

    // Set the content type to application/json
    w.Header().Set("Content-Type", "application/json")
    
    // Write the JSON response
    if err := json.NewEncoder(w).Encode(data); err != nil {
        http.Error(w, "Unable to encode to JSON", http.StatusInternalServerError)
        return
    }
}

func main() {
    http.HandleFunc("/api/data", jsonHandler) // Set up the route
    log.Println("Server is listening on port :8080")
    log.Fatal(http.ListenAndServe(":8080", nil)) // Start the server
}

