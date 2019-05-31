package context

import (
	bcontext "github.com/astaxie/beego/context"
	"math/rand"
	"time"
	//	"manager/logger"
)

var (
	noSpan      = "------"
	spanIDLenth = 6
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Context struct {
	BeegoCtx *bcontext.Context
	spanID   string
	//gcontext.Context
	//	Logger logger.Logger
}

func (c *Context) GetSpanID() string {
	return c.spanID
}

func newSpanID(n int) string {
	//random number
	var letterRunes = []rune("1234567890")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)

}

func NewContext(ctx *bcontext.Context) *Context {
	c := &Context{
		BeegoCtx: ctx,
		//Context: gcontext.Background(),
	}
	//uuidstr, _ := uuid.NewUUID()
	c.spanID = newSpanID(spanIDLenth)
	//	c.Logger = logger.NewLogger()
	return c
}
