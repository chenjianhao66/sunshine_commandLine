package cmd

import (
	blogService "github.com/chenjianhao66/sunshine_commandLine/internal/blog"
	"github.com/spf13/cobra"
)

var blog = &cobra.Command{
	Use:   "blog",
	Short: "博客相关命令",
	Long:  "博客相关命令",
	Run: func(cmd *cobra.Command, args []string) {
		blogService.Blog().Run(cmd, args)
	},
}

func init() {
	blogService.Blog().InitFlag(blog.Flags())
	rootCmd.AddCommand(blog)
}
