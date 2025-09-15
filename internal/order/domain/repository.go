package domain

import "context"


type Repository interface {
	Create(context.Context, *Order)(*Order,error)
	Get(ctx context.Context, id string, customerID string)(*Order,error)
	Update(
		ctx context.Context,
		o *Order,
		updateFn func(context.Context,*Order)(*Order,error),
	) error
}