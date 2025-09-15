package order

import order2pb "github.com/looksaw/go-orderv2/common/genproto/orderpb"

type Order struct {
	ID          string
	CustomerID  string
	Status      string
	PaymentLink string
	Item        []*order2pb.Item
}
