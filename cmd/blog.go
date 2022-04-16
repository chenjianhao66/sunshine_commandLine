package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var (
	flags               = blog.Flags()
	sourceFilesLocation []string // 博客源文件，可选多个
	targetFilesLocation []string // 生成的目标博客文件，可选多个
	sourceDirLocation   string   // 博客源文件夹目录
	targetDirLocation   string   // 博客源文件夹目录
)

var blog = &cobra.Command{
	Use:   "blog",
	Short: "博客相关命令",
	Long:  "博客相关命令",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("执行", len(args))
		fmt.Println(sourceFilesLocation)
		fmt.Println(targetFilesLocation)
		fmt.Println(sourceDirLocation)
		fmt.Println(targetDirLocation)
	},
}

func init() {
	initSourceFlag()
	initTargetFlag()
	initSourceDirFlag()
	initTargetDirFlag()
	rootCmd.AddCommand(blog)
}

func initSourceFlag() {
	abstract := strings.Join([]string{
		"指定要生成hexo头部消息的源博客文章地址",
		"该选型与target选项成对出现，如果没有指定target选项的话，那么生成位置则是源博客文章地址;",
		"该选项可以指定多个源文件，用英文逗号隔开;",
		"示例：sunshine blog -s v1.md ",
		"sunshine blog -s v1.md,v2.md",
		"sunshine blog -s .  //代表当下目录所有markdown文件都将生成hexo头部消息",
	}, "\n")
	flags.StringSliceVarP(&sourceFilesLocation, "source", "s", []string{}, abstract)
}

func initSourceDirFlag() {

	abstract := strings.Join([]string{
		"指定要生成hexo头部消息的源博客文章文件夹地址，该文件夹里面的所有文章都将添加hexo头部消息;",
		"该选型与target选项成对出现，如果没有指定target选项的话，那么生成位置则是源博客文章地址",
	}, "\n")
	flags.StringVar(&sourceDirLocation, "source-dir", "", abstract)
}

func initTargetFlag() {
	abstract := `生成hexo头部消息后的博客文章地址`
	flags.StringSliceVarP(&targetFilesLocation, "target", "t", []string{}, abstract)
}

func initTargetDirFlag() {
	abstract := `生成hexo头部消息后的博客文章地址文件夹地址，该文件夹里面的所有文章都将添加hexo头部消息；`
	flags.StringVar(&targetDirLocation, "target-dir", "", abstract)
}
