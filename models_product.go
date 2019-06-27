package stock_landed_costs

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

//SPLIT_METHOD = [
//    ('equal', 'Equal'),
//    ('by_quantity', 'By Quantity'),
//    ('by_current_cost_price', 'By Current Cost'),
//    ('by_weight', 'By Weight'),
//    ('by_volume', 'By Volume'),
//]
func init() {
	h.ProductTemplate().DeclareModel()

	h.ProductTemplate().AddFields(map[string]models.FieldDefinition{
		"LandedCostOk": models.BooleanField{
			String: "Landed Costs",
		},
		"SplitMethod": models.SelectionField{
			//selection=SPLIT_METHOD
			String:  "Split Method",
			Default: models.DefaultValue("equal"),
			Help: "Equal : Cost will be equally divided." +
				"By Quantity : Cost will be divided according to product's quantity." +
				"By Current cost : Cost will be divided according to product's" +
				"current cost." +
				"By Weight : Cost will be divided depending on its weight." +
				"By Volume : Cost will be divided depending on its volume.",
		},
	})
}
