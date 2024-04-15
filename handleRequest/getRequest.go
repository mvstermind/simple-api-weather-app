package handlerequest

import (
	"fmt"
	"io"
	"net/http"
)

func ApiRequest(url string) string {

	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return ""
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return ""
	}

	return string(body)
}
