package adapters

import (
	"context"
	"sync"

	"github.com/looksaw/go-orderv2/common/genproto/orderpb"
	domain "github.com/looksaw/go-orderv2/stock/domain/stock"
)

type MeomoryStockRepository struct {
	lock  *sync.RWMutex
	store map[string]*order2pb.Item
}

func NewMemoryStockRepository() *MeomoryStockRepository {
	return &MeomoryStockRepository{
		lock:  &sync.RWMutex{},
		store: make(map[string]*order2pb.Item),
	}
}

var stub = map[string]*order2pb.Item{
	"item_id": {
		ID:       "foo_item",
		Name:     "stub item",
		Quantity: 10000,
		PriceID:  "stub_item_price_id",
	},
}

func (m MeomoryStockRepository) GetItems(ctx context.Context, ids []string) ([]*order2pb.Item, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	var res []*order2pb.Item
	var missing []string
	for _, id := range ids {
		if item, exist := m.store[id]; exist {
			res = append(res, item)
		} else {
			missing = append(missing, id)
		}
	}
	if len(ids) == len(res) {
		return res, nil
	}
	return res, domain.NotFoundError{Missing: missing}
}
