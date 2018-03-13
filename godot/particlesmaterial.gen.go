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

// ParticlesMaterialEmissionShape is an enum for EmissionShape values.
type ParticlesMaterialEmissionShape int

const (
	ParticlesMaterialEmissionShapeBox            ParticlesMaterialEmissionShape = 2
	ParticlesMaterialEmissionShapeDirectedPoints ParticlesMaterialEmissionShape = 4
	ParticlesMaterialEmissionShapePoint          ParticlesMaterialEmissionShape = 0
	ParticlesMaterialEmissionShapePoints         ParticlesMaterialEmissionShape = 3
	ParticlesMaterialEmissionShapeSphere         ParticlesMaterialEmissionShape = 1
)

// ParticlesMaterialFlags is an enum for Flags values.
type ParticlesMaterialFlags int

const (
	ParticlesMaterialFlagAlignYToVelocity ParticlesMaterialFlags = 0
	ParticlesMaterialFlagMax              ParticlesMaterialFlags = 4
	ParticlesMaterialFlagRotateY          ParticlesMaterialFlags = 1
)

// ParticlesMaterialParameter is an enum for Parameter values.
type ParticlesMaterialParameter int

const (
	ParticlesMaterialParamAngle                 ParticlesMaterialParameter = 7
	ParticlesMaterialParamAngularVelocity       ParticlesMaterialParameter = 1
	ParticlesMaterialParamAnimOffset            ParticlesMaterialParameter = 11
	ParticlesMaterialParamAnimSpeed             ParticlesMaterialParameter = 10
	ParticlesMaterialParamDamping               ParticlesMaterialParameter = 6
	ParticlesMaterialParamHueVariation          ParticlesMaterialParameter = 9
	ParticlesMaterialParamInitialLinearVelocity ParticlesMaterialParameter = 0
	ParticlesMaterialParamLinearAccel           ParticlesMaterialParameter = 3
	ParticlesMaterialParamMax                   ParticlesMaterialParameter = 12
	ParticlesMaterialParamOrbitVelocity         ParticlesMaterialParameter = 2
	ParticlesMaterialParamRadialAccel           ParticlesMaterialParameter = 4
	ParticlesMaterialParamScale                 ParticlesMaterialParameter = 8
	ParticlesMaterialParamTangentialAccel       ParticlesMaterialParameter = 5
)

//func NewParticlesMaterialFromPointer(ptr gdnative.Pointer) ParticlesMaterial {
func newParticlesMaterialFromPointer(ptr gdnative.Pointer) ParticlesMaterial {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ParticlesMaterial{}
	obj.SetBaseObject(owner)

	return obj
}

/*
ParticlesMaterial defines particle properties and behavior. It is used in the [code]process_material[/code] of [Particles] and [Particles2D] emitter nodes. Some of this material's properties are applied to each particle when emitted, while others can have a [CurveTexture] applied to vary values over the lifetime of the particle.
*/
type ParticlesMaterial struct {
	Material
	owner gdnative.Object
}

func (o *ParticlesMaterial) BaseClass() string {
	return "ParticlesMaterial"
}

