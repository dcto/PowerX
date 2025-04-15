package wecom

import (
	"PowerX/internal/model/scene"
	"PowerX/internal/model/scrm/wechat/wecom/customer"
	organization2 "PowerX/internal/model/scrm/wechat/wecom/organization"
	"PowerX/internal/model/scrm/wechat/wecom/resource"
	tag2 "PowerX/internal/model/scrm/wechat/wecom/tag"
	"PowerX/internal/types"
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	kresp "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	agentResp "github.com/ArtisanCloud/PowerWeChat/v3/src/work/agent/response"
	customerGroupReq "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/groupChat/request"
	customerGroupResp "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/groupChat/response"
	creq "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/messageTemplate/request"
	crsp "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/messageTemplate/response"
	customerResp "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/response"
	tagReq "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/tag/request"
	tagResp "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/tag/response"
	botReq "github.com/ArtisanCloud/PowerWeChat/v3/src/work/groupRobot/request"
	botResp "github.com/ArtisanCloud/PowerWeChat/v3/src/work/groupRobot/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/message/appChat/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/message/appChat/response"
	appReq "github.com/ArtisanCloud/PowerWeChat/v3/src/work/message/request"
	appResp "github.com/ArtisanCloud/PowerWeChat/v3/src/work/message/response"
	"mime/multipart"
)

type IWechatInterface interface {
	//
	//  @Description:  Department
	//
	iWeComDepartmentInterface

	//
	//  @Description: User
	//
	iWeComUserInterface

	//
	//  @Description: Customer
	//
	iWeComCustomerInterface

	//
	//  @Description: App
	//
	iWeComAppInterface

	//
	//  @Description:  Bot
	//
	iWeComBotInterface

	//
	//  @Description: invoke
	//
	iMakeInvokeInterface

	//
	//  @Description: common
	//
	iCommonInterface

	//
	//  @Description: qrcode
	//
	iQRCodeInterface
	//
	//  @Description: tag
	//
	iTagInterface
}

// iWeComDepartmentInterface
// @Description: 部门
type iWeComDepartmentInterface interface {
	//
	// CreateWechatDepartment
	//  @Description: 创建部门
	//  @param ctx
	//  @param dep
	//  @return err
	//
	//CreateWechatDepartment(ctx context.Context, dep *organization.WeComDepartment) (err error)

	//
	// FindManyWeComDepartmentsPage
	//  @Description: 查询部门
	//  @param ctx
	//  @param option
	//  @return *types.Page[*organization.Department]
	//
	FindManyWeComDepartmentsPage(ctx context.Context, option *types.PageOption[FindManyWeComDepartmentsOption]) (*types.Page[*organization2.WeComDepartment], error)
}

// iWeComUserInterface
// @Description: 员工
type iWeComUserInterface interface {
	//
	// PullSyncDepartmentsAndUsersRequest
	//  @Description: 同步组织架构
	//  @param ctx
	//  @return error
	//
	PullSyncDepartmentsAndUsersRequest(ctx context.Context) error
	//
	// FindManyWechatUsersPage
	//  @Description: 查询员工
	//  @param ctx
	//  @param opt
	//  @return *types.Page[*organization.WeComUser]
	//
	FindManyWechatUsersPage(ctx context.Context, opt *types.PageOption[FindManyWeComUsersOption]) (*types.Page[*organization2.WeComUser], error)
}

// iWeComCustomerInterface
// @Description: 客户
type iWeComCustomerInterface interface {
	//
	// PullListWeComCustomerRequest
	//  @Description: 拉取客户列表
	//  @param userID
	//  @return []*customerResp.ResponseExternalContact
	//  @return error
	//
	PullListWeComCustomerRequest(userID ...string) ([]*customerResp.ResponseExternalContact, error)
	//
	// PullListWeComCustomerGroupRequest
	//  @Description: 拉取客户群列表
	//  @param opt
	//  @return list
	//  @return err
	//
	PullListWeComCustomerGroupRequest(opt *customerGroupReq.RequestGroupChatList) (list []*customerGroupResp.ResponseGroupChatGet, err error)

	//
	// FindManyWeComCustomerPage
	//  @Description: 所有客户
	//  @param ctx
	//  @param opt
	//  @param sync
	//  @return *types.Page[*customer.WeComExternalContact]
	//  @return error
	//
	FindManyWeComCustomerPage(ctx context.Context, opt *types.PageOption[FindManyWeComCustomerOption], sync int) (*types.Page[*customer.WeComExternalContact], error)

	//
	// PushWoWorkCustomerTemplateRequest
	//  @Description: 发送客户群信息1/day
	//  @param opt
	//  @param sendTime
	//  @return *crsp.ResponseAddMessageTemplate
	//  @return error
	//
	PushWoWorkCustomerTemplateRequest(opt *creq.RequestAddMsgTemplate, sendTime int64) (*crsp.ResponseAddMessageTemplate, error)
}

