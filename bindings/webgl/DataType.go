//go:build wasm

package webgl

type DataType uint

const (
	DataTypeBool            DataType = 0x8B56
	DataTypeBoolVec2        DataType = 0x8B57
	DataTypeBoolVec3        DataType = 0x8B58
	DataTypeBoolVec4        DataType = 0x8B59
	DataTypeByte            DataType = 0x1400
	DataTypeFloat           DataType = 0x1406
	DataTypeFloatMat2       DataType = 0x8B5A
	DataTypeFloatMat3       DataType = 0x8B5B
	DataTypeFloatMat4       DataType = 0x8B5C
	DataTypeFloatVec2       DataType = 0x8B50
	DataTypeFloatVec3       DataType = 0x8B51
	DataTypeFloatVec4       DataType = 0x8B52
	DataTypeHalfFloat       DataType = 0x140B
	DataTypeInt             DataType = 0x1404
	DataTypeIntVec2         DataType = 0x8B53
	DataTypeIntVec3         DataType = 0x8B54
	DataTypeIntVec4         DataType = 0x8B55
	DataTypeSampler2D       DataType = 0x8B5E
	DataTypeSampler2DArray  DataType = 0x8DC1
	DataTypeSamplerCube     DataType = 0x8B60
	DataTypeShort           DataType = 0x1402
	DataTypeUnsignedByte    DataType = 0x1401
	DataTypeUnsignedInt     DataType = 0x1405
	DataTypeUnsignedIntVec2 DataType = 0x8DC6
	DataTypeUnsignedIntVec3 DataType = 0x8DC7
	DataTypeUnsignedIntVec4 DataType = 0x8DC8
	DataTypeUnsignedShort   DataType = 0x1403
)
