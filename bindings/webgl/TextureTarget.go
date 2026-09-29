//go:build wasm

package webgl

type TextureTarget uint

const (
	TextureTarget2D               TextureTarget = 0x0DE1
	TextureTarget2DArray          TextureTarget = 0x8C1A
	TextureTarget3D               TextureTarget = 0x806F
	TextureTargetCubeMap          TextureTarget = 0x8513
	TextureTargetCubeMapNegativeX TextureTarget = 0x8516
	TextureTargetCubeMapNegativeY TextureTarget = 0x8518
	TextureTargetCubeMapNegativeZ TextureTarget = 0x851A
	TextureTargetCubeMapPositiveX TextureTarget = 0x8515
	TextureTargetCubeMapPositiveY TextureTarget = 0x8517
	TextureTargetCubeMapPositiveZ TextureTarget = 0x8519
)
