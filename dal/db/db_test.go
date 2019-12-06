package db_test

import (
	"testing"

	"github.com/BlogByFourMan/Server/dal/db"
	"github.com/BlogByFourMan/Server/dal/model"
	"github.com/boltdb/bolt"
)

func TestInit(t *testing.T) {
	db.Init()
	d, err := bolt.Open(db.GetDBPATH(), 0600, nil)
	if err != nil {
		t.Error("open error:", err)
	}
	defer d.Close()
	d.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("article"))
		if b == nil {
			t.Error("bucket article doesn't exist")
		}
		b = tx.Bucket([]byte("user"))
		if b == nil {
			t.Error("bucket user doesn't exist")
		}
		return nil
	})
}

func TestPutGetArticles(t *testing.T) {
	db.Init()
	a0 := model.Artitle{0, "title0", nil, "2019-16-6", "content0", nil}
	a1 := model.Artitle{1, "title0", nil, "2020-01-03", "content1", nil}
	articles := []model.Artitle{a0, a1}
	err := db.PutArticles(articles)
	if err != nil {
		t.Error(err)
	}

	dbArticles := db.GetArticles(1)
	if len(dbArticles) != 1 {
		t.Error("len(dbArticles) != 1")
	}
	if dbArticles[0].Id != 1 {
		t.Error("dbArticles[0].Id != 1")
	}

	dbArticles = db.GetArticles(-1)
	if len(dbArticles) != 2 {
		t.Error("len(dbArticles) != 2")
	}
	if dbArticles[0].Id != 0 {
		t.Error("dbArticles[0].Id != 0")
	}
	if dbArticles[1].Id != 1 {
		t.Error("dbArticles[1].Id != 1")

	}
}

func TestGetNULLArticle(t *testing.T) {
	articles := db.GetArticles(12)
	if len(articles) != 0 {
		t.Error("len(articles) != 0")
	}
}
