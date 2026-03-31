<p align="center">
    <img src="https://raw.githubusercontent.com/vulncheck-oss/sdk-go-v2/refs/heads/main/logo-sdk.png" align="center" alt="VulnCheck Logo" width="150" />
</p>

# The VulnCheck SDK For Go

Bring the VulnCheck API to your Go applications.

## Installation

```bash
go get github.com/vulncheck-oss/sdk-go-v2/v2
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
```


### Available Methods

### PURL

Get the CVE's for a given PURL

```go
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

  req := client.EndpointsAPI.PurlGet(auth).Purl("pkg:hex/coherence@0.1.2")
  resp, httpRes, err := req.Execute()
  if err != nil || httpRes.StatusCode != 200 {
    log.Fatal(err)
  }

  data := resp.GetData()

  fmt.Println(data.GetCves())
}
```


### CPE

Get all CPE's related to a CVE

```go
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

  req := client.EndpointsAPI.CpeGet(auth).Cpe("cpe:/a:microsoft:internet_explorer:8.0.6001:beta")
  resp, httpRes, err := req.Execute()
  if err != nil || httpRes.StatusCode != 200 {
    log.Fatal(err)
  }

  for _, v := range resp.GetData() {
    fmt.Println(v)
  }
}
```


### Advisory

List advisory feeds and query advisories

```go
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
```


### Backup

Download the backup for an index

```go
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
```


### Backup V4

List and download v4 backups by feed name

```go
package main

import (
  "context"
  "fmt"
  "io"
  "log"
  "net/http"
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

  // List available backups (/v4/backup)
  available, httpRes, err := client.BackupAPI.V4ListBackups(auth).Execute()
  if err != nil || httpRes.StatusCode != 200 {
    log.Fatal(err)
  }

  for _, b := range available.GetData() {
    fmt.Printf("Found backup: %s\n", b.GetName())
  }

  // Get backup for the wolfi feed (/v4/backup/wolfi)
  feed := "wolfi"
  resp, httpRes, err := client.BackupAPI.V4GetBackupByName(auth, feed).Execute()
  if err != nil || httpRes.StatusCode != 200 {
    log.Fatal(err)
  }

  fmt.Printf("Downloading %s backup\n", feed)
  filePath := feed + ".zip"
  fileResp, err := http.Get(resp.GetUrl())
  if err != nil {
    log.Fatal(err)
  }
  defer fileResp.Body.Close()

  f, err := os.Create(filePath)
  if err != nil {
    log.Fatal(err)
  }
  defer f.Close()

  if _, err := io.Copy(f, fileResp.Body); err != nil {
    log.Fatal(err)
  }

  fmt.Printf("Successfully saved to %s\n", filePath)
}
```


### Indices

Get all available indices

```go
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

  req := client.EndpointsAPI.IndexGet(auth)
  resp, httpRes, err := req.Execute()
  if err != nil || httpRes.StatusCode != 200 {
    log.Fatal(err)
  }

  for _, v := range resp.GetData() {
    fmt.Println(v.GetName())
  }
}
```


### Index

Query VulnCheck-NVD2 for `CVE-2019-19781`

```go
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
```


### Pagination

Paginate over results for a query to VulnCheck-KEV using `cursor`

```go
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
```


## License

Apache License 2.0. Please see [License File](./LICENSE) for more information.
