package resource

import (
	"PowerX/internal/model"
	"PowerX/internal/model/powerModel"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WeComResource struct {
	powerModel.PowerModel

	Url          string `gorm:"comment:微信地址;column:url" json:"url"`
	FileName     string `gorm:"unique;comment:文件名;column:file_name" json:"file_name"`
	Remark       string `gorm:"comment:备注;column:remark" json:"remark"`
	BucketName   string `gorm:"comment:桶;column:bucket_name" json:"bucket_name"`
	Size         int    `gorm:"comment:大小;column:size" json:"size"`
	ResourceType string `gorm:"comment:资源类型：image,voice,file, video, other;column:resource_type" json:"resource_type"`
}

func (mdl *WeComResource) TableName() string {
	return model.PowerXSchema + "." + model.TableNameWeComResource
}

func (mdl *WeComResource) GetTableName(needFull bool) string {
	tableName := model.TableNameWeComResource
	if needFull {
		tableName = mdl.TableName()
	}
	return tableName
}

// Query
//
//	@Description:
//	@receiver e
//	@param db
//	@return departments
func (e WeComResource) Query(db *gorm.DB) (resources []*WeComResource) {

	err := db.Model(e).Find(&resources).Error
	if err != nil {
		panic(err)
	}
	return resources

}

// Action
//
//	@Description:
//	@receiver e
//	@param db
//	@param contacts
func (e *WeComResource) Action(db *gorm.DB, resources []*WeComResource) {

	err := db.Table(e.TableName()).Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "file_name"}}, UpdateAll: true}).CreateInBatches(&resources, 100).Error
	if err != nil {
		panic(err)
	}

}