// iWeComBotInterface
// @Description:
type iWeComBotInterface interface {
	//
	// PushWeComBotArticlesRequest
	//  @Description: 机器人发送图文
	//  @param key
	//  @param articles
	//  @return resp
	//  @return error
	//
	PushWeComBotArticlesRequest(key string, articles []*botReq.GroupRobotMsgNewsArticles) (resp *botResp.ResponseGroupRobotSend, err error)
}

// iWeComAppInterface
// @Description:
type iWeComAppInterface interface {
	//
	// PullDetailWeComAppRequest
	//  @Description: 应用详情
	//  @param agentID
	//  @return reply
	//  @return err
	//
	PullDetailWeComAppRequest(agentID int) (reply *agentResp.ResponseAgentGet, err error)

	//
	// PullListWeComAppRequest
	//  @Description: 应用列表
	//  @return reply
	//  @return err
	//
	PullListWeComAppRequest() (reply *agentResp.ResponseAgentList, err error)

	//
	// PushAppWeComMessageArticlesRequest
	//  @Description: 发送应用图文信息
	//  @param opt
	//  @return *appResp.ResponseMessageSend
	//  @return error
	//
	PushAppWeComMessageArticlesRequest(opt *appReq.RequestMessageSendNews, sendTime int64) (reply *appResp.ResponseMessageSend, err error)

	iWeComAppGroupInterface
}

type iWeComAppGroupInterface interface {
	//
	// PullListWeComAppGroupRequest
	//  @Description: 获取应用群聊
	//  @param chatID
	//  @return reply
	//  @return err
	//
	PullListWeComAppGroupRequest(chatIDs ...string) (replys []*power.HashMap, err error)
	//
	// AppWechatGroupCreate
	//  @Description: 创建应用群聊
	//  @param option
	//  @return reply
	//  @return err
	//
	CreateWeComAppGroupRequest(option *request.RequestAppChatCreate) (reply *response.ResponseAppChatCreate, err error)

	//
	// PushAppWeComGroupMessageArticlesRequest
	//  @Description: 群内推送
	//  @param messages
	//  @return resp
	//  @return err
	//
	PushAppWeComGroupMessageArticlesRequest(messages *power.HashMap, sendTime int64) (resp *kresp.ResponseWork, err error)
}

// iMakeInvokeInterface
// @Description: 消费信息
type iMakeInvokeInterface interface {
	//
	// InvokeTimerMessageGrabUniteSend
	//  @Description:
	//  @param ttp
	//  @param sendTime
	//  @return count
	//
	InvokeTimerMessageGrabUniteSend(ttp TimerTypeByte, sendTime int64) error
}

// iCommonInterface
// @Description:
type iCommonInterface interface {
	//
	// UploadImageToResourceRequest
	//  @Description: 上传图片到微信
	//  @param path
	//  @param handle
	//  @return link
	//  @return err
	//
	UploadImageToResourceRequest(path string, handle *multipart.FileHeader) (link string, err error)

	//
	// FindWeComResourceListFromLocalPage
	//  @Description: FindWeComResourceListFromLocalPage
	//  @param opt
	//  @return *types.Page[*resource.WeComResource]
	//  @return error
	//
	FindWeComResourceListFromLocalPage(opt *types.ListWeComResourceImageRequest) (*types.Page[*resource.WeComResource], error)
}

