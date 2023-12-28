// Fetch displays the response to the request and the response code at the given URL

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		str := "http://"
		if strings.HasPrefix(url, str) == false {
			str += url
			url = str
		}
		fmt.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: error copying response body: %v\n", err)
			os.Exit(1)
		}
		resp, err = http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		fmt.Println("\nThe status code we got is: ", resp.StatusCode)
	}

}
