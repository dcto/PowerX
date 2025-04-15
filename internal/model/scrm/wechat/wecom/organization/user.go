package organization

import (
	"PowerX/internal/model"
	"PowerX/internal/model/powerModel"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WeComUser struct {
	powerModel.PowerModel

	WeComUserId           string `gorm:"comment:员工ID;column:wecom_user_id;unique" json:"wecom_user_id"`
	Name                  string `gorm:"comment:员工名称;column:name" json:"name"`
	Position              string `gorm:"comment:员工位置;column:position" json:"position"`
	Mobile                string `gorm:"comment:员工电话;column:mobile" json:"mobile"`
	Gender                string `gorm:"comment:员工性别;column:gender" json:"gender"`
	Email                 string `gorm:"comment:邮箱;column:email" json:"email"`
	BizMail               string `gorm:"comment:商务邮箱;column:biz_mail" json:"biz_mail"`
	Avatar                string `gorm:"comment:头像;column:avatar" json:"avatar"`
	ThumbAvatar           string `gorm:"comment:ThumbAvatar;column:thumb_avatar" json:"thumb_avatar"`
	Telephone             string `gorm:"comment:电话;column:telephone" json:"telephone"`
	Alias                 string `gorm:"comment:别称;column:alias" json:"alias"`
	Address               string `gorm:"comment:地址;column:address" json:"address"`
	OpenUserId            string `gorm:"comment:开放ID;column:open_user_id" json:"open_user_id"`
	WeComMainDepartmentId int    `gorm:"comment:部门ID;column:wecom_main_department_id" json:"wecom_main_department_id"`
	Status                int    `gorm:"comment:状态;column:status" json:"status"`
	QrCode                string `gorm:"comment:二维码;column:qr_code" json:"qr_code"`
	Department            string `gorm:"comment:部门;column:department" json:"department"`
	RefUserId             int64  `gorm:"comment:RefUserId;column:ref_user_id" json:"ref_user_id"`
}

func (mdl *WeComUser) TableName() string {
	return model.PowerXSchema + "." + model.TableNameWeComUser
}

func (mdl *WeComUser) GetTableName(needFull bool) string {
	tableName := model.TableNameWeComUser
	if needFull {
		tableName = mdl.TableName()
	}
	return tableName
}

type (
	AdapterUserSliceUserIDs func(user []*WeComUser) (ids []string)
)

// Query
//
//	@Description:
//	@receiver this
//	@param db
//	@return users
func (e WeComUser) Query(db *gorm.DB) (users []*WeComUser) {

	err := db.Model(e).Find(&users).Error
	if err != nil {
		panic(err)
	}
	return users

}

// Action
//
//	@Description:
//	@receiver e
//	@param db
//	@param users
func (e WeComUser) Action(db *gorm.DB, users []*WeComUser) {

	err := db.Table(e.TableName()).
		//Debug().
		Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "wecom_user_id"}}, UpdateAll: true}).CreateInBatches(&users, 100).Error
	if err != nil {
		panic(err)
	}

}