// iQRCodeInterface
// @Description: 活码
type iQRCodeInterface interface {

	//
	// CreateWeComCustomerGroupQRCodeRequest
	//  @Description: 创建群活码
	//  @param opt
	//  @return err
	//
	CreateWeComCustomerGroupQRCodeRequest(opt *types.QRCodeActiveRequest) (err error)
	//
	// UpdateWeComCustomerGroupQRCodeRequest
	//  @Description: 更新群活码
	//  @param opt
	//  @return err
	//
	UpdateWeComCustomerGroupQRCodeRequest(opt *types.QRCodeActiveRequest) (err error)
	//
	// FindWeComCustomerGroupQRCodePage
	//  @Description: 客户群活码
	//  @param opt
	//  @return reply
	//  @return err
	//
	FindWeComCustomerGroupQRCodePage(opt *types.PageOption[types.ListWeComGroupQRCodeActiveReqeust]) (reply *types.Page[*scene.SceneQRCode], err error)

	//
	// ActionCustomerGroupQRCode
	//  @Description: 启用，禁用，删除
	//  @param qid
	//  @param action
	//  @return error
	//
	ActionCustomerGroupQRCode(qid string, action int) error

	//
	// UpdateSceneQRCodeLink
	//  @Description: 更新场景码
	//  @param qid
	//  @param link
	//  @return error
	//
	UpdateSceneQRCodeLink(qid string, link string) error
}

// iTagInterface
// @Description: TAG
type iTagInterface interface {
	//
	// FindListWeComTagGroupOption
	//  @Description: 标签组
	//  @return reply
	//  @return err
	//
	FindListWeComTagGroupOption() (reply []*tag2.WeComTagGroup, err error)
	//
	// FindListWeComTagGroupPage
	//  @Description: 标签组分页
	//  @param option
	//  @return reply
	//  @return err
	//
	FindListWeComTagGroupPage(option *types.PageOption[types.ListWeComTagGroupPageRequest]) (reply *types.Page[*tag2.WeComTagGroup], err error)

	//
	// ActionWeComCorpTagGroupRequest
	//  @Description: 添加，删除标签组内的标签
	//  @param options
	//  @return work
	//  @return err
	//
	ActionWeComCorpTagGroupRequest(options *types.ActionCorpTagGroupRequest) (work *kresp.ResponseWork, err error)

	//
	// FindListWeComTagOption
	//  @Description: 标签
	//  @return reply
	//  @return err
	//
	FindListWeComTagOption() (reply []*tag2.WeComTag, err error)
	//
	// FindListWeComTagPage
	//  @Description: 企业标签查询
	//  @param option
	//  @return reply
	//  @return err
	//
	FindListWeComTagPage(option *types.PageOption[types.ListWeComTagReqeust]) (reply *types.Page[*tag2.WeComTag], err error)
	//
	// PullListWeComCorpTagRequest
	//  @Description: 企业标签
	//  @param tagIds
	//  @param groupIds
	//  @param sync
	//  @return reply
	//  @return err
	//
	PullListWeComCorpTagRequest(tagIds []string, groupIds []string, sync int) (reply *tagResp.ResponseTagGetCorpTagList, err error)

	//
	// PullListWeComStrategyTagRequest
	//  @Description: 策略标签
	//  @param options
	//  @return reply
	//  @return err
	//
	PullListWeComStrategyTagRequest(options *tagReq.RequestTagGetStrategyTagList) (reply *tagResp.ResponseTagGetStrategyTagList, err error)

	//
	// CreateWeComCorpTagRequest
	//  @Description: 创建企业标签
	//  @param options
	//  @return *tagResp.ResponseTagAddCorpTag
	//  @return error
	//
	CreateWeComCorpTagRequest(options *tagReq.RequestTagAddCorpTag) (*tagResp.ResponseTagAddCorpTag, error)

	//
	// UpdateWeComCorpTagRequest
	//  @Description: 更新企业标签
	//  @param options
	//  @return *kresp.ResponseWork
	//  @return error
	//
	UpdateWeComCorpTagRequest(options *tagReq.RequestTagEditCorpTag) (*kresp.ResponseWork, error)

	//
	// DeleteWeComCorpTagRequest
	//  @Description: 删除标签
	//  @param options
	//  @return *kresp.ResponseWork
	//  @return error
	//
	DeleteWeComCorpTagRequest(options *tagReq.RequestTagDelCorpTag) (*kresp.ResponseWork, error)

	//
	// ActionWeComCustomerTagRequest
	//  @Description: 添加/移除客户标签
	//  @param option
	//  @return *kresp.ResponseWork
	//  @return error
	//
	ActionWeComCustomerTagRequest(option *tagReq.RequestTagMarkTag) (*kresp.ResponseWork, error)
}
