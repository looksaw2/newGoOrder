package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/looksaw/go-orderv2/order/app"
	"github.com/looksaw/go-orderv2/order/app/query"
)

type HTTPServer struct{
	app app.Application
}

func(h HTTPServer)PostCustomerCustomerIdOrders(c *gin.Context, customerId string){

}


func(h HTTPServer)GetCustomerCustomerIdOrdersOrderId(c *gin.Context, customerId string, orderId string){
	o , err := h.app.Queries.GetCustomerOrder.Handle(c,query.GetCustomerOrder{
		CustomerID: "fake-customer-ID",
		OrderID: "fake-ID",
	})
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"error" : err})
		return 
	}
	c.JSON(http.StatusOK,gin.H{"message" : "success","data" :  o})		
}