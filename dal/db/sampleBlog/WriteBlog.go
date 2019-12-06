package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/BlogByFourMan/Server/dal/db"

	"github.com/BlogByFourMan/Server/dal/model"
)

func main() {
	articles := make([]model.Article, 3)
	titles := []string{"TCP三次握手原理", "GitHub 标星 1.6w+，我发现了一个宝藏项目，作为编程新手有福了！", "有哪些让程序员受益终生的建议"}
	for i := 0; i < 3; i++ {
		f, err := os.Open(strconv.FormatInt(int64(i), 10) + ".html")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		content, _ := ioutil.ReadAll(f)
		a1 := model.Article{int64(i + 1), titles[i], "magic_1024", nil, "2019-10-22 09:06:54", string(content), nil}
		articles = append(articles, a1)
	}
	db.PutArticles(articles)

}
