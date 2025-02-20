package category

import (
	infoOrganizatoin "PowerX/internal/model/infoOrganization"
	"PowerX/internal/types"
)

func TransformCategoryRequestToCategory(req *types.CreateCategoryRequest) *infoOrganizatoin.Category {
	return &infoOrganizatoin.Category{
		PId:          req.PId,
		Name:         req.Name,
		Scene:        req.Scene,
		CustomerId:   req.CustomerId,
		Sort:         req.Sort,
		ViceName:     req.ViceName,
		Description:  req.Description,
		CoverImageId: req.CoverImageId,
	}
}
