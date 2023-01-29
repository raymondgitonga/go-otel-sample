package instrumentation

//type ctxKey struct{}
//
//func LoggerNewContext(parent context.Context, z *zap.Logger) context.Context {
//	return context.WithValue(parent, ctxKey{}, z)
//}
//
//func NewLogger(level zapcore.Level, structured bool, options ...zap.Option) (*zap.Logger, error) {
//	var cfg zap.Config
//	if structured {
//		cfg = zap.Config{
//			Level:       zap.NewAtomicLevelAt(level),
//			Development: false,
//			Sampling: &zap.SamplingConfig{
//				Initial:    100,
//				Thereafter: 100,
//			},
//			Encoding:         "json",
//			EncoderConfig:    zap.NewProductionEncoderConfig(),
//			OutputPaths:      []string{"stderr"},
//			ErrorOutputPaths: []string{"stderr"},
//		}
//	} else {
//		cfg = zap.Config{
//			Level:            zap.NewAtomicLevelAt(level),
//			Development:      true,
//			Encoding:         "console",
//			EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
//			OutputPaths:      []string{"stderr"},
//			ErrorOutputPaths: []string{"stderr"},
//		}
//	}
//
//	return cfg.Build(options...)
//}
//
//func LoggerFromContext(ctx context.Context) *zap.Logger {
//	childLogger, _ := ctx.Value(ctxKey{}).(*zap.Logger)
//
//	if traceID := trace.SpanFromContext(ctx).SpanContext().TraceID(); traceID.IsValid() {
//		childLogger = childLogger.With(zap.String("trace-id", traceID.String()))
//	}
//
//	if spanID := trace.SpanFromContext(ctx).SpanContext().SpanID(); spanID.IsValid() {
//		childLogger = childLogger.With(zap.String("span-id", spanID.String()))
//	}
//
//	return childLogger
//}
