package models

import "time"

var (
	Comment1 = Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "nice article",
		CreatedAt: time.Now(),
	}

	Comment2 = Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "beautiful photo!",
		CreatedAt: time.Now(),
	}
)

var (
	Article1 = Article{
		ID:          1,
		Title:       "my photo",
		Contents:    "this is my photo",
		UserName:    "takum",
		NiceNum:     23,
		CommentList: []Comment{Comment1, Comment2},
		CreatedAt:   time.Now(),
	}

	Article2 = Article{
		ID:        2,
		Title:     "exam result",
		Contents:  "0 point",
		UserName:  "takum",
		NiceNum:   100,
		CreatedAt: time.Now(),
	}
)
