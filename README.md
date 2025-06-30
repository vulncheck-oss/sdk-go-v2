<p align="center">
    <img src="/logo-sdk.png" align="center" alt="VulnCheck Logo" width="150" />
</p>

# The VulnCheck SDK For Go

Bring the VulnCheck API to your Go applications.

## Installation

```bash
go get github.com/vulncheck-oss/sdk-go-v2/vulncheck
```

## Examples

### Connecting to the API

```go
package main

import (
  "context"
  "fmt"
  "log"
  "os"

  "github.com/vulncheck/sdk-go-v2/vulncheck"
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
```

### Available Methods

### PURL

Get the CVE's for a given PURL

```go
req := client.EndpointsAPI.PurlGet(auth).Purl("pkg:hex/coherence@0.1.2")
resp, httpRes, err := req.Execute()
if err != nil {
  log.Fatal(err)
}

data := resp.GetData()

fmt.Println(data.GetCves())
```

### CPE

Get all CPE's related to a CVE

```go
req := client.EndpointsAPI.CpeGet(auth).Cpe("cpe:/a:microsoft:internet_explorer:8.0.6001:beta")
resp, httpRes, err := req.Execute()
if err != nil {
  log.Fatal(err)
}

for _, v := range resp.GetData() {
  fmt.Println(v)
}
```

### Backup

Download the backup for an index

```go
req := client.EndpointsAPI.BackupIndexGet(auth, "initial-access")
resp, httpRes, err := req.Execute()
if err != nil {
  log.Fatal(err)
}

data := resp.GetData()

for _, v := range data {
  fmt.Println(v.GetUrl())
}
```

### Indices

Get all available indices

```go
req := client.EndpointsAPI.IndexGet(auth)
resp, httpRes, err := req.Execute()
if err != nil {
  log.Fatal(err)
}

for _, v := range resp.GetData() {
  fmt.Println(v.GetName())
}
```

### Index

Query VulnCheck-NVD2 for `CVE-2019-19781`

```go
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
```

### Pagination

Paginate over results for a query to VulnCheck-KEV using `cursor`

```go
var queryLimit int32 = 300
req := client.
  IndicesAPI.
  IndexVulncheckKevGet(auth).
  StartCursor("true").
  Limit(queryLimit)

resp, httpRes, err := req.Execute()
if err != nil {
  log.Fatal(err)
}

var nextCursor string
var count int
var total int32

nextCursor = resp.Meta.GetNextCursor()
total = resp.Meta.GetTotalDocuments()
count += len(resp.Data)
fmt.Printf("Total Items: %d/%d\n", count, total)

count += len(resp.Data)

for range 10 {
  req := client.
    IndicesAPI.
    IndexVulncheckKevGet(auth).
    Cursor(nextCursor).
    Limit(queryLimit)

  resp, httpRes, err := req.Execute()
  if err != nil {
    log.Fatal(err)
  }

  nextCursor = resp.Meta.GetNextCursor()
  count += len(resp.Data)
  fmt.Printf("Total Items: %d/%d\n", count, total)
}
```

## License

Apache License 2.0. Please see License File for more information.
