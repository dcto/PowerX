package datadictionary

import (
	"PowerX/internal/model"
	"PowerX/internal/model/crm/customerDomain"
)

func defaultCustomerTypeDataDictionary() *model.DataDictionaryType {
	return &model.DataDictionaryType{
		Items: []*model.DataDictionaryItem{
			&model.DataDictionaryItem{
				Key:   customerDomain.CustomerPersonal,
				Type:  customerDomain.TypeCustomerType,
				Name:  "个人",
				Value: customerDomain.CustomerPersonal,
				Sort:  0,
			},
			&model.DataDictionaryItem{
				Key:   customerDomain.CustomerCompany,
				Type:  customerDomain.TypeCustomerType,
				Name:  "公司",
				Value: customerDomain.CustomerCompany,
				Sort:  0,
			},
		},
		Type:        customerDomain.TypeCustomerType,
		Name:        "客户类型",
		Description: "客户类型分个人，公司",
	}
}
