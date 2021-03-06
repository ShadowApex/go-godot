package godot

import (
	"github.com/shadowapex/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

//func NewMarginContainerFromPointer(ptr gdnative.Pointer) MarginContainer {
func newMarginContainerFromPointer(ptr gdnative.Pointer) MarginContainer {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := MarginContainer{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Simple margin container. Adds a left margin to anything contained.
*/
type MarginContainer struct {
	Container
	owner gdnative.Object
}

func (o *MarginContainer) BaseClass() string {
	return "MarginContainer"
}

// MarginContainerImplementer is an interface that implements the methods
// of the MarginContainer class.
type MarginContainerImplementer interface {
	ContainerImplementer
}
