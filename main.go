package main

// OAuth1
import (
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// Get my environment credentials
const consumerSecret = os.Getenv("API_SECRET")
const consumerSecret = os.Getenv("API_KEY")
const accessToken = os.Getenv("ACCESS_TOKEN")
const accessSecret = os.Getenv("ACCESS_SECRET")

var config = oauth1.NewConfig(consumerKey, consumerSecret)
var token = oauth1.NewToken(accessToken, accessSecret)

// http.Client will automatically authorize Requests
var httpClient = config.Client(oauth1.NoContext, token)

// twitter client
var client = twitter.NewClient(httpClient)
