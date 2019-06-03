package controllers

import (
	"encoding/json"
	"fmt"
		"net/http"
	common_scode "scode"
	common_status "status"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

const (
	httpGetConst    = "GET"
	httpPostConst   = "POST"
	httpPutConst    = "PUT"
	httpDeleteConst = "DELETE"
	httpPatchConst  = "PATCH"
	httpOptionConst = "OPTIONS"
)

// Output 输出
type Output struct {
	StatusCode  int         `json:"scode"`
	Description string      `json:"desc"`
	Message     string      `json:"message"`    //信息英文
	MessageCN   string      `json:"message_cn"` //信息中文
	Stack       []string    `json:"stack"`
	Data        interface{} `json:"data"`
}

// ------------------- Controller -----------------------
// Controller
type Controller struct {
	beego.Controller

	MyOutput   Output
}

//初始化controller.
func (c *Controller) Init(ctx *context.Context, controllerName, actionName string, app interface{}) {
	c.Controller.Init(ctx, controllerName, actionName, app)

}

// Prepare 在所有请求之前做操作
func (c *Controller) Prepare() {
	beego.Debug( "Request Begin:-----------------------------------------------------------------------------------")
	beego.Info( fmt.Sprintf("Request Begin: uri: %v, method: %v, ClientIP:%v, RemoteAddr:%v",
		c.Ctx.Input.URI(), c.Ctx.Input.Method(), c.GetClientIP(), c.Ctx.Request.RemoteAddr))

	switch c.Ctx.Input.Method() {
	case httpPutConst:
		beego.Debug("PUT...")
	case httpDeleteConst:
		beego.Debug("DELETE...")
	case httpPatchConst:
		beego.Debug("PATCH...")
	case httpGetConst:
		beego.Debug("GET...")
	case httpPostConst:
		beego.Debug("POST...")
	case httpOptionConst:
		beego.Debug("OPTIONS...")
	default:
		beego.Debug("Method Not Support")
	}
}

// Options 跨域
func (c *Controller) Options() {
	c.Data["json"] = map[string]interface{}{"common_scode": 0, "message": "ok"}
}

func (c *Controller) parseRequestForm(req interface{}) error {
	if err := c.ParseForm(req); err != nil {
		return common_status.NewStatusDesc(common_scode.ScodeManagerCommonParameterError, err.Error())
	}

	if req == nil {
		return common_status.NewStatusDesc(common_scode.ScodeManagerCommonParameterError, "Request Form Is Empty")
	}

	return nil
}

func (c *Controller) unmarshalRequestBodyJSON(req interface{}) error {
	if req == nil {
		return common_status.NewStatusDesc(common_scode.ScodeManagerCommonParameterError, "Request Body Is Nil")
	}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, req); err != nil {
		return common_status.NewStatusDesc(common_scode.ScodeManagerCommonParameterError, "Unmarshal Request Body JSON Error: "+err.Error())
	}

	return nil
}

// AllowCross 跨域
func (c *Controller) AllowCross() {
	c.Ctx.Output.Header("Access-Control-Allow-Origin", c.Ctx.Input.Header("Origin"))
	// c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	c.Ctx.Output.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	c.Ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type,Authorization,Cookie,UserName,Token")
	c.Ctx.Output.Header("Access-Control-Max-Age", "1728000")
	c.Ctx.Output.Header("Content-Type", "application/json")
}

func (c *Controller) setOutput(data interface{}, err error) {

	if err != nil {
		//models层必须保证所有返回的err都是status 对象
		c.MyOutput.StatusCode = int(err.(*common_status.Status).Code)
		c.MyOutput.Message = err.(*common_status.Status).Message
		c.MyOutput.MessageCN = err.(*common_status.Status).MessageCn
		c.MyOutput.Stack = append(c.MyOutput.Stack, err.(*common_status.Status).Stack...)
		c.MyOutput.Description = err.(*common_status.Status).Desc

		//	innerlog.Error("Response Error: ", err.(*common_status.Status).Code, err.(*common_status.Status).Message, err.(*common_status.Status).Stack)
	} else {
		c.MyOutput.StatusCode = int(common_status.SuccessStatus.Code)
		c.MyOutput.Message = "success"
		c.MyOutput.MessageCN = "成功"
		c.Ctx.Output.SetStatus(http.StatusOK)
	}
	if data != nil {
		c.MyOutput.Data = data
	}

}

