package repositories

import (
	"database/sql"
	"fmt"

	"github.com/takumi-pro/go-intermediate-api/models"
)

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
		insert into articles (title, contents, username, nice, created_at) values
		($1, $2, $3, 0, now());
	`

	newArticle := models.Article{
		Title:    article.Title,
		Contents: article.Contents,
		UserName: article.UserName,
	}

	_, err := db.Exec(sqlStr, newArticle.Title, newArticle.Contents, newArticle.UserName)
	if err != nil {
		fmt.Println(err)
		return models.Article{}, err
	}

	return newArticle, nil
}

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `select article_id, title, contents, username, nice from articles limit $1 offset $2;`

	articleArray := make([]models.Article, 0)
	rows, err := db.Query(sqlStr, 10, 0)
	if err != nil {
		return articleArray, err
	}
	defer rows.Close()

	for rows.Next() {
		var article models.Article
		err = rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &article.CreatedAt)
		if err != nil {
			fmt.Println(err)
			return []models.Article{}, err
		}

		articleArray = append(articleArray, article)
	}

	return articleArray, nil
}

func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `select * from articles where article_id = $1;`

	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		return models.Article{}, err
	}

	var article models.Article
	err := row.Scan(&articleID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &article.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return models.Article{}, err
	}

	return article, nil
}

func UpdateNiceNum(db *sql.DB, articleID int) error {
	const sqlGetNice = `
		select nice
		from articles
		where aricle_id = $1;
	`

	const sqlUpdateNice = `
		update articles
		set nice = $1
		where article_id = $2;
	`

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	row := tx.QueryRow(sqlGetNice, articleID)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return err
	}

	// 現在のいいね数を読み込む
	var nicenum int
	err = row.Scan(&nicenum)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(sqlUpdateNice, nicenum+1, articleID)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
