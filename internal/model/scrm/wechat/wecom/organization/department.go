package organization

import (
	"PowerX/internal/model"
	"PowerX/internal/model/powerModel"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WeComDepartment struct {
	powerModel.PowerModel

	// Leader         *WeComUser `gorm:"foreignKey:LeaderId"`
	WeComDepId       int                `gorm:"comment:部门ID;column:wecom_dep_id;unique" json:"wecom_dep_id"`
	Name             string             `gorm:"comment:部门名称;column:name" json:"name"`
	NameEn           string             `gorm:"comment:部门英文名称;column:name_en" json:"name_en"`
	WeComParentId    int                `gorm:"comment:上级部门ID;column:wecom_parent_id" json:"wecom_parent_id"`
	Order            int                `gorm:"comment:Order;column:order" json:"order"`
	DepartmentLeader string             `gorm:"comment:部门Leader;column:department_leader" json:"department_leader"`
	RefDepartmentId  int64              `gorm:"comment:-;column:ref_department_id" json:"ref_department_id"`
	Children         []*WeComDepartment `gorm:"-" json:"children,omitempty"`
}

func (mdl *WeComDepartment) TableName() string {
	return model.PowerXSchema + "." + model.TableNameWeComDepartment
}

func (mdl *WeComDepartment) GetTableName(needFull bool) string {
	tableName := model.TableNameWeComDepartment
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
func (e WeComDepartment) Query(db *gorm.DB) (departments []*WeComDepartment) {

	err := db.Model(e).Find(&departments).Error
	if err != nil {
		panic(err)
	}
	return departments

}

// Action
//
//	@Description:
//	@receiver e
//	@param db
//	@param contacts
func (e *WeComDepartment) Action(db *gorm.DB, contacts []*WeComDepartment) {

	err := db.Table(e.TableName()).Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "wecom_dep_id"}}, UpdateAll: true}).CreateInBatches(&contacts, 100).Error
	if err != nil {
		panic(err)
	}

}
