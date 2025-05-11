package zhconv

import (
	"github.com/suwakei/go-zhconv/tables"
)

var convTables *tables.ConversionTables

func init() {
	convTables = tables.New()
}