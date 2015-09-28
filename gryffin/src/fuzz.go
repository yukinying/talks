package main

import "github.com/bitly/go-nsq"

func main() {
	producer, err := nsq.NewProducer("127.0.0.1:4150", nsq.NewConfig()) // HL
	// ...
	scan.CrawlAsync(r)
	go func() {
		if s := <-r.GetRequestBody(); s != nil { // using channel to get the scan result.
			err := producer.Publish("fuzz", s.Json()) // HL
			if err != nil {
				fmt.Println("Could not publish", "fuzz", err)
			}
		}
	}()
}