/*
        Undocumented
	Args: [], Returns: Color
*/
func (o *ParticlesMaterial) GetColor() gdnative.Color {
	//log.Println("Calling ParticlesMaterial.GetColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_color")

	// Call the parent method.
	// Color
	retPtr := gdnative.NewEmptyColor()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewColorFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Texture
*/
func (o *ParticlesMaterial) GetColorRamp() TextureImplementer {
	//log.Println("Calling ParticlesMaterial.GetColorRamp()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_color_ramp")

	// Call the parent method.
	// Texture
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newTextureFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(TextureImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Texture" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(TextureImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: Vector3
*/
func (o *ParticlesMaterial) GetEmissionBoxExtents() gdnative.Vector3 {
	//log.Println("Calling ParticlesMaterial.GetEmissionBoxExtents()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_emission_box_extents")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Texture
*/
func (o *ParticlesMaterial) GetEmissionColorTexture() TextureImplementer {
	//log.Println("Calling ParticlesMaterial.GetEmissionColorTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_emission_color_texture")

	// Call the parent method.
	// Texture
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newTextureFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(TextureImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Texture" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(TextureImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: Texture
*/
func (o *ParticlesMaterial) GetEmissionNormalTexture() TextureImplementer {
	//log.Println("Calling ParticlesMaterial.GetEmissionNormalTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_emission_normal_texture")

	// Call the parent method.
	// Texture
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newTextureFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(TextureImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Texture" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(TextureImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *ParticlesMaterial) GetEmissionPointCount() gdnative.Int {
	//log.Println("Calling ParticlesMaterial.GetEmissionPointCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_emission_point_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Texture
*/
func (o *ParticlesMaterial) GetEmissionPointTexture() TextureImplementer {
	//log.Println("Calling ParticlesMaterial.GetEmissionPointTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_emission_point_texture")

	// Call the parent method.
	// Texture
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newTextureFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(TextureImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Texture" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(TextureImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: enum.ParticlesMaterial::EmissionShape
*/
func (o *ParticlesMaterial) GetEmissionShape() ParticlesMaterialEmissionShape {
	//log.Println("Calling ParticlesMaterial.GetEmissionShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_emission_shape")

	// Call the parent method.
	// enum.ParticlesMaterial::EmissionShape
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ParticlesMaterialEmissionShape(ret)
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *ParticlesMaterial) GetEmissionSphereRadius() gdnative.Float {
	//log.Println("Calling ParticlesMaterial.GetEmissionSphereRadius()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_emission_sphere_radius")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false flag int}], Returns: bool
*/
func (o *ParticlesMaterial) GetFlag(flag gdnative.Int) gdnative.Bool {
	//log.Println("Calling ParticlesMaterial.GetFlag()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(flag)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_flag")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *ParticlesMaterial) GetFlatness() gdnative.Float {
	//log.Println("Calling ParticlesMaterial.GetFlatness()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_flatness")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Vector3
*/
func (o *ParticlesMaterial) GetGravity() gdnative.Vector3 {
	//log.Println("Calling ParticlesMaterial.GetGravity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_gravity")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false param int}], Returns: float
*/
func (o *ParticlesMaterial) GetParam(param gdnative.Int) gdnative.Float {
	//log.Println("Calling ParticlesMaterial.GetParam()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(param)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_param")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false param int}], Returns: float
*/
func (o *ParticlesMaterial) GetParamRandomness(param gdnative.Int) gdnative.Float {
	//log.Println("Calling ParticlesMaterial.GetParamRandomness()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(param)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_param_randomness")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false param int}], Returns: Texture
*/
func (o *ParticlesMaterial) GetParamTexture(param gdnative.Int) TextureImplementer {
	//log.Println("Calling ParticlesMaterial.GetParamTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(param)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_param_texture")

	// Call the parent method.
	// Texture
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newTextureFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(TextureImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Texture" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(TextureImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *ParticlesMaterial) GetSpread() gdnative.Float {
	//log.Println("Calling ParticlesMaterial.GetSpread()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_spread")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: GradientTexture
*/
func (o *ParticlesMaterial) GetTrailColorModifier() GradientTextureImplementer {
	//log.Println("Calling ParticlesMaterial.GetTrailColorModifier()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_trail_color_modifier")

	// Call the parent method.
	// GradientTexture
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newGradientTextureFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(GradientTextureImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "GradientTexture" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(GradientTextureImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *ParticlesMaterial) GetTrailDivisor() gdnative.Int {
	//log.Println("Calling ParticlesMaterial.GetTrailDivisor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_trail_divisor")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: CurveTexture
*/
func (o *ParticlesMaterial) GetTrailSizeModifier() CurveTextureImplementer {
	//log.Println("Calling ParticlesMaterial.GetTrailSizeModifier()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_trail_size_modifier")

	// Call the parent method.
	// CurveTexture
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newCurveTextureFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(CurveTextureImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "CurveTexture" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(CurveTextureImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [{ false color Color}], Returns: void
*/
func (o *ParticlesMaterial) SetColor(color gdnative.Color) {
	//log.Println("Calling ParticlesMaterial.SetColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_color")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false ramp Texture}], Returns: void
*/
func (o *ParticlesMaterial) SetColorRamp(ramp TextureImplementer) {
	//log.Println("Calling ParticlesMaterial.SetColorRamp()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(ramp.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_color_ramp")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false extents Vector3}], Returns: void
*/
func (o *ParticlesMaterial) SetEmissionBoxExtents(extents gdnative.Vector3) {
	//log.Println("Calling ParticlesMaterial.SetEmissionBoxExtents()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(extents)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_emission_box_extents")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture Texture}], Returns: void
*/
func (o *ParticlesMaterial) SetEmissionColorTexture(texture TextureImplementer) {
	//log.Println("Calling ParticlesMaterial.SetEmissionColorTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_emission_color_texture")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture Texture}], Returns: void
*/
func (o *ParticlesMaterial) SetEmissionNormalTexture(texture TextureImplementer) {
	//log.Println("Calling ParticlesMaterial.SetEmissionNormalTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_emission_normal_texture")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false point_count int}], Returns: void
*/
func (o *ParticlesMaterial) SetEmissionPointCount(pointCount gdnative.Int) {
	//log.Println("Calling ParticlesMaterial.SetEmissionPointCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(pointCount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_emission_point_count")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture Texture}], Returns: void
*/
func (o *ParticlesMaterial) SetEmissionPointTexture(texture TextureImplementer) {
	//log.Println("Calling ParticlesMaterial.SetEmissionPointTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_emission_point_texture")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false shape int}], Returns: void
*/
func (o *ParticlesMaterial) SetEmissionShape(shape gdnative.Int) {
	//log.Println("Calling ParticlesMaterial.SetEmissionShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(shape)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_emission_shape")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false radius float}], Returns: void
*/
func (o *ParticlesMaterial) SetEmissionSphereRadius(radius gdnative.Float) {
	//log.Println("Calling ParticlesMaterial.SetEmissionSphereRadius()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(radius)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_emission_sphere_radius")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false flag int} { false enable bool}], Returns: void
*/
func (o *ParticlesMaterial) SetFlag(flag gdnative.Int, enable gdnative.Bool) {
	//log.Println("Calling ParticlesMaterial.SetFlag()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(flag)
	ptrArguments[1] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_flag")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false amount float}], Returns: void
*/
func (o *ParticlesMaterial) SetFlatness(amount gdnative.Float) {
	//log.Println("Calling ParticlesMaterial.SetFlatness()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(amount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_flatness")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false accel_vec Vector3}], Returns: void
*/
func (o *ParticlesMaterial) SetGravity(accelVec gdnative.Vector3) {
	//log.Println("Calling ParticlesMaterial.SetGravity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(accelVec)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_gravity")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false param int} { false value float}], Returns: void
*/
func (o *ParticlesMaterial) SetParam(param gdnative.Int, value gdnative.Float) {
	//log.Println("Calling ParticlesMaterial.SetParam()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(param)
	ptrArguments[1] = gdnative.NewPointerFromFloat(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_param")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false param int} { false randomness float}], Returns: void
*/
func (o *ParticlesMaterial) SetParamRandomness(param gdnative.Int, randomness gdnative.Float) {
	//log.Println("Calling ParticlesMaterial.SetParamRandomness()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(param)
	ptrArguments[1] = gdnative.NewPointerFromFloat(randomness)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_param_randomness")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false param int} { false texture Texture}], Returns: void
*/
func (o *ParticlesMaterial) SetParamTexture(param gdnative.Int, texture TextureImplementer) {
	//log.Println("Calling ParticlesMaterial.SetParamTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(param)
	ptrArguments[1] = gdnative.NewPointerFromObject(texture.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_param_texture")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false degrees float}], Returns: void
*/
func (o *ParticlesMaterial) SetSpread(degrees gdnative.Float) {
	//log.Println("Calling ParticlesMaterial.SetSpread()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(degrees)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_spread")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture GradientTexture}], Returns: void
*/
func (o *ParticlesMaterial) SetTrailColorModifier(texture GradientTextureImplementer) {
	//log.Println("Calling ParticlesMaterial.SetTrailColorModifier()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_trail_color_modifier")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false divisor int}], Returns: void
*/
func (o *ParticlesMaterial) SetTrailDivisor(divisor gdnative.Int) {
	//log.Println("Calling ParticlesMaterial.SetTrailDivisor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(divisor)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_trail_divisor")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture CurveTexture}], Returns: void
*/
func (o *ParticlesMaterial) SetTrailSizeModifier(texture CurveTextureImplementer) {
	//log.Println("Calling ParticlesMaterial.SetTrailSizeModifier()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_trail_size_modifier")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// ParticlesMaterialImplementer is an interface that implements the methods
// of the ParticlesMaterial class.
type ParticlesMaterialImplementer interface {
	MaterialImplementer
	GetColor() gdnative.Color
	GetColorRamp() TextureImplementer
	GetEmissionBoxExtents() gdnative.Vector3
	GetEmissionColorTexture() TextureImplementer
	GetEmissionNormalTexture() TextureImplementer
	GetEmissionPointCount() gdnative.Int
	GetEmissionPointTexture() TextureImplementer
	GetEmissionSphereRadius() gdnative.Float
	GetFlag(flag gdnative.Int) gdnative.Bool
	GetFlatness() gdnative.Float
	GetGravity() gdnative.Vector3
	GetParam(param gdnative.Int) gdnative.Float
	GetParamRandomness(param gdnative.Int) gdnative.Float
	GetParamTexture(param gdnative.Int) TextureImplementer
	GetSpread() gdnative.Float
	GetTrailColorModifier() GradientTextureImplementer
	GetTrailDivisor() gdnative.Int
	GetTrailSizeModifier() CurveTextureImplementer
	SetColor(color gdnative.Color)
	SetColorRamp(ramp TextureImplementer)
	SetEmissionBoxExtents(extents gdnative.Vector3)
	SetEmissionColorTexture(texture TextureImplementer)
	SetEmissionNormalTexture(texture TextureImplementer)
	SetEmissionPointCount(pointCount gdnative.Int)
	SetEmissionPointTexture(texture TextureImplementer)
	SetEmissionShape(shape gdnative.Int)
	SetEmissionSphereRadius(radius gdnative.Float)
	SetFlag(flag gdnative.Int, enable gdnative.Bool)
	SetFlatness(amount gdnative.Float)
	SetGravity(accelVec gdnative.Vector3)
	SetParam(param gdnative.Int, value gdnative.Float)
	SetParamRandomness(param gdnative.Int, randomness gdnative.Float)
	SetParamTexture(param gdnative.Int, texture TextureImplementer)
	SetSpread(degrees gdnative.Float)
	SetTrailColorModifier(texture GradientTextureImplementer)
	SetTrailDivisor(divisor gdnative.Int)
	SetTrailSizeModifier(texture CurveTextureImplementer)
}