package ipconfig

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/oim/ipconfig/domain"
)

func getIp(c context.Context, ctx *app.RequestContext) {
	// 捕捉Panic
	defer func() {
		if err := recover(); err != nil {
			ctx.JSON(consts.StatusBadRequest, utils.H{"err": err})
		}
	}()
	ipContext := domain.BuildGetIpContext(&c, ctx)
	dispatch := domain.Dispatch(ipContext)
	ctx.JSON(200, dispatch)
}
