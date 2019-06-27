package stock_landed_costs

import (
	"github.com/hexya-erp/hexya/src/server"
)

const MODULE_NAME string = "stock_landed_costs"

func init() {
	server.RegisterModule(&server.Module{
		Name:     MODULE_NAME,
		PreInit:  func() {},
		PostInit: func() {},
	})

}
