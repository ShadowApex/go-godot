package class

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

//func NewCollisionShape2DFromPointer(ptr gdnative.Pointer) CollisionShape2D {
func NewCollisionShape2DFromPointer(ptr gdnative.Pointer) CollisionShape2D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := CollisionShape2D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Editor facility for creating and editing collision shapes in 2D space. You can use this node to represent all sorts of collision shapes, for example, add this to an [Area2D] to give it a detection shape, or add it to a [PhysicsBody2D] to give create solid object. [b]IMPORTANT[/b]: this is an Editor-only helper to create shapes, use [method get_shape] to get the actual shape.
*/
type CollisionShape2D struct {
	Node2D
	owner gdnative.Object
}

func (o *CollisionShape2D) BaseClass() string {
	return "CollisionShape2D"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *CollisionShape2D) X_ShapeChanged() {
	//log.Println("Calling CollisionShape2D.X_ShapeChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionShape2D", "_shape_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: Shape2D
*/
func (o *CollisionShape2D) GetShape() Shape2D {
	//log.Println("Calling CollisionShape2D.GetShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionShape2D", "get_shape")

	// Call the parent method.
	// Shape2D
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := NewShape2DFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *CollisionShape2D) IsDisabled() gdnative.Bool {
	//log.Println("Calling CollisionShape2D.IsDisabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionShape2D", "is_disabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *CollisionShape2D) IsOneWayCollisionEnabled() gdnative.Bool {
	//log.Println("Calling CollisionShape2D.IsOneWayCollisionEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionShape2D", "is_one_way_collision_enabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [{ false disabled bool}], Returns: void
*/
func (o *CollisionShape2D) SetDisabled(disabled gdnative.Bool) {
	//log.Println("Calling CollisionShape2D.SetDisabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(disabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionShape2D", "set_disabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *CollisionShape2D) SetOneWayCollision(enabled gdnative.Bool) {
	//log.Println("Calling CollisionShape2D.SetOneWayCollision()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionShape2D", "set_one_way_collision")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false shape Shape2D}], Returns: void
*/
func (o *CollisionShape2D) SetShape(shape Shape2D) {
	//log.Println("Calling CollisionShape2D.SetShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(shape.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionShape2D", "set_shape")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}