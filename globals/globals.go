package globals

import (
	"os"
)

var (
	SERVER_HOST string
	SERVER_PORT string
	TOKEN       string
	MONGO_PASS  string
)

func InitConfig() {
	TOKEN = os.Getenv("CALIBOT_TOKEN")
	MONGO_PASS = os.Getenv("MONGO_PASS")
	SERVER_PORT = "58839"
	SERVER_HOST = "mongodb://mongo:" + MONGO_PASS + "@viaduct.proxy.rlwy.net:" + SERVER_PORT + "/?tlsCertificateKeyFilePassword=" + MONGO_PASS
}
