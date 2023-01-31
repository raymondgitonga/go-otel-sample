package instrumentation

import (
	"context"
	"go.uber.org/zap"
)

type LoggerWithCtx struct {
	*zap.Logger
	context *context.Context
}
