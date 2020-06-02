package main

import (
	"encoding/json"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

var configuration struct {
	ConsumerKey       string `json:"ConsumerKey"`
	ConsumerSecret    string `json:"ConsumerSecret"`
	AccessToken       string `json:"AccessToken"`
	AccessTokenSecret string `json:"AccessTokenSecret"`
}

func printErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	args := os.Args[1:]

	configFile, err := os.Open(args[0])
	printErr(err)

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&configuration)
	printErr(err)

	anaconda.SetConsumerKey(configuration.ConsumerKey)
	anaconda.SetConsumerSecret(configuration.ConsumerSecret)
	client := anaconda.NewTwitterApi(configuration.AccessToken, configuration.AccessTokenSecret)

	stream := client.PublicStreamSample(nil)

	file, _ := os.Create(args[1])
	defer file.Close()

	for {
		status := <-stream.C

		jsonTweet, err := json.Marshal(status)
		printErr(err)

		file.Write(jsonTweet)
		file.WriteString("\n")
		file.Sync()
	}
}
