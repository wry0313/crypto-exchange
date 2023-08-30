package main

import (
	"log"
	"net/http"
)


func main() {
	// Kafka configuration
	brokerList := []string{"localhost:9092"}
	producer, err := newProducer(brokerList)
	if err != nil {
		log.Fatalf("Could not create producer: %v", err)
	}

	ex := NewExchange()

	go ex.consumeAndPlaceOrders(brokerList) // consume and place orders to the orderbook
	go ex.FetchAndStoreMarketPrice()
	go ex.FetchAndStoreBestBids()
	go ex.FetchAndStoreBestAsks()

	http.HandleFunc("/order", ex.PlaceOrderHandler(producer))
	http.HandleFunc("/price", ex.StreamMarketPrice)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
