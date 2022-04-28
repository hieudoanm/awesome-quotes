package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Quote struct {
  Tags []string `json:"tags"`
  Content string `json:"content"`
  Author string `json:"author"`
  AuthorSlug string `json:"authorSlug"`
}

type ResponseBody struct {
  Results []Quote `json:"results"`
}

func main() {
  var quotes []Quote = []Quote{}
  for i := 1; i <= 13; i++ {
    fmt.Println("Page " + strconv.Itoa(i))
    var url string = fmt.Sprintf("https://api.quotable.io/quotes?page=%d&limit=150", i)
    response, httpGetError := http.Get(url)
    if httpGetError != nil {
      log.Fatalln(httpGetError.Error())
    }
    defer response.Body.Close()

    body, readBodyError := ioutil.ReadAll(response.Body)
    if readBodyError != nil {
      log.Fatalln(readBodyError.Error())
    }

    var responseBody ResponseBody
    unmarshalError := json.Unmarshal(body, &responseBody)
    if unmarshalError != nil {
      log.Fatalln(unmarshalError.Error())
    }

    var results = responseBody.Results
    quotes = append(quotes, results...)
  }

  // Update content
  var list []string
  for _, quote := range quotes {
    var item string = fmt.Sprintf(
      "- `%s` - %s",
      quote.Author,
      quote.Content,
    )
    list = append(list, item)
  }
  sort.Strings(list)

  // Write to README
  var readme string = fmt.Sprintf("# Quotes (%d)\n\n%s\n", len(quotes), strings.Join(list, "\n"))
  writeMdError := os.WriteFile("./README.md", []byte(readme), 0666);
  if writeMdError != nil {
    log.Fatal(writeMdError.Error())
  }
}
