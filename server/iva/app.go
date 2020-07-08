package iva

import (
	"iris-vue-admin/infra"
	"iris-vue-admin/infra/base"
)

func init()  {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.LogWriterStarter{})
}