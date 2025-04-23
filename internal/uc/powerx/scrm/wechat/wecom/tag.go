package wecom

import (
	"PowerX/internal/model/powerModel"
	"PowerX/internal/model/scrm/wechat/wecom/customer"
	tag2 "PowerX/internal/model/scrm/wechat/wecom/tag"
	"PowerX/internal/types"
	baseResp "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/tag/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/tag/response"
	"strings"
	"time"
)

// FindListWeComTagPage
//
//	@Description:
//	@receiver this
//	@param option
//	@return reply
//	@return err
func (uc *WeComUseCase) FindListWeComTagGroupOption() (reply []*tag2.WeComTagGroup, err error) {

	reply = uc.modelWeComTag.group.Query(uc.db)

	return reply, err

}

// FindListWeComTagGroupPage
//
//	@Description:
//	@receiver this
//	@param option
//	@return reply
//	@return err
func (uc *WeComUseCase) FindListWeComTagGroupPage(option *types.PageOption[types.ListWeComTagGroupPageRequest]) (reply *types.Page[*tag2.WeComTagGroup], err error) {

	var tagGroups []*tag2.WeComTagGroup
	var count int64
	query := uc.db.WithContext(uc.ctx).Model(tag2.WeComTagGroup{}).Where(`is_delete = ?`, false)

	if v := option.Option.GroupId; v != `` {
		query.Where(`group_id = ?`, v)
	}

	if v := option.Option.GroupName; v != `` {
		query.Where(`name like ?`, "%"+v+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, err
	}
	if option.PageIndex != 0 && option.PageSize != 0 {
		query.Offset((option.PageIndex - 1) * option.PageSize).Limit(option.PageSize)
	}

	err = query.Preload(`WeComGroupTags`).Find(&tagGroups).Error

	return &types.Page[*tag2.WeComTagGroup]{
		List:      tagGroups,
		PageIndex: option.PageIndex,
		PageSize:  option.PageSize,
		Total:     count,
	}, err

}

// FindListWeComTagOption
//
//	@Description:
//	@receiver this
//	@return reply
//	@return err
func (uc *WeComUseCase) FindListWeComTagOption() (reply []*tag2.WeComTag, err error) {

	reply = uc.modelWeComTag.tag.Query(uc.db)

	return reply, err

}

// FindListWeComTagPage
//
//	@Description:
//	@receiver this
//	@param option
//	@return reply
//	@return err
func (uc *WeComUseCase) FindListWeComTagPage(option *types.PageOption[types.ListWeComTagReqeust]) (reply *types.Page[*tag2.WeComTag], err error) {

	var tags []*tag2.WeComTag
	var count int64
	query := uc.db.WithContext(uc.ctx).
		//Debug().
		Model(tag2.WeComTag{}).Where(`is_delete = ?`, false)

	if v := option.Option.TagIds; len(v) > 0 {
		query.Where(`tag_id in ?`, v)
	}

	if v := option.Option.GroupIds; len(v) > 0 {
		query.Where(`group_id in ?`, v)
	}
	if v := option.Option.Name; v != `` {
		query.Where(`name like ?`, "%"+v+"%")
	}
	if err := query.Count(&count).Error; err != nil {
		return nil, err
	}
	if option.PageIndex != 0 && option.PageSize != 0 {
		query.Offset((option.PageIndex - 1) * option.PageSize).Limit(option.PageSize)
	}

	err = query.Preload(`WeComGroup`).Find(&tags).Error

	return &types.Page[*tag2.WeComTag]{
		List:      tags,
		PageIndex: option.PageIndex,
		PageSize:  option.PageSize,
		Total:     count,
	}, err

}

// PullListWeComCorpTagRequest
//
//	@Description:
//	@receiver this
//	@param tagIds
//	@param groupIds
//	@param sync
//	@return reply
//	@return err
func (uc *WeComUseCase) PullListWeComCorpTagRequest(tagIds []string, groupIds []string, sync int) (reply *response.ResponseTagGetCorpTagList, err error) {

	reply, err = uc.Client.ExternalContactTag.GetCorpTagList(uc.ctx, tagIds, groupIds)
	if err != nil {
		panic(err)
	} else {
		err = uc.help.error(`scrm.pull.wecom.crop.tag.error`, reply.ResponseWork)
	}

	if err == nil && sync > 0 {
		// sync to local
		groups, tags := uc.transferWeComToModel(reply.TagGroups, nil, 0)
		if groups != nil {
			uc.modelWeComTag.group.Action(uc.db, groups)
		}
		if tags != nil {
			uc.modelWeComTag.tag.Action(uc.db, tags)
		}

	}

	return reply, err

}

// transferWeComToModel
//
//	@Description:
//	@receiver this
//	@param data
//	@param agentId
//	@return groups
//	@return tags
func (uc *WeComUseCase) transferWeComToModel(data []*response.CorpTagGroup, agentId *int64, isSelf int) (groups []*tag2.WeComTagGroup, tags []*tag2.WeComTag) {

	if data != nil {
		for _, val := range data {
			groups = append(groups, &tag2.WeComTagGroup{
				PowerModel: powerModel.PowerModel{
					CreatedAt: time.Unix(int64(val.CreateTime), 0),
				},
				//AgentId:  int(*agentId),
				GroupId:  val.GroupID,
				Name:     val.GroupName,
				Sort:     val.Order,
				IsDelete: val.Deleted,
			})
			if val.Tags != nil {
				for _, value := range val.Tags {
					tags = append(tags, &tag2.WeComTag{
						PowerModel: powerModel.PowerModel{
							CreatedAt: time.Unix(int64(value.CreateTime), 0),
						},
						Type:     1,
						IsSelf:   isSelf,
						TagId:    value.ID,
						GroupId:  val.GroupID,
						Name:     value.Name,
						Sort:     value.Order,
						IsDelete: value.Deleted,
					})
				}

			}
		}
	}
	return groups, tags

}

// PullListWeComStrategyTagRequest
//
//	@Description:
//	@receiver this
//	@param options
//	@return reply
//	@return err
func (uc *WeComUseCase) PullListWeComStrategyTagRequest(options *request.RequestTagGetStrategyTagList) (reply *response.ResponseTagGetStrategyTagList, err error) {

	reply, err = uc.Client.ExternalContactTag.GetStrategyTagList(uc.ctx, options)
	if err != nil {
		panic(err)
	} else {
		err = uc.help.error(`scrm.pull.wecom.strategy.tag.error`, reply.ResponseWork)
	}
	return reply, err

}

// ActionWeComCorpTagGroupRequest
//
//	@Description:
//	@receiver this
//	@param options
//	@return work
//	@return error
func (uc *WeComUseCase) ActionWeComCorpTagGroupRequest(options *types.ActionCorpTagGroupRequest) (work *baseResp.ResponseWork, err error) {

	//tags := uc.modelWeComTag.tag.FindOneByTagGroupId(uc.db, *options.GroupId)
	var addTagGroup []request.RequestTagAddCorpTagFieldTag
	var delTag []string

	for _, newTag := range options.Tags {
		if newTag.TagId == `` {
			addTagGroup = append(addTagGroup, request.RequestTagAddCorpTagFieldTag{
				Name: newTag.TagName,
			})
		} else {
			delTag = append(delTag, newTag.TagId)
		}
	}
	if delTag != nil {
		work, err = uc.DeleteWeComCorpTagRequest(&request.RequestTagDelCorpTag{TagID: delTag})
	}
	if len(addTagGroup) > 0 {
		add, er := uc.CreateWeComCorpTagRequest(&request.RequestTagAddCorpTag{
			GroupID:   options.GroupId,
			GroupName: options.GroupName,
			Tag:       addTagGroup,
			AgentID:   options.AgentId,
		})
		err = er
		work = &add.ResponseWork
	}

	return work, err

}

// CreateWeComCorpTagRequest
//
//	@Description:
//	@receiver this
//	@param options
//	@return *response.ResponseTagAddCorpTag
//	@return error
func (uc *WeComUseCase) CreateWeComCorpTagRequest(options *request.RequestTagAddCorpTag) (*response.ResponseTagAddCorpTag, error) {

	corpTag, err := uc.Client.ExternalContactTag.AddCorpTag(uc.ctx, options)
	if err != nil {
		panic(err)
	} else {
		err = uc.help.error(`scrm.create.wecom.corp.tag.error`, corpTag.ResponseWork)
	}

	if err == nil {
		groups, tags := uc.transferWeComToModel([]*response.CorpTagGroup{corpTag.TagGroups}, options.AgentID, 1)
		if groups != nil {
			uc.modelWeComTag.group.Action(uc.db, groups)
		}
		if tags != nil {
			uc.modelWeComTag.tag.Action(uc.db, tags)
		}
	}

	return corpTag, err

}

// UpdateWeComCorpTagRequest
//
//	@Description:
//	@receiver this
//	@param options
//	@return *baseResp.ResponseWork
//	@return error
func (uc *WeComUseCase) UpdateWeComCorpTagRequest(options *request.RequestTagEditCorpTag) (*baseResp.ResponseWork, error) {

	corpTag, err := uc.Client.ExternalContactTag.EditCorpTag(uc.ctx, options)

	if err != nil {
		panic(err)
	} else {
		err = uc.help.error(`scrm.update.wecom.corp.tag.error`, *corpTag)
	}
	if err == nil {
		info := uc.modelWeComTag.tag.FindOneByTagId(uc.db, options.ID)
		if info != nil {
			info.Name = options.Name
			info.Sort = options.Order
			uc.modelWeComTag.tag.Action(uc.db, []*tag2.WeComTag{info})
		}
	}

	return corpTag, err

}

// DeleteWeComCorpTagRequest
//
//	@Description:
//	@receiver this
//	@param options
//	@return *baseResp.ResponseWork
//	@return error
func (uc *WeComUseCase) DeleteWeComCorpTagRequest(options *request.RequestTagDelCorpTag) (*baseResp.ResponseWork, error) {

	corpTag, err := uc.Client.ExternalContactTag.DelCorpTag(uc.ctx, options)
	if err != nil {
		panic(err)
	} else {
		err = uc.help.error(`scrm.delete.wecom.corp.tag.error`, *corpTag)
	}

	err = uc.modelWeComTag.tag.Delete(uc.db, options.GroupID, options.TagID)

	return corpTag, err

}

// ActionWeComCustomerTagRequest
//
//	@Description:
//	@receiver this
//	@param options
//	@return *baseResp.ResponseWork
//	@return error
func (uc *WeComUseCase) ActionWeComCustomerTagRequest(option *request.RequestTagMarkTag) (*baseResp.ResponseWork, error) {

	customerTag, err := uc.Client.ExternalContactTag.MarkTag(uc.ctx, option)
	if err != nil {
		panic(err)
	} else {
		err = uc.help.error(`scrm.update.wecom.customer.tag.error`, *customerTag)
	}

	if err == nil {
		uc.updateCustomerFolowTagIds(option)

	}

	return customerTag, err

}

// updateCustomerFolowTagIds
//
//	@Description:
//	@receiver this
//	@param option
func (uc *WeComUseCase) updateCustomerFolowTagIds(option *request.RequestTagMarkTag) {

	follow := uc.modelWeComCustomer.follow.FindFollowByExternalUserId(uc.db, option.ExternalUserID)
	column := make(map[string]string)
	if follow.TagIds != `` {
		for _, val := range strings.Split(follow.TagIds, `,`) {
			column[val] = val
		}
	}

	if option.AddTag != nil {
		for _, val := range option.AddTag {
			if _, ok := column[val]; !ok {
				column[val] = val
			}
		}
	}
	if option.RemoveTag != nil {
		for _, val := range option.RemoveTag {
			if _, ok := column[val]; ok {
				delete(column, val)
			}
		}
	}
	if column != nil {
		var tagIds []string
		for _, val := range column {
			tagIds = append(tagIds, val)
		}
		follow.TagIds = strings.Join(tagIds, `,`)
		uc.modelWeComCustomer.follow.Action(uc.db, []*customer.WeComExternalContactFollow{follow})
	}

}
