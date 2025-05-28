package canvas2d

type CompositeOperation string

const (
	CompositeOperationSourceOver CompositeOperation = "source-over"
	CompositeOperationSourceIn   CompositeOperation = "source-in"
	CompositeOperationSourceOut  CompositeOperation = "source-out"
	CompositeOperationSourceAtop CompositeOperation = "source-atop"

	CompositeOperationDestinationOver CompositeOperation = "destination-over"
	CompositeOperationDestinationIn   CompositeOperation = "destination-in"
	CompositeOperationDestinationOut  CompositeOperation = "destination-out"
	CompositeOperationDestinationAtop CompositeOperation = "destination-atop"

	CompositeOperationLighter  CompositeOperation = "lighter"
	CompositeOperationCopy     CompositeOperation = "copy"
	CompositeOperationXOR      CompositeOperation = "xor"
	CompositeOperationMultiply CompositeOperation = "multiply"
	CompositeOperationScreen   CompositeOperation = "screen"
	CompositeOperationOverlay  CompositeOperation = "overlay"
	CompositeOperationDarken   CompositeOperation = "darken"
	CompositeOperationLighten  CompositeOperation = "lighten"

	CompositeOperationColorDodge CompositeOperation = "color-dodge"
	CompositeOperationColorBurn  CompositeOperation = "color-burn"

	CompositeOperationHardLight CompositeOperation = "hard-light"
	CompositeOperationSoftLight CompositeOperation = "soft-light"

	CompositeOperationDifference CompositeOperation = "difference"
	CompositeOperationExclusion  CompositeOperation = "exclusion"
	CompositeOperationHue        CompositeOperation = "hue"
	CompositeOperationSaturation CompositeOperation = "saturation"
	CompositeOperationColor      CompositeOperation = "color"
	CompositeOperationLuminosity CompositeOperation = "luminosity"
)
