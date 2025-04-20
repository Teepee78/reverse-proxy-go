package utils

import "teepee78/reverse-proxy-go/config"

func GetPort() int {
	if config.Vars.Port == 0 {
		return 80
	}

	return config.Vars.Port
}
