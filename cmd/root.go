package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sunshine",
	Short: "阳光工具集",
	Long:  `这是一个个人命令行工具集，目前打算将博客的文章自动添加hexo的markdown头部信息，并且发布到个人博客服务器中，编译文件并运行`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	rootCmd.Execute()
}
