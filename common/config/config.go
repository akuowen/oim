package config

import (
	"time"

	"github.com/spf13/viper"
)

const DISCOVERY string = "discovery.endpoints"
const TIMEOUT string = "discovery.timeout"
const SERVICE_PATH string = "ip_conf.service_path"
const ENV string = "global.env"
const DEBUG string = "debug"

func GetGatewayTCPServerPort() int {
	return 8900
}

func InitGateway(path string) {}

func InitConfig(path string) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func GetEndpointsForDiscovery() []string {
	return viper.GetStringSlice(DISCOVERY)
}

func GetServicePathForIPConf() string {
	return viper.GetString(SERVICE_PATH)
}

func GetTimeoutForDiscovery() time.Duration {
	return viper.GetDuration(TIMEOUT) * time.Second
}

func IsDebug() bool {
	return getEnv() == DEBUG
}

func getEnv() string {
	return viper.GetString(ENV)
}

func Init(path string) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
