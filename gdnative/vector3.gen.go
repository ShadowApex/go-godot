package gdnative

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "types.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

/*
#include "gdnative.gen.h"
#include <gdnative/vector3.h>
// Include all headers for now. TODO: Look up all the required
// headers we need to import based on the method arguments and return types.
#include <gdnative/aabb.h>
#include <gdnative/array.h>
#include <gdnative/basis.h>
#include <gdnative/color.h>
#include <gdnative/dictionary.h>
#include <gdnative/gdnative.h>
#include <gdnative/node_path.h>
#include <gdnative/plane.h>
#include <gdnative/pool_arrays.h>
#include <gdnative/quat.h>
#include <gdnative/rect2.h>
#include <gdnative/rid.h>
#include <gdnative/string.h>
#include <gdnative/string_name.h>
#include <gdnative/transform.h>
#include <gdnative/transform2d.h>
#include <gdnative/variant.h>
#include <gdnative/vector2.h>
#include <gdnative/vector3.h>
#include <gdnative_api_struct.gen.h>
*/
import "C"
import "unsafe"

// NewEmptyVector3 will return a pointer to an empty
// initialized Vector3. This is primarily used in
// conjunction with MethodBindPtrCall.
func NewEmptyVector3() Pointer {
	var obj C.godot_vector3
	return Pointer{base: unsafe.Pointer(&obj)}
}

// NewPointerFromVector3 will return an unsafe pointer to the given
// object. This is primarily used in conjunction with MethodBindPtrCall.
func NewPointerFromVector3(obj Vector3) Pointer {
	return Pointer{base: unsafe.Pointer(obj.getBase())}
}

// NewVector3FromPointer will return a Vector3 from the
// given unsafe pointer. This is primarily used in conjunction with MethodBindPtrCall.
func NewVector3FromPointer(ptr Pointer) Vector3 {

	return Vector3{base: (*C.godot_vector3)(ptr.getBase())}
}

type Vector3 struct {
	base *C.godot_vector3
}

func (gdt Vector3) getBase() *C.godot_vector3 {
	return gdt.base
}

// NewVector3 godot_vector3_new [[godot_vector3 * r_dest] [const godot_real p_x] [const godot_real p_y] [const godot_real p_z]] void
func NewVector3(x Real, y Real, z Real) Vector3 {
	var dest C.godot_vector3
	arg1 := x.getBase()
	arg2 := y.getBase()
	arg3 := z.getBase()
	C.go_godot_vector3_new(GDNative.api, &dest, arg1, arg2, arg3)
	return Vector3{base: &dest}
}

// AsString godot_vector3_as_string [[const godot_vector3 * p_self]] godot_string
func (gdt *Vector3) AsString() String {
	arg0 := gdt.getBase()

	ret := C.go_godot_vector3_as_string(GDNative.api, arg0)

	wchar := C.go_godot_string_wide_str(GDNative.api, &ret)
	goWchar := newWcharT(wchar)
	return String(goWchar.AsString())

}

// MinAxis godot_vector3_min_axis [[const godot_vector3 * p_self]] godot_int
func (gdt *Vector3) MinAxis() Int {
	arg0 := gdt.getBase()

	ret := C.go_godot_vector3_min_axis(GDNative.api, arg0)

	return Int(ret)
}

// MaxAxis godot_vector3_max_axis [[const godot_vector3 * p_self]] godot_int
func (gdt *Vector3) MaxAxis() Int {
	arg0 := gdt.getBase()

	ret := C.go_godot_vector3_max_axis(GDNative.api, arg0)

	return Int(ret)
}

// Length godot_vector3_length [[const godot_vector3 * p_self]] godot_real
func (gdt *Vector3) Length() Real {
	arg0 := gdt.getBase()

	ret := C.go_godot_vector3_length(GDNative.api, arg0)

	return Real(ret)
}

// LengthSquared godot_vector3_length_squared [[const godot_vector3 * p_self]] godot_real
func (gdt *Vector3) LengthSquared() Real {
	arg0 := gdt.getBase()

	ret := C.go_godot_vector3_length_squared(GDNative.api, arg0)

	return Real(ret)
}

