package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/ubuntupunk/obsidian-cli/config"
	"github.com/ubuntupunk/obsidian-cli/pkg/obsidian"
	"github.com/spf13/cobra"
)

// VaultWrapper wraps the obsidian.Vault and provides additional methods
type VaultWrapper struct {
	Vault obsidian.Vault
}

// Exists checks if the vault exists in the directory structure
func (v *VaultWrapper) Exists() bool {
	path, err := v.Vault.Path()
	if err != nil {
		return false
	}
	_, err = os.Stat(path)
	return err == nil
}

var setDefaultCmd = &cobra.Command{
	Use:     "set-default",
	Aliases: []string{"sd"},
	Short:   "Sets default vault",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vaultName := args[0]

		cfg := GetConfig().(*config.Config)

		// Find the vault in config
		var targetVault *config.VaultConfig
		for _, vault := range cfg.Vaults {
			if vault.Name == vaultName {
				targetVault = &vault
				break
			}
		}

		if targetVault == nil {
			log.Fatalf("Vault '%s' not found in configuration", vaultName)
		}

		wrapper := &VaultWrapper{
			Vault: obsidian.Vault{
				Name: targetVault.Name,
			},
		}

		if !wrapper.Exists() {
			path, _ := wrapper.Vault.Path()
			log.Fatalf("Vault directory '%s' does not exist", path)
		}

		// Update default vault in config
		cfg.DefaultVault = vaultName
		if err := config.SaveConfig(); err != nil {
			log.Fatalf("Error saving config: %v", err)
		}

		fmt.Printf("Default vault set to: %s\n", vaultName)
	},
}

func init() {
	rootCmd.AddCommand(setDefaultCmd)
}
