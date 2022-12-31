package apps

import (
	"github.com/hyponet/outline-workflow/exporter"
	"github.com/spf13/cobra"
)

var (
	queryContent string
)

var documentCmd = &cobra.Command{
	Use:   "document",
	Short: "outline document tools",
	Long:  `integration your Outline data into a workflow`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	searchDocumentCmd.Flags().StringVar(&queryContent, "query", "", "query content")
}

var searchDocumentCmd = &cobra.Command{
	Use:   "search",
	Short: "search outline document",
	Long:  `integration your Outline data into a workflow`,
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := MustNewOutline().SearchDocuments(cmd.Context(), queryContent)
		if err != nil {
			return err
		}

		e := exporter.New()
		e.AddResource(result...)
		return e.Flush()
	},
}
