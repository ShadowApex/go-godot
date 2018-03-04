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

//func NewAudioEffectDistortionFromPointer(ptr gdnative.Pointer) AudioEffectDistortion {
func NewAudioEffectDistortionFromPointer(ptr gdnative.Pointer) AudioEffectDistortion {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AudioEffectDistortion{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Modify the sound and make it dirty. Different types are available : clip, tan, lofi (bit crushing), overdrive, or waveshape. By distorting the waveform the frequency content change, which will often make the sound "crunchy" or "abrasive". For games, it can simulate sound coming from some saturated device or speaker very efficiently.
*/
type AudioEffectDistortion struct {
	AudioEffect
	owner gdnative.Object
}

func (o *AudioEffectDistortion) BaseClass() string {
	return "AudioEffectDistortion"
}

// SetBaseObject will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *AudioEffectDistortion) SetBaseObject(object gdnative.Object) {
	o.owner = object
}

func (o *AudioEffectDistortion) GetBaseObject() gdnative.Object {
	return o.owner
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AudioEffectDistortion) GetDrive() gdnative.Float {
	log.Println("Calling AudioEffectDistortion.GetDrive()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDistortion", "get_drive")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AudioEffectDistortion) GetKeepHfHz() gdnative.Float {
	log.Println("Calling AudioEffectDistortion.GetKeepHfHz()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDistortion", "get_keep_hf_hz")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: enum.AudioEffectDistortion::Mode
*/

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AudioEffectDistortion) GetPostGain() gdnative.Float {
	log.Println("Calling AudioEffectDistortion.GetPostGain()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDistortion", "get_post_gain")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AudioEffectDistortion) GetPreGain() gdnative.Float {
	log.Println("Calling AudioEffectDistortion.GetPreGain()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDistortion", "get_pre_gain")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [{ false drive float}], Returns: void
*/
func (o *AudioEffectDistortion) SetDrive(drive gdnative.Float) {
	log.Println("Calling AudioEffectDistortion.SetDrive()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(drive)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDistortion", "set_drive")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false keep_hf_hz float}], Returns: void
*/
func (o *AudioEffectDistortion) SetKeepHfHz(keepHfHz gdnative.Float) {
	log.Println("Calling AudioEffectDistortion.SetKeepHfHz()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(keepHfHz)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDistortion", "set_keep_hf_hz")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *AudioEffectDistortion) SetMode(mode gdnative.Int) {
	log.Println("Calling AudioEffectDistortion.SetMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDistortion", "set_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false post_gain float}], Returns: void
*/
func (o *AudioEffectDistortion) SetPostGain(postGain gdnative.Float) {
	log.Println("Calling AudioEffectDistortion.SetPostGain()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(postGain)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDistortion", "set_post_gain")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false pre_gain float}], Returns: void
*/
func (o *AudioEffectDistortion) SetPreGain(preGain gdnative.Float) {
	log.Println("Calling AudioEffectDistortion.SetPreGain()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(preGain)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectDistortion", "set_pre_gain")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}