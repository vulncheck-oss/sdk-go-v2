//go:build ignore

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
  fileResp, err := http.Get(resp.GetUrlMrap())
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
