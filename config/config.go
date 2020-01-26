package config

import (
	"os"
)

type option struct {
	MongoURL string
	MongoDB string
	MongoUser string
	MongoPass string
}

var Option *option

func init() {

	if os.Getenv("GO_ENV") == "development" {
		Option = optionDev

		return
	}
	Option = optionPro
}
