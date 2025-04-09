package category

import (
	infoOrganization "PowerX/internal/model/infoOrganization"
	"PowerX/internal/types"
)

func TransformCategoryRequestToCategory(req *types.CreateCategoryRequest) *infoOrganization.Category {
	return &infoOrganization.Category{
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
