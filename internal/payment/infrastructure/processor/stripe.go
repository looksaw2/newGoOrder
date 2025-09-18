package processor

import (
	"context"
	"encoding/json"

	"github.com/looksaw/go-orderv2/common/genproto/orderpb"
	"github.com/sirupsen/logrus"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/checkout/session"
)

type StripeProcessor struct {
	apiKey string
}

func NewStrpeProcessor(apiKey string) *StripeProcessor {
	if apiKey == ""{
		panic("empty api key")
	}
	logrus.Infof("api-key is %s",apiKey)
	stripe.Key = apiKey
	return &StripeProcessor{
		apiKey:apiKey,
	}
}
var(
	successURL = "http://localhost:8282"
)

func (s StripeProcessor)CreatePaymentLink(ctx context.Context,order *orderpb.Order)(string,error){
	var items []*stripe.CheckoutSessionLineItemParams
	for _ , item := range order.Items {
		items = append(items, &stripe.CheckoutSessionLineItemParams{
			Price: stripe.String("price_1S7ub20r6AiEyTXk7JXoO2Ty"),
			Quantity: stripe.Int64(int64(item.Quantity)),
		})
	}
	marshalledItems , _ := json.Marshal(order.Items)
	metadata := map[string]string {
		"orderID" : order.ID,
		"customerID" : order.CustomerID,
		"status" : order.Status,
		"items" : string(marshalledItems),
	}
	params := &stripe.CheckoutSessionParams{
		Metadata: metadata,
		LineItems: items,
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(successURL),
	}
	result ,err := session.New(params)
	if err != nil {
		return "" , nil
	}
	return result.URL , nil
}