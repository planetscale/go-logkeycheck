package logs

import (
	"github.com/planetscale/log"
	"go.uber.org/zap"
)

func Foo() {
	logger := log.NewPlanetScaleLogger()
	defer logger.Sync()

	logger.Info("some log message",
		log.String("userId", "12345678"), // want "log key 'userId' should be snake_case."
		zap.Int("connId", 1),             // want "log key 'connId' should be snake_case."
		log.String("branch_id", "xzy12345678"),
	)
}
