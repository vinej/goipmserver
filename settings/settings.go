package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	//"flag"
)

var environments = map[string]string{
	"production":    "settings/prod.json",
	"preproduction": "settings/pre.json",
	"tests":         "../../settings/tests.json",
}

const keyMongoHost = "MONGOHOST"
const keyMongoDatabase = "MONGODB"

type Settings struct {
	PrivateKeyPath     string
	PublicKeyPath      string
	JWTExpirationDelta int
	MongoHost string
	MongoDatabase string
}

var settings Settings
var env = "preproduction"

func init() {
	ReInit()
}

func ReInit() {
	env = os.Getenv("GO_ENV")
	if env == "" {
		fmt.Println("Warning: Setting preproduction environment due to lack of GO_ENV value")
		env = "preproduction"
	}
	LoadSettingsByEnv(env)
}

func LoadSettingsByEnv(env string) {
	content, err := ioutil.ReadFile(environments[env])
	if err != nil {
		fmt.Println("Error while reading config file", err)
	}
	settings = Settings{}
	jsonErr := json.Unmarshal(content, &settings)
	if jsonErr != nil {
		fmt.Println("Error while parsing config file", jsonErr)
	}
	settings.MongoHost = getEnv("localhost", keyMongoHost)
	settings.MongoDatabase = getEnv("restdb", keyMongoDatabase)
}

func getEnv(defaultVal string, key string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	} else {
		fmt.Printf("Warning: Environment variable %s is not set, using default value: %s\n", key, defaultVal)
		return defaultVal
	}
}

func Get() Settings {
	return settings
}

func IsTestEnvironment() bool {
	return env == "tests"
}
