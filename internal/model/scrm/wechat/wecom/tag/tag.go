package tag

import (
	"PowerX/internal/model"
	"PowerX/internal/model/powerModel"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type WeComTag struct {
	powerModel.PowerModel

	WeComGroup *WeComTagGroup `gorm:"foreignKey:GroupId;references:group_id" json:"WeComGroup"`
	//
	IsSelf   int    `gorm:"comment:是否自建:1:平台创建:其他：微信创建;column:is_self" json:"is_self"`
	Type     int    `gorm:"comment:类型:1:企业标签2:策略标签;column:type" json:"type"`
	TagId    string `gorm:"comment:标签ID;column:tag_id;unique" json:"tag_id"`
	GroupId  string `gorm:"index:idx_group_id;not null;;comment:标签组ID;column:group_id" json:"group_id"`
	Name     string `gorm:"comment:标签名称;column:name" json:"name"`
	Sort     int    `gorm:"comment:排序;column:sort" json:"sort"`
	IsDelete bool   `gorm:"comment:是否删除;column:is_delete" json:"is_delete"`
}

func (mdl *WeComTag) TableName() string {
	return model.PowerXSchema + "." + model.TableNameWeComTag
}

func (mdl *WeComTag) GetTableName(needFull bool) string {
	tableName := model.TableNameWeComTag
	if needFull {
		tableName = mdl.TableName()
	}
	return tableName
}

// Query
//
//	@Description:
//	@receiver this
//	@param db
//	@return groups
//	@return err
func (e *WeComTag) Query(db *gorm.DB) (tags []*WeComTag) {

	err := db.Model(e).Where(`is_delete = ?`, false).Preload(`WeComGroup`).Order(`sort ASC`).Find(&tags).Error
	if err != nil {
		panic(err)
	}
	return tags

}

// Action
//
//	@Description:
//	@receiver this
//	@param db
//	@param group
//	@return []*WeComAppGroup
func (e *WeComTag) Action(db *gorm.DB, tags []*WeComTag) {

	err := db.Table(e.TableName()).Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "tag_id"}},
		DoUpdates: clause.AssignmentColumns([]string{`type`, `group_id`, `name`, `sort`, `is_delete`}),
	}).Create(&tags).Error
	if err != nil {
		panic(err)
	}

}

// FindOneByTagGroupId
//
//	@Description:
//	@receiver e
//	@param db
//	@param groupId
//	@return tags
func (e *WeComTag) FindOneByTagGroupId(db *gorm.DB, groupId string) (tags []*WeComTag) {

	err := db.Model(e).Where(`is_delete = ? AND group_id = ?`, false, groupId).Find(&tags).Error
	if err != nil {
		panic(err)
	}
	return tags

}

// FindOneByTagId
//
//	@Description:
//	@receiver e
//	@param db
//	@param tagId
//	@return tag
func (e *WeComTag) FindOneByTagId(db *gorm.DB, tagId string) (tag *WeComTag) {

	err := db.Model(e).Where(`is_delete = ? AND tag_id = ?`, false, tagId).Find(&tag).Error
	if err != nil {
		panic(err)
	}
	return tag

}

// Delete
//
//	@Description:
//	@receiver e
//	@param db
//	@param groupIds
//	@param tagIds
//	@return error
func (e *WeComTag) Delete(db *gorm.DB, groupIds, tagIds []string) error {

	query := db.Model(e)
	if v := len(tagIds); v > 0 {
		query = query.Where(`tag_id IN ?`, tagIds)
	}
	if v := len(groupIds); v > 0 {
		query = query.Where(`group_id IN ?`, groupIds)
	}
	column := make(map[string]interface{})
	column[`is_delete`] = true
	column[`deleted_at`] = time.Now()
	err := query.
		//Debug().
		UpdateColumns(&column).Error
	if err != nil {
		panic(err)
	}
	return err

}
