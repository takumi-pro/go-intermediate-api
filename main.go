package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/takumi-pro/go-intermediate-api/models"
)

func main() {
	// r := mux.NewRouter()
	// r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	// r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	// r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	// r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	// r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
	// r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	// log.Println("server start at http://localhost:8080")
	// log.Fatal(http.ListenAndServe(":8080", r))

	dbUser := "takumi"
	dbPassword := "pass"
	dbDatabase := "go-api-db"
	dbConn := fmt.Sprintf("host=localhost port=5434 user=%s "+
		"password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbDatabase)
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	articleID := 100
	const sqlStr = `select * from articles where article_id = $1`

	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		return
	}

	var article models.Article
	var createdTime sql.NullTime
	err = row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		fmt.Println(err)
		return
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	fmt.Printf("%+v\n", article)
}
