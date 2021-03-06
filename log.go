package zigbee

import (
	"github.com/Traliaa/chlab/logger"
)

var (
	log = logger.GetLogger("go-openzwave") // usually replaced by driver
)

func SetLogger(logger *logger.Logger) {
	log = logger
}
