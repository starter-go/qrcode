package main

import (
	"os"

	"github.com/starter-go/qrcode/modules/qrcode"
	"github.com/starter-go/starter"
)

func main() {
	m := qrcode.Module()
	i := starter.Init(os.Args)
	i.MainModule(m)
	i.WithPanic(true).Run()
}
