//go:build ignore

package main

import (
  "context"
  "fmt"
  "log"
  "os"

  vulncheck "github.com/vulncheck-oss/sdk-go-v2/v2"
)

func main() {
  configuration := vulncheck.NewConfiguration()
  client := vulncheck.NewAPIClient(configuration)

  token := os.Getenv("VULNCHECK_API_TOKEN")
  auth := context.WithValue(
    context.Background(),
    vulncheck.ContextAPIKeys,
    map[string]vulncheck.APIKey{
      "Bearer": {Key: token},
    },
  )

  // List all available advisory feeds (/v4/advisory)
  feedsResp, httpRes, err := client.AdvisoryAPI.V4ListAdvisoryFeeds(auth).Execute()
  if err != nil || httpRes.StatusCode != 200 {
    log.Fatal(err)
  }

  fmt.Println("Available feeds:")
  for _, feed := range feedsResp.GetData() {
    fmt.Println(feed.GetName())
  }

  // Query advisories filtered by feed=wolfi (/v4/advisory?feed=wolfi)
  feed := "wolfi"
  advisoriesResp, httpRes, err := client.AdvisoryAPI.V4QueryAdvisories(auth).Name(feed).Execute()
  if err != nil || httpRes.StatusCode != 200 {
    log.Fatal(err)
  }

  fmt.Printf("%s advisories (page 1): %d results\n", feed, len(advisoriesResp.GetData()))
  for _, advisory := range advisoriesResp.GetData() {
    meta := advisory.GetCveMetadata()
    fmt.Println(meta.GetCveId())
  }
}
