package coinsuper

import (
	"strconv"
	"testing"
	"time"

	"github.com/hexoul/go-coinsuper/types"
)

func init() {
	GetInstanceWithKey("YOUR_ACCESS_KEY", "YOUR_SECRET_KEY")
}

func TestUserAssetInfo(t *testing.T) {
	if info, err := GetInstance().UserAssetInfo(nil); err != nil {
		t.Fatal(err)
	} else if info.Assets["BTC"] == nil {
		t.FailNow()
	}
}

func TestTradeHistory(t *testing.T) {
	past1hour := strconv.FormatInt(time.Now().Add(-time.Hour).Unix()*1000, 10)
	if info, err := GetInstance().TradeHistory(&types.Options{
		Symbol:    "ETC/BTC",
		UtcStart:  past1hour,
		WithTrade: "true",
	}); err != nil {
		t.Fatal(err)
	} else if len(info.TransactionInfoList) == 0 {
		t.FailNow()
	}
}