// IsNormalized godot_vector3_is_normalized [[const godot_vector3 * p_self]] godot_bool
func (gdt *Vector3) IsNormalized() Bool {
	arg0 := gdt.getBase()

	ret := C.go_godot_vector3_is_normalized(GDNative.api, arg0)

	return Bool(ret)
}

// Normalized godot_vector3_normalized [[const godot_vector3 * p_self]] godot_vector3
func (gdt *Vector3) Normalized() Vector3 {
	arg0 := gdt.getBase()

	ret := C.go_godot_vector3_normalized(GDNative.api, arg0)

	return Vector3{base: &ret}

}

// Inverse godot_vector3_inverse [[const godot_vector3 * p_self]] godot_vector3
func (gdt *Vector3) Inverse() Vector3 {
	arg0 := gdt.getBase()

	ret := C.go_godot_vector3_inverse(GDNative.api, arg0)

	return Vector3{base: &ret}

}

// Snapped godot_vector3_snapped [[const godot_vector3 * p_self] [const godot_vector3 * p_by]] godot_vector3
func (gdt *Vector3) Snapped(by Vector3) Vector3 {
	arg0 := gdt.getBase()
	arg1 := by.getBase()

	ret := C.go_godot_vector3_snapped(GDNative.api, arg0, arg1)

	return Vector3{base: &ret}

}

// Rotated godot_vector3_rotated [[const godot_vector3 * p_self] [const godot_vector3 * p_axis] [const godot_real p_phi]] godot_vector3
func (gdt *Vector3) Rotated(axis Vector3, phi Real) Vector3 {
	arg0 := gdt.getBase()
	arg1 := axis.getBase()
	arg2 := phi.getBase()

	ret := C.go_godot_vector3_rotated(GDNative.api, arg0, arg1, arg2)

	return Vector3{base: &ret}

}

// LinearInterpolate godot_vector3_linear_interpolate [[const godot_vector3 * p_self] [const godot_vector3 * p_b] [const godot_real p_t]] godot_vector3
func (gdt *Vector3) LinearInterpolate(b Vector3, t Real) Vector3 {
	arg0 := gdt.getBase()
	arg1 := b.getBase()
	arg2 := t.getBase()

	ret := C.go_godot_vector3_linear_interpolate(GDNative.api, arg0, arg1, arg2)

	return Vector3{base: &ret}

}

// CubicInterpolate godot_vector3_cubic_interpolate [[const godot_vector3 * p_self] [const godot_vector3 * p_b] [const godot_vector3 * p_pre_a] [const godot_vector3 * p_post_b] [const godot_real p_t]] godot_vector3
func (gdt *Vector3) CubicInterpolate(b Vector3, preA Vector3, postB Vector3, t Real) Vector3 {
	arg0 := gdt.getBase()
	arg1 := b.getBase()
	arg2 := preA.getBase()
	arg3 := postB.getBase()
	arg4 := t.getBase()

	ret := C.go_godot_vector3_cubic_interpolate(GDNative.api, arg0, arg1, arg2, arg3, arg4)

	return Vector3{base: &ret}

}

// Dot godot_vector3_dot [[const godot_vector3 * p_self] [const godot_vector3 * p_b]] godot_real
func (gdt *Vector3) Dot(b Vector3) Real {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_vector3_dot(GDNative.api, arg0, arg1)

	return Real(ret)
}

// Cross godot_vector3_cross [[const godot_vector3 * p_self] [const godot_vector3 * p_b]] godot_vector3
func (gdt *Vector3) Cross(b Vector3) Vector3 {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_vector3_cross(GDNative.api, arg0, arg1)

	return Vector3{base: &ret}

}

// Outer godot_vector3_outer [[const godot_vector3 * p_self] [const godot_vector3 * p_b]] godot_basis
func (gdt *Vector3) Outer(b Vector3) Basis {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_vector3_outer(GDNative.api, arg0, arg1)

	return Basis{base: &ret}

}