// Finish 在所有请求之前做操作
func (c *Controller) Finish() {
	c.AllowCross() // 允许跨域
	c.Data["json"] = c.MyOutput
	c.ServeJSON()
	beego.Info(fmt.Sprintf("Controller Finish :uri: %v,  method: %v, ClientIP:%v, RemoteAddr:%v, status code:%v, message:%v, stack:%v",
		c.Ctx.Input.URI(), c.Ctx.Input.Method(), c.GetClientIP(), c.Ctx.Request.RemoteAddr, c.MyOutput.StatusCode, c.MyOutput.Message, c.MyOutput.Stack))
}

// GetClientIP 获取客户端ip
func (c *Controller) GetClientIP() string {
	clientIP := c.Ctx.Request.Header.Get("X-Forwarded-For")
	if clientIP == "" {
		return ""
	}

	return strings.Split(clientIP, ",")[0]
}

// IsPublicAddress 检查请求地址是否为外网地址
func (c *Controller) IsPublicAddress() bool {
	return false
}

// ---------------------------- error Controller -------------------------
// ErrorController 错误
type ErrorController struct {
	Controller
}

// Error404 404 error
func (e *ErrorController) Error404() {
	err := common_status.NewStatusDesc(common_scode.ScodeManagerCommonHTTPError, "404 Not Found")
	beego.Error("Error404: ", err.Error())
	e.Controller.setOutput(nil, err)
}

// Error403 403 error
func (e *ErrorController) Error403() {
	err := common_status.NewStatusDesc(common_scode.ScodeManagerCommonHTTPError, "403 Forbidden")
	beego.Error("Error403: ", err.Error())
	e.Controller.setOutput(nil, err)
}

// Error405 405 error
func (e *ErrorController) Error405() {
	err := common_status.NewStatusDesc(common_scode.ScodeManagerCommonHTTPError, "405 Method Not Allowed")
	beego.Error("Error405: ", err.Error())
	e.Controller.setOutput(nil, err)
}

// Error501 501 error
func (e *ErrorController) Error501() {
	err := common_status.NewStatusDesc(common_scode.ScodeManagerCommonHTTPError, "501 Method Not Implemented")
	beego.Error("Error501: ", err.Error())
	e.Controller.setOutput(nil, err)
}

// ---------------------- base  Controller---------------------
// OperateLog 操作日志
type OperateLog struct {
}

// BaseController Controller
// only allow server mode manager request passing
type BaseController struct {
	ErrorController
}

//初始化controller.
func (c *BaseController) Init(ctx *context.Context, controllerName, actionName string, app interface{}) {
	c.Controller.Init(ctx, controllerName, actionName, app)
}

// Prepare 在所有请求之前做操作
// 检测leader, 生成token
func (c *BaseController) Prepare() {
	beego.Debug( "Request Begin:-----------------------------------------------------------------------------------")
	beego.Info( fmt.Sprintf("Request Begin: uri: %v, method: %v, ClientIP:%v, RemoteAddr:%v",
		c.Ctx.Input.URI(), c.Ctx.Input.Method(), c.GetClientIP(), c.Ctx.Request.RemoteAddr))

	switch c.Ctx.Input.Method() {
	case httpPutConst:
		fallthrough
	case httpDeleteConst:
		fallthrough
	case httpPatchConst:
		c.Abort("501")
	case httpGetConst:
		beego.Info("GET - Request Params: ", c.Ctx.Input.URL())
	case httpPostConst:
		beego.Info( "POST - Request Body: ", strings.Replace(string(c.Ctx.Input.RequestBody), "\n", "\\n", -1))
	case httpOptionConst:
		beego.Info( "OPTIONS...")
	default:
		beego.Error("Method Not Support")
	}

}

// Finish 在所有请求之前做操作
func (c *BaseController) Finish() {
	c.AllowCross() // 允许跨域
	c.Data["json"] = c.MyOutput
	c.ServeJSON()

	beego.Info( fmt.Sprintf("Request End:uri: %v,  method: %v, ClientIP:%v, RemoteAddr:%v, status code:%v, message:%v, stack:%v",
		c.Ctx.Input.URI(), c.Ctx.Input.Method(), c.GetClientIP(), c.Ctx.Request.RemoteAddr, c.MyOutput.StatusCode, c.MyOutput.Message, c.MyOutput.Stack))

	beego.Info( "Request End:-----------------------------------------------------------------------------------")
}

//继承Basecontroller method返回都会调用setOutput
func (c *BaseController) setOutput(data interface{}, err error) {

	c.Controller.setOutput(data, err)

}
