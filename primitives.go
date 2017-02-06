package serializer

import (
	"bytes"
	"encoding/binary"
	"strconv"

	"github.com/unixpickle/essentials"
)

func init() {
	RegisterTypedDeserializer(Bytes(nil).SerializerType(), DeserializeBytes)
	RegisterTypedDeserializer(Int(0).SerializerType(), DeserializeInt)
	RegisterTypedDeserializer(Float64(0).SerializerType(), DeserializeFloat64)
	RegisterTypedDeserializer(Float64Slice(nil).SerializerType(), DeserializeFloat64Slice)
	RegisterTypedDeserializer(Float32Slice(nil).SerializerType(), DeserializeFloat32Slice)
}

// Bytes is a Serializer wrapper for []byte.
type Bytes []byte

// DeserializeBytes deserializes a Bytes.
func DeserializeBytes(d []byte) (Bytes, error) {
	return d, nil
}

// Serialize serializes the object.
func (b Bytes) Serialize() ([]byte, error) {
	return b, nil
}

// SerializerType returns the unique ID used to serialize
// Bytes.
func (b Bytes) SerializerType() string {
	return "[]byte"
}

// Int is a Serializer wrapper for an int.
type Int int

// DeserializeInt deserialize an Int.
func DeserializeInt(d []byte) (Int, error) {
	num, err := strconv.Atoi(string(d))
	if err != nil {
		return 0, essentials.AddCtx("deserialize int", err)
	}
	return Int(num), nil
}

// Serialize serializes the object.
func (i Int) Serialize() ([]byte, error) {
	return []byte(strconv.Itoa(int(i))), nil
}

// SerializerType returns the unique ID used to serialize
// an Int.
func (i Int) SerializerType() string {
	return "int"
}

// Float64 is a Serializer for a float64.
type Float64 float64

// DeserializeFloat64 deserializes a Float64.
func DeserializeFloat64(d []byte) (Float64, error) {
	buf := bytes.NewBuffer(d)
	var value float64
	if err := binary.Read(buf, binary.LittleEndian, &value); err != nil {
		return 0, essentials.AddCtx("deserialize float64", err)
	}
	return Float64(value), nil
}

// Serialize serializes the object.
func (f Float64) Serialize() ([]byte, error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, float64(f))
	return buf.Bytes(), nil
}

// SerializerType returns the unique ID used to serialize
// a Float64.
func (f Float64) SerializerType() string {
	return "float64"
}

// A Float64Slice is a Serializer for a []float64.
type Float64Slice []float64

// DeserializeFloat64Slice deserializes a Float64Slice.
func DeserializeFloat64Slice(d []byte) (Float64Slice, error) {
	reader := bytes.NewBuffer(d)
	var size uint64
	if err := binary.Read(reader, binary.LittleEndian, &size); err != nil {
		return nil, essentials.AddCtx("deserialize []float64", err)
	}
	vec := make(Float64Slice, int(size))
	for i := range vec {
		if err := binary.Read(reader, binary.LittleEndian, &vec[i]); err != nil {
			return nil, essentials.AddCtx("deserialize []float64", err)
		}
	}
	return vec, nil
}

// Serialize serializes the object.
func (f Float64Slice) Serialize() ([]byte, error) {
	var w bytes.Buffer
	binary.Write(&w, binary.LittleEndian, uint64(len(f)))
	for _, x := range f {
		binary.Write(&w, binary.LittleEndian, x)
	}
	return w.Bytes(), nil
}

// SerializerType returns the unique ID used to serialize
// a Float64Slice.
func (f Float64Slice) SerializerType() string {
	return "[]float64"
}

// A Float32Slice is a Serializer for a []float32.
type Float32Slice []float32

// DeserializeFloat32Slice deserializes a Float32Slice.
func DeserializeFloat32Slice(d []byte) (Float32Slice, error) {
	reader := bytes.NewBuffer(d)
	var size uint64
	if err := binary.Read(reader, binary.LittleEndian, &size); err != nil {
		return nil, essentials.AddCtx("deserialize []float32", err)
	}
	vec := make(Float32Slice, int(size))
	for i := range vec {
		if err := binary.Read(reader, binary.LittleEndian, &vec[i]); err != nil {
			return nil, essentials.AddCtx("deserialize []float32", err)
		}
	}
	return vec, nil
}

// Serialize serializes the object.
func (f Float32Slice) Serialize() ([]byte, error) {
	var w bytes.Buffer
	binary.Write(&w, binary.LittleEndian, uint64(len(f)))
	for _, x := range f {
		binary.Write(&w, binary.LittleEndian, x)
	}
	return w.Bytes(), nil
}

// SerializerType returns the unique ID used to serialize
// a Float32Slice.
func (f Float32Slice) SerializerType() string {
	return "[]float32"
}
