package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	url := "https://integrations.wallib.dev/third/v1/strike/paymentmethods?phone=phone-number"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("x-api-key", "api-key")
	req.Header.Add("Authorization", "Bearer jwt-token")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
