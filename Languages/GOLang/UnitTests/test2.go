package main

import "fmt"

func monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) int {
	var lastMonthBill int
	var thisMonthBill int
	getBillForMonth(lastMonthBill, costPerSend, numLastMonth)
	getBillForMonth(thisMonthBill, costPerSend, numThisMonth)
	return thisMonthBill - lastMonthBill
}

func getBillForMonth(bill, costPerSend, messagesSent int) int {
	// returning the bill of given month 
	bill = costPerSend * messagesSent 
	return bill
}


func main(){
	costPerSend := 10
	numThisMonth :=1000
	numLastMonth := 10000
	bill:= 1000000
monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth)

	getBillForMonth(bill, costPerSend, messagesSent)

}
