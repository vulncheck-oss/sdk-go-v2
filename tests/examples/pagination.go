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
  configuration.Scheme = "https"
  configuration.Host = "api.vulncheck.com"

  client := vulncheck.NewAPIClient(configuration)

  token := os.Getenv("VULNCHECK_API_TOKEN")
  auth := context.WithValue(
    context.Background(),
    vulncheck.ContextAPIKeys,
    map[string]vulncheck.APIKey{
      "Bearer": {Key: token},
    },
  )

  var queryLimit int32 = 300
  req := client.
    IndicesAPI.
    IndexVulncheckKevGet(auth).
    StartCursor("true").
    Limit(queryLimit)

  resp, httpRes, err := req.Execute()
  if err != nil || httpRes.StatusCode != 200 {
    log.Fatal(err)
  }

  var nextCursor string
  var count int
  var total int32

  nextCursor = resp.Meta.GetNextCursor()
  total = resp.Meta.GetTotalDocuments()
  count += len(resp.Data)
  fmt.Printf("Total Items: %d/%d\n", count, total)

  for range 10 {
    req := client.
      IndicesAPI.
      IndexVulncheckKevGet(auth).
      Cursor(nextCursor).
      Limit(queryLimit)

    resp, httpRes, err := req.Execute()
    if err != nil || httpRes.StatusCode != 200 {
      log.Fatal(err)
    }

    nextCursor = resp.Meta.GetNextCursor()
    count += len(resp.Data)
    fmt.Printf("Total Items: %d/%d\n", count, total)
  }
}
