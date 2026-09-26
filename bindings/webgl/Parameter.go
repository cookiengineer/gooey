//go:build wasm

package webgl

type Parameter uint

const (
	ParameterAliasedLineWidthRange        Parameter = 0x846E
	ParameterAliasedPointSizeRange        Parameter = 0x846D
	ParameterMaxCombinedTextureImageUnits Parameter = 0x8B4D
	ParameterMaxCubeMapTextureSize        Parameter = 0x851C
	ParameterMaxFragmentUniformVectors    Parameter = 0x8DFD
	ParameterMaxRenderbufferSize          Parameter = 0x84E8
	ParameterMaxTextureImageUnits         Parameter = 0x8872
	ParameterMaxTextureSize               Parameter = 0x0D33
	ParameterMaxVaryingVectors            Parameter = 0x8DFC
	ParameterMaxVertexAttribs             Parameter = 0x8869
	ParameterMaxVertexUniformVectors      Parameter = 0x8DFB
	ParameterMaxViewportDims              Parameter = 0x0D3A
	ParameterPackAlignment                Parameter = 0x0D05
	ParameterRenderer                     Parameter = 0x1F01
	ParameterSamples                      Parameter = 0x80A9
	ParameterShadingLanguageVersion       Parameter = 0x8B8C
	ParameterUnpackAlignment              Parameter = 0x0CF5
	ParameterVendor                       Parameter = 0x1F00
	ParameterVersion                      Parameter = 0x1F02
	ParameterViewport                     Parameter = 0x0BA2
)
