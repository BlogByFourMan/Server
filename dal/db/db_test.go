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

func TestPutUsers(t *testing.T) {
	db.Init()
	u0 := model.User{"testUser1", "123456"}
	u1 := model.User{"testUser2", "123456"}

	users := []model.User{u0, u1}

	err := db.PutUsers(users)
	if err != nil {
		t.Error(err)
	}

	u2 := db.GetUser("testUser1")

	if u2.Password != u0.Password {
		t.Error("data error")
	}

	u3 := db.GetUser("testUser2")

	if u3.Password != u1.Password {
		t.Error("data error")
	}
}

func TestPutGetArticles(t *testing.T) {
	db.Init()
	a0 := model.Article{0, "title0", "", nil, "2019-16-6", "content0", nil}
	a1 := model.Article{1, "title0", "", nil, "2020-01-03", "content1", nil}
	articles := []model.Article{a0, a1}
	err := db.PutArticles(articles)
	if err != nil {
		t.Error(err)
	}

	dbArticles := db.GetArticles(1, 0)
	if len(dbArticles) != 1 {
		t.Error("len(dbArticles) != 1")
	}
	if dbArticles[0].Id != 1 {
		t.Error("dbArticles[0].Id != 1")
	}

	dbArticles = db.GetArticles(-1, 2)
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
	articles := db.GetArticles(12, 0)
	if len(articles) != 0 {
		t.Error("len(articles) != 0")
	}
}
