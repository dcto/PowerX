package wecom

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/agent/response"
)

// PullDetailWeComAppRequest
//
//	@Description:
//	@receiver this
//	@param agentID
//	@return reply
//	@return err
func (uc *WeComUseCase) PullDetailWeComAppRequest(agentID int) (reply *response.ResponseAgentGet, err error) {

	reply, err = uc.Client.Agent.Get(uc.ctx, agentID)
	if err != nil {
		panic(err)
	} else {
		err = uc.help.error(`scrm.pull.wecom.app.detail.error`, reply.ResponseWork)
	}
	return reply, err

}

// PullListWeComAppRequest
//
//	@Description:
//	@receiver this
//	@return reply
//	@return err
func (uc *WeComUseCase) PullListWeComAppRequest() (reply *response.ResponseAgentList, err error) {

	reply, err = uc.Client.Agent.List(uc.ctx)
	if err != nil {
		panic(err)
	} else {
		err = uc.help.error(`scrm.pull.wecom.app.list.error`, reply.ResponseWork)
	}
	return reply, err

}
