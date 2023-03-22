package main

import (
	"go.uber.org/zap"
)

func main() {
	logger := zap.NewExample()
	defer logger.Sync()

	if ce := logger.Check(zap.DebugLevel, "debugging"); ce != nil {
		// If debug-level log output isn't enabled or if zap's sampling would have
		// dropped this log entry, we don't allocate the slice that holds these
		// fields.
		ce.Write(
			zap.String("foo", "bar"),
			zap.String("baz", "quux"),
		)
	}

}
