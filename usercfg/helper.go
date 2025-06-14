package usercfg

// ensureConfigLoaded checks if the global LoadedConfig is nil.
// If it is, it initializes the configuration by calling InitConfig.
// This function ensures that configuration is loaded before proceeding.
func ensureConfigLoaded() {
	if LoadedConfig == nil {
		InitConfig()
	}
}