// ToDiagonalMatrix godot_vector3_to_diagonal_matrix [[const godot_vector3 * p_self]] godot_basis
func (gdt *Vector3) ToDiagonalMatrix() Basis {
	arg0 := gdt.getBase()

	ret := C.go_godot_vector3_to_diagonal_matrix(GDNative.api, arg0)

	return Basis{base: &ret}

}

// Abs godot_vector3_abs [[const godot_vector3 * p_self]] godot_vector3
func (gdt *Vector3) Abs() Vector3 {
	arg0 := gdt.getBase()

	ret := C.go_godot_vector3_abs(GDNative.api, arg0)

	return Vector3{base: &ret}

}

// Floor godot_vector3_floor [[const godot_vector3 * p_self]] godot_vector3
func (gdt *Vector3) Floor() Vector3 {
	arg0 := gdt.getBase()

	ret := C.go_godot_vector3_floor(GDNative.api, arg0)

	return Vector3{base: &ret}

}

// Ceil godot_vector3_ceil [[const godot_vector3 * p_self]] godot_vector3
func (gdt *Vector3) Ceil() Vector3 {
	arg0 := gdt.getBase()

	ret := C.go_godot_vector3_ceil(GDNative.api, arg0)

	return Vector3{base: &ret}

}

// DistanceTo godot_vector3_distance_to [[const godot_vector3 * p_self] [const godot_vector3 * p_b]] godot_real
func (gdt *Vector3) DistanceTo(b Vector3) Real {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_vector3_distance_to(GDNative.api, arg0, arg1)

	return Real(ret)
}

// DistanceSquaredTo godot_vector3_distance_squared_to [[const godot_vector3 * p_self] [const godot_vector3 * p_b]] godot_real
func (gdt *Vector3) DistanceSquaredTo(b Vector3) Real {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_vector3_distance_squared_to(GDNative.api, arg0, arg1)

	return Real(ret)
}

// AngleTo godot_vector3_angle_to [[const godot_vector3 * p_self] [const godot_vector3 * p_to]] godot_real
func (gdt *Vector3) AngleTo(to Vector3) Real {
	arg0 := gdt.getBase()
	arg1 := to.getBase()

	ret := C.go_godot_vector3_angle_to(GDNative.api, arg0, arg1)

	return Real(ret)
}

// Slide godot_vector3_slide [[const godot_vector3 * p_self] [const godot_vector3 * p_n]] godot_vector3
func (gdt *Vector3) Slide(n Vector3) Vector3 {
	arg0 := gdt.getBase()
	arg1 := n.getBase()

	ret := C.go_godot_vector3_slide(GDNative.api, arg0, arg1)

	return Vector3{base: &ret}

}

// Bounce godot_vector3_bounce [[const godot_vector3 * p_self] [const godot_vector3 * p_n]] godot_vector3
func (gdt *Vector3) Bounce(n Vector3) Vector3 {
	arg0 := gdt.getBase()
	arg1 := n.getBase()

	ret := C.go_godot_vector3_bounce(GDNative.api, arg0, arg1)

	return Vector3{base: &ret}

}

// Reflect godot_vector3_reflect [[const godot_vector3 * p_self] [const godot_vector3 * p_n]] godot_vector3
func (gdt *Vector3) Reflect(n Vector3) Vector3 {
	arg0 := gdt.getBase()
	arg1 := n.getBase()

	ret := C.go_godot_vector3_reflect(GDNative.api, arg0, arg1)

	return Vector3{base: &ret}

}

// OperatorAdd godot_vector3_operator_add [[const godot_vector3 * p_self] [const godot_vector3 * p_b]] godot_vector3
func (gdt *Vector3) OperatorAdd(b Vector3) Vector3 {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_vector3_operator_add(GDNative.api, arg0, arg1)

	return Vector3{base: &ret}

}

// OperatorSubtract godot_vector3_operator_subtract [[const godot_vector3 * p_self] [const godot_vector3 * p_b]] godot_vector3
func (gdt *Vector3) OperatorSubtract(b Vector3) Vector3 {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_vector3_operator_subtract(GDNative.api, arg0, arg1)

	return Vector3{base: &ret}

}

