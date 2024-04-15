package handlerequest

import (
	"fmt"
	"io"
	"net/http"
)

func ApiRequest(url string) string {

	res, err := http.Get(url) // get response
	if err != nil {
		fmt.Printf("Error: %v", err)
		return ""
	}
	defer res.Body.Close() // close after reading all content of request

	body, err := io.ReadAll(res.Body) // read all json content
	if err != nil {
		fmt.Printf("Error: %v", err)
		return ""
	}

	return string(body) // return json as string isntead of slice
}
