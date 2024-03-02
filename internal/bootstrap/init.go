package bootstrap

import (
	"log"
	"sre-tool/internal/config"
	"sre-tool/internal/config/driver"
	"sre-tool/internal/constants"
	"sre-tool/internal/logger"
)

func init() {
	var err error
	if constants.Config, err = config.New(driver.New(), config.Options{
		BasePath: constants.BasePath,
	}); err != nil {
		log.Fatal(constants.ErrorInitConfig)
	}
	if constants.Log, err = logger.New(
		logger.WithDebug(true),
		logger.WithEncode("json"),
		logger.WithFilename(constants.BasePath+"/storage/logs/system.log"),
	); err != nil {
		log.Fatal(constants.ErrorInitLogger)
	}
}
