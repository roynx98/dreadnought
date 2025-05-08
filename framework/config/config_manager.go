package config

type ConfigDTO struct {
	TargetHost string
	Port       string
	File       string
}

type ConfigManager struct {
	Config  *ConfigDTO
	loaders []ConfigLoader
}

func (configManager ConfigManager) Load() {
	for _, loader := range configManager.loaders {
		overrideConfig := loader.Load()
		mergeConfigs(configManager.Config, overrideConfig)
	}
}

func mergeConfigs(base *ConfigDTO, override ConfigDTO) {
	if override.TargetHost != "" {
		base.TargetHost = override.TargetHost
	}
	if override.Port != "" {
		base.Port = override.Port
	}
	if override.File != "" {
		base.File = override.File
	}
}

func ProvideConfigManager() ConfigManager {
	loaders := []ConfigLoader{FlagLoader{}}
	config := ConfigDTO{}
	return ConfigManager{Config: &config, loaders: loaders}
}
