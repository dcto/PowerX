package category

import (
	"PowerX/internal/logic/admin/mediaResource"
	infoOrganization "PowerX/internal/model/infoOrganization"
	"PowerX/internal/types"
)

func TransformCategoryToReplyForWeb(category *infoOrganization.Category) *types.Category {

	node := &types.Category{
		Id:          category.Id,
		PId:         category.PId,
		Name:        category.Name,
		Sort:        category.Sort,
		ViceName:    category.ViceName,
		Description: category.Description,
		CreatedAt:   category.CreatedAt.String(),
		CoverImage:  mediaResource.TransformMediaResourceToReply(category.CoverImage),
		ImageAbleInfo: types.ImageAbleInfo{
			Icon:            category.Icon,
			BackgroundColor: category.BackgroundColor,
		},
		Children: nil,
	}
	if len(category.Children) > 0 {
		node.Children = TransformCategoriesToReplyForWeb(category.Children)

	}

	return node
}

func TransformCategoriesToReplyForWeb(productCategoryList []*infoOrganization.Category) []*types.Category {
	uniqueIds := make(map[int64]bool)
	var productCategoryReplyList []*types.Category
	for _, c := range productCategoryList {
		if !uniqueIds[c.Id] {
			node := TransformCategoryToReplyForWeb(c)
			if len(c.Children) > 0 {
				node.Children = TransformCategoriesToReplyForWeb(c.Children)
			}

			productCategoryReplyList = append(productCategoryReplyList, node)
			uniqueIds[c.Id] = true

		}
	}

	return productCategoryReplyList
}
