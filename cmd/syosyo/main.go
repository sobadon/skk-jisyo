package main

import (
	"fmt"
	"os"

	"github.com/sobadon/skk-jisyo/cmd/syosyo/convert"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "syosyo",
		Short: "generate jisyo file",
	}

	rootCmd.AddCommand(convert.RootCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}
