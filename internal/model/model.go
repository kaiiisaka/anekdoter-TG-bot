package model

import "time"

type Item struct {
	Title      string
	Categories []string
	Link       string
	Date       time.Time
	Summary    string
	SourceName string
}

type Source struct {
	Name      string
	FeedURL   string
	ID        int64
	CreatedAt time.Time
}

type Article struct {
	Title       string
	Link        string
	ID          int64
	SourceId    int64
	PublishedAt time.Time
	PostedAt    time.Time
	CreatedAt   time.Time
	Summary     string
}
