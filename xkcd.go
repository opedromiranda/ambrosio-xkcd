package ambrosio_xkcd

import (
	"github.com/opedromiranda/ambrosio"
	"encoding/json"
	"io/ioutil"
	"bytes"
	"net/http"
)

const ERROR_MESSAGE = "Couldn't retrieve XCD image"

func requestXkcd(url string) (string, error) {
	var result string
	var error error

	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, requestError := ioutil.ReadAll(resp.Body)

	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err == nil && requestError == nil {
		var buffer bytes.Buffer
		buffer.WriteString(dat["title"].(string))
		buffer.WriteString("\n")
		buffer.WriteString(dat["img"].(string))

		result = string(buffer.String())
	} else {
		result = ERROR_MESSAGE
	}

	return result, error
}

var Latest = ambrosio.Behaviour{
	"^xkcd latest$",
	func(matches []string) (string, error) {
		var result string
		var error error

		result, error = requestXkcd("http://xkcd.com/info.0.json")

		return result, error
	},
}

var Specific = ambrosio.Behaviour{
	"^xkcd ([0-9]*)$",
	func(matches []string) (string, error) {
		var result string
		var error error
		var bUrl bytes.Buffer

		xkcdNumber := matches[1]

		bUrl.WriteString("http://xkcd.com/")
		bUrl.WriteString(xkcdNumber)
		bUrl.WriteString("/info.0.json")

		result, error = requestXkcd(bUrl.String())

		return result, error
	},
}

var Behaviours = []ambrosio.Behaviour{
	Latest,
	Specific,
}
