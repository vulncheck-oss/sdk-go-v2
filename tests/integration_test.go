package tests

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vulncheck-oss/sdk-go-v2/vulncheck"
)

const (
	SCHEME = "https"
	HOST   = "api.vulncheck.com"
)

func TestSDK(t *testing.T) {
	configuration := vulncheck.NewConfiguration()
	configuration.Scheme = SCHEME
	configuration.Host = HOST

	client := vulncheck.NewAPIClient(configuration)

	token := os.Getenv("VULNCHECK_API_TOKEN")
	auth := context.WithValue(
		context.Background(),
		vulncheck.ContextAPIKeys,
		map[string]vulncheck.APIKey{
			"Bearer": {Key: token},
		},
	)

	t.Run("openapi spec", func(t *testing.T) {
		res, httpRes, err := client.EndpointsAPI.OpenapiGet(
			context.Background(),
		).Execute()

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, httpRes.StatusCode)
		assert.NotNil(t, res)
	})

	validateResponse := func(t testing.TB, data any, ok bool, httpRes *http.Response, err error) {
		t.Helper()
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, httpRes.StatusCode)

		assert.True(t, ok)
		assert.NotNil(t, data)
	}

	t.Run("endpoints api", func(t *testing.T) {
		t.Run("purl", func(t *testing.T) {
			const purl = "pkg:hex/coherence@0.1.2"
			req := client.EndpointsAPI.PurlGet(auth).Purl(purl)
			res, httpRes, err := req.Execute()

			data, ok := res.GetDataOk()
			validateResponse(t, data, ok, httpRes, err)
		})

		t.Run("cpe", func(t *testing.T) {
			const cpe = "cpe:/a:microsoft:internet_explorer:8.0.6001:beta"
			req := client.EndpointsAPI.CpeGet(auth).Cpe(cpe)
			res, httpRes, err := req.Execute()

			data, ok := res.GetDataOk()
			validateResponse(t, data, ok, httpRes, err)
		})

		t.Run("backups", func(t *testing.T) {
			req := client.EndpointsAPI.BackupIndexGet(auth, "initial-access")
			res, httpRes, err := req.Execute()

			data, ok := res.GetDataOk()
			validateResponse(t, data, ok, httpRes, err)
		})

		t.Run("get all indices", func(t *testing.T) {
			req := client.EndpointsAPI.IndexGet(auth)
			res, httpRes, err := req.Execute()

			data, ok := res.GetDataOk()
			validateResponse(t, data, ok, httpRes, err)
		})
	})

	t.Run("indices api", func(t *testing.T) {
		t.Run("nist-nvd2", func(t *testing.T) {
			req := client.
				IndicesAPI.
				IndexVulncheckNvd2Get(auth).
				Cve("CVE-2019-19781")

			res, httpRes, err := req.Execute()
			data, ok := res.GetDataOk()

			validateResponse(t, data, ok, httpRes, err)
		})

		t.Run("ipintel-3d", func(t *testing.T) {
			req := client.
				IndicesAPI.
				IndexIpintel3dGet(auth)

			res, httpRes, err := req.Execute()
			data, ok := res.GetDataOk()

			validateResponse(t, data, ok, httpRes, err)
		})

		t.Run("vulncheck-kev", func(t *testing.T) {
			req := client.
				IndicesAPI.
				IndexVulncheckKevGet(auth)

			res, httpRes, err := req.Execute()
			data, ok := res.GetDataOk()

			validateResponse(t, data, ok, httpRes, err)
		})

		t.Run("cisa-kev", func(t *testing.T) {
			req := client.
				IndicesAPI.
				IndexCisaKevGet(auth)

			res, httpRes, err := req.Execute()
			data, ok := res.GetDataOk()

			validateResponse(t, data, ok, httpRes, err)
		})
	})
}
