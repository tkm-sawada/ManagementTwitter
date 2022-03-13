package tweetapi

import (
	"ManagementTwitter/config"
	"fmt"
	"net/url"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
)

type APIAccount struct {
	key         string
	secret      string
	token       string
	tokenSecret string
}

type Tweet struct {
	User       string `json:"user"`
	Text       string `json:"text"`
	ScreenName string `json:"screenName"`
	Id         string `json:"id"`
	Date       string `json:"date"`
	TweetId    string `json:"tweetId"`
}

type Trend struct {
	No              string
	Name            string `json:"name"`
	Query           string `json:"query"`
	Url             string `json:"url"`
	PromotedContent string `json:"promoted_content"`
}

func New(key, secret, token, tokenSecret string) *APIAccount {
	apiAccount := &APIAccount{key, secret, token, tokenSecret}
	return apiAccount
}

func connectTwitterAPI() *anaconda.TwitterApi {
	api := anaconda.NewTwitterApiWithCredentials(
		config.Config.AccessToken,
		config.Config.AccessTokenSecret,
		config.Config.ApiKey,
		config.Config.ApiSecret)
	return api
}

//ツイート
func PostTweet(word string) *Tweet {
	api := connectTwitterAPI()
	// 検索
	v := url.Values{}
	v.Set("count", "20")
	//searchResult, err := api.GetSearch(word+" exclude:retweets filter:images", v)
	result, err := api.PostTweet(word, v)
	if err != nil {
		fmt.Println("エラー発生！")
		fmt.Println(err.Error())
	}

	tweet := new(Tweet)
	tweet.Text = result.FullText
	tweet.User = result.User.Name
	tweet.Id = result.User.IdStr
	tweet.ScreenName = result.User.ScreenName
	tweet.Date = result.CreatedAt
	tweet.TweetId = result.IdStr

	return tweet
}

//ツイート検索
func GetTweetSearch(word string) []*Tweet {
	api := connectTwitterAPI()
	// 検索
	v := url.Values{}
	v.Set("count", "20")
	//searchResult, err := api.GetSearch(word+" exclude:retweets filter:images", v)
	result, err := api.GetSearch(word, v)
	if err != nil {
		fmt.Println("エラー発生！")
		fmt.Println(err.Error())
	}

	//var text []string
	tweets := make([]*Tweet, 0)
	for _, tweetdata := range result.Statuses {
		//text += tweet.Text
		//text = append(text, "["+strconv.Itoa(i)+"]"+tweet.Text+"\n")
		tweet := new(Tweet)
		tweet.Text = tweetdata.FullText
		tweet.User = tweetdata.User.Name
		tweet.Id = tweetdata.User.IdStr
		tweet.ScreenName = tweetdata.User.ScreenName
		tweet.Date = tweetdata.CreatedAt
		tweet.TweetId = tweetdata.IdStr
		tweets = append(tweets, tweet)
		//break
	}
	return tweets
}

//トレンド取得
func GetTweetTrend() []*Trend {
	api := connectTwitterAPI()
	// 検索
	// v := url.Values{}
	// v.Set("exclude", "hashtags") //ハッシュタグを除く
	result, err := api.GetTrendsByPlace(23424856, nil) //日本のトレンド
	if err != nil {
		fmt.Println("エラー発生！")
		fmt.Println(err.Error())
	}
	trends := make([]*Trend, 0)
	for i, trenddata := range result.Trends {
		trend := new(Trend)
		trend.No = strconv.Itoa(i + 1)
		trend.Name = trenddata.Name
		trend.Query = trenddata.Query
		trend.Url = trenddata.Url
		trend.PromotedContent = trenddata.PromotedContent
		trends = append(trends, trend)
	}
	return trends
}

// func PostTweet(text string) anaconda.Tweet {
// 	api := connectTwitterAPI()
// 	tweet, err := api.PostTweet(text, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return tweet
// }