// OperatorMultiplyVector godot_vector3_operator_multiply_vector [[const godot_vector3 * p_self] [const godot_vector3 * p_b]] godot_vector3
func (gdt *Vector3) OperatorMultiplyVector(b Vector3) Vector3 {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_vector3_operator_multiply_vector(GDNative.api, arg0, arg1)

	return Vector3{base: &ret}

}

// OperatorMultiplyScalar godot_vector3_operator_multiply_scalar [[const godot_vector3 * p_self] [const godot_real p_b]] godot_vector3
func (gdt *Vector3) OperatorMultiplyScalar(b Real) Vector3 {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_vector3_operator_multiply_scalar(GDNative.api, arg0, arg1)

	return Vector3{base: &ret}

}

// OperatorDivideVector godot_vector3_operator_divide_vector [[const godot_vector3 * p_self] [const godot_vector3 * p_b]] godot_vector3
func (gdt *Vector3) OperatorDivideVector(b Vector3) Vector3 {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_vector3_operator_divide_vector(GDNative.api, arg0, arg1)

	return Vector3{base: &ret}

}

// OperatorDivideScalar godot_vector3_operator_divide_scalar [[const godot_vector3 * p_self] [const godot_real p_b]] godot_vector3
func (gdt *Vector3) OperatorDivideScalar(b Real) Vector3 {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_vector3_operator_divide_scalar(GDNative.api, arg0, arg1)

	return Vector3{base: &ret}

}

// OperatorEqual godot_vector3_operator_equal [[const godot_vector3 * p_self] [const godot_vector3 * p_b]] godot_bool
func (gdt *Vector3) OperatorEqual(b Vector3) Bool {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_vector3_operator_equal(GDNative.api, arg0, arg1)

	return Bool(ret)
}

// OperatorLess godot_vector3_operator_less [[const godot_vector3 * p_self] [const godot_vector3 * p_b]] godot_bool
func (gdt *Vector3) OperatorLess(b Vector3) Bool {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_vector3_operator_less(GDNative.api, arg0, arg1)

	return Bool(ret)
}

// OperatorNeg godot_vector3_operator_neg [[const godot_vector3 * p_self]] godot_vector3
func (gdt *Vector3) OperatorNeg() Vector3 {
	arg0 := gdt.getBase()

	ret := C.go_godot_vector3_operator_neg(GDNative.api, arg0)

	return Vector3{base: &ret}

}

// SetAxis godot_vector3_set_axis [[godot_vector3 * p_self] [const godot_vector3_axis p_axis] [const godot_real p_val]] void
func (gdt *Vector3) SetAxis(axis Vector3Axis, val Real) {
	arg0 := gdt.getBase()
	arg1 := axis.getBase()
	arg2 := val.getBase()

	C.go_godot_vector3_set_axis(GDNative.api, arg0, arg1, arg2)
}

// GetAxis godot_vector3_get_axis [[const godot_vector3 * p_self] [const godot_vector3_axis p_axis]] godot_real
func (gdt *Vector3) GetAxis(axis Vector3Axis) Real {
	arg0 := gdt.getBase()
	arg1 := axis.getBase()

	ret := C.go_godot_vector3_get_axis(GDNative.api, arg0, arg1)

	return Real(ret)
}

// Vector3Axis is a Go wrapper for the C.godot_vector3_axis enum type.
type Vector3Axis int

func (e Vector3Axis) getBase() C.godot_vector3_axis {
	return C.godot_vector3_axis(e)
}

const (
	Vector3AxisX Vector3Axis = iota
	Vector3AxisY
	Vector3AxisZ
)

// Vector3AxisLookupMap is a string-based lookup table of constants for Vector3Axis.
var Vector3AxisLookupMap = map[string]Vector3Axis{
	"Vector3AxisX": Vector3AxisX,
	"Vector3AxisY": Vector3AxisY,
	"Vector3AxisZ": Vector3AxisZ,
}
