// Code generated by thriftrw v1.4.0
// @generated

package services

import (
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type Cache_Clear_Args struct{}

func (v *Cache_Clear_Args) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *Cache_Clear_Args) FromWire(w wire.Value) error {
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}
	return nil
}

func (v *Cache_Clear_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [0]string
	i := 0
	return fmt.Sprintf("Cache_Clear_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *Cache_Clear_Args) Equals(rhs *Cache_Clear_Args) bool {
	return true
}

func (v *Cache_Clear_Args) MethodName() string {
	return "clear"
}

func (v *Cache_Clear_Args) EnvelopeType() wire.EnvelopeType {
	return wire.OneWay
}

var Cache_Clear_Helper = struct{ Args func() *Cache_Clear_Args }{}

func init() {
	Cache_Clear_Helper.Args = func() *Cache_Clear_Args {
		return &Cache_Clear_Args{}
	}
}
