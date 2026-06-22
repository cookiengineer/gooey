//go:build wasm

package canvas2d

type Path2DOperation string

const (
	Path2DOperationMoveTo                 Path2DOperation = "M"
	Path2DOperationLineTo                 Path2DOperation = "L"
	Path2DOperationHorizontalLineTo       Path2DOperation = "H"
	Path2DOperationVerticalLineTo         Path2DOperation = "V"
	Path2DOperationBezierCurveTo          Path2DOperation = "C"
	Path2DOperationSmoothCurveTo          Path2DOperation = "S"
	Path2DOperationQuadraticCurveTo       Path2DOperation = "Q"
	Path2DOperationSmoothQuadraticCurveTo Path2DOperation = "T"
	Path2DOperationArc                    Path2DOperation = "A"
	Path2DOperationClosePath              Path2DOperation = "Z"
	Path2DOperationArcTo                  Path2DOperation = "Ar"
	Path2DOperationEllipse                Path2DOperation = "E"
	Path2DOperationRect                   Path2DOperation = "R"
	Path2DOperationRoundRect              Path2DOperation = "r"
	Path2DOperationAddPath                Path2DOperation = "P"
)
