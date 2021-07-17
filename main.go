package main

// OAuth1
import (
	"os"
	"fmt"
	"reflect"
	"encoding/json"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// Get my environment credentials
var consumerSecret = os.Getenv("API_SECRET")
var consumerKey = os.Getenv("API_KEY")
var accessToken = os.Getenv("ACCESS_TOKEN")
var accessSecret = os.Getenv("ACCESS_SECRET")

var config = oauth1.NewConfig(consumerKey, consumerSecret)
var token = oauth1.NewToken(accessToken, accessSecret)

// http.Client will automatically authorize Requests
var httpClient = config.Client(oauth1.NoContext, token)

// twitter client
var client = twitter.NewClient(httpClient)



func main()  {
	followers, _, _ := client.Followers.List(&twitter.FollowerListParams{})
	fmt.Println(reflect.TypeOf(followers))
	users, _ := json.Marshal(followers)
	fmt.Println(string(users))
}