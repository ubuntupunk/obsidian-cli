package obsidian

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/ubuntupunk/obsidian-cli/config"
)

var JsonMarshal = json.Marshal

func (v *Vault) DefaultName() (string, error) {
	if v.Name != "" {
		return v.Name, nil
	}

	// get cliConfig path
	cliConfigFile, err := config.CliPath()
	if err != nil {
		return "", err
	}

	// read file
	content, err := os.ReadFile(cliConfigFile)
	if err != nil {
		return "", errors.New(ObsidianCLIConfigReadError)
	}

	// unmarshal json
	cliConfig := CliConfig{}
	err = json.Unmarshal(content, &cliConfig)

	if err != nil {
		return "", errors.New(ObsidianCLIConfigParseError)
	}

	if cliConfig.DefaultVaultName == "" {
		return "", errors.New(ObsidianCLIConfigParseError)
	}

	v.Name = cliConfig.DefaultVaultName
	return cliConfig.DefaultVaultName, nil
}

func (v *Vault) SetDefaultName(name string) error {
	viper.Set("default_vault", name)
	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("failed to set default vault name: %w", err)
	}
	v.Name = name
	return nil
}
