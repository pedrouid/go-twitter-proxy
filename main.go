package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/gorilla/mux"
)

const (
	consumerKey       = "hFxyPg1JEue2785ZTAmZtnHKT"
	consumerSecret    = "scCFcETWKhKqSacEOfgSPVraSl34lg8DVFJeEm5Ly5vexsyFja"
	accessToken       = "823546213-UPV9Qcv3UNY5cuTKArDfMk8vY7ziWWZvYmZpWliC"
	accessTokenSecret = "fGPYdHDYBLVqsMrh4yP4uCZ4WqOjNG42631y88ZRW79WX"
)

var api *anaconda.TwitterApi

func initAPI() {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)

	_api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	api = _api
}

func getHelp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Twitter Proxy\n\n1. Gets tweets by screen_name\n/tweets/{screen_name}\n\n2. Gets top 10 tweets by screen_name\n/tweets/{screen_name}/top-10\n")
}

func getTweets(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	screenName := vars["screen_name"]

	v := url.Values{}
	v.Set("screen_name", screenName)

	timeline, _ := api.GetUserTimeline(v)

	json.NewEncoder(w).Encode(timeline)

}

func router() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", getHelp)
	myRouter.HandleFunc("/help", getHelp)
	myRouter.HandleFunc("/tweets/{screen_name}", getTweets)
	// myRouter.HandleFunc("/tweets/{screen_name}/top-10", getTop10Tweets)
	log.Fatal(http.ListenAndServe(":5001", myRouter))
}

func main() {
	initAPI()
	router()
}
