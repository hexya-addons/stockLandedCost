package stock_landed_costs

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/q"
)

func init() {
	h.StockLandedCost().DeclareModel()

	h.StockLandedCost().AddFields(map[string]models.FieldDefinition{
		"Name": models.CharField{
			String:   "Name",
			Default:  func(env models.Environment) interface{} { return odoo._() },
			NoCopy:   true,
			ReadOnly: true,
			//track_visibility='always'
		},
		"Date": models.DateField{
			String:   "Date",
			Default:  func(env models.Environment) interface{} { return odoo.fields.Date.context_today },
			NoCopy:   true,
			Required: true,
			//states={'done': [('readonly', True)]}
			//track_visibility='onchange'
		},
		"PickingIds": models.Many2ManyField{
			RelationModel: h.StockPicking(),
			String:        "Pickings",
			NoCopy:        true,
			//states={'done': [('readonly', True)]}
		},
		"CostLines": models.One2ManyField{
			RelationModel: h.StockLandedCostLines(),
			ReverseFK:     "",
			String:        "Cost Lines",
			NoCopy:        false,
			//states={'done': [('readonly', True)]}
		},
		"ValuationAdjustmentLines": models.One2ManyField{
			RelationModel: h.StockValuationAdjustmentLines(),
			ReverseFK:     "",
			String:        "Valuation Adjustments",
			//states={'done': [('readonly', True)]}
		},
		"Description": models.TextField{
			String: "Item Description",
			//states={'done': [('readonly', True)]}
		},
		"AmountTotal": models.FloatField{
			String:  "Total",
			Compute: h.StockLandedCost().Methods().ComputeTotalAmount(),
			//digits=0
			Stored: true,
			//track_visibility='always'
		},
		"State": models.SelectionField{
			Selection: types.Selection{
				"draft":  "Draft",
				"done":   "Posted",
				"cancel": "Cancelled",
			},
			String:   "State",
			Default:  models.DefaultValue("draft"),
			NoCopy:   true,
			ReadOnly: true,
			//track_visibility='onchange'
		},
		"AccountMoveId": models.Many2OneField{
			RelationModel: h.AccountMove(),
			String:        "Journal Entry",
			NoCopy:        true,
			ReadOnly:      true,
		},
		"AccountJournalId": models.Many2OneField{
			RelationModel: h.AccountJournal(),
			String:        "Account Journal",
			Required:      true,
			//states={'done': [('readonly', True)]}
		},
	})
	h.StockLandedCost().Methods().ComputeTotalAmount().DeclareMethod(
		`ComputeTotalAmount`,
		func(rs h.StockLandedCostSet) h.StockLandedCostData {
			//        self.amount_total = sum(line.price_unit for line in self.cost_lines)
		})
	h.StockLandedCost().Methods().Create().Extend(
		`Create`,
		func(rs m.StockLandedCostSet, vals models.RecordData) {
			//        if vals.get('name', _('New')) == _('New'):
			//            vals['name'] = self.env['ir.sequence'].next_by_code(
			//                'stock.landed.cost')
			//        return super(LandedCost, self).create(vals)
		})
	h.StockLandedCost().Methods().Unlink().Extend(
		`Unlink`,
		func(rs m.StockLandedCostSet) {
			//        self.button_cancel()
			//        return super(LandedCost, self).unlink()
		})
	h.StockLandedCost().Methods().TrackSubtype().DeclareMethod(
		`TrackSubtype`,
		func(rs m.StockLandedCostSet, init_values interface{}) {
			//        if 'state' in init_values and self.state == 'done':
			//            return 'stock_landed_costs.mt_stock_landed_cost_open'
			//        return super(LandedCost, self)._track_subtype(init_values)
		})
	h.StockLandedCost().Methods().ButtonCancel().DeclareMethod(
		`ButtonCancel`,
		func(rs m.StockLandedCostSet) {
			//        if any(cost.state == 'done' for cost in self):
			//            raise UserError(
			//                _('Validated landed costs cannot be cancelled, but you could create negative landed costs to reverse them'))
			//        return self.write({'state': 'cancel'})
		})
	h.StockLandedCost().Methods().ButtonValidate().DeclareMethod(
		`ButtonValidate`,
		func(rs m.StockLandedCostSet) {
			//        if any(cost.state != 'draft' for cost in self):
			//            raise UserError(_('Only draft landed costs can be validated'))
			//        if any(not cost.valuation_adjustment_lines for cost in self):
			//            raise UserError(
			//                _('No valuation adjustments lines. You should maybe recompute the landed costs.'))
			//        if not self._check_sum():
			//            raise UserError(
			//                _('Cost and adjustments lines do not match. You should maybe recompute the landed costs.'))
			//        for cost in self:
			//            move = self.env['account.move']
			//            move_vals = {
			//                'journal_id': cost.account_journal_id.id,
			//                'date': cost.date,
			//                'ref': cost.name,
			//                'line_ids': [],
			//            }
			//            for line in cost.valuation_adjustment_lines.filtered(lambda line: line.move_id):
			//                per_unit = line.final_cost / line.quantity
			//                diff = per_unit - line.former_cost_per_unit
			//
			//                # If the precision required for the variable diff is larger than the accounting
			//                # precision, inconsistencies between the stock valuation and the accounting entries
			//                # may arise.
			//                # For example, a landed cost of 15 divided in 13 units. If the products leave the
			//                # stock one unit at a time, the amount related to the landed cost will correspond to
			//                # round(15/13, 2)*13 = 14.95. To avoid this case, we split the quant in 12 + 1, then
			//                # record the difference on the new quant.
			//                # We need to make sure to able to extract at least one unit of the product. There is
			//                # an arbitrary minimum quantity set to 2.0 from which we consider we can extract a
			//                # unit and adapt the cost.
			//                curr_rounding = line.move_id.company_id.currency_id.rounding
			//                diff_rounded = tools.float_round(
			//                    diff, precision_rounding=curr_rounding)
			//                diff_correct = diff_rounded
			//                quants = line.move_id.quant_ids.sorted(
			//                    key=lambda r: r.qty, reverse=True)
			//                quant_correct = False
			//                if quants\
			//                        and tools.float_compare(quants[0].product_id.uom_id.rounding, 1.0, precision_digits=1) == 0\
			//                        and tools.float_compare(line.quantity * diff, line.quantity * diff_rounded, precision_rounding=curr_rounding) != 0\
			//                        and tools.float_compare(quants[0].qty, 2.0, precision_rounding=quants[0].product_id.uom_id.rounding) >= 0:
			//                    # Search for existing quant of quantity = 1.0 to avoid creating a new one
			//                    quant_correct = quants.filtered(lambda r: tools.float_compare(
			//                        r.qty, 1.0, precision_rounding=quants[0].product_id.uom_id.rounding) == 0)
			//                    if not quant_correct:
			//                        quant_correct = quants[0]._quant_split(
			//                            quants[0].qty - 1.0)
			//                    else:
			//                        quant_correct = quant_correct[0]
			//                        quants = quants - quant_correct
			//                    diff_correct += (line.quantity * diff) - \
			//                        (line.quantity * diff_rounded)
			//                    diff = diff_rounded
			//
			//                quant_dict = {}
			//                for quant in quants:
			//                    quant_dict[quant] = quant.cost + diff
			//                if quant_correct:
			//                    quant_dict[quant_correct] = quant_correct.cost + \
			//                        diff_correct
			//                for quant, value in quant_dict.items():
			//                    quant.sudo().write({'cost': value})
			//                qty_out = 0
			//                for quant in line.move_id.quant_ids:
			//                    if quant.location_id.usage != 'internal':
			//                        qty_out += quant.qty
			//                move_vals['line_ids'] += line._create_accounting_entries(
			//                    move, qty_out)
			//
			//            move = move.create(move_vals)
			//            cost.write({'state': 'done', 'account_move_id': move.id})
			//            move.post()
			//        return True
		})
	h.StockLandedCost().Methods().CheckSum().DeclareMethod(
		` Check if each cost line its valuation lines sum to the correct amount
        and if the overall total amount is correct also `,
		func(rs m.StockLandedCostSet) {
			//        prec_digits = self.env['decimal.precision'].precision_get('Account')
			//        for landed_cost in self:
			//            total_amount = sum(landed_cost.valuation_adjustment_lines.mapped(
			//                'additional_landed_cost'))
			//            if not tools.float_compare(total_amount, landed_cost.amount_total, precision_digits=prec_digits) == 0:
			//                return False
			//
			//            val_to_cost_lines = defaultdict(lambda: 0.0)
			//            for val_line in landed_cost.valuation_adjustment_lines:
			//                val_to_cost_lines[val_line.cost_line_id] += val_line.additional_landed_cost
			//            if any(tools.float_compare(cost_line.price_unit, val_amount, precision_digits=prec_digits) != 0
			//                   for cost_line, val_amount in val_to_cost_lines.iteritems()):
			//                return False
			//        return True
		})
	h.StockLandedCost().Methods().GetValuationLines().DeclareMethod(
		`GetValuationLines`,
		func(rs m.StockLandedCostSet) {
			//        lines = []
			//        for move in self.mapped('picking_ids').mapped('move_lines'):
			//            # it doesn't make sense to make a landed cost for a product that isn't set as being valuated in real time at real cost
			//            if move.product_id.valuation != 'real_time' or move.product_id.cost_method != 'real':
			//                continue
			//            vals = {
			//                'product_id': move.product_id.id,
			//                'move_id': move.id,
			//                'quantity': move.product_qty,
			//                'former_cost': sum(quant.cost * quant.qty for quant in move.quant_ids),
			//                'weight': move.product_id.weight * move.product_qty,
			//                'volume': move.product_id.volume * move.product_qty
			//            }
			//            lines.append(vals)
			//        if not lines and self.mapped('picking_ids'):
			//            raise UserError(_('The selected picking does not contain any move that would be impacted by landed costs. Landed costs are only possible for products configured in real time valuation with real price costing method. Please make sure it is the case, or you selected the correct picking'))
			//        return lines
		})
	h.StockLandedCost().Methods().ComputeLandedCost().DeclareMethod(
		`ComputeLandedCost`,
		func(rs m.StockLandedCostSet) {
			//        AdjustementLines = self.env['stock.valuation.adjustment.lines']
			//        AdjustementLines.search([('cost_id', 'in', self.ids)]).unlink()
			//        digits = dp.get_precision('Product Price')(self._cr)
			//        towrite_dict = {}
			//        for cost in self.filtered(lambda cost: cost.picking_ids):
			//            total_qty = 0.0
			//            total_cost = 0.0
			//            total_weight = 0.0
			//            total_volume = 0.0
			//            total_line = 0.0
			//            all_val_line_values = cost.get_valuation_lines()
			//            for val_line_values in all_val_line_values:
			//                for cost_line in cost.cost_lines:
			//                    val_line_values.update(
			//                        {'cost_id': cost.id, 'cost_line_id': cost_line.id})
			//                    self.env['stock.valuation.adjustment.lines'].create(
			//                        val_line_values)
			//                total_qty += val_line_values.get('quantity', 0.0)
			//                total_weight += val_line_values.get('weight', 0.0)
			//                total_volume += val_line_values.get('volume', 0.0)
			//
			//                former_cost = val_line_values.get('former_cost', 0.0)
			//                # round this because former_cost on the valuation lines is also rounded
			//                total_cost += tools.float_round(
			//                    former_cost, precision_digits=digits[1]) if digits else former_cost
			//
			//                total_line += 1
			//
			//            for line in cost.cost_lines:
			//                value_split = 0.0
			//                for valuation in cost.valuation_adjustment_lines:
			//                    value = 0.0
			//                    if valuation.cost_line_id and valuation.cost_line_id.id == line.id:
			//                        if line.split_method == 'by_quantity' and total_qty:
			//                            per_unit = (line.price_unit / total_qty)
			//                            value = valuation.quantity * per_unit
			//                        elif line.split_method == 'by_weight' and total_weight:
			//                            per_unit = (line.price_unit / total_weight)
			//                            value = valuation.weight * per_unit
			//                        elif line.split_method == 'by_volume' and total_volume:
			//                            per_unit = (line.price_unit / total_volume)
			//                            value = valuation.volume * per_unit
			//                        elif line.split_method == 'equal':
			//                            value = (line.price_unit / total_line)
			//                        elif line.split_method == 'by_current_cost_price' and total_cost:
			//                            per_unit = (line.price_unit / total_cost)
			//                            value = valuation.former_cost * per_unit
			//                        else:
			//                            value = (line.price_unit / total_line)
			//
			//                        if digits:
			//                            value = tools.float_round(
			//                                value, precision_digits=digits[1], rounding_method='UP')
			//                            fnc = min if line.price_unit > 0 else max
			//                            value = fnc(value, line.price_unit - value_split)
			//                            value_split += value
			//
			//                        if valuation.id not in towrite_dict:
			//                            towrite_dict[valuation.id] = value
			//                        else:
			//                            towrite_dict[valuation.id] += value
			//        if towrite_dict:
			//            for key, value in towrite_dict.items():
			//                AdjustementLines.browse(key).write(
			//                    {'additional_landed_cost': value})
			//        return True
		})
	h.StockLandedCostLines().DeclareModel()

	h.StockLandedCostLines().AddFields(map[string]models.FieldDefinition{
		"Name": models.CharField{
			String: "Description",
		},
		"CostId": models.Many2OneField{
			RelationModel: h.StockLandedCost(),
			String:        "Landed Cost",
			Required:      true,
			OnDelete:      `cascade`,
		},
		"ProductId": models.Many2OneField{
			RelationModel: h.ProductProduct(),
			String:        "Product",
			Required:      true,
		},
		"PriceUnit": models.FloatField{
			String: "Cost",
			//digits=dp.get_precision('Product Price')
			Required: true,
		},
		"SplitMethod": models.SelectionField{
			Selection: product.SPLIT_METHOD,
			String:    "Split Method",
			Required:  true,
		},
		"AccountId": models.Many2OneField{
			RelationModel: h.AccountAccount(),
			String:        "Account",
			Filter:        q.Deprecated().Equals(False),
		},
	})
	h.StockLandedCostLines().Methods().OnchangeProductId().DeclareMethod(
		`OnchangeProductId`,
		func(rs m.StockLandedCostLinesSet) {
			//        if not self.product_id:
			//            self.quantity = 0.0
			//        self.name = self.product_id.name or ''
			//        self.split_method = self.product_id.split_method or 'equal'
			//        self.price_unit = self.product_id.standard_price or 0.0
			//        self.account_id = self.product_id.property_account_expense_id.id or self.product_id.categ_id.property_account_expense_categ_id.id
		})
	h.StockValuationAdjustmentLines().DeclareModel()

	h.StockValuationAdjustmentLines().AddFields(map[string]models.FieldDefinition{
		"Name": models.CharField{
			String:  "Description",
			Compute: h.StockValuationAdjustmentLines().Methods().ComputeName(),
			Stored:  true,
		},
		"CostId": models.Many2OneField{
			RelationModel: h.StockLandedCost(),
			String:        "Landed Cost",
			OnDelete:      `cascade`,
			Required:      true,
		},
		"CostLineId": models.Many2OneField{
			RelationModel: h.StockLandedCostLines(),
			String:        "Cost Line",
			ReadOnly:      true,
		},
		"MoveId": models.Many2OneField{
			RelationModel: h.StockMove(),
			String:        "Stock Move",
			ReadOnly:      true,
		},
		"ProductId": models.Many2OneField{
			RelationModel: h.ProductProduct(),
			String:        "Product",
			Required:      true,
		},
		"Quantity": models.FloatField{
			String:  "Quantity",
			Default: models.DefaultValue(1),
			//digits=0
			Required: true,
		},
		"Weight": models.FloatField{
			String:  "Weight",
			Default: models.DefaultValue(1),
			//digits=dp.get_precision('Stock Weight')
		},
		"Volume": models.FloatField{
			String:  "Volume",
			Default: models.DefaultValue(1),
		},
		"FormerCost": models.FloatField{
			String: "Former Cost",
			//digits=dp.get_precision('Product Price')
		},
		"FormerCostPerUnit": models.FloatField{
			String:  "Former Cost(Per Unit)",
			Compute: h.StockValuationAdjustmentLines().Methods().ComputeFormerCostPerUnit(),
			//digits=0
			Stored: true,
		},
		"AdditionalLandedCost": models.FloatField{
			String: "Additional Landed Cost",
			//digits=dp.get_precision('Product Price')
		},
		"FinalCost": models.FloatField{
			String:  "Final Cost",
			Compute: h.StockValuationAdjustmentLines().Methods().ComputeFinalCost(),
			//digits=0
			Stored: true,
		},
	})
	h.StockValuationAdjustmentLines().Methods().ComputeName().DeclareMethod(
		`ComputeName`,
		func(rs h.StockValuationAdjustmentLinesSet) h.StockValuationAdjustmentLinesData {
			//        name = '%s - ' % (self.cost_line_id.name if self.cost_line_id else '')
			//        self.name = name + (self.product_id.code or self.product_id.name or '')
		})
	h.StockValuationAdjustmentLines().Methods().ComputeFormerCostPerUnit().DeclareMethod(
		`ComputeFormerCostPerUnit`,
		func(rs h.StockValuationAdjustmentLinesSet) h.StockValuationAdjustmentLinesData {
			//        self.former_cost_per_unit = self.former_cost / (self.quantity or 1.0)
		})
	h.StockValuationAdjustmentLines().Methods().ComputeFinalCost().DeclareMethod(
		`ComputeFinalCost`,
		func(rs h.StockValuationAdjustmentLinesSet) h.StockValuationAdjustmentLinesData {
			//        self.final_cost = self.former_cost + self.additional_landed_cost
		})
	h.StockValuationAdjustmentLines().Methods().CreateAccountingEntries().DeclareMethod(
		`CreateAccountingEntries`,
		func(rs m.StockValuationAdjustmentLinesSet, move interface{}, qty_out interface{}) {
			//        cost_product = self.cost_line_id.product_id
			//        if not cost_product:
			//            return False
			//        accounts = self.product_id.product_tmpl_id.get_product_accounts()
			//        debit_account_id = accounts.get(
			//            'stock_valuation') and accounts['stock_valuation'].id or False
			//        already_out_account_id = accounts['stock_output'].id
			//        credit_account_id = self.cost_line_id.account_id.id or cost_product.property_account_expense_id.id or cost_product.categ_id.property_account_expense_categ_id.id
			//        if not credit_account_id:
			//            raise UserError(_('Please configure Stock Expense Account for product: %s.') % (
			//                cost_product.name))
			//        return self._create_account_move_line(move, credit_account_id, debit_account_id, qty_out, already_out_account_id)
		})
	h.StockValuationAdjustmentLines().Methods().CreateAccountMoveLine().DeclareMethod(
		`
        Generate the account.move.line values to track the landed cost.
        Afterwards, for the goods that are already out
of stock, we should create the out moves
        `,
		func(rs m.StockValuationAdjustmentLinesSet, move interface{}, credit_account_id interface{}, debit_account_id interface{}, qty_out interface{}, already_out_account_id interface{}) {
			//        AccountMoveLine = []
			//        base_line = {
			//            'name': self.name,
			//            'product_id': self.product_id.id,
			//            'quantity': self.quantity,
			//        }
			//        debit_line = dict(base_line, account_id=debit_account_id)
			//        credit_line = dict(base_line, account_id=credit_account_id)
			//        diff = self.additional_landed_cost
			//        if diff > 0:
			//            debit_line['debit'] = diff
			//            credit_line['credit'] = diff
			//        else:
			//            # negative cost, reverse the entry
			//            debit_line['credit'] = -diff
			//            credit_line['debit'] = -diff
			//        AccountMoveLine.append([0, 0, debit_line])
			//        AccountMoveLine.append([0, 0, credit_line])
			//        if qty_out > 0:
			//            debit_line = dict(base_line,
			//                              name=(self.name + ": " +
			//                                    str(qty_out) + _(' already out')),
			//                              quantity=qty_out,
			//                              account_id=already_out_account_id)
			//            credit_line = dict(base_line,
			//                               name=(self.name + ": " +
			//                                     str(qty_out) + _(' already out')),
			//                               quantity=qty_out,
			//                               account_id=debit_account_id)
			//            diff = diff * qty_out / self.quantity
			//            if diff > 0:
			//                debit_line['debit'] = diff
			//                credit_line['credit'] = diff
			//            else:
			//                # negative cost, reverse the entry
			//                debit_line['credit'] = -diff
			//                credit_line['debit'] = -diff
			//            AccountMoveLine.append([0, 0, debit_line])
			//            AccountMoveLine.append([0, 0, credit_line])
			//
			//            # TDE FIXME: oh dear
			//            if self.env.user.company_id.anglo_saxon_accounting:
			//                debit_line = dict(base_line,
			//                                  name=(self.name + ": " +
			//                                        str(qty_out) + _(' already out')),
			//                                  quantity=qty_out,
			//                                  account_id=credit_account_id)
			//                credit_line = dict(base_line,
			//                                   name=(self.name + ": " +
			//                                         str(qty_out) + _(' already out')),
			//                                   quantity=qty_out,
			//                                   account_id=already_out_account_id)
			//
			//                if diff > 0:
			//                    debit_line['debit'] = diff
			//                    credit_line['credit'] = diff
			//                else:
			//                    # negative cost, reverse the entry
			//                    debit_line['credit'] = -diff
			//                    credit_line['debit'] = -diff
			//                AccountMoveLine.append([0, 0, debit_line])
			//                AccountMoveLine.append([0, 0, credit_line])
			//        return AccountMoveLine
		})
}
