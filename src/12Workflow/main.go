package main

import (
	"encoding/json"
	"time"
)

var inventory = map[string]float32{
	"half-siezed sprat":         12.20,
	"brass-fitted nickel slit":  3.42,
	"bracketed cap":             8.23,
	"splay-flexed brace column": 0.43,
	"damper crown":              7.28,
	"spurv plinth":              34.21,
	"husk nut":                  3.21,
	"girdle jerry":              5.21,
}

type orderRequest struct {
	customerID        string
	requestedItem     string
	requestedQuantity int
}

type invoice struct {
	customerID   string
	invoiceItems []*invoiceItem
	date         time.Time
}

type invoiceItem struct {
	invoiceID    string
	itemKey      string
	perUnitPrice float32
	unitQuantity int
}

func main() {
	parsingChan := make(chan string)
	inventoryReservationChan := make(chan *orderRequest)
	paymentRoutingChan := make(chan *invoice)
	creditCheckChan := make(chan *invoice)
	paymentProcessingChan := make(chan *invoice)
	completionChan := make(chan *invoice)
	alertChan := make(chan *invoice)

	go parsing(parsingChan, inventoryReservationChan)
	go reservingInventory(inventoryReservationChan, paymentRoutingChan)
	go paymentRouting(paymentRoutingChan, creditCheckChan, paymentProcessingChan)
	go checkingCredit(creditCheckChan, completionChan)
	go processingPayment(paymentProcessingChan, completionChan)
	go markingCompleted(completionChan, alertChan)
	go alerting(alertChan)
}

func parsing(parsingChan <-chan string, inventoryReservationChan chan<- *orderRequest) {
	for {
		req := <-parsingChan

		var order *orderRequest

		parseErr := json.Unmarshal([]byte(req), order)

		if parseErr != nil {
			//Log error
		} else {
			inventoryReservationChan <- order
		}
	}
}

func reservingInventory(reservationChan <-chan *orderRequest, paymentRoutingChan chan<- *invoice) {
	for {
		reservationReq := <-reservationChan

		//perform necessary reservation

		invoice := &invoice{
			customerID: reservationReq.customerID,
		}

		paymentRoutingChan <- invoice
	}
}

func paymentRouting(paymentRoutingChan <-chan *invoice, creditCheckChan chan<- *invoice, paymentProcessingChan chan<- *invoice) {
	for {
		invoiceNeedingPayment := <-paymentRoutingChan

		if invoiceNeedingPayment.customerID == "" {
			creditCheckChan <- invoiceNeedingPayment
		} else {
			paymentProcessingChan <- invoiceNeedingPayment
		}
	}
}

func checkingCredit(creditCheckChan <-chan *invoice, completionChan chan<- *invoice) {
	for {
		checkableInvoice := <-creditCheckChan

		// do credit checks

		completionChan <- checkableInvoice
	}
}

func processingPayment(pmtProcChan <-chan *invoice, completionChan chan<- *invoice) {}

func markingCompleted(completionChan <-chan *invoice, alertsChan chan<- *invoice) {}
func alerting(alertsChan <-chan *invoice)                                         {}
