package config

import "fmt"

type Config struct {
	TargetHost string
	Port       string
}

type ConfigManager struct {
	Config
}

func (configManager ConfigManager) Load() {
	fmt.Println("Load config")
}

func ProvideConfigManager() ConfigManager {
	config := Config{TargetHost: "https://pokeapi.co/", Port: "8080"}
	return ConfigManager{Config: config}
}
