package huobikorea_test

import (
	"testing"

	"github.com/investing-kr/xhuobi/huobikorea"
)

func TestPrice(t *testing.T) {
	tick, err := huobikorea.MarketClient.GetLatestTrade("btckrw")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(tick.Data[0].Price.String())
}
