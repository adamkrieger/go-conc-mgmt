package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"
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
	//Integrate with surrounding operating system
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		_ = <-sigs
		//initiate any cleanup here
		done <- true
	}()

	//Set up workflow channels (queues)
	inventoryReservationChan := make(chan *orderRequest)
	paymentRoutingChan := make(chan *invoice)
	creditCheckChan := make(chan *invoice)
	paymentProcessingChan := make(chan *invoice)
	completionChan := make(chan *invoice)
	alertChan := make(chan *invoice)

	//Launch go routines to do work
	go parsing(inventoryReservationChan)
	go reservingInventory(inventoryReservationChan, paymentRoutingChan)
	go paymentRouting(paymentRoutingChan, creditCheckChan, paymentProcessingChan)
	go checkingCredit(creditCheckChan, completionChan)
	go processingPayment(paymentProcessingChan, completionChan)
	go markingCompleted(completionChan, alertChan)
	go alerting(alertChan)

	<-done
}

func parsing(inventoryReservationChan chan<- *orderRequest) {
	for {
		//read next input from source
		req := ""

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

func processingPayment(pmtProcChan <-chan *invoice, completionChan chan<- *invoice) {
	for {
		processableInvoice := <-pmtProcChan

		//do payment processing

		completionChan <- processableInvoice
	}
}

func markingCompleted(completionChan <-chan *invoice, alertsChan chan<- *invoice) {
	for {
		completableInvoice := <-completionChan

		//do final completions

		alertsChan <- completableInvoice
	}
}
func alerting(alertsChan <-chan *invoice) {
	for {
		completedInvoice := <-alertsChan

		//Perform necessary alerting
		log.Println(completedInvoice.customerID)
	}
}
