/***
When fetching JSON over a network with the header, 
"Accept-Encoding":"gzip, br" ... etc.
Because gzip is selected first the JSON data is often compressed into gzip.
The following code snippet allows decompression for further usage,	

Original Link:
https://stackoverflow.com/questions/64460591/unmarshalling-json-returns-an-error-invalid-character-x1f-looking-for-beginn
***/
    // Check if the response is encoded in gzip format
    if resp.Header.Get("Content-Encoding") == "gzip" {
        reader, err := gzip.NewReader(resp.Body)
        if err != nil {
            panic(err)
        }
        defer reader.Close()
        // Read the decompressed response body
        body, err := io.ReadAll(reader)
        if err != nil {
            panic(err)
        }
        // Do something with the response body
        fmt.Println(string(body))
    } else {
        // The response is not gzip encoded, so read it directly
        body, err := io.ReadAll(resp.Body)
        if err != nil {
            panic(err)
        }
        // Do something with the response body
        fmt.Println(string(body))
    }
}
	


