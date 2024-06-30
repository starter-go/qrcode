package main4qrcode
import (
    pd1a916a20 "github.com/starter-go/libgin"
    pcb8abc8da "github.com/starter-go/qrcode/app/web/controllers"
     "github.com/starter-go/application"
)

// type pcb8abc8da.QRCodeController in package:github.com/starter-go/qrcode/app/web/controllers
//
// id:com-cb8abc8da1f2391b-controllers-QRCodeController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pcb8abc8da1_controllers_QRCodeController struct {
}

func (inst* pcb8abc8da1_controllers_QRCodeController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-cb8abc8da1f2391b-controllers-QRCodeController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pcb8abc8da1_controllers_QRCodeController) new() any {
    return &pcb8abc8da.QRCodeController{}
}

func (inst* pcb8abc8da1_controllers_QRCodeController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pcb8abc8da.QRCodeController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)


    return nil
}


func (inst*pcb8abc8da1_controllers_QRCodeController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


