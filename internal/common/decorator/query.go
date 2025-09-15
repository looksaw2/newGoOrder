package decorator

import (
	"context"

	"github.com/sirupsen/logrus"
)

type QueryHandler[Q any, R any] interface {
	Handle(ctx context.Context, query Q) (R, error)
}

func ApplyQueryDecorators[H any, R any](handler QueryHandler[H, R], logger *logrus.Entry, matricsClient MetricsClient) QueryHandler[H, R] {
	return queryLoggingDecorator[H, R]{
		logger: logger,
		base: queryMetricsDecorator[H, R]{
			base:   handler,
			client: matricsClient,
		},
	}
}
