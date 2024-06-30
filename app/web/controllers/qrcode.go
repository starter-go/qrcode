package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/libgin"
	"github.com/starter-go/rbac"

	goqrcode "github.com/skip2/go-qrcode"
)

// QRCodeController ...
type QRCodeController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Sender libgin.Responder //starter:inject("#")

}

func (inst *QRCodeController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *QRCodeController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{Route: inst.route}
}

func (inst *QRCodeController) route(rp libgin.RouterProxy) error {
	const path = "/qrcode/:content"
	rp.GET(path, inst.handleGetImage)
	return nil
}

func (inst *QRCodeController) handle(c *gin.Context) {
	req := &myQRCodeRequest{
		context:    c,
		controller: inst,
	}
	req.execute(req.doNOP)
}

func (inst *QRCodeController) handleGetImage(c *gin.Context) {
	req := &myQRCodeRequest{
		context:            c,
		controller:         inst,
		wantRequestContent: true,
		wantQueryParams:    true,
	}
	req.execute(req.doGetImage)
}

////////////////////////////////////////////////////////////////////////////////

func normalizeSize(size int) int {
	const (
		min = 10
		max = 1024
	)
	if size < min {
		size = min
	}
	if size > max {
		size = max
	}
	return size
}

func normalizeLevel(level goqrcode.RecoveryLevel) goqrcode.RecoveryLevel {
	const (
		min = goqrcode.Low
		max = goqrcode.Highest
	)
	if level < min {
		level = min
	}
	if level > max {
		level = max
	}
	return level
}

func parseLevel(str string) goqrcode.RecoveryLevel {
	str = strings.ToUpper(str)
	str = strings.TrimSpace(str)
	switch str {
	case "L":
		return goqrcode.Low
	case "M":
		return goqrcode.Medium
	case "Q":
		return goqrcode.High
	case "H":
		return goqrcode.Highest
	}
	return goqrcode.Medium
}

////////////////////////////////////////////////////////////////////////////////

type myQRCodeResult struct {
	content     []byte
	contentType string
	status      int
}

////////////////////////////////////////////////////////////////////////////////

type myQRCodeRequest struct {
	context    *gin.Context
	controller *QRCodeController

	wantRequestContent bool
	wantQueryParams    bool

	contentRaw      string
	contentEncoding string // [string|hex|base64]
	size            int
	level           goqrcode.RecoveryLevel

	body1  rbac.BaseVO
	body2  rbac.BaseVO
	result myQRCodeResult
}

func (inst *myQRCodeRequest) open() error {

	c := inst.context

	if inst.wantRequestContent {
		str := c.Param("content")
		inst.contentRaw = str
	}

	if inst.wantQueryParams {

		enc := c.Query("encoding")
		size := c.Query("size")
		level := c.Query("level") // [L,M,Q,H]

		sizeInt, _ := strconv.Atoi(size)
		levelNum := parseLevel(level)

		inst.contentEncoding = enc
		inst.size = sizeInt
		inst.level = levelNum
	}

	return nil
}

func (inst *myQRCodeRequest) send(err error) {
	if err == nil {
		inst.sendImage(&inst.result)
	} else {
		inst.sendError(err)
	}
	return
}

func (inst *myQRCodeRequest) sendImage(res *myQRCodeResult) {
	c := inst.context
	data := res.content
	ctype := res.contentType
	code := res.status
	c.Data(code, ctype, data)
}

func (inst *myQRCodeRequest) sendError(err error) {
	data := &inst.body2
	code := inst.body2.Status
	resp := new(libgin.Response)
	resp.Context = inst.context
	resp.Error = err
	resp.Data = data
	resp.Status = code
	inst.controller.Sender.Send(resp)
}

func (inst *myQRCodeRequest) execute(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myQRCodeRequest) doNOP() error {
	return nil
}

func (inst *myQRCodeRequest) doGetImage() error {

	size := normalizeSize(inst.size)
	level := normalizeLevel(inst.level)

	content, err := inst.getQRCodeContent()
	if err != nil {
		return err
	}

	png, err := goqrcode.Encode(content, level, size)
	if err != nil {
		return err
	}

	res := &inst.result
	res.content = png
	res.contentType = "image/png"
	res.status = http.StatusOK
	return nil
}

func (inst *myQRCodeRequest) getQRCodeContent() (string, error) {
	enc := inst.contentEncoding
	switch enc {
	case "string":
		return inst.getQRCodeContentString()
	case "hex":
		return inst.getQRCodeContentHex()
	case "base64":
		return inst.getQRCodeContentBase64()
	}
	return inst.getQRCodeContentString()
}

func (inst *myQRCodeRequest) getQRCodeContentString() (string, error) {
	raw := inst.contentRaw
	return raw, nil
}

func (inst *myQRCodeRequest) getQRCodeContentHex() (string, error) {
	raw := inst.contentRaw
	hex := lang.Hex(raw)
	bin := hex.Bytes()
	return string(bin), nil
}

func (inst *myQRCodeRequest) getQRCodeContentBase64() (string, error) {
	raw := inst.contentRaw
	b64 := lang.Base64(raw)
	bin := b64.Bytes()
	return string(bin), nil
}
