package huobikorea

import (
	"os"

	"github.com/huobirdcenter/huobi_golang/pkg/client"
)

var (
	Host      = "api-cloud.huobi.co.kr"
	AccessKey = os.Getenv("HUOBI_KOREA_API_KEY")
	SecretKey = os.Getenv("HUOBI_KOREA_SECRET_KEY")

	CommonClient  = new(client.CommonClient).Init(Host)
	MarketClient  = new(client.MarketClient).Init(Host)
	AccountClient = new(client.AccountClient).Init(Host, AccessKey, SecretKey)
	OrderClinet   = new(client.OrderClient).Init(Host, AccessKey, SecretKey)
	WalletClient  = new(client.WalletClient).Init(Host, AccessKey, SecretKey)
)
