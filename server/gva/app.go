package iva

import (
	"gva/infra"
	"gva/infra/base"
)

func init()  {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.LogWriterStarter{})
	//infra.Register(&base.IrisServerStarter{})
	infra.Register(&base.GinServerStarter{})
	infra.Register(&infra.WebApiStarter{})
	infra.Register(&base.HookStarter{})
}