package main

import (
	"github.com/gin-gonic/gin"
	"github.com/looksaw/go-orderv2/order/app"
)

type HTTPServer struct{
	app app.Application
}

func(s HTTPServer)PostCustomerCustomerIdOrders(c *gin.Context, customerId string){

}


func(s HTTPServer)GetCustomerCustomerIdOrdersOrderId(c *gin.Context, customerId string, orderId string){
	
}