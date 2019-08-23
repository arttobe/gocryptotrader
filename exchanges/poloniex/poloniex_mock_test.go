//+build !mock_test_off

// This will build if build tag mock_test_off is not parsed and will try to mock
// all tests in _test.go
package poloniex

import (
	"log"
	"os"
	"testing"

	"github.com/thrasher-corp/gocryptotrader/config"
	"github.com/thrasher-corp/gocryptotrader/exchanges/mock"
	"github.com/thrasher-corp/gocryptotrader/exchanges/sharedtestvalues"
)

const mockfile = "../../testdata/http_mock/poloniex/poloniex.json"

var mockTests = true

func TestMain(m *testing.M) {
	cfg := config.GetConfig()
	cfg.LoadConfig("../../testdata/configtest.json")
	poloniexConfig, err := cfg.GetExchangeConfig("Poloniex")
	if err != nil {
		log.Fatal("Test Failed - Poloniex Setup() init error", err)
	}
	p.SkipAuthCheck = true
	poloniexConfig.API.AuthenticatedSupport = true
	poloniexConfig.API.Credentials.Key = apiKey
	poloniexConfig.API.Credentials.Secret = apiSecret
	p.SetDefaults()
	p.Setup(poloniexConfig)

	serverDetails, newClient, err := mock.NewVCRServer(mockfile)
	if err != nil {
		log.Fatalf("Test Failed - Mock server error %s", err)
	}

	p.HTTPClient = newClient
	p.API.Endpoints.URL = serverDetails

	log.Printf(sharedtestvalues.MockTesting, p.GetName(), p.API.Endpoints.URL)
	os.Exit(m.Run())
}
