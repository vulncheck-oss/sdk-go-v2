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

  req := client.EndpointsAPI.BackupGet(auth)
  resp, httpRes, err := req.Execute()
  if err != nil || httpRes.StatusCode != 200 {
    log.Fatal(err)
  }

  for _, v := range resp.GetData() {
    fmt.Println(v.GetHref())
  }
}
