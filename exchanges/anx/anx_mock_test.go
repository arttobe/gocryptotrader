//+build !mock_test_off

// This will build if build tag mock_test_off is not parsed and will try to mock
// all tests in _test.go
package anx

import (
	"log"
	"os"
	"testing"

	"github.com/thrasher-corp/gocryptotrader/config"
	"github.com/thrasher-corp/gocryptotrader/exchanges/mock"
	"github.com/thrasher-corp/gocryptotrader/exchanges/sharedtestvalues"
)

const mockFile = "../../testdata/http_mock/anx/anx.json"

var mockTests = true

func TestMain(m *testing.M) {
	cfg := config.GetConfig()
	cfg.LoadConfig("../../testdata/configtest.json")
	anxConfig, err := cfg.GetExchangeConfig("ANX")
	if err != nil {
		log.Fatal("Test Failed - Mock server error", err)
	}
	a.SkipAuthCheck = true
	anxConfig.API.AuthenticatedSupport = true
	anxConfig.API.Credentials.Key = apiKey
	anxConfig.API.Credentials.Secret = apiSecret
	a.SetDefaults()
	a.Setup(anxConfig)

	serverDetails, newClient, err := mock.NewVCRServer(mockFile)
	if err != nil {
		log.Fatalf("Test Failed - Mock server error %s", err)
	}

	a.HTTPClient = newClient
	a.API.Endpoints.URL = serverDetails + "/"

	log.Printf(sharedtestvalues.MockTesting, a.GetName(), a.API.Endpoints.URL)
	os.Exit(m.Run())
}
