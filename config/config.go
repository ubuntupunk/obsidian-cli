package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	DefaultVault string        `mapstructure:"default_vault"`
	Vaults      []VaultConfig `mapstructure:"vaults"`
}

type VaultConfig struct {
	Name     string `mapstructure:"name"`
	Path     string `mapstructure:"path"`
	IsActive bool   `mapstructure:"is_active"`
}

// InitConfig initializes the configuration system
func InitConfig() (*Config, error) {
	// Set default configuration name and type
		viper.SetConfigName("settings")
		viper.SetConfigType("json")

	// Add configuration path lookup
	configDir, err := getConfigDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get config directory: %w", err)
	}

	viper.AddConfigPath(configDir)
	viper.AddConfigPath(".")

	// Set defaults
	viper.SetDefault("default_vault", "")

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Create default config if not found
			if err := createDefaultConfig(configDir); err != nil {
				return nil, fmt.Errorf("failed to create default config: %w", err)
			}
			if err := viper.ReadInConfig(); err != nil {
				return nil, fmt.Errorf("failed to read newly created config: %w", err)
			}
		} else {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}

// getConfigDir returns the configuration directory path
func getConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configDir := filepath.Join(homeDir, ".config", "obsidian-cli")

	// Create config directory if it doesn't exist
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", err
	}

	return configDir, nil
}

// createDefaultConfig creates a default configuration file
func createDefaultConfig(configDir string) error {
	defaultConfig := Config{
		DefaultVault: "",
		Vaults:      []VaultConfig{},
	}

	viper.Set("default_vault", defaultConfig.DefaultVault)
	viper.Set("vaults", defaultConfig.Vaults)

	configFile := filepath.Join(configDir, "config.yaml")
	return viper.WriteConfigAs(configFile)
}

// SaveConfig saves the current configuration to file
func SaveConfig() error {
	return viper.WriteConfig()
}

// GetConfig returns the current configuration
func GetConfig() (*Config, error) {
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to get config: %w", err)
	}
	return &config, nil
}

// CliPath returns the path to the obsidian-cli config file
func CliPath() (string, error) {
	configDir, err := getConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "config.json"), nil
}

// ObsidianFile returns the path to the Obsidian vault configuration file
func ObsidianFile() (string, error) {
	// This assumes the obsidian config is in the vault's .obsidian directory
	// You might need to adjust this based on the user's Obsidian setup
	cfg, err := GetConfig()
	if err != nil {
		return "", err
	}
	if cfg.DefaultVault == "" {
		return "", fmt.Errorf("default vault not set")
	}
	for _, vault := range cfg.Vaults {
		if vault.Name == cfg.DefaultVault {
			return filepath.Join(vault.Path, ".obsidian", "vault.json"), nil
		}
	}
	return "", fmt.Errorf("default vault not found in config")
}
