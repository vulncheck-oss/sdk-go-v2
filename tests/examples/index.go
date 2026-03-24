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

  req := client.
    IndicesAPI.
    IndexVulncheckNvd2Get(auth).
    Cve("CVE-2019-19781")

  resp, _, err := req.Execute()
  if err != nil {
    log.Fatal(err)
  }

  vuln := resp.GetData()[0]

  fmt.Printf("Name: %q\n", vuln.GetCisaVulnerabilityName())

  descriptions := vuln.GetDescriptions()
  for _, desc := range descriptions {
    if desc.GetLang() == "en" {
      fmt.Printf("Description: %s\n", desc.GetValue())
    }
  }

  if metrics, ok := vuln.GetMetricsOk(); !ok {
    fmt.Println("")
  } else {
    fmt.Printf("Base Score: %f\n", metrics.CvssMetricV31[0].CvssData.GetBaseScore())
  }
}
