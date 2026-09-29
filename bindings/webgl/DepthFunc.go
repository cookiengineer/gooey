//go:build wasm

package webgl

type DepthFunc uint

const (
	DepthFuncAlways   DepthFunc = 0x0207
	DepthFuncEqual    DepthFunc = 0x0202
	DepthFuncGreater  DepthFunc = 0x0204
	DepthFuncGEqual   DepthFunc = 0x0206
	DepthFuncLess     DepthFunc = 0x0201
	DepthFuncLEqual   DepthFunc = 0x0203
	DepthFuncNever    DepthFunc = 0x0200
	DepthFuncNotEqual DepthFunc = 0x0205
)
