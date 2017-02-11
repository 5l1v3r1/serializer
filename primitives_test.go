package serializer

import (
	"reflect"
	"testing"
)

func TestPrimitives(t *testing.T) {
	objects := []interface{}{
		Bytes([]byte("hello, world")),
		Int(1337),
		Int(-(1 << 30)),
		Float64(3.1415),
		Float64Slice([]float64{1, 0.5, -5.15e20}),
		Float32Slice([]float32{-1337, 0.5, 5e10, -5e10}),
		Float64Slice([]float64{}),
		Float32Slice([]float32{}),
	}
	data, err := SerializeAny(objects...)
	if err != nil {
		t.Fatal(err)
	}
	var obj1 Bytes
	var obj2, obj3 Int
	var obj4 Float64
	var obj5, obj7 Float64Slice
	var obj6, obj8 Float32Slice
	err = DeserializeAny(data, &obj1, &obj2, &obj3, &obj4, &obj5, &obj6, &obj7, &obj8)
	if err != nil {
		t.Fatal(err)
	}
	newObjs := []interface{}{obj1, obj2, obj3, obj4, obj5, obj6, obj7, obj8}
	for i, x := range objects {
		if !reflect.DeepEqual(x, newObjs[i]) {
			t.Errorf("object %d: expected %v but got %v", i, x, newObjs[i])
		}
	}
}