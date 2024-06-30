package qrcode

import (
	"github.com/starter-go/application"
	"github.com/starter-go/libgin/modules/libgin"
	"github.com/starter-go/qrcode"
	"github.com/starter-go/qrcode/gen/main4qrcode"
	"github.com/starter-go/qrcode/gen/test4qrcode"
)

// Module  ...
func Module() application.Module {
	mb := qrcode.NewMainModule()
	mb.Components(main4qrcode.ExportComponents)
	mb.Depend(libgin.Module())
	return mb.Create()
}

// ModuleForTest ...
func ModuleForTest() application.Module {
	mb := qrcode.NewTestModule()
	mb.Components(test4qrcode.ExportComponents)
	mb.Depend(Module())
	return mb.Create()
}
