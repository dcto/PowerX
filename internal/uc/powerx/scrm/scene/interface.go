package scene

import "PowerX/internal/model/scene"

type IsceneInterface interface {

	//
	//  @Description: qrcode
	//
	iQRCodeInterface
}

// iQRCodeInterface
// @Description:
type iQRCodeInterface interface {
	//
	// FindOneSceneQRCodeDetail
	//  @Description: 场景码详情
	//  @param qid
	//  @return *qrcode.QRCodeActive
	//
	FindOneSceneQRCodeDetail(qid string) *scene.SceneQRCode
	//
	// IncreaseSceneCpaNumber
	//  @Description: CPA+1
	//  @param qid
	//
	IncreaseSceneCpaNumber(qid string)
}
