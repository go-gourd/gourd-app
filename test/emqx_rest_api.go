package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	username := "ee26b0dd4af7e749"
	password := "aO9BSnoR9CyXFbOy03GWl0IVRE1uurr9Cu1tPaoNKqUFHK"

	url := "http://127.0.0.1:18083/api/v5/nodes"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(buf.String())
}
