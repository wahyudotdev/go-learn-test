package config

import "time"

func TimeOut() time.Duration {
	return time.Second * 10
}
