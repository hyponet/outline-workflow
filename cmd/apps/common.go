package apps

import (
	"github.com/hyponet/outline-workflow/outline"
	"github.com/spf13/cobra"
	"time"
)

var (
	Host     string
	ApiToken string
)

var RootCmd = &cobra.Command{
	Use:   "outline-workflow",
	Short: "outline workflow tools",
	Long:  `integration your Outline data into a workflow`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	RootCmd.PersistentFlags().StringVar(&Host, "host", "", "outline host")
	RootCmd.PersistentFlags().StringVar(&ApiToken, "apiToken", "", "outline api token")

	documentCmd.AddCommand(searchDocumentCmd)

	RootCmd.AddCommand(documentCmd)
}

func MustNewOutline() *outline.Outline {
	if Host != "" {
		outline.Host = Host
		outline.ApiToken = ApiToken
	}

	o, err := outline.New(outline.ConnConfig{
		InsecureSkipVerify: true,
		Timeout:            time.Minute,
	})
	if err != nil {
		panic(err)
	}
	return o
}
