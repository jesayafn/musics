package configs

import (
	"os"
	"strings"
)

func DatabaseConf() (fullUriDb string, authSourceDb string, usernameDb string, passwordUserDb string) {
	uriDb := os.Getenv("MONGO_URI")
	portDb := os.Getenv("MONGO_PORT")
	authSourceDb = os.Getenv("MONGO_AUTH_DB")
	usernameDb = os.Getenv("MONGO_AUTH_USER")
	passwordUserDb = os.Getenv("MONGO_AUTH_PASS")
	fullUriSlice := []string{"mongodb://", uriDb, ":", portDb}
	fullUriDb = strings.Join(fullUriSlice, "")
	return fullUriDb, authSourceDb, usernameDb, passwordUserDb
}

func RouterConf() (mode string, env string) {
	mode = os.Getenv("MODE")
	env = os.Getenv("ENV_DEPLOY")
	return mode, env
}
