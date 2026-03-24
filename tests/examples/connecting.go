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
  resp, httpRes, err := client.EndpointsAPI.OpenapiGet(auth).Execute()

  if err != nil || httpRes.StatusCode != 200 {
    log.Fatal(err)
  }

  fmt.Printf("%+v", resp)
}
