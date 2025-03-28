package trade

import (
	"PowerX/internal/model"
	"PowerX/internal/model/powerModel"
)

type PivotProductToPromotion struct {
	*powerModel.PowerPivot

	PromotionRuleId int64 `gorm:"comment:促销规则d; not null;index:idx_promotion_rule_id" json:"promotionRuleId"`
	ProductId       int64 `gorm:"comment:商品Id; not null;index:idx_product_id" json:"productId"`
}

func (mdl *PivotProductToPromotion) TableName() string {
	return model.PowerXSchema + "." + model.TableNamePivotProductToPromotion
}

func (mdl *PivotProductToPromotion) GetTableName(needFull bool) string {
	tableName := model.TableNamePivotProductToPromotion
	if needFull {
		tableName = mdl.TableName()
	}
	return tableName
}
