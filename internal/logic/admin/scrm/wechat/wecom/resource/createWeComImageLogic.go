package resource

import (
	"PowerX/internal/types/errorx"
	"context"
	"net/http"
	"os"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateWeComImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 上传图片到微信
func NewCreateWeComImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateWeComImageLogic {
	return &CreateWeComImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateWeComImageLogic) CreateWeComImage(r *http.Request) (resp *types.CreateWeComSourceImageReply, err error) {

	err = r.ParseMultipartForm(2 << 20)
	if err != nil {
		return nil, errorx.WithCause(errorx.ErrBadRequest, err.Error())
	}
	var uri string
	file, handler, err := r.FormFile("resource")

	if err != nil || handler == nil {
		uri = r.FormValue("link")
		uri = `.` + uri
		if _, err := os.Stat(uri); os.IsNotExist(err) {
			return nil, errorx.WithCause(errorx.ErrBadRequest, err.Error())
		}
	}

	link, err := l.svcCtx.PowerX.SCRM.WeCom.UploadImageToResourceRequest(uri, handler)

	if err != nil {
		return nil, errorx.WithCause(errorx.ErrBadRequest, err.Error())
	}
	if file != nil {
		file.Close()
	}

	return &types.CreateWeComSourceImageReply{
		Link: link,
	}, err

}
