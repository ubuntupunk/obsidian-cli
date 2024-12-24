package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/ubuntupunk/obsidian-cli/pkg/actions"
	"github.com/ubuntupunk/obsidian-cli/pkg/obsidian"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d"},
	Short:   "Delete note in vault",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vault := obsidian.Vault{Name: vaultName}
		note := obsidian.Note{}
		notePath := args[0]
		params := actions.DeleteParams{NotePath: notePath}
		err := actions.DeleteNote(&vault, &note, params)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	deleteCmd.Flags().BoolVarP(&shouldOpen, "open", "o", false, "open new note")
	deleteCmd.Flags().StringVarP(&vaultName, "vault", "v", "", "vault name")
	rootCmd.AddCommand(deleteCmd)
}
