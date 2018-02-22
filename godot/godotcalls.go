package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <gdnative/gdnative.h>
#include <nativescript/godot_nativescript.h>

void **build_array(int length);
void **build_array(int length) {
	void *ptr;
	void **arr = malloc(sizeof(void *) * length);
	for (int i = 0; i < length; i++) {
	    arr[i] = ptr;
	}

	return arr;
}

void add_element(void**, void*, int);
void add_element(void **array, void *element, int index) {
	printf("CGO: Array %p %p %p %p %p\n", &array, array, &array[index], *array, array[index]);
    array[index] = element;
	printf("CGO: Index %i %p\n", index, element);
	printf("CGO: Array %p %p %p %p %p\n", &array, array, &array[index], *array, array[index]);
}
*/
import "C"

import (
	"log"
	"unsafe"
)

func getGodotMethod(baseClass string, methodName string) *C.godot_method_bind {
	// Convert the base class and method names to C strings.
	log.Println("  Using base class: ", baseClass)
	classCString := C.CString(baseClass)
	defer C.free(unsafe.Pointer(classCString))

	log.Println("  Using method name: ", methodName)
	methodCString := C.CString(methodName)
	defer C.free(unsafe.Pointer(methodCString))

	// Get the Godot method bind pointer so we can pass it to godot_method_bind_ptrcall.
	var methodBind *C.godot_method_bind
	methodBind = C.godot_method_bind_get_method(classCString, methodCString)
	log.Println("  Using method bind pointer: ", methodBind)
	return methodBind
}

