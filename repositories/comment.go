package repositories

import (
	"database/sql"

	"github.com/takumi-pro/go-intermediate-api/models"
)

func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
		inset into comments
		(article_id, message, created_at) values
		($1, $2, now());
	`

	newComment := models.Comment{
		CommentID: comment.CommentID,
		Message:   comment.Message,
		ArticleID: comment.ArticleID,
		CreatedAt: comment.CreatedAt,
	}

	_, err := db.Exec(sqlStr, newComment.ArticleID, newComment.Message)
	if err != nil {
		return models.Comment{}, err
	}
	return newComment, nil
}

func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		select *
		from comments
		where article_id = $1;
	`

	commentArray := make([]models.Comment, 0)
	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return commentArray, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment models.Comment
		err = rows.Scan(&comment.ArticleID, &comment.CommentID, &comment.Message, &comment.CreatedAt)
		if err != nil {
			return []models.Comment{}, err
		}

		commentArray = append(commentArray, comment)
	}

	return commentArray, nil
}
