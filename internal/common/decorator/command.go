package decorator

import (
	"context"

	"github.com/sirupsen/logrus"
)

type CommandHandler[C any, R any] interface {
	Handle(ctx context.Context , cmd C)(R , error)
}

func ApplyCommandDecorators[C any, R any](handler QueryHandler[C, R], logger *logrus.Entry, matricsClient MetricsClient) QueryHandler[C, R] {
	return queryLoggingDecorator[C, R]{
		logger: logger,
		base: queryMetricsDecorator[C, R]{
			base:   handler,
			client: matricsClient,
		},
	}
}