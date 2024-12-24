package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/ubuntupunk/obsidian-cli/pkg/actions"
	"github.com/ubuntupunk/obsidian-cli/pkg/obsidian"
)

var searchCmd = &cobra.Command{
	Use:     "search",
	Aliases: []string{"s"},
	Short:   "Searches note in vault",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vault := obsidian.Vault{Name: vaultName}
		uri := obsidian.Uri{}
		searchText := args[0]
		params := actions.SearchParams{SearchText: searchText}
		err := actions.SearchNotes(&vault, &uri, params)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	searchCmd.Flags().StringVarP(&vaultName, "vault", "v", "", "vault name")
	rootCmd.AddCommand(searchCmd)
}
