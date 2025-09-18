package app

import "github.com/looksaw/go-orderv2/payment/app/command"


type Application struct {
	Command Command
}

type Command struct {
	CreatePayment command.CreatePaymentHandler
}
