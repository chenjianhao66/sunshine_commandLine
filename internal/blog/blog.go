package blog

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/chenjianhao66/sunshine_commandLine/internal/model"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

type blog struct{}

var (
	blogService         = blog{}
	sourceFilesLocation []string               // 博客源文件，可选多个
	targetFilesLocation = "D:\\File\\hexo博客文章" // 生成的目标博客文件夹位置
	sourceDirLocation   string                 // 博客源文件夹目录
	targetDirLocation   string                 // 博客源文件夹目录
)

func Blog() *blog {
	return &blogService
}

// Run 执行函数
func (receiver *blog) Run(cmd *cobra.Command, args []string) {
	now := time.Now()
	fmt.Println(now)
	defer func() {
		fmt.Printf("花费时间 -> %v", time.Since(now).Milliseconds())
	}()
	var (
		target *os.File
		source *os.File
		err    error
	)
	if sourceFilesLocation == nil {
		fmt.Println("source是必需选项")
		return
	}
	// 判断源文件是否都是markdown文件
	for i := 0; i < len(sourceFilesLocation); i++ {
		if err := checkFileSuffix(sourceFilesLocation[i]); err != nil {
			fmt.Println(err.Error())
			return
		}
		source, err = os.Open(sourceFilesLocation[i])
		// 判断文件是否存在
		if exist := os.IsNotExist(err); exist {
			fmt.Println(sourceFilesLocation[i], "文件不存在")
			return
		}

		// 通过反射拿到输入对象的字段数，通过for循环从控制台输入每一个字段的信息
		typeOf := reflect.TypeOf(model.BlogHead{})
		blogHead := &model.BlogHead{}
		valueOf := reflect.ValueOf(blogHead)
		reader := bufio.NewReader(os.Stdin)
		// 根据用户输入信息来生成hexo的头部信息
		for i := 0; i < typeOf.NumField(); i++ {
			// 跳过创建时间输入，让系统自动生成
			if typeOf.Field(i).Name == "Date" {
				continue
			}
			fmt.Printf("请输入 %v,回车结束输入: ", typeOf.Field(i).Name)
			input, _ := reader.ReadString('\n')
			fmt.Printf("获取到的值是 ->  %v", input)
			// 通过反射设置值
			valueOf.Elem().FieldByName(typeOf.Field(i).Name).SetString(input)
		}
		fmt.Printf("最终输入的对象 \n%+v", blogHead)
		target, err = os.OpenFile(filepath.Join(targetFilesLocation, source.Name()), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err.Error())
		}
		target.WriteString(blogHead.String())
		readAll, _ := ioutil.ReadAll(source)
		target.Write(readAll)
	}

}

// 检查文件是否是md文件
func checkFileSuffix(fileName string) error {
	if !strings.HasSuffix(fileName, ".md") {
		return errors.New(fileName + " 文件不是markdown文件")
	}
	return nil
}

func (receiver *blog) InitFlag(flags *pflag.FlagSet) {
	initSourceFlag(flags)
	initSourceDirFlag(flags)
}

// 初始化source选项
func initSourceFlag(flags *pflag.FlagSet) {
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

// 初始化source-dir选项
func initSourceDirFlag(flags *pflag.FlagSet) {
	abstract := strings.Join([]string{
		"指定要生成hexo头部消息的源博客文章文件夹地址，该文件夹里面的所有文章都将添加hexo头部消息;",
		"该选型与target选项成对出现，如果没有指定target选项的话，那么生成位置则是源博客文章地址",
	}, "\n")
	flags.StringVar(&sourceDirLocation, "source-dir", "", abstract)
}
