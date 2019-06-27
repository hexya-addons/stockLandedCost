package stock_landed_costs

import (
	"github.com/hexya-addons/base"
	"github.com/hexya-erp/hexya/src/models/security"
	"github.com/hexya-erp/pool/h"
)

//vars

//rights
func init() {
	h."ModelStockLandedCost"().Methods().AllowAllToGroup(GroupStockManager")
	h."ModelStockLandedCostLines"().Methods().AllowAllToGroup(GroupStockManager")
	h."ModelStockValuationAdjustmentLines"().Methods().AllowAllToGroup(GroupStockManager")
}
