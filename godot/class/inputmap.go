package class

import (
	"log"

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

//func NewinputMapFromPointer(ptr gdnative.Pointer) inputMap {
func NewInputMapFromPointer(ptr gdnative.Pointer) inputMap {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := inputMap{}
	obj.SetBaseObject(owner)

	return obj
}

func newSingletonInputMap() *inputMap {
	return &inputMap{}
}

/*
   Manages all [InputEventAction] which can be created/modified from the project settings menu [code]Project > Project Settings > Input Map[/code] or in code with [method add_action] and [method action_add_event]. See [method Node._input].
*/
var InputMap = newSingletonInputMap()

/*
Manages all [InputEventAction] which can be created/modified from the project settings menu [code]Project > Project Settings > Input Map[/code] or in code with [method add_action] and [method action_add_event]. See [method Node._input].
*/
type inputMap struct {
	Object
	owner       gdnative.Object
	initialized bool
}

// EnsureSingleton will check to see if we have an object for it. If not, it will fetch its
// GDNative object and set it.
func (o *inputMap) ensureSingleton() {
	if o.initialized == true {
		return
	}
	log.Println("Singleton not found. Fetching from GDNative...")
	base := gdnative.GetSingleton("InputMap")
	o.SetBaseObject(base)
	o.initialized = true
}

func (o *inputMap) BaseClass() string {
	return "InputMap"
}

// SetBaseObject will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *inputMap) SetBaseObject(object gdnative.Object) {
	o.owner = object
}

func (o *inputMap) GetBaseObject() gdnative.Object {
	return o.owner
}

/*
        Adds an [InputEvent] to an action. This [InputEvent] will trigger the action.
	Args: [{ false action String} { false event InputEvent}], Returns: void
*/
func (o *inputMap) ActionAddEvent(action gdnative.String, event InputEvent) {
	o.ensureSingleton()
	log.Println("Calling InputMap.ActionAddEvent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(action)
	ptrArguments[1] = gdnative.NewPointerFromObject(event.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputMap", "action_add_event")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Removes an [InputEvent] from an action.
	Args: [{ false action String} { false event InputEvent}], Returns: void
*/
func (o *inputMap) ActionEraseEvent(action gdnative.String, event InputEvent) {
	o.ensureSingleton()
	log.Println("Calling InputMap.ActionEraseEvent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(action)
	ptrArguments[1] = gdnative.NewPointerFromObject(event.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputMap", "action_erase_event")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns [true] if an action has an [InputEvent] associated with it.
	Args: [{ false action String} { false event InputEvent}], Returns: bool
*/
func (o *inputMap) ActionHasEvent(action gdnative.String, event InputEvent) gdnative.Bool {
	o.ensureSingleton()
	log.Println("Calling InputMap.ActionHasEvent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(action)
	ptrArguments[1] = gdnative.NewPointerFromObject(event.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputMap", "action_has_event")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Adds an (empty) action to the [code]InputMap[/code]. An [InputEvent] can then be added to this action with [method action_add_event].
	Args: [{ false action String}], Returns: void
*/
func (o *inputMap) AddAction(action gdnative.String) {
	o.ensureSingleton()
	log.Println("Calling InputMap.AddAction()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(action)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputMap", "add_action")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Removes an action from the [code]InputMap[/code].
	Args: [{ false action String}], Returns: void
*/
func (o *inputMap) EraseAction(action gdnative.String) {
	o.ensureSingleton()
	log.Println("Calling InputMap.EraseAction()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(action)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputMap", "erase_action")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns [true] if the given event is part of an existing action. This method ignores keyboard modifiers if the given [InputEvent] is not pressed (for proper release detection). See [method action_has_event] if you don't want this behavior.
	Args: [{ false event InputEvent} { false action String}], Returns: bool
*/
func (o *inputMap) EventIsAction(event InputEvent, action gdnative.String) gdnative.Bool {
	o.ensureSingleton()
	log.Println("Calling InputMap.EventIsAction()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromObject(event.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromString(action)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputMap", "event_is_action")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns an array of [InputEvent]s associated with a given action.
	Args: [{ false action String}], Returns: Array
*/
func (o *inputMap) GetActionList(action gdnative.String) gdnative.Array {
	o.ensureSingleton()
	log.Println("Calling InputMap.GetActionList()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(action)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputMap", "get_action_list")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns an array of all actions in the [code]InputMap[/code].
	Args: [], Returns: Array
*/
func (o *inputMap) GetActions() gdnative.Array {
	o.ensureSingleton()
	log.Println("Calling InputMap.GetActions()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputMap", "get_actions")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns [code]true[/code] if the [code]InputMap[/code] has a registered action with the given name.
	Args: [{ false action String}], Returns: bool
*/
func (o *inputMap) HasAction(action gdnative.String) gdnative.Bool {
	o.ensureSingleton()
	log.Println("Calling InputMap.HasAction()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(action)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputMap", "has_action")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Clears all [InputEventAction] in the [code]InputMap[/code] and load it anew from [ProjectSettings].
	Args: [], Returns: void
*/
func (o *inputMap) LoadFromGlobals() {
	o.ensureSingleton()
	log.Println("Calling InputMap.LoadFromGlobals()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputMap", "load_from_globals")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}
