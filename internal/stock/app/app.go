package app

import "github.com/looksaw/go-orderv2/stock/app/query"

type Application struct {
	Commands  Commands
	Queries Queries
}

type Commands struct {

}

type Queries struct {
	CheckIfItemInStock query.CheckIfItemsInStockHandler
	GetItems query.GetItemsHandler
}