package main

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Println("Error in intializing Zap logger")
	}
	defer logger.Sync()

	logger.Info("This is an info message.")
	logger.Info("User logged in", zap.String("username", "Ahmed Hazem"), zap.String("method", "GET"))

	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05") // Custom time format
	logger2, _ := cfg.Build()
	defer logger2.Sync()

	logger2.Info("this is the new logger with new ts")

}

/*
Why Zap requires Sync() but Logrus doesn't:
Zap:

Uses asynchronous, buffered log writing internally for performance.

Logs may be kept in internal buffers before flushing to the actual output (file, stdout, etc.).

Sync() explicitly flushes these buffers to prevent log loss on shutdown.​

Logrus:

Primarily uses synchronous writes by default, writing logs immediately to the output destination.​

Since it writes directly without buffering, there's no need for an explicit flush call like Sync().

If buffering is added (e.g., via setting a custom buffered io.Writer), flushing has to be managed externally by the user.​

| Feature           | Zap                               | Logrus                      |
| ----------------- | --------------------------------- | --------------------------- |
| Writes            | Buffered, asynchronous            | Synchronous by default      |
| Requires Sync()   | Yes, to flush buffers before exit | No, writes immediately      |
| Performance focus | High performance with buffers     | Flexibility and ease of use |

*/
