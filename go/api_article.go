/*
 * Swagger Blog
 *
 * Simple Blog
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/BlogByFourMan/Server/dal/db"
)

func ArticleIdCommentsGet(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	s := strings.Split(r.URL.Path, "/")
	articleID, err := strconv.ParseInt((s[len(s)-2]), 10, 64)
	fmt.Println(articleID)
	if err != nil {
		log.Fatal(err)
	}

	articles := db.GetArticles(articleID)
	if len(articles) != 0 {
		Response(MyResponse{
			articles[0].Comments,
			nil,
		}, w, http.StatusOK)
	}
}

func ArticleIdGet(w http.ResponseWriter, r *http.Request) {

	s := strings.Split(r.URL.Path, "/")
	articleID, err := strconv.ParseInt((s[len(s)-1]), 10, 64)
	fmt.Println(articleID)
	if err != nil {
		log.Fatal(err)
	}

	articles := db.GetArticles(articleID)
	if len(articles) != 0 {
		Response(MyResponse{
			articles[0],
			nil,
		}, w, http.StatusOK)
	}
}

func ArticlesGet(w http.ResponseWriter, r *http.Request) {
	articles := db.GetArticles(-1)

	Response(MyResponse{
		articles,
		nil,
	}, w, http.StatusOK)
}
