package main 
import "fmt"

func DiscountRate(messagesSent int ) float64 {
	if messagesSent > 1000{
	return 0.1
}
	if messagesSent > 5000{
	return 0.5
}
	return 0.0
}

func calculateFinalBill(costPerMessage float64, numMessage int) float64 {
	basebill := calculateBaseBill(costPerMessage, numMessage)
	DiscountRate := DiscountRate(messagesSent)
	DiscountAmount := basebill * DiscountRate
	return  basebill - DiscountAmount
}

func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
    return costPerMessage * float64(messagesSent)
}

func main(){

}
