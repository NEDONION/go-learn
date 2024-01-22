package model

import "fmt"

type PaymentArgs struct {
	AppID       string
	MerchantID  string
	Key         string
	CallBackUrl string
}

func (paymentArgs *PaymentArgs) Info() {
	fmt.Printf("Info = %v\n", paymentArgs)
}
