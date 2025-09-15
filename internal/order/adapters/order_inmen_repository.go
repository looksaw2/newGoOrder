package adapters

import (
	"context"
	"strconv"
	"sync"
	"time"

	domain "github.com/looksaw/go-orderv2/order/domain/order"
	"github.com/sirupsen/logrus"
)

type MemoryOrderRepository struct {
	lock  *sync.RWMutex
	store []*domain.Order
}

func NewMemoryOrderRepository() *MemoryOrderRepository {
	return &MemoryOrderRepository{
		lock:  &sync.RWMutex{},
		store: make([]*domain.Order, 0),
	}
}

func (m MemoryOrderRepository) Create(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	newOrder := &domain.Order{
		ID:          strconv.FormatInt(time.Now().Unix(), 10),
		CustomerID:  order.CustomerID,
		Status:      order.Status,
		PaymentLink: order.PaymentLink,
		Item:        order.Item,
	}
	m.store = append(m.store, newOrder)
	logrus.WithFields(logrus.Fields{
		"input_order":        order,
		"store_after_create": m.store,
	}).Debug("memory_order_repo_create")
	return newOrder, nil
}

func (m MemoryOrderRepository) Get(ctx context.Context, id string, customerID string) (*domain.Order, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	for _, o := range m.store {
		if o.ID == id && o.CustomerID == customerID {
			logrus.Debugf("memory_order_get id = %s , customerID = %s ,res = %s", id, customerID, o)
			return o, nil
		}
	}
	return nil, &domain.NotFoundError{OrderID: id}

}

func (m MemoryOrderRepository) Update(ctx context.Context, order *domain.Order, updateFn func(context.Context, *domain.Order) (*domain.Order, error)) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	found := false
	for i, o := range m.store {
		if o.ID == order.ID && o.CustomerID == order.CustomerID {
			found = true
			updateOrder, err := updateFn(ctx, order)
			if err != nil {
				return err
			}
			m.store[i] = updateOrder
		}
	}
	if !found {
		return &domain.NotFoundError{OrderID: order.ID}
	}
	return nil
}