func godotCallPoolIntArrayIntFloatFloat(o Class, methodName string, arg0 int64, arg1 float64, arg2 float64) *PoolIntArray {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_real(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_real(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_pool_int_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolIntArray{&ret}

}
func godotCallIntIntFloatBool(o Class, methodName string, arg0 int64, arg1 float64, arg2 bool) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_real(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallVector2Int(o Class, methodName string, arg0 int64) *Vector2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_vector2
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector2{&ret}

}
func godotCallVoidRidVector2StringColorInt(o Class, methodName string, arg0 *RID, arg1 *Vector2, arg2 string, arg3 *Color, arg4 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(stringAsGodotString(arg2))

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.color)

	C.add_element(cArgsArray, cArg3, C.int(3))

	primArg4 := C.godot_int(arg4)
	cArg4 := unsafe.Pointer(&primArg4)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidTransform2DVector2(o Class, methodName string, arg0 *Transform2D, arg1 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.transform2d)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidRidFloat(o Class, methodName string, arg0 *RID, arg1 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_real(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidRidRect2BoolColorBoolObject(o Class, methodName string, arg0 *RID, arg1 *Rect2, arg2 bool, arg3 *Color, arg4 bool, arg5 *Object) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(6))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.rect2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.color)

	C.add_element(cArgsArray, cArg3, C.int(3))

	primArg4 := C.godot_bool(arg4)
	cArg4 := unsafe.Pointer(&primArg4)

	C.add_element(cArgsArray, cArg4, C.int(4))

	cArg5 := unsafe.Pointer(arg5.owner)

	C.add_element(cArgsArray, cArg5, C.int(5))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidIntObjectString(o Class, methodName string, arg0 int64, arg1 *Object, arg2 string) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(stringAsGodotString(arg2))

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidFloatFloatFloat(o Class, methodName string, arg0 float64, arg1 float64, arg2 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_real(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_real(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_real(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallArrayString(o Class, methodName string, arg0 string) *Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Array{&ret}

}
func godotCallBoolFloat(o Class, methodName string, arg0 float64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_real(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallVoidIntIntObject(o Class, methodName string, arg0 int64, arg1 int64, arg2 *Object) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.owner)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidPoolStringArrayInt(o Class, methodName string, arg0 *PoolStringArray, arg1 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.array)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallTransform2DRidInt(o Class, methodName string, arg0 *RID, arg1 int64) *Transform2D {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_transform2d
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Transform2D{&ret}

}
func godotCallBoolStringIntIntIntInt(o Class, methodName string, arg0 string, arg1 int64, arg2 int64, arg3 int64, arg4 int64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	primArg4 := C.godot_int(arg4)
	cArg4 := unsafe.Pointer(&primArg4)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallBoolStringIntIntInt(o Class, methodName string, arg0 string, arg1 int64, arg2 int64, arg3 int64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallBoolStringInt(o Class, methodName string, arg0 string, arg1 int64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallVoidStringStringColor(o Class, methodName string, arg0 string, arg1 string, arg2 *Color) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.color)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVector2StringInt(o Class, methodName string, arg0 string, arg1 int64) *Vector2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_vector2
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector2{&ret}

}
func godotCallObjectTransform2DVector2(o Class, methodName string, arg0 *Transform2D, arg1 *Vector2) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.transform2d)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_object
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{&ret}

}
func godotCallVoidIntPoolRealArray(o Class, methodName string, arg0 int64, arg1 *PoolRealArray) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.array)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidRidRidTransform(o Class, methodName string, arg0 *RID, arg1 *RID, arg2 *Transform) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.rid)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.transform)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidVector3Float(o Class, methodName string, arg0 *Vector3, arg1 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector3)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_real(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallPoolVector2Array(o Class, methodName string) *PoolVector2Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_pool_vector2_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolVector2Array{&ret}

}
func godotCallVector3Vector3Vector3Bool(o Class, methodName string, arg0 *Vector3, arg1 *Vector3, arg2 bool) *Vector3 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector3)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector3)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_vector3
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector3{&ret}

}
func godotCallNodePathIntBool(o Class, methodName string, arg0 int64, arg1 bool) *NodePath {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_bool(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_node_path
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &NodePath{&ret}

}
func godotCallVoidRect2(o Class, methodName string, arg0 *Rect2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rect2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidIntObjectVector2(o Class, methodName string, arg0 int64, arg1 *Object, arg2 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.vector2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVector2Vector2Vector2FloatIntFloat(o Class, methodName string, arg0 *Vector2, arg1 *Vector2, arg2 float64, arg3 int64, arg4 float64) *Vector2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_real(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	primArg4 := C.godot_real(arg4)
	cArg4 := unsafe.Pointer(&primArg4)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_vector2
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector2{&ret}

}
func godotCallVoidIntFloatFloatFloatBool(o Class, methodName string, arg0 int64, arg1 float64, arg2 float64, arg3 float64, arg4 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_real(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_real(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_real(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	primArg4 := C.godot_bool(arg4)
	cArg4 := unsafe.Pointer(&primArg4)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallIntStringString(o Class, methodName string, arg0 string, arg1 string) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallVector2FloatBool(o Class, methodName string, arg0 float64, arg1 bool) *Vector2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_real(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_bool(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_vector2
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector2{&ret}

}
func godotCallVoidObjectRect2Vector2(o Class, methodName string, arg0 *Object, arg1 *Rect2, arg2 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.rect2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.vector2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidStringIntObject(o Class, methodName string, arg0 string, arg1 int64, arg2 *Object) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.owner)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidIntVector3Float(o Class, methodName string, arg0 int64, arg1 *Vector3, arg2 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector3)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_real(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidIntIntFloat(o Class, methodName string, arg0 int64, arg1 int64, arg2 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_real(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidStringInt(o Class, methodName string, arg0 string, arg1 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidStringDictionary(o Class, methodName string, arg0 string, arg1 *Dictionary) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.dictionary)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidIntInt(o Class, methodName string, arg0 int64, arg1 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallPoolVector3ArrayIntInt(o Class, methodName string, arg0 int64, arg1 int64) *PoolVector3Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_pool_vector3_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolVector3Array{&ret}

}
func godotCallVoidIntFloatVariantFloat(o Class, methodName string, arg0 int64, arg1 float64, arg2 *Variant, arg3 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_real(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.variant)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_real(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallBoolRidIntInt(o Class, methodName string, arg0 *RID, arg1 int64, arg2 int64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallVoidStringIntString(o Class, methodName string, arg0 string, arg1 int64, arg2 string) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(stringAsGodotString(arg2))

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidPoolVector2ArrayPoolColorArrayPoolVector2ArrayObjectFloatObject(o Class, methodName string, arg0 *PoolVector2Array, arg1 *PoolColorArray, arg2 *PoolVector2Array, arg3 *Object, arg4 float64, arg5 *Object) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(6))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.array)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.array)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.array)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.owner)

	C.add_element(cArgsArray, cArg3, C.int(3))

	primArg4 := C.godot_real(arg4)
	cArg4 := unsafe.Pointer(&primArg4)

	C.add_element(cArgsArray, cArg4, C.int(4))

	cArg5 := unsafe.Pointer(arg5.owner)

	C.add_element(cArgsArray, cArg5, C.int(5))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallIntStringStringDictionaryArrayArray(o Class, methodName string, arg0 string, arg1 string, arg2 *Dictionary, arg3 *Array, arg4 *Array) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.dictionary)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.array)

	C.add_element(cArgsArray, cArg3, C.int(3))

	cArg4 := unsafe.Pointer(arg4.array)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallBoolTransform2DVector2(o Class, methodName string, arg0 *Transform2D, arg1 *Vector2) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.transform2d)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallIntInt(o Class, methodName string, arg0 int64) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallVoidDictionary(o Class, methodName string, arg0 *Dictionary) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.dictionary)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidVector3Vector3Vector3Int(o Class, methodName string, arg0 *Vector3, arg1 *Vector3, arg2 *Vector3, arg3 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector3)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector3)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.vector3)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallBoolTransformVector3(o Class, methodName string, arg0 *Transform, arg1 *Vector3) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.transform)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector3)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallPoolVector3ArrayVector3Vector3Bool(o Class, methodName string, arg0 *Vector3, arg1 *Vector3, arg2 bool) *PoolVector3Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector3)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector3)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_pool_vector3_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolVector3Array{&ret}

}
func godotCallVoidString(o Class, methodName string, arg0 string) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallStringIntInt(o Class, methodName string, arg0 int64, arg1 int64) string {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_string
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return godotStringAsString(&ret)

}
func godotCallPoolVector3Array(o Class, methodName string) *PoolVector3Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_pool_vector3_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolVector3Array{&ret}

}
func godotCallColorInt(o Class, methodName string, arg0 int64) *Color {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_color
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Color{&ret}

}
func godotCallVoidRidVector2Vector2(o Class, methodName string, arg0 *RID, arg1 *Vector2, arg2 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.vector2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallArrayRidInt(o Class, methodName string, arg0 *RID, arg1 int64) *Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Array{&ret}

}
func godotCallPoolColorArray(o Class, methodName string) *PoolColorArray {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_pool_color_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolColorArray{&ret}

}
func godotCallVoidStringIntStringInt(o Class, methodName string, arg0 string, arg1 int64, arg2 string, arg3 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(stringAsGodotString(arg2))

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallFloatVector2(o Class, methodName string, arg0 *Vector2) float64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_real
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return float64(ret)

}
func godotCallVoidRidRect2Int(o Class, methodName string, arg0 *RID, arg1 *Rect2, arg2 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.rect2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVariantStringBool(o Class, methodName string, arg0 string, arg1 bool) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_bool(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_variant
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{&ret}

}
func godotCallArrayObject(o Class, methodName string, arg0 *Object) *Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Array{&ret}

}
func godotCallObjectFloatBool(o Class, methodName string, arg0 float64, arg1 bool) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_real(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_bool(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_object
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{&ret}

}
func godotCallVoidStringVector2(o Class, methodName string, arg0 string, arg1 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallRidInt(o Class, methodName string, arg0 int64) *RID {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_rid
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &RID{&ret}

}
func godotCallPlaneInt(o Class, methodName string, arg0 int64) *Plane {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_plane
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Plane{&ret}

}
func godotCallVoidRidVector3Vector3(o Class, methodName string, arg0 *RID, arg1 *Vector3, arg2 *Vector3) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector3)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.vector3)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallBoolVector3(o Class, methodName string, arg0 *Vector3) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector3)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallVoidIntColorBool(o Class, methodName string, arg0 int64, arg1 *Color, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.color)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidObjectRect2(o Class, methodName string, arg0 *Object, arg1 *Rect2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.rect2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidStringVariantBool(o Class, methodName string, arg0 string, arg1 *Variant, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.variant)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidIntArray(o Class, methodName string, arg0 int64, arg1 *Array) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.array)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallRidVector2Vector2RidRid(o Class, methodName string, arg0 *Vector2, arg1 *Vector2, arg2 *RID, arg3 *RID) *RID {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.rid)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.rid)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_rid
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &RID{&ret}

}
func godotCallVariantTransform2DObjectTransform2D(o Class, methodName string, arg0 *Transform2D, arg1 *Object, arg2 *Transform2D) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.transform2d)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.transform2d)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_variant
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{&ret}

}
func godotCallVariant(o Class, methodName string) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_variant
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{&ret}

}
func godotCallDictionaryObject(o Class, methodName string, arg0 *Object) *Dictionary {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_dictionary
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Dictionary{&ret}

}
func godotCallObjectRid(o Class, methodName string, arg0 *RID) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_object
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{&ret}

}
func godotCallVoidObjectObjectIntBool(o Class, methodName string, arg0 *Object, arg1 *Object, arg2 int64, arg3 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_bool(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallBoolObjectStringVariantVariantFloatIntIntFloat(o Class, methodName string, arg0 *Object, arg1 string, arg2 *Variant, arg3 *Variant, arg4 float64, arg5 int64, arg6 int64, arg7 float64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(8))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.variant)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.variant)

	C.add_element(cArgsArray, cArg3, C.int(3))

	primArg4 := C.godot_real(arg4)
	cArg4 := unsafe.Pointer(&primArg4)

	C.add_element(cArgsArray, cArg4, C.int(4))

	primArg5 := C.godot_int(arg5)
	cArg5 := unsafe.Pointer(&primArg5)

	C.add_element(cArgsArray, cArg5, C.int(5))

	primArg6 := C.godot_int(arg6)
	cArg6 := unsafe.Pointer(&primArg6)

	C.add_element(cArgsArray, cArg6, C.int(6))

	primArg7 := C.godot_real(arg7)
	cArg7 := unsafe.Pointer(&primArg7)

	C.add_element(cArgsArray, cArg7, C.int(7))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallTransformInt(o Class, methodName string, arg0 int64) *Transform {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_transform
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Transform{&ret}

}
func godotCallVector3IntFloat(o Class, methodName string, arg0 int64, arg1 float64) *Vector3 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_real(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_vector3
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector3{&ret}

}
func godotCallVoidIntVariant(o Class, methodName string, arg0 int64, arg1 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.variant)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVariantVector2Object(o Class, methodName string, arg0 *Vector2, arg1 *Object) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_variant
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{&ret}

}
func godotCallIntString(o Class, methodName string, arg0 string) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallIntStringObjectStringArrayInt(o Class, methodName string, arg0 string, arg1 *Object, arg2 string, arg3 *Array, arg4 int64) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(stringAsGodotString(arg2))

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.array)

	C.add_element(cArgsArray, cArg3, C.int(3))

	primArg4 := C.godot_int(arg4)
	cArg4 := unsafe.Pointer(&primArg4)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallVariantNodePath(o Class, methodName string, arg0 *NodePath) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.nodePath)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_variant
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{&ret}

}
func godotCallVector3Rid(o Class, methodName string, arg0 *RID) *Vector3 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_vector3
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector3{&ret}

}
func godotCallVoidStringIntVector2(o Class, methodName string, arg0 string, arg1 int64, arg2 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.vector2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallIntStringStringInt(o Class, methodName string, arg0 string, arg1 string, arg2 int64) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallVariantString(o Class, methodName string, arg0 string) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_variant
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{&ret}

}
func godotCallVoidPoolVector3ArrayPoolVector2ArrayPoolColorArrayPoolVector2ArrayPoolVector3ArrayArray(o Class, methodName string, arg0 *PoolVector3Array, arg1 *PoolVector2Array, arg2 *PoolColorArray, arg3 *PoolVector2Array, arg4 *PoolVector3Array, arg5 *Array) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(6))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.array)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.array)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.array)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.array)

	C.add_element(cArgsArray, cArg3, C.int(3))

	cArg4 := unsafe.Pointer(arg4.array)

	C.add_element(cArgsArray, cArg4, C.int(4))

	cArg5 := unsafe.Pointer(arg5.array)

	C.add_element(cArgsArray, cArg5, C.int(5))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidStringObjectIntStringVariant(o Class, methodName string, arg0 string, arg1 *Object, arg2 int64, arg3 string, arg4 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(stringAsGodotString(arg3))

	C.add_element(cArgsArray, cArg3, C.int(3))

	cArg4 := unsafe.Pointer(arg4.variant)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidPoolRealArray(o Class, methodName string, arg0 *PoolRealArray) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.array)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidStringString(o Class, methodName string, arg0 string, arg1 string) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallStringObject(o Class, methodName string, arg0 *Object) string {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_string
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return godotStringAsString(&ret)

}
func godotCallVoidStringBool(o Class, methodName string, arg0 string, arg1 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_bool(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallStringInt(o Class, methodName string, arg0 int64) string {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_string
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return godotStringAsString(&ret)

}
func godotCallVoidAabb(o Class, methodName string, arg0 *AABB) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.aabb)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidObjectRect2Rect2ColorBoolObjectBool(o Class, methodName string, arg0 *Object, arg1 *Rect2, arg2 *Rect2, arg3 *Color, arg4 bool, arg5 *Object, arg6 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(7))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.rect2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.rect2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.color)

	C.add_element(cArgsArray, cArg3, C.int(3))

	primArg4 := C.godot_bool(arg4)
	cArg4 := unsafe.Pointer(&primArg4)

	C.add_element(cArgsArray, cArg4, C.int(4))

	cArg5 := unsafe.Pointer(arg5.owner)

	C.add_element(cArgsArray, cArg5, C.int(5))

	primArg6 := C.godot_bool(arg6)
	cArg6 := unsafe.Pointer(&primArg6)

	C.add_element(cArgsArray, cArg6, C.int(6))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallPoolStringArrayString(o Class, methodName string, arg0 string) *PoolStringArray {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_pool_string_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolStringArray{&ret}

}
func godotCallVoidStringStringVariant(o Class, methodName string, arg0 string, arg1 string, arg2 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.variant)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVariantIntString(o Class, methodName string, arg0 int64, arg1 string) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_variant
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{&ret}

}
func godotCallVoidIntStringInt(o Class, methodName string, arg0 int64, arg1 string, arg2 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallTransform(o Class, methodName string) *Transform {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_transform
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Transform{&ret}

}
func godotCallVoidPoolByteArray(o Class, methodName string, arg0 *PoolByteArray) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.array)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallIntStringPoolStringArrayBoolIntString(o Class, methodName string, arg0 string, arg1 *PoolStringArray, arg2 bool, arg3 int64, arg4 string) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.array)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	cArg4 := unsafe.Pointer(stringAsGodotString(arg4))

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallVoidNodePath(o Class, methodName string, arg0 *NodePath) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.nodePath)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidVector2IntBoolBoolBool(o Class, methodName string, arg0 *Vector2, arg1 int64, arg2 bool, arg3 bool, arg4 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_bool(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	primArg4 := C.godot_bool(arg4)
	cArg4 := unsafe.Pointer(&primArg4)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallNodePath(o Class, methodName string) *NodePath {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_node_path
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &NodePath{&ret}

}
func godotCallVoidRidIntIntBool(o Class, methodName string, arg0 *RID, arg1 int64, arg2 int64, arg3 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_bool(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallBoolObjectNodePathVariantObjectNodePathFloatIntIntFloat(o Class, methodName string, arg0 *Object, arg1 *NodePath, arg2 *Variant, arg3 *Object, arg4 *NodePath, arg5 float64, arg6 int64, arg7 int64, arg8 float64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(9))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.nodePath)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.variant)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.owner)

	C.add_element(cArgsArray, cArg3, C.int(3))

	cArg4 := unsafe.Pointer(arg4.nodePath)

	C.add_element(cArgsArray, cArg4, C.int(4))

	primArg5 := C.godot_real(arg5)
	cArg5 := unsafe.Pointer(&primArg5)

	C.add_element(cArgsArray, cArg5, C.int(5))

	primArg6 := C.godot_int(arg6)
	cArg6 := unsafe.Pointer(&primArg6)

	C.add_element(cArgsArray, cArg6, C.int(6))

	primArg7 := C.godot_int(arg7)
	cArg7 := unsafe.Pointer(&primArg7)

	C.add_element(cArgsArray, cArg7, C.int(7))

	primArg8 := C.godot_real(arg8)
	cArg8 := unsafe.Pointer(&primArg8)

	C.add_element(cArgsArray, cArg8, C.int(8))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallIntIntInt(o Class, methodName string, arg0 int64, arg1 int64) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallVector3Vector2(o Class, methodName string, arg0 *Vector2) *Vector3 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_vector3
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector3{&ret}

}
func godotCallVoidStringPoolByteArrayBool(o Class, methodName string, arg0 string, arg1 *PoolByteArray, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.array)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidIntVariantBool(o Class, methodName string, arg0 int64, arg1 *Variant, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.variant)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidRidTransform(o Class, methodName string, arg0 *RID, arg1 *Transform) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.transform)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallIntRidInt(o Class, methodName string, arg0 *RID, arg1 int64) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallIntStringObject(o Class, methodName string, arg0 string, arg1 *Object) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallVoidPoolIntArray(o Class, methodName string, arg0 *PoolIntArray) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.array)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidStringStringObject(o Class, methodName string, arg0 string, arg1 string, arg2 *Object) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.owner)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidObjectColorBool(o Class, methodName string, arg0 *Object, arg1 *Color, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.color)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallStringRid(o Class, methodName string, arg0 *RID) string {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_string
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return godotStringAsString(&ret)

}
func godotCallVoidIntRidIntIntInt(o Class, methodName string, arg0 int64, arg1 *RID, arg2 int64, arg3 int64, arg4 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.rid)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	primArg4 := C.godot_int(arg4)
	cArg4 := unsafe.Pointer(&primArg4)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVariantStringStringArray(o Class, methodName string, arg0 string, arg1 string, arg2 *Array) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.array)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_variant
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{&ret}

}
func godotCallVoidNodePathVariant(o Class, methodName string, arg0 *NodePath, arg1 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.nodePath)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.variant)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidRidRect2Rect2ColorBoolObjectBool(o Class, methodName string, arg0 *RID, arg1 *Rect2, arg2 *Rect2, arg3 *Color, arg4 bool, arg5 *Object, arg6 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(7))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.rect2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.rect2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.color)

	C.add_element(cArgsArray, cArg3, C.int(3))

	primArg4 := C.godot_bool(arg4)
	cArg4 := unsafe.Pointer(&primArg4)

	C.add_element(cArgsArray, cArg4, C.int(4))

	cArg5 := unsafe.Pointer(arg5.owner)

	C.add_element(cArgsArray, cArg5, C.int(5))

	primArg6 := C.godot_bool(arg6)
	cArg6 := unsafe.Pointer(&primArg6)

	C.add_element(cArgsArray, cArg6, C.int(6))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidRidRect2RidBoolColorBoolRid(o Class, methodName string, arg0 *RID, arg1 *Rect2, arg2 *RID, arg3 bool, arg4 *Color, arg5 bool, arg6 *RID) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(7))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.rect2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.rid)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_bool(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	cArg4 := unsafe.Pointer(arg4.color)

	C.add_element(cArgsArray, cArg4, C.int(4))

	primArg5 := C.godot_bool(arg5)
	cArg5 := unsafe.Pointer(&primArg5)

	C.add_element(cArgsArray, cArg5, C.int(5))

	cArg6 := unsafe.Pointer(arg6.rid)

	C.add_element(cArgsArray, cArg6, C.int(6))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidRidPoolVector2ArrayBool(o Class, methodName string, arg0 *RID, arg1 *PoolVector2Array, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.array)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidRidStringVariant(o Class, methodName string, arg0 *RID, arg1 string, arg2 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.variant)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidIntFloat(o Class, methodName string, arg0 int64, arg1 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_real(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallIntVector2(o Class, methodName string, arg0 *Vector2) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallVoidIntColor(o Class, methodName string, arg0 int64, arg1 *Color) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.color)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidRid(o Class, methodName string, arg0 *RID) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidObjectObjectBool(o Class, methodName string, arg0 *Object, arg1 *Object, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidRidBoolRect2(o Class, methodName string, arg0 *RID, arg1 bool, arg2 *Rect2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_bool(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.rect2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVector3(o Class, methodName string) *Vector3 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_vector3
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector3{&ret}

}
func godotCallBoolObjectString(o Class, methodName string, arg0 *Object, arg1 string) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallVoidObjectStringInt(o Class, methodName string, arg0 *Object, arg1 string, arg2 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallIntRid(o Class, methodName string, arg0 *RID) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallVoidRidBool(o Class, methodName string, arg0 *RID, arg1 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_bool(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallPoolByteArrayRidInt(o Class, methodName string, arg0 *RID, arg1 int64) *PoolByteArray {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_pool_byte_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolByteArray{&ret}

}
func godotCallNodePathObject(o Class, methodName string, arg0 *Object) *NodePath {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_node_path
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &NodePath{&ret}

}
func godotCallFloatRid(o Class, methodName string, arg0 *RID) float64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_real
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return float64(ret)

}
func godotCallRidRidTransformRidTransform(o Class, methodName string, arg0 *RID, arg1 *Transform, arg2 *RID, arg3 *Transform) *RID {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.transform)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.rid)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.transform)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_rid
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &RID{&ret}

}
func godotCallVoidIntObject(o Class, methodName string, arg0 int64, arg1 *Object) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVector2IntFloat(o Class, methodName string, arg0 int64, arg1 float64) *Vector2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_real(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_vector2
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector2{&ret}

}
func godotCallBoolObjectObject(o Class, methodName string, arg0 *Object, arg1 *Object) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallBoolRidTransform2DVector2FloatObject(o Class, methodName string, arg0 *RID, arg1 *Transform2D, arg2 *Vector2, arg3 float64, arg4 *Object) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.transform2d)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.vector2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_real(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	cArg4 := unsafe.Pointer(arg4.owner)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallPoolIntArrayInt(o Class, methodName string, arg0 int64) *PoolIntArray {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_pool_int_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolIntArray{&ret}

}
func godotCallVoidFloatBool(o Class, methodName string, arg0 float64, arg1 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_real(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_bool(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallObjectVector2(o Class, methodName string, arg0 *Vector2) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_object
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{&ret}

}
func godotCallObjectStringString(o Class, methodName string, arg0 string, arg1 string) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_object
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{&ret}

}
func godotCallVoidRidVariant(o Class, methodName string, arg0 *RID, arg1 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.variant)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidStringFloat(o Class, methodName string, arg0 string, arg1 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_real(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVector2Float(o Class, methodName string, arg0 float64) *Vector2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_real(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_vector2
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector2{&ret}

}
func godotCallVariantVariant(o Class, methodName string, arg0 *Variant) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.variant)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_variant
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{&ret}

}
func godotCallVoidInt(o Class, methodName string, arg0 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallPoolVector2ArrayVector2Vector2Bool(o Class, methodName string, arg0 *Vector2, arg1 *Vector2, arg2 bool) *PoolVector2Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_pool_vector2_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolVector2Array{&ret}

}
func godotCallVoidRidRect2RidRect2ColorBoolRidBool(o Class, methodName string, arg0 *RID, arg1 *Rect2, arg2 *RID, arg3 *Rect2, arg4 *Color, arg5 bool, arg6 *RID, arg7 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(8))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.rect2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.rid)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.rect2)

	C.add_element(cArgsArray, cArg3, C.int(3))

	cArg4 := unsafe.Pointer(arg4.color)

	C.add_element(cArgsArray, cArg4, C.int(4))

	primArg5 := C.godot_bool(arg5)
	cArg5 := unsafe.Pointer(&primArg5)

	C.add_element(cArgsArray, cArg5, C.int(5))

	cArg6 := unsafe.Pointer(arg6.rid)

	C.add_element(cArgsArray, cArg6, C.int(6))

	primArg7 := C.godot_bool(arg7)
	cArg7 := unsafe.Pointer(&primArg7)

	C.add_element(cArgsArray, cArg7, C.int(7))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallPoolByteArray(o Class, methodName string) *PoolByteArray {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_pool_byte_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolByteArray{&ret}

}
func godotCallArrayNodePath(o Class, methodName string, arg0 *NodePath) *Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.nodePath)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Array{&ret}

}
func godotCallVoidStringArrayBool(o Class, methodName string, arg0 string, arg1 *Array, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.array)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallBool(o Class, methodName string) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallStringVector2(o Class, methodName string, arg0 *Vector2) string {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_string
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return godotStringAsString(&ret)

}
func godotCallVoidIntTransform(o Class, methodName string, arg0 int64, arg1 *Transform) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.transform)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidVariantObject(o Class, methodName string, arg0 *Variant, arg1 *Object) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.variant)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidIntObjectIntBoolString(o Class, methodName string, arg0 int64, arg1 *Object, arg2 int64, arg3 bool, arg4 string) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_bool(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	cArg4 := unsafe.Pointer(stringAsGodotString(arg4))

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallFloatIntInt(o Class, methodName string, arg0 int64, arg1 int64) float64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_real
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return float64(ret)

}
func godotCallVoidStringFloatFloatBool(o Class, methodName string, arg0 string, arg1 float64, arg2 float64, arg3 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_real(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_real(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_bool(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallIntPoolByteArray(o Class, methodName string, arg0 *PoolByteArray) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.array)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallVoidStringStringColorBool(o Class, methodName string, arg0 string, arg1 string, arg2 *Color, arg3 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.color)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_bool(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallArrayArrayInt(o Class, methodName string, arg0 *Array, arg1 int64) *Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.array)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Array{&ret}

}
func godotCallVoidIntIntFloatBool(o Class, methodName string, arg0 int64, arg1 int64, arg2 float64, arg3 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_real(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_bool(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallIntIntStringPoolStringArrayPoolByteArray(o Class, methodName string, arg0 int64, arg1 string, arg2 *PoolStringArray, arg3 *PoolByteArray) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.array)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.array)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallVariantRidInt(o Class, methodName string, arg0 *RID, arg1 int64) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_variant
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{&ret}

}
func godotCallVoidBoolVector2Vector2(o Class, methodName string, arg0 bool, arg1 *Vector2, arg2 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_bool(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.vector2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidIntIntInt(o Class, methodName string, arg0 int64, arg1 int64, arg2 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallObjectObject(o Class, methodName string, arg0 *Object) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_object
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{&ret}

}
func godotCallVector3Vector3Vector3FloatIntFloat(o Class, methodName string, arg0 *Vector3, arg1 *Vector3, arg2 float64, arg3 int64, arg4 float64) *Vector3 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector3)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector3)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_real(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	primArg4 := C.godot_real(arg4)
	cArg4 := unsafe.Pointer(&primArg4)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_vector3
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector3{&ret}

}
func godotCallIntStringIntIntInt(o Class, methodName string, arg0 string, arg1 int64, arg2 int64, arg3 int64) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallObject(o Class, methodName string) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_object
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{&ret}

}
func godotCallStringStringInt(o Class, methodName string, arg0 string, arg1 int64) string {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_string
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return godotStringAsString(&ret)

}
func godotCallVoidIntObjectInt(o Class, methodName string, arg0 int64, arg1 *Object, arg2 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidArray(o Class, methodName string, arg0 *Array) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.array)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallColorFloat(o Class, methodName string, arg0 float64) *Color {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_real(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_color
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Color{&ret}

}
func godotCallVariantIntStringString(o Class, methodName string, arg0 int64, arg1 string, arg2 string) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(stringAsGodotString(arg2))

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_variant
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{&ret}

}
func godotCallVoidRidStringRid(o Class, methodName string, arg0 *RID, arg1 string, arg2 *RID) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.rid)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidFloat(o Class, methodName string, arg0 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_real(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallAabbRidInt(o Class, methodName string, arg0 *RID, arg1 int64) *AABB {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_aabb
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &AABB{&ret}

}
func godotCallRect2Int(o Class, methodName string, arg0 int64) *Rect2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_rect2
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Rect2{&ret}

}
func godotCallObjectBool(o Class, methodName string, arg0 bool) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_bool(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_object
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{&ret}

}
func godotCallIntVector2FloatFloatIntInt(o Class, methodName string, arg0 *Vector2, arg1 float64, arg2 float64, arg3 int64, arg4 int64) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_real(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_real(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	primArg4 := C.godot_int(arg4)
	cArg4 := unsafe.Pointer(&primArg4)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallVoidObjectInt(o Class, methodName string, arg0 *Object, arg1 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallRidRid(o Class, methodName string, arg0 *RID) *RID {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_rid
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &RID{&ret}

}
func godotCallVoidIntBoolBool(o Class, methodName string, arg0 int64, arg1 bool, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_bool(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVector2Vector2Bool(o Class, methodName string, arg0 *Vector2, arg1 bool) *Vector2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_bool(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_vector2
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector2{&ret}

}
func godotCallVoidRidPoolVector2Array(o Class, methodName string, arg0 *RID, arg1 *PoolVector2Array) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.array)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidRidString(o Class, methodName string, arg0 *RID, arg1 string) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidObjectBoolRid(o Class, methodName string, arg0 *Object, arg1 bool, arg2 *RID) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_bool(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.rid)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallIntStringIntBoolBool(o Class, methodName string, arg0 string, arg1 int64, arg2 bool, arg3 bool) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_bool(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallBoolRid(o Class, methodName string, arg0 *RID) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallInt(o Class, methodName string) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallVoidStringNodePathBool(o Class, methodName string, arg0 string, arg1 *NodePath, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.nodePath)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallString(o Class, methodName string) string {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_string
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return godotStringAsString(&ret)

}
func godotCallVoidIntVector3(o Class, methodName string, arg0 int64, arg1 *Vector3) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector3)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidPoolVector2ArrayPoolColorArrayFloatBool(o Class, methodName string, arg0 *PoolVector2Array, arg1 *PoolColorArray, arg2 float64, arg3 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.array)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.array)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_real(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_bool(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallDictionaryVector2Vector2ArrayInt(o Class, methodName string, arg0 *Vector2, arg1 *Vector2, arg2 *Array, arg3 int64) *Dictionary {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.array)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_dictionary
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Dictionary{&ret}

}
func godotCallIntNodePath(o Class, methodName string, arg0 *NodePath) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.nodePath)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallVoidObject(o Class, methodName string, arg0 *Object) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidPoolVector2ArrayColorFloatBool(o Class, methodName string, arg0 *PoolVector2Array, arg1 *Color, arg2 float64, arg3 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.array)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.color)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_real(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_bool(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidVector2Vector2Vector2Int(o Class, methodName string, arg0 *Vector2, arg1 *Vector2, arg2 *Vector2, arg3 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.vector2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVariantInt(o Class, methodName string, arg0 int64) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_variant
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{&ret}

}
func godotCallObjectStringBoolBool(o Class, methodName string, arg0 string, arg1 bool, arg2 bool) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_bool(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_object
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{&ret}

}
func godotCallArrayObjectInt(o Class, methodName string, arg0 *Object, arg1 int64) *Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Array{&ret}

}
func godotCallPlane(o Class, methodName string) *Plane {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_plane
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Plane{&ret}

}
func godotCallVector2IntInt(o Class, methodName string, arg0 int64, arg1 int64) *Vector2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_vector2
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector2{&ret}

}
func godotCallVoidPoolVector3ArrayBoolBool(o Class, methodName string, arg0 *PoolVector3Array, arg1 bool, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.array)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_bool(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidPoolVector2ArrayPoolIntArray(o Class, methodName string, arg0 *PoolVector2Array, arg1 *PoolIntArray) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.array)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.array)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallRidVector2Vector2Vector2RidRid(o Class, methodName string, arg0 *Vector2, arg1 *Vector2, arg2 *Vector2, arg3 *RID, arg4 *RID) *RID {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.vector2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.rid)

	C.add_element(cArgsArray, cArg3, C.int(3))

	cArg4 := unsafe.Pointer(arg4.rid)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_rid
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &RID{&ret}

}
func godotCallVariantStringString(o Class, methodName string, arg0 string, arg1 string) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_variant
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{&ret}

}
func godotCallVoidBoolBool(o Class, methodName string, arg0 bool, arg1 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_bool(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_bool(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallObjectObjectInt(o Class, methodName string, arg0 *Object, arg1 int64) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_object
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{&ret}

}
func godotCallVariantObjectString(o Class, methodName string, arg0 *Object, arg1 string) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_variant
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{&ret}

}
func godotCallRidRidString(o Class, methodName string, arg0 *RID, arg1 string) *RID {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_rid
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &RID{&ret}

}
func godotCallArrayRid(o Class, methodName string, arg0 *RID) *Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Array{&ret}

}
func godotCallArrayIntFloat(o Class, methodName string, arg0 int64, arg1 float64) *Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_real(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Array{&ret}

}
func godotCallVoidRidIntFloat(o Class, methodName string, arg0 *RID, arg1 int64, arg2 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_real(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVariantRid(o Class, methodName string, arg0 *RID) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_variant
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{&ret}

}
func godotCallBoolObjectNodePathObjectNodePathVariantFloatIntIntFloat(o Class, methodName string, arg0 *Object, arg1 *NodePath, arg2 *Object, arg3 *NodePath, arg4 *Variant, arg5 float64, arg6 int64, arg7 int64, arg8 float64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(9))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.nodePath)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.owner)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.nodePath)

	C.add_element(cArgsArray, cArg3, C.int(3))

	cArg4 := unsafe.Pointer(arg4.variant)

	C.add_element(cArgsArray, cArg4, C.int(4))

	primArg5 := C.godot_real(arg5)
	cArg5 := unsafe.Pointer(&primArg5)

	C.add_element(cArgsArray, cArg5, C.int(5))

	primArg6 := C.godot_int(arg6)
	cArg6 := unsafe.Pointer(&primArg6)

	C.add_element(cArgsArray, cArg6, C.int(6))

	primArg7 := C.godot_int(arg7)
	cArg7 := unsafe.Pointer(&primArg7)

	C.add_element(cArgsArray, cArg7, C.int(7))

	primArg8 := C.godot_real(arg8)
	cArg8 := unsafe.Pointer(&primArg8)

	C.add_element(cArgsArray, cArg8, C.int(8))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallVoidIntIntColor(o Class, methodName string, arg0 int64, arg1 int64, arg2 *Color) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.color)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidStringObjectString(o Class, methodName string, arg0 string, arg1 *Object, arg2 string) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(stringAsGodotString(arg2))

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallTransformRid(o Class, methodName string, arg0 *RID) *Transform {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_transform
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Transform{&ret}

}
func godotCallVoidPlane(o Class, methodName string, arg0 *Plane) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.plane)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVariantIntInt(o Class, methodName string, arg0 int64, arg1 int64) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_variant
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{&ret}

}
func godotCallVoidVector2FloatVector2(o Class, methodName string, arg0 *Vector2, arg1 float64, arg2 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_real(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.vector2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidPoolVector2Array(o Class, methodName string, arg0 *PoolVector2Array) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.array)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallObjectRect2(o Class, methodName string, arg0 *Rect2) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rect2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_object
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{&ret}

}
func godotCallVoidIntPoolVector2Array(o Class, methodName string, arg0 int64, arg1 *PoolVector2Array) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.array)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidRidVector2(o Class, methodName string, arg0 *RID, arg1 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidObjectStringBool(o Class, methodName string, arg0 *Object, arg1 string, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidRidRidRidRidIntInt(o Class, methodName string, arg0 *RID, arg1 *RID, arg2 *RID, arg3 *RID, arg4 int64, arg5 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(6))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.rid)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.rid)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.rid)

	C.add_element(cArgsArray, cArg3, C.int(3))

	primArg4 := C.godot_int(arg4)
	cArg4 := unsafe.Pointer(&primArg4)

	C.add_element(cArgsArray, cArg4, C.int(4))

	primArg5 := C.godot_int(arg5)
	cArg5 := unsafe.Pointer(&primArg5)

	C.add_element(cArgsArray, cArg5, C.int(5))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVector3Vector3(o Class, methodName string, arg0 *Vector3) *Vector3 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector3)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_vector3
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector3{&ret}

}
func godotCallVoidObjectStringPoolStringArray(o Class, methodName string, arg0 *Object, arg1 string, arg2 *PoolStringArray) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.array)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidRidRidRidRid(o Class, methodName string, arg0 *RID, arg1 *RID, arg2 *RID, arg3 *RID) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.rid)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.rid)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.rid)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidRect2ColorBool(o Class, methodName string, arg0 *Rect2, arg1 *Color, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rect2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.color)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallPoolIntArrayStringIntIntInt(o Class, methodName string, arg0 string, arg1 int64, arg2 int64, arg3 int64) *PoolIntArray {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_pool_int_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolIntArray{&ret}

}
func godotCallVoidObjectStringVariant(o Class, methodName string, arg0 *Object, arg1 string, arg2 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.variant)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallIntIntFloatVector3QuatVector3(o Class, methodName string, arg0 int64, arg1 float64, arg2 *Vector3, arg3 *Quat, arg4 *Vector3) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_real(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.vector3)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.quat)

	C.add_element(cArgsArray, cArg3, C.int(3))

	cArg4 := unsafe.Pointer(arg4.vector3)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallStringDictionary(o Class, methodName string, arg0 *Dictionary) string {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.dictionary)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_string
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return godotStringAsString(&ret)

}
func godotCallVoidIntIntBoolInt(o Class, methodName string, arg0 int64, arg1 int64, arg2 bool, arg3 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallFloatInt(o Class, methodName string, arg0 int64) float64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_real
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return float64(ret)

}
func godotCallRect2(o Class, methodName string) *Rect2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_rect2
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Rect2{&ret}

}
func godotCallRidVector2RidRid(o Class, methodName string, arg0 *Vector2, arg1 *RID, arg2 *RID) *RID {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.rid)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.rid)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_rid
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &RID{&ret}

}
func godotCallVoidStringIntIntInt(o Class, methodName string, arg0 string, arg1 int64, arg2 int64, arg3 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidVector3(o Class, methodName string, arg0 *Vector3) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector3)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallIntObjectInt(o Class, methodName string, arg0 *Object, arg1 int64) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallVoidIntStringVariant(o Class, methodName string, arg0 int64, arg1 string, arg2 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.variant)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallObjectStringInt(o Class, methodName string, arg0 string, arg1 int64) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_object
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{&ret}

}
func godotCallVoidIntBool(o Class, methodName string, arg0 int64, arg1 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_bool(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidStringIntBool(o Class, methodName string, arg0 string, arg1 int64, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallDictionaryVector3Vector3ArrayInt(o Class, methodName string, arg0 *Vector3, arg1 *Vector3, arg2 *Array, arg3 int64) *Dictionary {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector3)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector3)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.array)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_dictionary
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Dictionary{&ret}

}
func godotCallVoidRidVector2ColorBoolObject(o Class, methodName string, arg0 *RID, arg1 *Vector2, arg2 *Color, arg3 bool, arg4 *Object) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.color)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_bool(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	cArg4 := unsafe.Pointer(arg4.owner)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallRidIntIntFloat(o Class, methodName string, arg0 int64, arg1 int64, arg2 float64) *RID {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_real(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_rid
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &RID{&ret}

}
func godotCallFloat(o Class, methodName string) float64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_real
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return float64(ret)

}
func godotCallVoidRidIntVariant(o Class, methodName string, arg0 *RID, arg1 int64, arg2 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.variant)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidVector2Vector2(o Class, methodName string, arg0 *Vector2, arg1 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallIntObjectBoolString(o Class, methodName string, arg0 *Object, arg1 bool, arg2 string) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_bool(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(stringAsGodotString(arg2))

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallVoidRidVector2Vector2ColorFloatBool(o Class, methodName string, arg0 *RID, arg1 *Vector2, arg2 *Vector2, arg3 *Color, arg4 float64, arg5 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(6))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.vector2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.color)

	C.add_element(cArgsArray, cArg3, C.int(3))

	primArg4 := C.godot_real(arg4)
	cArg4 := unsafe.Pointer(&primArg4)

	C.add_element(cArgsArray, cArg4, C.int(4))

	primArg5 := C.godot_bool(arg5)
	cArg5 := unsafe.Pointer(&primArg5)

	C.add_element(cArgsArray, cArg5, C.int(5))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallObjectFloat(o Class, methodName string, arg0 float64) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_real(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_object
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{&ret}

}
func godotCallVoidStringStringInt(o Class, methodName string, arg0 string, arg1 string, arg2 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidVector2VariantObject(o Class, methodName string, arg0 *Vector2, arg1 *Variant, arg2 *Object) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.variant)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.owner)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidObjectStringArray(o Class, methodName string, arg0 *Object, arg1 string, arg2 *Array) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.array)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidRidIntIntIntInt(o Class, methodName string, arg0 *RID, arg1 int64, arg2 int64, arg3 int64, arg4 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	primArg4 := C.godot_int(arg4)
	cArg4 := unsafe.Pointer(&primArg4)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidStringColor(o Class, methodName string, arg0 string, arg1 *Color) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.color)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallDictionaryBool(o Class, methodName string, arg0 bool) *Dictionary {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_bool(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_dictionary
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Dictionary{&ret}

}
func godotCallVoidStringArray(o Class, methodName string, arg0 string, arg1 *Array) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.array)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallBoolTransform2DObjectTransform2D(o Class, methodName string, arg0 *Transform2D, arg1 *Object, arg2 *Transform2D) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.transform2d)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.transform2d)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallNodePathInt(o Class, methodName string, arg0 int64) *NodePath {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_node_path
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &NodePath{&ret}

}
func godotCallObjectVector3(o Class, methodName string, arg0 *Vector3) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector3)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_object
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{&ret}

}
func godotCallTransformRidInt(o Class, methodName string, arg0 *RID, arg1 int64) *Transform {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_transform
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Transform{&ret}

}
func godotCallStringStringStringBoolIntInt(o Class, methodName string, arg0 string, arg1 string, arg2 bool, arg3 int64, arg4 int64) string {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	primArg4 := C.godot_int(arg4)
	cArg4 := unsafe.Pointer(&primArg4)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_string
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return godotStringAsString(&ret)

}
func godotCallIntIntIntInt(o Class, methodName string, arg0 int64, arg1 int64, arg2 int64) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallIntBool(o Class, methodName string, arg0 bool) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_bool(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallRidObjectInt(o Class, methodName string, arg0 *Object, arg1 int64) *RID {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_rid
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &RID{&ret}

}
func godotCallIntVector3(o Class, methodName string, arg0 *Vector3) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector3)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallVoidColor(o Class, methodName string, arg0 *Color) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.color)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallColorStringString(o Class, methodName string, arg0 string, arg1 string) *Color {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_color
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Color{&ret}

}
func godotCallPoolVector3ArrayIntFloat(o Class, methodName string, arg0 int64, arg1 float64) *PoolVector3Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_real(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_pool_vector3_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolVector3Array{&ret}

}
func godotCallVoidObjectFloat(o Class, methodName string, arg0 *Object, arg1 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_real(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallIntIntIntFloat(o Class, methodName string, arg0 int64, arg1 int64, arg2 float64) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_real(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallVoidRidRidTransform2D(o Class, methodName string, arg0 *RID, arg1 *RID, arg2 *Transform2D) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.rid)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.transform2d)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallObjectRidInt(o Class, methodName string, arg0 *RID, arg1 int64) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_object
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{&ret}

}
func godotCallFloatStringString(o Class, methodName string, arg0 string, arg1 string) float64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_real
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return float64(ret)

}
func godotCallIntIntFloat(o Class, methodName string, arg0 int64, arg1 float64) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_real(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallVoidStringObjectVariant(o Class, methodName string, arg0 string, arg1 *Object, arg2 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.variant)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallPoolStringArrayInt(o Class, methodName string, arg0 int64) *PoolStringArray {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_pool_string_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolStringArray{&ret}

}
func godotCallRidRidVector3RidVector3(o Class, methodName string, arg0 *RID, arg1 *Vector3, arg2 *RID, arg3 *Vector3) *RID {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector3)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.rid)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.vector3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_rid
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &RID{&ret}

}
func godotCallVoidObjectStringIntInt(o Class, methodName string, arg0 *Object, arg1 string, arg2 int64, arg3 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallTransform2DIntInt(o Class, methodName string, arg0 int64, arg1 int64) *Transform2D {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_transform2d
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Transform2D{&ret}

}
func godotCallObjectString(o Class, methodName string, arg0 string) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_object
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{&ret}

}
func godotCallPoolIntArrayIntInt(o Class, methodName string, arg0 int64, arg1 int64) *PoolIntArray {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_pool_int_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolIntArray{&ret}

}
func godotCallVoidStringObject(o Class, methodName string, arg0 string, arg1 *Object) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallBoolBool(o Class, methodName string, arg0 bool) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_bool(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallVoidIntNodePath(o Class, methodName string, arg0 int64, arg1 *NodePath) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.nodePath)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallBoolStringIntStringInt(o Class, methodName string, arg0 string, arg1 int64, arg2 string, arg3 int64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(stringAsGodotString(arg2))

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallVoidRidRid(o Class, methodName string, arg0 *RID, arg1 *RID) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.rid)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidIntIntVector2Float(o Class, methodName string, arg0 int64, arg1 int64, arg2 *Vector2, arg3 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.vector2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_real(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallBoolNodePath(o Class, methodName string, arg0 *NodePath) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.nodePath)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallRidIntBool(o Class, methodName string, arg0 int64, arg1 bool) *RID {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_bool(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_rid
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &RID{&ret}

}
func godotCallVoidRidRidInt(o Class, methodName string, arg0 *RID, arg1 *RID, arg2 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.rid)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallObjectInt(o Class, methodName string, arg0 int64) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_object
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{&ret}

}
func godotCallRidRidInt(o Class, methodName string, arg0 *RID, arg1 int64) *RID {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_rid
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &RID{&ret}

}
func godotCallVoidVector3Vector3Vector3(o Class, methodName string, arg0 *Vector3, arg1 *Vector3, arg2 *Vector3) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector3)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector3)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.vector3)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidRidRidRid(o Class, methodName string, arg0 *RID, arg1 *RID, arg2 *RID) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.rid)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.rid)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallPoolVector2ArrayInt(o Class, methodName string, arg0 int64) *PoolVector2Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_pool_vector2_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolVector2Array{&ret}

}
func godotCallBoolVector2VariantObject(o Class, methodName string, arg0 *Vector2, arg1 *Variant, arg2 *Object) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.variant)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.owner)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallTransformBool(o Class, methodName string, arg0 bool) *Transform {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_bool(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_transform
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Transform{&ret}

}
func godotCallBoolObject(o Class, methodName string, arg0 *Object) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallVoidIntFloatBoolBool(o Class, methodName string, arg0 int64, arg1 float64, arg2 bool, arg3 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_real(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_bool(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidObjectObjectStringVariant(o Class, methodName string, arg0 *Object, arg1 *Object, arg2 string, arg3 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(stringAsGodotString(arg2))

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.variant)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallObjectObjectString(o Class, methodName string, arg0 *Object, arg1 string) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_object
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{&ret}

}
func godotCallBoolStringObjectString(o Class, methodName string, arg0 string, arg1 *Object, arg2 string) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(stringAsGodotString(arg2))

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallVoidRidRect2Rect2RidVector2Vector2IntIntBoolColorRid(o Class, methodName string, arg0 *RID, arg1 *Rect2, arg2 *Rect2, arg3 *RID, arg4 *Vector2, arg5 *Vector2, arg6 int64, arg7 int64, arg8 bool, arg9 *Color, arg10 *RID) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(11))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.rect2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.rect2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.rid)

	C.add_element(cArgsArray, cArg3, C.int(3))

	cArg4 := unsafe.Pointer(arg4.vector2)

	C.add_element(cArgsArray, cArg4, C.int(4))

	cArg5 := unsafe.Pointer(arg5.vector2)

	C.add_element(cArgsArray, cArg5, C.int(5))

	primArg6 := C.godot_int(arg6)
	cArg6 := unsafe.Pointer(&primArg6)

	C.add_element(cArgsArray, cArg6, C.int(6))

	primArg7 := C.godot_int(arg7)
	cArg7 := unsafe.Pointer(&primArg7)

	C.add_element(cArgsArray, cArg7, C.int(7))

	primArg8 := C.godot_bool(arg8)
	cArg8 := unsafe.Pointer(&primArg8)

	C.add_element(cArgsArray, cArg8, C.int(8))

	cArg9 := unsafe.Pointer(arg9.color)

	C.add_element(cArgsArray, cArg9, C.int(9))

	cArg10 := unsafe.Pointer(arg10.rid)

	C.add_element(cArgsArray, cArg10, C.int(10))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallBoolVector2Variant(o Class, methodName string, arg0 *Vector2, arg1 *Variant) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.variant)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallVoidPoolStringArray(o Class, methodName string, arg0 *PoolStringArray) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.array)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidIntIntIntInt(o Class, methodName string, arg0 int64, arg1 int64, arg2 int64, arg3 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallBoolVector2Rect2(o Class, methodName string, arg0 *Vector2, arg1 *Rect2) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.rect2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallVoidIntIntVariant(o Class, methodName string, arg0 int64, arg1 int64, arg2 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.variant)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidObjectString(o Class, methodName string, arg0 *Object, arg1 string) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidRidObjectString(o Class, methodName string, arg0 *RID, arg1 *Object, arg2 string) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(stringAsGodotString(arg2))

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVariantTransform2DVector2ObjectTransform2DVector2(o Class, methodName string, arg0 *Transform2D, arg1 *Vector2, arg2 *Object, arg3 *Transform2D, arg4 *Vector2) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.transform2d)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.owner)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.transform2d)

	C.add_element(cArgsArray, cArg3, C.int(3))

	cArg4 := unsafe.Pointer(arg4.vector2)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_variant
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{&ret}

}
func godotCallFloatString(o Class, methodName string, arg0 string) float64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_real
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return float64(ret)

}
func godotCallDictionary(o Class, methodName string) *Dictionary {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_dictionary
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Dictionary{&ret}

}
func godotCallIntObject(o Class, methodName string, arg0 *Object) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallIntVector2Object(o Class, methodName string, arg0 *Vector2, arg1 *Object) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallVoidPoolVector2ArrayInt(o Class, methodName string, arg0 *PoolVector2Array, arg1 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.array)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidObjectIntTransform(o Class, methodName string, arg0 *Object, arg1 int64, arg2 *Transform) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.transform)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidIntIntTransform2D(o Class, methodName string, arg0 int64, arg1 int64, arg2 *Transform2D) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.transform2d)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallDictionaryString(o Class, methodName string, arg0 string) *Dictionary {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_dictionary
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Dictionary{&ret}

}
func godotCallVoidRidPoolVector2ArrayPoolColorArrayPoolVector2ArrayRidFloatRid(o Class, methodName string, arg0 *RID, arg1 *PoolVector2Array, arg2 *PoolColorArray, arg3 *PoolVector2Array, arg4 *RID, arg5 float64, arg6 *RID) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(7))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.array)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.array)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.array)

	C.add_element(cArgsArray, cArg3, C.int(3))

	cArg4 := unsafe.Pointer(arg4.rid)

	C.add_element(cArgsArray, cArg4, C.int(4))

	primArg5 := C.godot_real(arg5)
	cArg5 := unsafe.Pointer(&primArg5)

	C.add_element(cArgsArray, cArg5, C.int(5))

	cArg6 := unsafe.Pointer(arg6.rid)

	C.add_element(cArgsArray, cArg6, C.int(6))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidBoolFloat(o Class, methodName string, arg0 bool, arg1 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_bool(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_real(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallIntObjectTransform2DObject(o Class, methodName string, arg0 *Object, arg1 *Transform2D, arg2 *Object) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.transform2d)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.owner)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallDictionaryInt(o Class, methodName string, arg0 int64) *Dictionary {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_dictionary
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Dictionary{&ret}

}
func godotCallVoidRidObjectInt(o Class, methodName string, arg0 *RID, arg1 *Object, arg2 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallPoolStringArray(o Class, methodName string) *PoolStringArray {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_pool_string_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolStringArray{&ret}

}
func godotCallObjectStringIntInt(o Class, methodName string, arg0 string, arg1 int64, arg2 int64) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_object
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{&ret}

}
func godotCallVector2(o Class, methodName string) *Vector2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_vector2
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector2{&ret}

}
func godotCallBoolStringVariant(o Class, methodName string, arg0 string, arg1 *Variant) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.variant)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallFloatRidIntInt(o Class, methodName string, arg0 *RID, arg1 int64, arg2 int64) float64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_real
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return float64(ret)

}
func godotCallVoidRidPoolVector2ArrayPoolColorArrayFloatBool(o Class, methodName string, arg0 *RID, arg1 *PoolVector2Array, arg2 *PoolColorArray, arg3 float64, arg4 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.array)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.array)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_real(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	primArg4 := C.godot_bool(arg4)
	cArg4 := unsafe.Pointer(&primArg4)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallAabbRid(o Class, methodName string, arg0 *RID) *AABB {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_aabb
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &AABB{&ret}

}
func godotCallVoidIntIntPoolByteArray(o Class, methodName string, arg0 int64, arg1 int64, arg2 *PoolByteArray) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.array)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidVariant(o Class, methodName string, arg0 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.variant)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallStringStringObject(o Class, methodName string, arg0 string, arg1 *Object) string {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_string
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return godotStringAsString(&ret)

}
func godotCallIntObjectBool(o Class, methodName string, arg0 *Object, arg1 bool) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_bool(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallTransform2D(o Class, methodName string) *Transform2D {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_transform2d
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Transform2D{&ret}

}
func godotCallVoidRidTransform2D(o Class, methodName string, arg0 *RID, arg1 *Transform2D) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.transform2d)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidVector2(o Class, methodName string, arg0 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidRidIntTransform(o Class, methodName string, arg0 *RID, arg1 int64, arg2 *Transform) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.transform)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidStringIntObjectVector2(o Class, methodName string, arg0 string, arg1 int64, arg2 *Object, arg3 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.owner)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.vector2)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidBool(o Class, methodName string, arg0 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_bool(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidIntVector2(o Class, methodName string, arg0 int64, arg1 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidStringPoolStringArray(o Class, methodName string, arg0 string, arg1 *PoolStringArray) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.array)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVector3Int(o Class, methodName string, arg0 int64) *Vector3 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_vector3
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector3{&ret}

}
func godotCallVoidRect2Bool(o Class, methodName string, arg0 *Rect2, arg1 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rect2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_bool(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallPoolRealArray(o Class, methodName string) *PoolRealArray {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_pool_real_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolRealArray{&ret}

}
func godotCallBoolObjectStringVariantObjectStringFloatIntIntFloat(o Class, methodName string, arg0 *Object, arg1 string, arg2 *Variant, arg3 *Object, arg4 string, arg5 float64, arg6 int64, arg7 int64, arg8 float64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(9))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.variant)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.owner)

	C.add_element(cArgsArray, cArg3, C.int(3))

	cArg4 := unsafe.Pointer(stringAsGodotString(arg4))

	C.add_element(cArgsArray, cArg4, C.int(4))

	primArg5 := C.godot_real(arg5)
	cArg5 := unsafe.Pointer(&primArg5)

	C.add_element(cArgsArray, cArg5, C.int(5))

	primArg6 := C.godot_int(arg6)
	cArg6 := unsafe.Pointer(&primArg6)

	C.add_element(cArgsArray, cArg6, C.int(6))

	primArg7 := C.godot_int(arg7)
	cArg7 := unsafe.Pointer(&primArg7)

	C.add_element(cArgsArray, cArg7, C.int(7))

	primArg8 := C.godot_real(arg8)
	cArg8 := unsafe.Pointer(&primArg8)

	C.add_element(cArgsArray, cArg8, C.int(8))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallVoidRidAabb(o Class, methodName string, arg0 *RID, arg1 *AABB) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.aabb)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVector2String(o Class, methodName string, arg0 string) *Vector2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_vector2
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector2{&ret}

}
func godotCallVoidIntArrayArrayInt(o Class, methodName string, arg0 int64, arg1 *Array, arg2 *Array, arg3 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.array)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.array)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidVector2FloatColor(o Class, methodName string, arg0 *Vector2, arg1 float64, arg2 *Color) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_real(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.color)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVector3FloatBool(o Class, methodName string, arg0 float64, arg1 bool) *Vector3 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_real(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_bool(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_vector3
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector3{&ret}

}
func godotCallBoolStringObject(o Class, methodName string, arg0 string, arg1 *Object) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallVoidRidColor(o Class, methodName string, arg0 *RID, arg1 *Color) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.color)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallArrayIntInt(o Class, methodName string, arg0 int64, arg1 int64) *Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Array{&ret}

}
func godotCallVoidPoolVector3ArrayObjectBool(o Class, methodName string, arg0 *PoolVector3Array, arg1 *Object, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.array)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallArrayVector2IntArrayInt(o Class, methodName string, arg0 *Vector2, arg1 int64, arg2 *Array, arg3 int64) *Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.array)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Array{&ret}

}
func godotCallTransform2DInt(o Class, methodName string, arg0 int64) *Transform2D {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_transform2d
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Transform2D{&ret}

}
func godotCallPoolVector2ArrayIntFloat(o Class, methodName string, arg0 int64, arg1 float64) *PoolVector2Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_real(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_pool_vector2_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolVector2Array{&ret}

}
func godotCallVoidIntPoolIntArray(o Class, methodName string, arg0 int64, arg1 *PoolIntArray) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.array)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidRidIntInt(o Class, methodName string, arg0 *RID, arg1 int64, arg2 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidIntIntBool(o Class, methodName string, arg0 int64, arg1 int64, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallPoolIntArray(o Class, methodName string) *PoolIntArray {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_pool_int_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolIntArray{&ret}

}
func godotCallVector2Vector3(o Class, methodName string, arg0 *Vector3) *Vector2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector3)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_vector2
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector2{&ret}

}
func godotCallVoidObjectVector2ColorObject(o Class, methodName string, arg0 *Object, arg1 *Vector2, arg2 *Color, arg3 *Object) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.color)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.owner)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidStringStringPoolStringArray(o Class, methodName string, arg0 string, arg1 string, arg2 *PoolStringArray) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.array)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVariantStringArray(o Class, methodName string, arg0 string, arg1 *Array) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.array)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_variant
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{&ret}

}
func godotCallVoidIntBoolInt(o Class, methodName string, arg0 int64, arg1 bool, arg2 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_bool(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidStringIntIntIntInt(o Class, methodName string, arg0 string, arg1 int64, arg2 int64, arg3 int64, arg4 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	primArg4 := C.godot_int(arg4)
	cArg4 := unsafe.Pointer(&primArg4)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidStringStringFloat(o Class, methodName string, arg0 string, arg1 string, arg2 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_real(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallAabb(o Class, methodName string) *AABB {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_aabb
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &AABB{&ret}

}
func godotCallVoidVector2Bool(o Class, methodName string, arg0 *Vector2, arg1 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_bool(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallBoolStringDictionary(o Class, methodName string, arg0 string, arg1 *Dictionary) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.dictionary)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallVoidObjectBool(o Class, methodName string, arg0 *Object, arg1 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_bool(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidStringObjectBool(o Class, methodName string, arg0 string, arg1 *Object, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVariantStringStringVariant(o Class, methodName string, arg0 string, arg1 string, arg2 *Variant) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.variant)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_variant
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{&ret}

}
func godotCallVoidPoolStringArrayBoolStringInt(o Class, methodName string, arg0 *PoolStringArray, arg1 bool, arg2 string, arg3 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.array)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_bool(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(stringAsGodotString(arg2))

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidPoolVector2ArrayPoolColorArrayPoolVector2ArrayObjectObjectBool(o Class, methodName string, arg0 *PoolVector2Array, arg1 *PoolColorArray, arg2 *PoolVector2Array, arg3 *Object, arg4 *Object, arg5 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(6))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.array)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.array)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.array)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.owner)

	C.add_element(cArgsArray, cArg3, C.int(3))

	cArg4 := unsafe.Pointer(arg4.owner)

	C.add_element(cArgsArray, cArg4, C.int(4))

	primArg5 := C.godot_bool(arg5)
	cArg5 := unsafe.Pointer(&primArg5)

	C.add_element(cArgsArray, cArg5, C.int(5))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidObjectObjectRect2Vector2(o Class, methodName string, arg0 *Object, arg1 *Object, arg2 *Rect2, arg3 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.rect2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.vector2)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallColorIntInt(o Class, methodName string, arg0 int64, arg1 int64) *Color {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_color
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Color{&ret}

}
func godotCallVoidRidPoolIntArrayPoolVector2ArrayPoolColorArrayPoolVector2ArrayRidIntRid(o Class, methodName string, arg0 *RID, arg1 *PoolIntArray, arg2 *PoolVector2Array, arg3 *PoolColorArray, arg4 *PoolVector2Array, arg5 *RID, arg6 int64, arg7 *RID) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(8))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.array)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.array)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.array)

	C.add_element(cArgsArray, cArg3, C.int(3))

	cArg4 := unsafe.Pointer(arg4.array)

	C.add_element(cArgsArray, cArg4, C.int(4))

	cArg5 := unsafe.Pointer(arg5.rid)

	C.add_element(cArgsArray, cArg5, C.int(5))

	primArg6 := C.godot_int(arg6)
	cArg6 := unsafe.Pointer(&primArg6)

	C.add_element(cArgsArray, cArg6, C.int(6))

	cArg7 := unsafe.Pointer(arg7.rid)

	C.add_element(cArgsArray, cArg7, C.int(7))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidBasis(o Class, methodName string, arg0 *Basis) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.basis)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallPoolRealArrayInt(o Class, methodName string, arg0 int64) *PoolRealArray {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_pool_real_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolRealArray{&ret}

}
func godotCallVoidRidIntRid(o Class, methodName string, arg0 *RID, arg1 int64, arg2 *RID) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.rid)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidVector2Vector2ColorFloatBool(o Class, methodName string, arg0 *Vector2, arg1 *Vector2, arg2 *Color, arg3 float64, arg4 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.color)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_real(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	primArg4 := C.godot_bool(arg4)
	cArg4 := unsafe.Pointer(&primArg4)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidTransform2D(o Class, methodName string, arg0 *Transform2D) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.transform2d)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidStringStringObjectObject(o Class, methodName string, arg0 string, arg1 string, arg2 *Object, arg3 *Object) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.owner)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.owner)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallIntVariant(o Class, methodName string, arg0 *Variant) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.variant)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallVoidRidIntBool(o Class, methodName string, arg0 *RID, arg1 int64, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallBoolObjectNodePathVariantVariantFloatIntIntFloat(o Class, methodName string, arg0 *Object, arg1 *NodePath, arg2 *Variant, arg3 *Variant, arg4 float64, arg5 int64, arg6 int64, arg7 float64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(8))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.nodePath)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.variant)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.variant)

	C.add_element(cArgsArray, cArg3, C.int(3))

	primArg4 := C.godot_real(arg4)
	cArg4 := unsafe.Pointer(&primArg4)

	C.add_element(cArgsArray, cArg4, C.int(4))

	primArg5 := C.godot_int(arg5)
	cArg5 := unsafe.Pointer(&primArg5)

	C.add_element(cArgsArray, cArg5, C.int(5))

	primArg6 := C.godot_int(arg6)
	cArg6 := unsafe.Pointer(&primArg6)

	C.add_element(cArgsArray, cArg6, C.int(6))

	primArg7 := C.godot_real(arg7)
	cArg7 := unsafe.Pointer(&primArg7)

	C.add_element(cArgsArray, cArg7, C.int(7))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallVoidStringObjectStringVariant(o Class, methodName string, arg0 string, arg1 *Object, arg2 string, arg3 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(stringAsGodotString(arg2))

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.variant)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVector3IntIntInt(o Class, methodName string, arg0 int64, arg1 int64, arg2 int64) *Vector3 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_vector3
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector3{&ret}

}
func godotCallIntVector2Bool(o Class, methodName string, arg0 *Vector2, arg1 bool) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_bool(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallVoidVector3Vector3(o Class, methodName string, arg0 *Vector3, arg1 *Vector3) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector3)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector3)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallBoolObjectStringObjectStringVariantFloatIntIntFloat(o Class, methodName string, arg0 *Object, arg1 string, arg2 *Object, arg3 string, arg4 *Variant, arg5 float64, arg6 int64, arg7 int64, arg8 float64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(9))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.owner)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(stringAsGodotString(arg3))

	C.add_element(cArgsArray, cArg3, C.int(3))

	cArg4 := unsafe.Pointer(arg4.variant)

	C.add_element(cArgsArray, cArg4, C.int(4))

	primArg5 := C.godot_real(arg5)
	cArg5 := unsafe.Pointer(&primArg5)

	C.add_element(cArgsArray, cArg5, C.int(5))

	primArg6 := C.godot_int(arg6)
	cArg6 := unsafe.Pointer(&primArg6)

	C.add_element(cArgsArray, cArg6, C.int(6))

	primArg7 := C.godot_int(arg7)
	cArg7 := unsafe.Pointer(&primArg7)

	C.add_element(cArgsArray, cArg7, C.int(7))

	primArg8 := C.godot_real(arg8)
	cArg8 := unsafe.Pointer(&primArg8)

	C.add_element(cArgsArray, cArg8, C.int(8))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallVoidRidPoolVector2ArrayPoolColorArrayPoolVector2ArrayRidRidBool(o Class, methodName string, arg0 *RID, arg1 *PoolVector2Array, arg2 *PoolColorArray, arg3 *PoolVector2Array, arg4 *RID, arg5 *RID, arg6 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(7))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.array)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.array)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.array)

	C.add_element(cArgsArray, cArg3, C.int(3))

	cArg4 := unsafe.Pointer(arg4.rid)

	C.add_element(cArgsArray, cArg4, C.int(4))

	cArg5 := unsafe.Pointer(arg5.rid)

	C.add_element(cArgsArray, cArg5, C.int(5))

	primArg6 := C.godot_bool(arg6)
	cArg6 := unsafe.Pointer(&primArg6)

	C.add_element(cArgsArray, cArg6, C.int(6))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidTransform(o Class, methodName string, arg0 *Transform) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.transform)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidObjectIntVector2(o Class, methodName string, arg0 *Object, arg1 int64, arg2 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.vector2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidIntFloatFloatFloat(o Class, methodName string, arg0 int64, arg1 float64, arg2 float64, arg3 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_real(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_real(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_real(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidRidRect2Color(o Class, methodName string, arg0 *RID, arg1 *Rect2, arg2 *Color) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.rect2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.color)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidStringIntInt(o Class, methodName string, arg0 string, arg1 int64, arg2 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallStringVariant(o Class, methodName string, arg0 *Variant) string {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.variant)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_string
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return godotStringAsString(&ret)

}
func godotCallVoidFloatColor(o Class, methodName string, arg0 float64, arg1 *Color) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_real(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.color)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallIntStringInt(o Class, methodName string, arg0 string, arg1 int64) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallBoolObjectFloatStringVariantVariantVariantVariantVariant(o Class, methodName string, arg0 *Object, arg1 float64, arg2 string, arg3 *Variant, arg4 *Variant, arg5 *Variant, arg6 *Variant, arg7 *Variant) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(8))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_real(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(stringAsGodotString(arg2))

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.variant)

	C.add_element(cArgsArray, cArg3, C.int(3))

	cArg4 := unsafe.Pointer(arg4.variant)

	C.add_element(cArgsArray, cArg4, C.int(4))

	cArg5 := unsafe.Pointer(arg5.variant)

	C.add_element(cArgsArray, cArg5, C.int(5))

	cArg6 := unsafe.Pointer(arg6.variant)

	C.add_element(cArgsArray, cArg6, C.int(6))

	cArg7 := unsafe.Pointer(arg7.variant)

	C.add_element(cArgsArray, cArg7, C.int(7))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallVoidObjectRect2BoolColorBoolObject(o Class, methodName string, arg0 *Object, arg1 *Rect2, arg2 bool, arg3 *Color, arg4 bool, arg5 *Object) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(6))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.rect2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.color)

	C.add_element(cArgsArray, cArg3, C.int(3))

	primArg4 := C.godot_bool(arg4)
	cArg4 := unsafe.Pointer(&primArg4)

	C.add_element(cArgsArray, cArg4, C.int(4))

	cArg5 := unsafe.Pointer(arg5.owner)

	C.add_element(cArgsArray, cArg5, C.int(5))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidRidObjectStringVariant(o Class, methodName string, arg0 *RID, arg1 *Object, arg2 string, arg3 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(stringAsGodotString(arg2))

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.variant)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallBoolVector2FloatObject(o Class, methodName string, arg0 *Vector2, arg1 float64, arg2 *Object) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_real(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.owner)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallColor(o Class, methodName string) *Color {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_color
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Color{&ret}

}
func godotCallVector3Float(o Class, methodName string, arg0 float64) *Vector3 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_real(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_vector3
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector3{&ret}

}
func godotCallVoidStringVariant(o Class, methodName string, arg0 string, arg1 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.variant)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallIntIntStringPoolStringArrayString(o Class, methodName string, arg0 int64, arg1 string, arg2 *PoolStringArray, arg3 string) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.array)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(stringAsGodotString(arg3))

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallVoidIntRect2(o Class, methodName string, arg0 int64, arg1 *Rect2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.rect2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidRidIntTransform2D(o Class, methodName string, arg0 *RID, arg1 int64, arg2 *Transform2D) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.transform2d)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidStringObjectInt(o Class, methodName string, arg0 string, arg1 *Object, arg2 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallBoolVector2(o Class, methodName string, arg0 *Vector2) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallVoidVector2Variant(o Class, methodName string, arg0 *Vector2, arg1 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.variant)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallPoolVector2ArrayVector2Vector2(o Class, methodName string, arg0 *Vector2, arg1 *Vector2) *PoolVector2Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_pool_vector2_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &PoolVector2Array{&ret}

}
func godotCallArray(o Class, methodName string) *Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Array{&ret}

}
func godotCallBoolString(o Class, methodName string, arg0 string) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallFloatObjectVector2StringStringColor(o Class, methodName string, arg0 *Object, arg1 *Vector2, arg2 string, arg3 string, arg4 *Color) float64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(stringAsGodotString(arg2))

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(stringAsGodotString(arg3))

	C.add_element(cArgsArray, cArg3, C.int(3))

	cArg4 := unsafe.Pointer(arg4.color)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_real
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return float64(ret)

}
func godotCallBoolStringString(o Class, methodName string, arg0 string, arg1 string) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallVoidIntBoolIntColorBoolIntColorObjectObject(o Class, methodName string, arg0 int64, arg1 bool, arg2 int64, arg3 *Color, arg4 bool, arg5 int64, arg6 *Color, arg7 *Object, arg8 *Object) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(9))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_bool(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.color)

	C.add_element(cArgsArray, cArg3, C.int(3))

	primArg4 := C.godot_bool(arg4)
	cArg4 := unsafe.Pointer(&primArg4)

	C.add_element(cArgsArray, cArg4, C.int(4))

	primArg5 := C.godot_int(arg5)
	cArg5 := unsafe.Pointer(&primArg5)

	C.add_element(cArgsArray, cArg5, C.int(5))

	cArg6 := unsafe.Pointer(arg6.color)

	C.add_element(cArgsArray, cArg6, C.int(6))

	cArg7 := unsafe.Pointer(arg7.owner)

	C.add_element(cArgsArray, cArg7, C.int(7))

	cArg8 := unsafe.Pointer(arg8.owner)

	C.add_element(cArgsArray, cArg8, C.int(8))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallIntIntStringInt(o Class, methodName string, arg0 int64, arg1 string, arg2 int64) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallVoidRidVector3(o Class, methodName string, arg0 *RID, arg1 *Vector3) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector3)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallStringString(o Class, methodName string, arg0 string) string {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_string
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return godotStringAsString(&ret)

}
func godotCallIntStringIntStringInt(o Class, methodName string, arg0 string, arg1 int64, arg2 string, arg3 int64) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(stringAsGodotString(arg2))

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallIntIntIntIntInt(o Class, methodName string, arg0 int64, arg1 int64, arg2 int64, arg3 int64) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallVoidObjectIntBool(o Class, methodName string, arg0 *Object, arg1 int64, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallArrayInt(o Class, methodName string, arg0 int64) *Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Array{&ret}

}
func godotCallVoidIntIntRect2Vector2Float(o Class, methodName string, arg0 int64, arg1 int64, arg2 *Rect2, arg3 *Vector2, arg4 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.rect2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.vector2)

	C.add_element(cArgsArray, cArg3, C.int(3))

	primArg4 := C.godot_real(arg4)
	cArg4 := unsafe.Pointer(&primArg4)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidPoolVector3Array(o Class, methodName string, arg0 *PoolVector3Array) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.array)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidIntStringStringVariant(o Class, methodName string, arg0 int64, arg1 string, arg2 string, arg3 *Variant) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(stringAsGodotString(arg2))

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.variant)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidObjectObjectVector3Vector3Int(o Class, methodName string, arg0 *Object, arg1 *Object, arg2 *Vector3, arg3 *Vector3, arg4 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.vector3)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.vector3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	primArg4 := C.godot_int(arg4)
	cArg4 := unsafe.Pointer(&primArg4)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidPoolColorArray(o Class, methodName string, arg0 *PoolColorArray) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.array)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidIntPlane(o Class, methodName string, arg0 int64, arg1 *Plane) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.plane)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidIntObjectBool(o Class, methodName string, arg0 int64, arg1 *Object, arg2 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidIntIntIntBoolBoolBoolVector2(o Class, methodName string, arg0 int64, arg1 int64, arg2 int64, arg3 bool, arg4 bool, arg5 bool, arg6 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(7))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_bool(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	primArg4 := C.godot_bool(arg4)
	cArg4 := unsafe.Pointer(&primArg4)

	C.add_element(cArgsArray, cArg4, C.int(4))

	primArg5 := C.godot_bool(arg5)
	cArg5 := unsafe.Pointer(&primArg5)

	C.add_element(cArgsArray, cArg5, C.int(5))

	cArg6 := unsafe.Pointer(arg6.vector2)

	C.add_element(cArgsArray, cArg6, C.int(6))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallBoolIntInt(o Class, methodName string, arg0 int64, arg1 int64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallFloatFloat(o Class, methodName string, arg0 float64) float64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_real(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_real
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return float64(ret)

}
func godotCallVoidIntIntIntIntInt(o Class, methodName string, arg0 int64, arg1 int64, arg2 int64, arg3 int64, arg4 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	primArg4 := C.godot_int(arg4)
	cArg4 := unsafe.Pointer(&primArg4)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallObjectNodePath(o Class, methodName string, arg0 *NodePath) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.nodePath)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_object
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{&ret}

}
func godotCallVariantArray(o Class, methodName string, arg0 *Array) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.array)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_variant
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{&ret}

}
func godotCallArrayStringIntInt(o Class, methodName string, arg0 string, arg1 int64, arg2 int64) *Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Array{&ret}

}
func godotCallVoidObjectVector2StringColorInt(o Class, methodName string, arg0 *Object, arg1 *Vector2, arg2 string, arg3 *Color, arg4 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(stringAsGodotString(arg2))

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.color)

	C.add_element(cArgsArray, cArg3, C.int(3))

	primArg4 := C.godot_int(arg4)
	cArg4 := unsafe.Pointer(&primArg4)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidObjectAabb(o Class, methodName string, arg0 *Object, arg1 *AABB) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.aabb)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallFloatRidVector2IntIntColor(o Class, methodName string, arg0 *RID, arg1 *Vector2, arg2 int64, arg3 int64, arg4 *Color) float64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	cArg4 := unsafe.Pointer(arg4.color)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_real
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return float64(ret)

}
func godotCallVoidIntIntPoolStringArrayPoolByteArray(o Class, methodName string, arg0 int64, arg1 int64, arg2 *PoolStringArray, arg3 *PoolByteArray) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.array)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.array)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidIntBoolStringString(o Class, methodName string, arg0 int64, arg1 bool, arg2 string, arg3 string) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_bool(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(stringAsGodotString(arg2))

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(stringAsGodotString(arg3))

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallBoolRidInt(o Class, methodName string, arg0 *RID, arg1 int64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallVoidFloatFloatFloatFloat(o Class, methodName string, arg0 float64, arg1 float64, arg2 float64, arg3 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_real(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_real(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_real(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_real(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidObjectObjectInt(o Class, methodName string, arg0 *Object, arg1 *Object, arg2 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallBoolTransform2DVector2ObjectTransform2DVector2(o Class, methodName string, arg0 *Transform2D, arg1 *Vector2, arg2 *Object, arg3 *Transform2D, arg4 *Vector2) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.transform2d)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.owner)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.transform2d)

	C.add_element(cArgsArray, cArg3, C.int(3))

	cArg4 := unsafe.Pointer(arg4.vector2)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallVoidIntString(o Class, methodName string, arg0 int64, arg1 string) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallBoolStringStringInt(o Class, methodName string, arg0 string, arg1 string, arg2 int64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallVoidBoolBoolIntInt(o Class, methodName string, arg0 bool, arg1 bool, arg2 int64, arg3 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_bool(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_bool(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallTransform2DRid(o Class, methodName string, arg0 *RID) *Transform2D {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_transform2d
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Transform2D{&ret}

}
func godotCallVoid(o Class, methodName string) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallRid(o Class, methodName string) *RID {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_rid
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &RID{&ret}

}
func godotCallVoidIntTransform2D(o Class, methodName string, arg0 int64, arg1 *Transform2D) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.transform2d)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallRect2ObjectInt(o Class, methodName string, arg0 *Object, arg1 int64) *Rect2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_rect2
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Rect2{&ret}

}
func godotCallVoidRidVector2FloatColor(o Class, methodName string, arg0 *RID, arg1 *Vector2, arg2 float64, arg3 *Color) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_real(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.color)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidRidRidVector2(o Class, methodName string, arg0 *RID, arg1 *RID, arg2 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.rid)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.vector2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVariantRidString(o Class, methodName string, arg0 *RID, arg1 string) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_variant
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{&ret}

}
func godotCallObjectIntInt(o Class, methodName string, arg0 int64, arg1 int64) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_object
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{&ret}

}
func godotCallVoidIntIntBoolIntPoolByteArray(o Class, methodName string, arg0 int64, arg1 int64, arg2 bool, arg3 int64, arg4 *PoolByteArray) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_bool(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_int(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	cArg4 := unsafe.Pointer(arg4.array)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallIntObjectTransformObject(o Class, methodName string, arg0 *Object, arg1 *Transform, arg2 *Object) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.transform)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.owner)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallTransform2DObject(o Class, methodName string, arg0 *Object) *Transform2D {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_transform2d
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Transform2D{&ret}

}
func godotCallVoidIntFloatFloatBool(o Class, methodName string, arg0 int64, arg1 float64, arg2 float64, arg3 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_real(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_real(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_bool(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallFloatRidInt(o Class, methodName string, arg0 *RID, arg1 int64) float64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_real
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return float64(ret)

}
func godotCallArrayObjectVector3(o Class, methodName string, arg0 *Object, arg1 *Vector3) *Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.owner)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.vector3)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Array{&ret}

}
func godotCallVector2IntIntObjectVector2(o Class, methodName string, arg0 int64, arg1 int64, arg2 *Object, arg3 *Vector2) *Vector2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.owner)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.vector2)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_vector2
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector2{&ret}

}
func godotCallObjectStringBoolString(o Class, methodName string, arg0 string, arg1 bool, arg2 string) *Object {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(stringAsGodotString(arg0))

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_bool(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(stringAsGodotString(arg2))

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_object
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Object{&ret}

}
func godotCallVector2Vector2(o Class, methodName string, arg0 *Vector2) *Vector2 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.vector2)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_vector2
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Vector2{&ret}

}
func godotCallVoidIntObjectTransform2DBoolVector2(o Class, methodName string, arg0 int64, arg1 *Object, arg2 *Transform2D, arg3 bool, arg4 *Vector2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.transform2d)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_bool(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	cArg4 := unsafe.Pointer(arg4.vector2)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallBasis(o Class, methodName string) *Basis {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(0))
	log.Println("    C Array: ", cArgsArray)

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_basis
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Basis{&ret}

}
func godotCallBoolInt(o Class, methodName string, arg0 int64) bool {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_bool
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return bool(ret)

}
func godotCallVoidNodePathObjectInt(o Class, methodName string, arg0 *NodePath, arg1 *Object, arg2 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(3))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.nodePath)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.owner)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidRidRect2(o Class, methodName string, arg0 *RID, arg1 *Rect2) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.rect2)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallIntIntString(o Class, methodName string, arg0 int64, arg1 string) int64 {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	primArg0 := C.godot_int(arg0)
	cArg0 := unsafe.Pointer(&primArg0)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(stringAsGodotString(arg1))

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_int
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return int64(ret)

}
func godotCallVoidRidIntArrayArrayInt(o Class, methodName string, arg0 *RID, arg1 int64, arg2 *Array, arg3 *Array, arg4 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(5))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.array)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.array)

	C.add_element(cArgsArray, cArg3, C.int(3))

	primArg4 := C.godot_int(arg4)
	cArg4 := unsafe.Pointer(&primArg4)

	C.add_element(cArgsArray, cArg4, C.int(4))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidPoolVector2ArrayColorPoolVector2ArrayObjectObjectBool(o Class, methodName string, arg0 *PoolVector2Array, arg1 *Color, arg2 *PoolVector2Array, arg3 *Object, arg4 *Object, arg5 bool) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(6))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.array)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.color)

	C.add_element(cArgsArray, cArg1, C.int(1))

	cArg2 := unsafe.Pointer(arg2.array)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.owner)

	C.add_element(cArgsArray, cArg3, C.int(3))

	cArg4 := unsafe.Pointer(arg4.owner)

	C.add_element(cArgsArray, cArg4, C.int(4))

	primArg5 := C.godot_bool(arg5)
	cArg5 := unsafe.Pointer(&primArg5)

	C.add_element(cArgsArray, cArg5, C.int(5))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidRidInt(o Class, methodName string, arg0 *RID, arg1 int64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(2))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallVoidRidIntIntFloat(o Class, methodName string, arg0 *RID, arg1 int64, arg2 int64, arg3 float64) {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.rid)

	C.add_element(cArgsArray, cArg0, C.int(0))

	primArg1 := C.godot_int(arg1)
	cArg1 := unsafe.Pointer(&primArg1)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	primArg3 := C.godot_real(arg3)
	cArg3 := unsafe.Pointer(&primArg3)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	log.Println("  No return value.")

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		nil,        // void*
	)
	log.Println("  Finished calling method")

}
func godotCallArrayPoolByteArray(o Class, methodName string, arg0 *PoolByteArray) *Array {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(1))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.array)

	C.add_element(cArgsArray, cArg0, C.int(0))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_array
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Array{&ret}

}
func godotCallVariantArrayArrayIntArray(o Class, methodName string, arg0 *Array, arg1 *Array, arg2 int64, arg3 *Array) *Variant {

	methodBind := getGodotMethod(o.baseClass(), methodName)

	// Construct a C array that will contain pointers to our arguments.
	log.Println("  Allocating argument array in C.")
	cArgsArray := C.build_array(C.int(4))
	log.Println("    C Array: ", cArgsArray)

	cArg0 := unsafe.Pointer(arg0.array)

	C.add_element(cArgsArray, cArg0, C.int(0))

	cArg1 := unsafe.Pointer(arg1.array)

	C.add_element(cArgsArray, cArg1, C.int(1))

	primArg2 := C.godot_int(arg2)
	cArg2 := unsafe.Pointer(&primArg2)

	C.add_element(cArgsArray, cArg2, C.int(2))

	cArg3 := unsafe.Pointer(arg3.array)

	C.add_element(cArgsArray, cArg3, C.int(3))

	log.Println("  Built argument array from variant arguments: ", cArgsArray)

	// Construct our return object that will be populated by the method call.
	log.Println("  Building return value.")
	var ret C.godot_variant
	retPtr := unsafe.Pointer(&ret)

	// Get the Godot objects owner so we can pass it to godot_method_bind_ptrcall.
	log.Println("  Using godot object owner:", o.getOwner())
	objectOwner := unsafe.Pointer(o.getOwner())

	// Call the parent method. "ret" will be populated with the return value.
	log.Println("  Calling bind_ptrcall...")
	C.godot_method_bind_ptrcall(
		methodBind,
		objectOwner,
		cArgsArray, // void**
		retPtr,     // void*
	)
	log.Println("  Finished calling method")

	// Convert the return value based on the type.
	return &Variant{&ret}

}