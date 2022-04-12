package scraper

import (
	"context"
	twitterscraper "github.com/n0madic/twitter-scraper"
	"time"
)

type TweetEntry struct  {
	Text string
	User string
	Date string
}

func GetLatestTweets(query string) []TweetEntry{
	var tweets []TweetEntry
	scraper := twitterscraper.New()
	scraper.SetSearchMode(twitterscraper.SearchLatest)
	for tweet := range scraper.SearchTweets(context.Background(), query, 50) {
		if tweet.Error != nil {
			continue
		}
		tweets = append(tweets, TweetEntry{
			Text: tweet.Text,
			User: tweet.Username,
			Date: tweet.TimeParsed.Format(time.RFC3339),
		})
	}
	return tweets
}

