package stock_landed_costs

import (
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.PurchaseConfigSettings().DeclareModel()

	h.PurchaseConfigSettings().Methods().OnchangeCostingMethod().DeclareMethod(
		`OnchangeCostingMethod`,
		func(rs m.PurchaseConfigSettingsSet) {
			//        if self.group_costing_method == 0:
			//            self.group_costing_method = 1
			//            return {
			//                'warning': {
			//                    'title': _('Warning!'),
			//                    'message': _('Disabling the costing methods will prevent you to use the landed costs feature.'),
			//                }
			//            }
		})
}
