package migrate

import (
	"PowerX/cmd/ctl/database/custom/migrate"
	migratePro "PowerX/cmd/ctl/database/pro/migrate"
	"PowerX/internal/config"
	"PowerX/internal/model"
	"PowerX/internal/model/crm/customerDomain"
	"PowerX/internal/model/crm/market"
	"PowerX/internal/model/crm/operation"
	"PowerX/internal/model/crm/product"
	"PowerX/internal/model/crm/trade"
	infoOrganization2 "PowerX/internal/model/infoOrganization"
	"PowerX/internal/model/media"
	"PowerX/internal/model/organization"
	"PowerX/internal/model/permission"
	"PowerX/internal/model/scene"
	"PowerX/internal/model/scrm/wechat/wecom/app"
	customer2 "PowerX/internal/model/scrm/wechat/wecom/customer"
	organization3 "PowerX/internal/model/scrm/wechat/wecom/organization"
	"PowerX/internal/model/scrm/wechat/wecom/resource"
	tag2 "PowerX/internal/model/scrm/wechat/wecom/tag"
	"PowerX/internal/model/wechat"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PowerMigrator struct {
	db *gorm.DB
}

func NewPowerMigrator(conf *config.Config) (*PowerMigrator, error) {
	var dsn gorm.Dialector
	switch conf.PowerXDatabase.Driver {
	case config.DriverMysql:
		dsn = mysql.Open(conf.PowerXDatabase.DSN)
	case config.DriverPostgres:
		dsn = postgres.Open(conf.PowerXDatabase.DSN)
	}
	db, err := gorm.Open(dsn, &gorm.Config{
		//Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	return &PowerMigrator{
		db: db,
	}, err
}

func (m *PowerMigrator) InitSchema(schema string) error {
	// 检查 schema 是否存在
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM information_schema.schemata WHERE schema_name = ?);`
	err := m.db.Raw(query, schema).Scan(&exists).Error
	if err != nil {
		return fmt.Errorf("failed to check schema existence: %w", err)
	}

	// 如果 schema 不存在，创建它
	if !exists {
		createSchemaQuery := fmt.Sprintf("CREATE SCHEMA %s;", schema)
		err := m.db.Exec(createSchemaQuery).Error
		if err != nil {
			return fmt.Errorf("failed to create schema: %w", err)
		}
		fmt.Printf("Schema %s created successfully.\n", schema)
	} else {
		fmt.Printf("Schema %s already exists.\n", schema)
	}

	return nil
}

func (m *PowerMigrator) AutoMigrate() {

	_ = m.db.AutoMigrate(&model.DataDictionaryType{}, &model.DataDictionaryItem{}, &model.PivotDataDictionaryToObject{})
	_ = m.db.AutoMigrate(&organization.Department{}, &organization.User{}, &organization.Position{})
	_ = m.db.AutoMigrate(&permission.UserCasbinPolicy{}, permission.AdminRole{}, permission.AdminRoleMenuName{}, permission.AdminAPI{})

	// info organization
	_ = m.db.AutoMigrate(&infoOrganization2.Category{}, &infoOrganization2.Label{}, &infoOrganization2.Tag{})
	_ = m.db.AutoMigrate(&infoOrganization2.PivotCategoryToObject{})

	// customer domain
	_ = m.db.AutoMigrate(
		&customerDomain.Lead{}, &customerDomain.Contact{}, customerDomain.RegisterCode{},
		&customerDomain.Customer{},
	)
	// operation
	_ = m.db.AutoMigrate(
		&operation.Membership{}, &operation.TicketRecord{},
	)
	_ = m.db.AutoMigrate(&wechat.WechatOACustomer{}, &wechat.WechatMPCustomer{})
	_ = m.db.AutoMigrate(
		&product.PivotProductToProductCategory{},
	)
	// product
	_ = m.db.AutoMigrate(&product.Product{}, &product.ProductCategory{})
	_ = m.db.AutoMigrate(&product.ProductSpecific{}, &product.SpecificOption{}, &product.ProductStatistics{})
	_ = m.db.AutoMigrate(&product.SKU{}, &product.PivotSkuToSpecificOption{})
	_ = m.db.AutoMigrate(&product.PriceBookEntry{}, &product.PriceBook{}, &product.PriceConfig{})
	_ = m.db.AutoMigrate(&product.PivotStoreToArtisan{}, &market.Store{}, &product.Artisan{})

	// market
	_ = m.db.AutoMigrate(&market.Media{})
	_ = m.db.AutoMigrate(&market.MGMRule{}, market.InviteRecord{}, market.CommissionRecord{})

	// media
	_ = m.db.AutoMigrate(&media.MediaResource{}, &media.PivotMediaResourceToObject{})

	// trade
	_ = m.db.AutoMigrate(&trade.ShippingAddress{}, &trade.DeliveryAddress{}, &trade.BillingAddress{})
	_ = m.db.AutoMigrate(&trade.Warehouse{}, &trade.Inventory{}, &trade.Logistics{})
	_ = m.db.AutoMigrate(&trade.Cart{}, &trade.CartItem{}, &trade.Order{}, &trade.OrderItem{})
	_ = m.db.AutoMigrate(&trade.OrderStatusTransition{}, &trade.PivotOrderToInventoryLog{})
	_ = m.db.AutoMigrate(&trade.Payment{}, &trade.PaymentItem{})
	_ = m.db.AutoMigrate(&trade.RefundOrder{}, &trade.RefundOrderItem{})
	_ = m.db.AutoMigrate(&trade.TokenBalance{},
		&trade.TokenExchangeRatio{}, &trade.TokenExchangeRecord{},
		trade.TokenReservation{}, trade.TokenTransaction{},
	)

	// pro
	migratePro.AutoMigratePro(m.db)

	// custom
	migrate.AutoMigrateCustom(m.db)

	// wechat organization
	_ = m.db.AutoMigrate(&organization3.WeComUser{}, &organization3.WeComDepartment{})
	// wechat customer
	_ = m.db.AutoMigrate(&customer2.WeComExternalContact{}, &customer2.WeComExternalContactFollow{})
	// wechat resource
	_ = m.db.AutoMigrate(&resource.WeComResource{})
	// wechat app
	_ = m.db.AutoMigrate(&app.WeComAppGroup{})
	// wechat tag
	_ = m.db.AutoMigrate(&tag2.WeComTag{}, &tag2.WeComTagGroup{})
	// qrcode
	_ = m.db.AutoMigrate(&scene.SceneQRCode{})
}
