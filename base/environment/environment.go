package environment

import (
	"os"
)

func GetEnv() string {
	env, isEnv := os.LookupEnv("environment_tcp")
	if !isEnv {
		env = "tcp"
	}
	return env
}

func GetPort() string {
	port, isPort := os.LookupEnv("port_tcp")
	if !isPort {
		port = ":50051"
	}
	return port
}

func GetRestEnv() string {
	env, isEnv := os.LookupEnv("environment_rest")
	if !isEnv {
		env = "localhost"
	}
	return env
}

func GetRestPort() string {
	port, isPort := os.LookupEnv("port_rest")
	if !isPort {
		port = ":5051"
	}
	return port
}
