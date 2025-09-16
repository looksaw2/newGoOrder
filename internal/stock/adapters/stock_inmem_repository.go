package adapters

import (
	"context"
	"sync"

	"github.com/looksaw/go-orderv2/common/genproto/orderpb"
	domain "github.com/looksaw/go-orderv2/stock/domain/stock"
	"github.com/sirupsen/logrus"
)

type MeomoryStockRepository struct {
	lock  *sync.RWMutex
	store map[string]*order2pb.Item
}

func NewMemoryStockRepository() *MeomoryStockRepository {
	return &MeomoryStockRepository{
		lock:  &sync.RWMutex{},
		store: stub,
	}
}

var stub = map[string]*order2pb.Item{
	"item1": {
		ID:       "foo_item",
		Name:     "stub item",
		Quantity: 10000,
		PriceID:  "stub_item_price_id",
	},
	"item2" : {
		ID: "item2",
		Name: "stub item 2",
		Quantity: 10000,
		PriceID: "stub_item2_price_id",
	},
	"item3" : {
		ID: "item3",
		Name: "stub item 3",
		Quantity: 10000,
		PriceID: "stub_item3_price_id",
	},
	"item4" : {
		ID: "item4",
		Name: "stub item 4",
		Quantity: 10000,
		PriceID: "stub_item4_price_id",
	},
}

func (m MeomoryStockRepository) GetItems(ctx context.Context, ids []string) ([]*order2pb.Item, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	logrus.Info("start to the stock adapeters GetItems")
	var res []*order2pb.Item
	logrus.Infof("ids is +%v",ids)
	for key , _ := range m.store {
		logrus.Infof("m.store key is %s",key)
	}
	logrus.Info("m.Store is finish")
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
