package main

import (
	"fmt"
	"go.corp.yahoo.com/albertyu/html-distance"
	"net/http"
)

func compute() {
	url1 := "https://www.flickr.com/photos/flickr/20350469606/in/album-72157639868074114/"
	url2 := "https://www.flickr.com/photos/flickr/14145682444/in/album-72157639868074114/"
	s := 2 // HL

	resp1, _ := http.Get(url1)
	fmt.Printf("Fetching %s, Got %d\n", url1, resp1.StatusCode)
	resp2, _ := http.Get(url2)
	fmt.Printf("Fetching %s, Got %d\n", url2, resp2.StatusCode)

	f1 := distance.Fingerprint(resp1.Body, s) // HL
	f2 := distance.Fingerprint(resp2.Body, s) // HL
	d := distance.Distance(f1, f2)            // HL

	fmt.Printf("Fingerprint1 %064b\n", f1)
	fmt.Printf("Fingerprint2 %064b\n", f2)
	fmt.Printf("Feature distance is %d.\nShingle factor is %d.\nHTML Similarity is %3.2f%%\n", d, s, (1-float32(d)/64)*100)
}

func main() {
	compute()
}