package test4qrcode
import (
    p43b15cb1b "github.com/starter-go/qrcode/src/test/golang/unit"
     "github.com/starter-go/application"
)

// type p43b15cb1b.DemoUnit in package:github.com/starter-go/qrcode/src/test/golang/unit
//
// id:com-43b15cb1bf2faa20-unit-DemoUnit
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type p43b15cb1bf_unit_DemoUnit struct {
}

func (inst* p43b15cb1bf_unit_DemoUnit) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-43b15cb1bf2faa20-unit-DemoUnit"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p43b15cb1bf_unit_DemoUnit) new() any {
    return &p43b15cb1b.DemoUnit{}
}

func (inst* p43b15cb1bf_unit_DemoUnit) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p43b15cb1b.DemoUnit)
	nop(ie, com)

	


    return nil
}


