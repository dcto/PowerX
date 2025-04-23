package app

import (
	"PowerX/internal/model"
	"PowerX/internal/model/powerModel"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WeComAppGroup struct {
	powerModel.PowerModel

	Name     string `gorm:"comment:群名称;column:name" json:"name"`
	Owner    string `gorm:"comment:群主;column:owner" json:"owner"`
	UserList string `gorm:"comment:群用户;column:user_list" json:"user_list"`
	ChatID   string `gorm:"comment:群ID;unique"`
}

func (mdl *WeComAppGroup) TableName() string {
	return model.PowerXSchema + "." + model.TableNameWeComAppGroup
}

func (mdl *WeComAppGroup) GetTableName(needFull bool) string {
	tableName := model.TableNameWeComAppGroup
	if needFull {
		tableName = mdl.TableName()
	}
	return tableName
}

type (
	AdapterGroupSliceChatIDs func(groups []*WeComAppGroup) (ids []string)
)

// Query
//
//	@Description:
//	@receiver this
//	@param db
//	@return groups
//	@return err
func (e *WeComAppGroup) Query(db *gorm.DB) (groups []*WeComAppGroup) {

	err := db.Model(e).Find(&groups).Error
	if err != nil {
		panic(err)
	}
	return groups

}

// Action
//
//	@Description:
//	@receiver this
//	@param db
//	@param group
//	@return []*WeComAppGroup
func (e *WeComAppGroup) Action(db *gorm.DB, group []*WeComAppGroup) {

	err := db.Table(e.TableName()).Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "chat_id"}}, UpdateAll: true}).Create(&group).Error
	if err != nil {
		panic(err)
	}

}
