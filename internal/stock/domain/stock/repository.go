package stock

import (
	"context"
	"fmt"
	"strings"

	"github.com/looksaw/go-orderv2/common/genproto/orderpb"
)

type Repository interface {
	GetItems(ctx context.Context, ids []string) ([]*order2pb.Item, error)
}

type NotFoundError struct {
	Missing []string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("not found in stock %s", strings.Join(e.Missing, ","))
}
