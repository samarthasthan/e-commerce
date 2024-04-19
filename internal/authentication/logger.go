package authentication

import (
	constants "github.com/samarthasthan/e-commerce"
	"github.com/samarthasthan/e-commerce/pkg/logger"
)

var (
	Logrusc = logger.NewLogger(constants.BROKER_AUTHENTICATION_NAME)
)
