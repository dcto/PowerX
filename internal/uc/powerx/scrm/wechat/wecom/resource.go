package wecom

import (
	"PowerX/internal/model/powerModel"
	"PowerX/internal/model/scrm/wechat/wecom/resource"
	"PowerX/internal/types"
	"bytes"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/media/response"
	"mime/multipart"
	"strings"
)

// UploadImageToResourceRequest
//
//	@Description:
//	@receiver this
//	@param path
//	@return link
//	@return err
func (uc *WeComUseCase) UploadImageToResourceRequest(localPath string, handle *multipart.FileHeader) (link string, err error) {

	var reply *response.ResponseUploadImage
	var fileName string
	var fileSize int
	if localPath != `` {
		reply, fileName, err = uc.uploadLocalUriToWeComLink(localPath)
	} else {
		reply, err = uc.uploadLocalFileToWeComLink(handle)
		fileName = handle.Filename
		fileSize = int(handle.Size)
	}
	if err != nil {
		panic(err)
	} else {
		err = uc.help.error(`scrm.wecom.upload.image.error.`, reply.ResponseWork)
	}

	if err == nil {
		uc.modelWeComResource.resource.Action(uc.db, []*resource.WeComResource{
			{
				Url:          reply.URL,
				FileName:     fileName,
				Remark:       ``,
				BucketName:   ``,
				Size:         fileSize,
				ResourceType: `image`,
			},
		})
	}

	return reply.URL, err

}

// uploadLocalFileToWeComLink
//
//	@Description:
//	@receiver this
//	@param handle
//	@return *response.ResponseUploadImage
//	@return error
func (uc *WeComUseCase) uploadLocalFileToWeComLink(handle *multipart.FileHeader) (*response.ResponseUploadImage, error) {

	hms := power.HashMap{}
	bts := bytes.Buffer{}
	hms[`name`] = handle.Filename
	hms[`value`] = func(han *multipart.FileHeader) []byte {
		open, _ := han.Open()
		n, _ := bts.ReadFrom(open)
		if n > 0 {
			return bts.Bytes()
		}
		return nil
	}(handle)
	reply, err := uc.Client.Media.UploadImage(uc.ctx, ``, &hms)

	return reply, err
}

// uploadLocalUriToWeComLink
//
//	@Description:
//	@receiver this
//	@param localPath
//	@return *response.ResponseUploadImage
//	@return string
//	@return error
func (uc *WeComUseCase) uploadLocalUriToWeComLink(localPath string) (*response.ResponseUploadImage, string, error) {

	split := strings.Split(localPath, `/`)
	name := split[len(split)-1 : len(split)]
	reply, err := uc.Client.Media.UploadImage(uc.ctx, localPath, nil)

	return reply, name[0], err
}

// FindWeComResourceListFromLocalPage
//
//	@Description:
//	@receiver this
//	@param opt
//	@return *types.Page[*resource.WeComResource]
//	@return error
func (uc *WeComUseCase) FindWeComResourceListFromLocalPage(opt *types.ListWeComResourceImageRequest) (*types.Page[*resource.WeComResource], error) {

	var resources []*resource.WeComResource
	var count int64

	query := uc.db.WithContext(uc.ctx).Table(uc.modelWeComResource.resource.TableName())

	if opt.PageIndex == 0 {
		opt.PageIndex = 1
	}
	if opt.PageSize == 0 {
		opt.PageSize = powerModel.PageDefaultSize
	}
	if v := opt.ResourceType; v != `` {
		query.Where(`resource_type = ?`, v)
	}
	if err := query.Count(&count).Error; err != nil {
		return nil, err
	}
	if opt.PageIndex != 0 && opt.PageSize != 0 {
		query.Offset((opt.PageIndex - 1) * opt.PageSize).Limit(opt.PageSize)
	}

	err := query.Find(&resources).Error

	return &types.Page[*resource.WeComResource]{
		List:      resources,
		PageIndex: opt.PageIndex,
		PageSize:  opt.PageSize,
		Total:     count,
	}, err
}
