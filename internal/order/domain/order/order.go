package order

import (
	"errors"

	order2pb "github.com/looksaw/go-orderv2/common/genproto/orderpb"
)

type Order struct {
	ID          string
	CustomerID  string
	Status      string
	PaymentLink string
	Items        []*order2pb.Item
}


func NewOrder(id string, customerID string ,status string, paymentLink string , items []*order2pb.Item)(*Order ,error){
	if id == ""{
		return nil ,errors.New("empty id")
	}
	if customerID == ""{
		return nil , errors.New("empty customerID")
	}
	if status == ""{
		return nil ,errors.New("empty status")
	}
	if items == nil {
		return nil ,errors.New("empty items")
	}
	return &Order{
		ID: id,
		CustomerID: customerID,
		Status: status,
		PaymentLink: paymentLink,
		Items: items,
	},nil
}


func (o *Order)ToProto() *order2pb.Order {
	return &order2pb.Order{
		ID: o.ID,
		CustomerID: o.CustomerID,
		Status: o.Status,
		Items: o.Items,
		PaymentLink: o.PaymentLink,
	}
}