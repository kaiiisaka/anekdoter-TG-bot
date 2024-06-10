package source

import (
	"context"
	"github.com/samber/lo"
	"log"

	"github.com/SlyMarbo/rss"
	"snakefishAnekdot/internal/model"
)

type RSSSource struct {
	SourceName string
	URL        string
	SourceID   int64
}

func NewRSSSourceFromModel(m model.Source) RSSSource {
	return RSSSource{
		SourceName: m.Name,
		URL:        m.FeedURL,
		SourceID:   m.ID,
	}
}

func (s RSSSource) Fetch(ctx context.Context) ([]model.Item, error) {
	feed, err := s.loadFeed(ctx, s.URL)
	if err != nil {
		return nil, err
	}

	return lo.Map(feed.Items, func(item *rss.Item, _ int) model.Item {
		return model.Item{
			Title:      item.Title,
			Link:       item.Link,
			Categories: item.Categories,
			Date:       item.Date,
			Summary:    item.Summary,
			SourceName: s.SourceName,
		}
	}), nil
}

func (s RSSSource) ID() int64 {
	return s.SourceID
}

func (s RSSSource) Name() string {
	return s.SourceName
}

func (s RSSSource) loadFeed(ctx context.Context, url string) (*rss.Feed, error) {
	var (
		feedCh = make(chan *rss.Feed)
		errCh  = make(chan error)
	)

	go func() {
		feed, err := rss.Fetch(url)
		if err != nil {
			errCh <- err
			return
		}

		log.Printf("[INFO] fetched %d items from %s", len(feed.Items), url)

		feedCh <- feed
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case err := <-errCh:
		return nil, err
	case feed := <-feedCh:
		return feed, nil
	}
}
