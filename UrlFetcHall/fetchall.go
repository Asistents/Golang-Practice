// FetchHall perfomans a parallel URL fetc and report
// abot the time spent and size of the response for each of them

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err) // sending to channel
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // eliminate resource leakage
	if err != nil {
		ch <- fmt.Sprintf("while reading %s : %v", url, err)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		str := "http://"
		if strings.HasPrefix(url, str) == false {
			str += url
			url = str
		}
		go fetch(url, ch) // running go subroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // get from channels
	}
	fmt.Printf("%2fs elapsed\n", time.Since(start).Seconds())
}
