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

	for {
		input := "input"

		orderDetails := parsing(input)
		invoice := reservingInventory(orderDetails)
		paidInvoice := paymentRouting(invoice)
		completeInvoice := markingCompleted(paidInvoice)
		alerting(completeInvoice)

		//get next input
	}
}

func parsing(source string) *orderRequest {
	//read next input from source
	req := ""

	var order *orderRequest

	parseErr := json.Unmarshal([]byte(req), order)

	if parseErr != nil {
		//handle
	}

	return order
}

func reservingInventory(req *orderRequest) *invoice {
	//perform necessary reservation

	invoice := &invoice{
		customerID: req.customerID,
	}

	return invoice
}

func paymentRouting(invoiceNeedingPayment *invoice) *invoice {
	if invoiceNeedingPayment.customerID == "" {
		return checkingCredit(invoiceNeedingPayment)
	} else {
		return processingPayment(invoiceNeedingPayment)
	}
}

func checkingCredit(checkableInvoice *invoice) *invoice {
	// do credit checks

	return checkableInvoice
}

func processingPayment(processableInvoice *invoice) *invoice {
	//do payment processing

	return processableInvoice
}

func markingCompleted(completableInvoice *invoice) *invoice {
	//do final completions

	return completableInvoice
}
func alerting(completedInvoice *invoice) {
	//Perform necessary alerting
	log.Println(completedInvoice.customerID)
}
