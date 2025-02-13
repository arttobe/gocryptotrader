package engine

import (
	"errors"

	"github.com/thrasher-corp/gocryptotrader/communications/base"
	"github.com/thrasher-corp/gocryptotrader/currency"
	exchange "github.com/thrasher-corp/gocryptotrader/exchanges"
	"github.com/thrasher-corp/gocryptotrader/exchanges/asset"
	"github.com/thrasher-corp/gocryptotrader/exchanges/order"
	"github.com/thrasher-corp/gocryptotrader/exchanges/orderbook"
	"github.com/thrasher-corp/gocryptotrader/exchanges/stream"
	"github.com/thrasher-corp/gocryptotrader/exchanges/ticker"
	"github.com/thrasher-corp/gocryptotrader/portfolio"
)

const (
	// MsgSubSystemStarting message to return when subsystem is starting up
	MsgSubSystemStarting = "starting..."
	// MsgSubSystemStarted message to return when subsystem has started
	MsgSubSystemStarted = "started."
	// MsgSubSystemShuttingDown message to return when a subsystem is shutting down
	MsgSubSystemShuttingDown = "shutting down..."
	// MsgSubSystemShutdown message to return when a subsystem has shutdown
	MsgSubSystemShutdown = "shutdown."
)

var (
	// ErrSubSystemAlreadyStarted message to return when a subsystem is already started
	ErrSubSystemAlreadyStarted = errors.New("subsystem already started")
	// ErrSubSystemNotStarted message to return when subsystem not started
	ErrSubSystemNotStarted = errors.New("subsystem not started")
	// ErrNilSubsystem is returned when a subsystem hasn't had its Setup() func run
	ErrNilSubsystem       = errors.New("subsystem not setup")
	errNilWaitGroup       = errors.New("nil wait group received")
	errNilExchangeManager = errors.New("cannot start with nil exchange manager")
)

// iExchangeManager limits exposure of accessible functions to exchange manager
// so that subsystems can use some functionality
type iExchangeManager interface {
	GetExchanges() []exchange.IBotExchange
	GetExchangeByName(string) exchange.IBotExchange
}

// iCommsManager limits exposure of accessible functions to communication manager
type iCommsManager interface {
	PushEvent(evt base.Event)
}

// iOrderManager defines a limited scoped order manager
type iOrderManager interface {
	Exists(*order.Detail) bool
	Add(*order.Detail) error
	Cancel(*order.Cancel) error
	GetByExchangeAndID(string, string) (*order.Detail, error)
	UpdateExistingOrder(*order.Detail) error
}

// iPortfolioManager limits exposure of accessible functions to portfolio manager
type iPortfolioManager interface {
	GetPortfolioSummary() portfolio.Summary
	IsWhiteListed(string) bool
	IsExchangeSupported(string, string) bool
}

// iBot limits exposure of accessible functions to engine bot
type iBot interface {
	SetupExchanges() error
}

// iWebsocketDataReceiver limits exposure of accessible functions to websocket data receiver
type iWebsocketDataReceiver interface {
	IsRunning() bool
	WebsocketDataReceiver(ws *stream.Websocket)
	WebsocketDataHandler(string, interface{}) error
}

// iCurrencyPairSyncer defines a limited scoped currency pair syncer
type iCurrencyPairSyncer interface {
	IsRunning() bool
	PrintTickerSummary(*ticker.Price, string, error)
	PrintOrderbookSummary(*orderbook.Base, string, error)
	Update(string, currency.Pair, asset.Item, int, error) error
}
