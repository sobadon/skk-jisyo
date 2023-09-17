package generate

import (
	"github.com/sobadon/skk-jisyo/cmd/syosyo/generate/nogizaka46"
	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "generate",
		Short: "generate csv file",
	}

	rootCmd.AddCommand(nogizaka46.RootCmd())

	return rootCmd
}
