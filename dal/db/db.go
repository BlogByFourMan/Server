package db

import (
	"encoding/binary"
	"encoding/json"
	"log"
	"os"
	"path"

	"github.com/BlogByFourMan/Server/dal/model"

	"github.com/boltdb/bolt"
)

func GetDBPATH() string {
	return path.Join(os.Getenv("GOPATH"), "src", "github.com", "BlogByFourMan", "Server", "dal", "db", "Blog.db")
}
func Init() {
	db, err := bolt.Open(GetDBPATH(), 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("article"))
		if b == nil {
			_, err := tx.CreateBucket([]byte("article"))
			if err != nil {
				log.Fatal(err)
			}
		}

		b = tx.Bucket([]byte("user"))
		if b == nil {
			_, err := tx.CreateBucket([]byte("user"))
			if err != nil {
				log.Fatal(err)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

}

// PutArticles : put articles to article of blog.db
//
func PutArticles(articles []model.Artitle) error {
	db, err := bolt.Open(GetDBPATH(), 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("article"))
		if b != nil {
			for i := 0; i < len(articles); i++ {
				id := articles[i].Id
				key := make([]byte, 8)
				binary.LittleEndian.PutUint64(key, uint64(id))
				data, _ := json.Marshal(articles[i])
				b.Put(key, data)
			}
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// GetArticles 根据article_id 获取article
// 如果id == -1 表示获取所有articles
// return []Article. if not found, len(articles)==0
func GetArticles(id int64) []model.Artitle {
	articles := make([]model.Artitle, 0)
	db, err := bolt.Open(GetDBPATH(), 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("article"))
		if b != nil && id >= 0 {

			key := make([]byte, 8)
			binary.LittleEndian.PutUint64(key, uint64(id))
			data := b.Get(key)
			if data != nil {

				atc := model.Artitle{}
				err := json.Unmarshal(data, &atc)
				if err != nil {
					log.Fatal(err)
				}
				articles = append(articles, atc)
			}

		} else if b != nil && id == -1 {
			cursor := b.Cursor()
			for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
				atc := model.Artitle{}
				err := json.Unmarshal(v, &atc)
				if err != nil {
					log.Fatal(err)
				}
				articles = append(articles, atc)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	return articles
}
