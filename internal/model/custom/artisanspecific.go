package custom

import (
	"PowerX/internal/model"
	"PowerX/internal/model/powerModel"
)

type ArtisanSpecific struct {
	powerModel.PowerModel

	ArtisanId int64 `gorm:"comment:ArtisanId"`
}

const ArtisanSpecificUniqueId = powerModel.UniqueId

func (mdl *ArtisanSpecific) TableName() string {
	return model.PowerXSchema + "." + model.TableNameArtisanSpecific
}

func (mdl *ArtisanSpecific) GetTableName(needFull bool) string {
	tableName := model.TableNameArtisanSpecific
	if needFull {
		tableName = mdl.TableName()
	}
	return tableName
}
