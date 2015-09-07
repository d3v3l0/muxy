package middleware

import (
	"fmt"
	"github.com/mefellows/muxy/log"
	"github.com/mefellows/muxy/muxy"
	"github.com/mefellows/plugo/plugo"
	"time"
)

type LoggerMiddleware struct {
	HexOutput bool `mapstructure:"hex_output"`
	format    string
}

const DEFAULT_DELAY = 2 * time.Second
const bytesTab = "\n\t\t\t\t\t\t"

func init() {
	plugo.PluginFactories.Register(func() (interface{}, error) {
		return &LoggerMiddleware{}, nil
	}, "logger")
}

func (l *LoggerMiddleware) Setup() {
	if l.HexOutput {
		l.format = "%x"
	} else {
		l.format = "%s"
	}
}

func (LoggerMiddleware) Teardown() {
}

func (l *LoggerMiddleware) HandleEvent(e muxy.ProxyEvent, ctx *muxy.Context) {
	switch e {
	case muxy.EVENT_PRE_DISPATCH:
		if ctx.Request == nil {
			if len(ctx.Bytes) > 0 {
				log.Info("Handle TCP event " + log.Colorize(log.GREY, "PRE_DISPATCH") + fmt.Sprintf(" Received %d%s", len(ctx.Bytes), " bytes"))
				data := fmt.Sprintf(l.format, ctx.Bytes)
				log.Debug("Handle TCP event " + log.Colorize(log.GREY, "PRE_DISPATCH") + " Received request: " + bytesTab + log.Colorize(log.BLUE, data))
			}
		} else {
			log.Info("Handle HTTP event " + log.Colorize(log.GREY, "PRE_DISPATCH") + " Proxying request " +
				log.Colorize(log.LIGHTMAGENTA, ctx.Request.Method) +
				log.Colorize(log.BLUE, " \""+ctx.Request.URL.String()+"\""))
		}
	case muxy.EVENT_POST_DISPATCH:
	case muxy.EVENT_PRE_RESPONSE:
		if ctx.Request == nil {
			if len(ctx.Bytes) > 0 {
				log.Info("Handle TCP event " + log.Colorize(log.GREY, "POST_DISPATCH") + fmt.Sprintf(" Sent %d%s", len(ctx.Bytes), " bytes"))
				data := fmt.Sprintf(l.format, ctx.Bytes)
				log.Debug("Handle TCP event " + log.Colorize(log.GREY, "POST_DISPATCH") + " Sent response: " + bytesTab + log.Colorize(log.BLUE, data))
			}
		} else {
			log.Info("Handle HTTP event " + log.Colorize(log.GREY, "PRE_DISPATCH") + " Returning " +
				log.Colorize(log.GREY, ctx.Response.Status))
		}
	}
}
