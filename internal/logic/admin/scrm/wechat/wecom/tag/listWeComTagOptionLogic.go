package tag

import (
	"PowerX/internal/model/scrm/wechat/wecom/tag"
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListWeComTagOptionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 标签列表对象/key=>val
func NewListWeComTagOptionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListWeComTagOptionLogic {
	return &ListWeComTagOptionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListWeComTagOptionLogic) ListWeComTagOption() (resp *types.ListWeComTagOptionReply, err error) {

	reply, err := l.svcCtx.PowerX.SCRM.WeCom.FindListWeComTagOption()
	return &types.ListWeComTagOptionReply{
		List: l.DTO(reply),
	}, err
}

// DTO
//
//	@Description:
//	@receiver tag
//	@param opt
//	@return tags
func (l *ListWeComTagOptionLogic) DTO(opt []*tag.WeComTag) (tags map[string]*types.WeComTag) {

	tags = make(map[string]*types.WeComTag)
	if opt == nil {
		return nil
	}
	for _, val := range opt {
		tags[val.TagId] = &types.WeComTag{
			Type:      val.Type,
			TagId:     val.TagId,
			GroupId:   val.GroupId,
			GroupName: val.WeComGroup.Name,
			Name:      val.Name,
			Sort:      val.Sort,
		}
	}
	return tags

}
