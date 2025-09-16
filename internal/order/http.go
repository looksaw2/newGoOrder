package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	order2pb "github.com/looksaw/go-orderv2/common/genproto/orderpb"
	"github.com/looksaw/go-orderv2/order/app"
	"github.com/looksaw/go-orderv2/order/app/command"
	"github.com/looksaw/go-orderv2/order/app/query"
)

type HTTPServer struct{
	app app.Application
}

func(h HTTPServer)PostCustomerCustomerIdOrders(c *gin.Context, customerId string){
	var req order2pb.CreateOrderRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": err})
		return 
	}
	r ,err := h.app.Commands.CreateOrder.Handle(c,command.CreateOrder{
		CustomerID: req.CustomerID,
		Items: req.Item,
	})
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"error" : err})
		return 
	}
	c.JSON(http.StatusOK,gin.H{
		"customer_id" : req.CustomerID,
		"mesaage" : "success",
		"order_id" : r.OrderID,
	})
}


func(h HTTPServer)GetCustomerCustomerIdOrdersOrderId(c *gin.Context, customerId string, orderId string){
	o , err := h.app.Queries.GetCustomerOrder.Handle(c,query.GetCustomerOrder{
		CustomerID: customerId,
		OrderID: orderId,
	})
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"error" : err})
		return 
	}
	c.JSON(http.StatusOK,gin.H{"message" : "success","data" :  o})		
}