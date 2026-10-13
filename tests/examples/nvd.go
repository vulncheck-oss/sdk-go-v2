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

  // Look up a single CVE, shaped exactly as NVD's API returns it
  resp, httpRes, err := client.
    NvdAPI.
    NvdCvesGet(auth).
    CveId("CVE-2021-44228").
    Execute()
  if err != nil || httpRes.StatusCode != 200 {
    log.Fatal(err)
  }

  cve := resp.Vulnerabilities[0].GetCve()
  fmt.Printf("%s: %s\n", cve.GetId(), cve.GetVulnStatus())

  // Walk every CVE modified in a date range. Sending an empty cursor opts in
  // to cursor paging; each response then carries the cursor for the next page
  req := client.
    NvdAPI.
    NvdCvesGet(auth).
    LastModStartDate("2026-09-01T00:00:00.000").
    LastModEndDate("2026-09-02T00:00:00.000").
    ResultsPerPage(500)

  var count int
  cursor := ""
  for {
    resp, httpRes, err := req.Cursor(cursor).Execute()
    if err != nil || httpRes.StatusCode != 200 {
      log.Fatal(err)
    }

    count += len(resp.Vulnerabilities)
    fmt.Printf("Total Items: %d/%d\n", count, resp.TotalResults)

    cursor = resp.GetNextCursor()
    if cursor == "" {
      break
    }
  }
}
