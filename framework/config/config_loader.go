package config

type ConfigLoader interface {
	Load() ConfigDTO
}
