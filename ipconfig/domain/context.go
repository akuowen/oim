package domain

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetIpContext struct {
	c   *context.Context
	ctx *app.RequestContext
}

func BuildGetIpContext(c *context.Context, ctx *app.RequestContext) *GetIpContext {
	return &GetIpContext{
		c:   c,
		ctx: ctx,
	}
}
