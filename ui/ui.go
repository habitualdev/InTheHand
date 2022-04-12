package ui

import (
	"InTheHand/scraper"
	"encoding/json"
	g "github.com/AllenDang/giu"
	"github.com/atotto/clipboard"
)

var queryResults = ""
var query string

var wnd = g.NewMasterWindow("In The Hand", 800, 600, 0)

func getTweets() {
	queryResultsRaw, _ := json.Marshal(scraper.GetLatestTweets(query))
	queryResults = string(queryResultsRaw)
	}

func copyToClipboard() {
	err := clipboard.WriteAll(queryResults)
	if err != nil {
		println(err.Error())
	}
	}


func loop() {

	w1Layout := g.Layout{
		g.Label("Clear Skies"),
		g.Label(queryResults).Wrapped(true),
		g.InputText(&query),
		g.Button("Query").OnClick(getTweets),
		g.Button("Copy to Clipboard").OnClick(copyToClipboard),
	}


	w1 := g.SingleWindow()
	tabLayout := g.TabBar().TabItems(
		g.TabItem("Twitter Data").Layout(w1Layout),
	)

	w1.Layout(tabLayout)
}

func StartUi() {
	wnd.Run(loop)
}
