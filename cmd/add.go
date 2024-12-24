package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var vaultPath string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new Obsidian vault",
	Long: `Add a new Obsidian vault to the configuration.
Specify the vault name and path using the --name and --path flags.`,
	Run: func(cmd *cobra.Command, args []string) {
		if vaultName == "" || vaultPath == "" {
			fmt.Println("Please provide both vault name and path.")
			os.Exit(1)
		}

		cfg := viper.GetViper()

		var vaults []map[string]interface{}
		if cfg.IsSet("vaults") {
			vaults = cfg.Get("vaults").([]map[string]interface{})
		}

		newVault := map[string]interface{}{
			"name": vaultName,
			"path": vaultPath,
		}

		vaults = append(vaults, newVault)
		cfg.Set("vaults", vaults)

		if err := viper.WriteConfig(); err != nil {
			fmt.Printf("Error adding vault: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("Vault '%s' added successfully.\n", vaultName)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands.
	addCmd.PersistentFlags().StringVar(&vaultName, "name", "", "Vault name")
	addCmd.PersistentFlags().StringVar(&vaultPath, "path", "", "Vault path")

	// Cobra supports local flags which will only run when this command
	// is called directly.
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}