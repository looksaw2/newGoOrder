package app

import "github.com/looksaw/go-orderv2/order/app/query"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
}

type Queries struct {
	GetCustomerOrder query.GetCustomerOrderHandler
}
