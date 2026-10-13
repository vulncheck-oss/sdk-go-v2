package tests

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	vulncheck "github.com/vulncheck-oss/sdk-go-v2/v2"
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

	t.Run("nvd api", func(t *testing.T) {
		const cveID = "CVE-2021-44228"

		t.Run("cves", func(t *testing.T) {
			res, httpRes, err := client.NvdAPI.NvdCvesGet(auth).CveId(cveID).Execute()

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, httpRes.StatusCode)
			assert.Equal(t, int32(1), res.TotalResults)
			assert.Equal(t, cveID, res.Vulnerabilities[0].Cve.GetId())
		})

		t.Run("cves cursor", func(t *testing.T) {
			req := client.NvdAPI.NvdCvesGet(auth).CveTag("disputed").ResultsPerPage(5)

			first, httpRes, err := req.Cursor("").Execute()
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, httpRes.StatusCode)
			assert.NotEmpty(t, first.GetNextCursor())

			second, httpRes, err := req.Cursor(first.GetNextCursor()).Execute()
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, httpRes.StatusCode)
			assert.NotEqual(t,
				first.Vulnerabilities[0].Cve.GetId(),
				second.Vulnerabilities[0].Cve.GetId(),
			)
		})

		t.Run("cve history", func(t *testing.T) {
			res, httpRes, err := client.NvdAPI.NvdCveHistoryGet(auth).CveId(cveID).Execute()

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, httpRes.StatusCode)
			assert.NotEmpty(t, res.CveChanges)
		})

		t.Run("cpes", func(t *testing.T) {
			res, httpRes, err := client.NvdAPI.NvdCpesGet(auth).
				CpeMatchString("cpe:2.3:a:apache:log4j").
				ResultsPerPage(5).
				Execute()

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, httpRes.StatusCode)
			assert.NotEmpty(t, res.Products)
		})

		t.Run("cpe match", func(t *testing.T) {
			res, httpRes, err := client.NvdAPI.NvdCpeMatchGet(auth).CveId(cveID).ResultsPerPage(5).Execute()

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, httpRes.StatusCode)
			assert.NotEmpty(t, res.MatchStrings)
		})

		t.Run("source", func(t *testing.T) {
			res, httpRes, err := client.NvdAPI.NvdSourceGet(auth).ResultsPerPage(5).Execute()

			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, httpRes.StatusCode)
			assert.NotEmpty(t, res.Sources)
		})
	})
}
