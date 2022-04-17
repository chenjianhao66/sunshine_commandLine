package model

import (
	"fmt"
	"strings"
	"time"
)

type BlogHead struct {
	Title      string
	Date       string
	Img        string
	Summary    string
	Categories string
	Tag        string
}

func (b BlogHead) String() string {
	join := strings.Join([]string{
		"---\n",
		"title: %v",
		"date: %v",
		"img: %v",
		"summary: %v",
		"categories: %v",
		"tags: \n",
	}, "")
	split := strings.Split(b.Tag, ",")
	for i := 0; i < len(split); i++ {
		sprintf := fmt.Sprintf(" - %v\n", split[i])
		join = join + sprintf
	}
	join = join + "---\n"

	return fmt.Sprintf(join, b.Title, time.Now().Format("2006-01-02 15:04:05\n"), b.Img, b.Summary, b.Categories)
}
