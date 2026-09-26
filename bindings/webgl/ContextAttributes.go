//go:build wasm

package webgl

import "syscall/js"

type ContextAttributes struct {
	Alpha                        bool            `json:"alpha"`
	Antialias                    bool            `json:"antialias"`
	Depth                        bool            `json:"depth"`
	Desynchronized               bool            `json:"desynchronized"`
	FailIfMajorPerformanceCaveat bool            `json:"failIfMajorPerformanceCaveat"`
	PowerPreference              PowerPreference `json:"powerPreference"`
	PremultipliedAlpha           bool            `json:"premultipliedAlpha"`
	PreserveDrawingBuffer        bool            `json:"preserveDrawingBuffer"`
	Stencil                      bool            `json:"stencil"`
}

// Returns new ContextAttributes instance, filled with the defaults the
// specification gives for a WebGL2RenderingContext.
func NewContextAttributes() ContextAttributes {

	var attributes ContextAttributes

	attributes.Alpha = true
	attributes.Antialias = true
	attributes.Depth = true
	attributes.Desynchronized = false
	attributes.FailIfMajorPerformanceCaveat = false
	attributes.PowerPreference = PowerPreferenceDefault
	attributes.PremultipliedAlpha = true
	attributes.PreserveDrawingBuffer = false
	attributes.Stencil = false

	return attributes

}

func ToContextAttributes(value js.Value) *ContextAttributes {

	if value.IsNull() || value.IsUndefined() {
		return nil
	}

	attributes := NewContextAttributes()

	attributes.Alpha = bool_or(value, "alpha", attributes.Alpha)
	attributes.Antialias = bool_or(value, "antialias", attributes.Antialias)
	attributes.Depth = bool_or(value, "depth", attributes.Depth)
	attributes.Desynchronized = bool_or(value, "desynchronized", attributes.Desynchronized)
	attributes.FailIfMajorPerformanceCaveat = bool_or(value, "failIfMajorPerformanceCaveat", attributes.FailIfMajorPerformanceCaveat)
	attributes.PremultipliedAlpha = bool_or(value, "premultipliedAlpha", attributes.PremultipliedAlpha)
	attributes.PreserveDrawingBuffer = bool_or(value, "preserveDrawingBuffer", attributes.PreserveDrawingBuffer)
	attributes.Stencil = bool_or(value, "stencil", attributes.Stencil)

	power_preference := value.Get("powerPreference")

	if !power_preference.IsNull() && !power_preference.IsUndefined() {
		attributes.PowerPreference = PowerPreference(power_preference.String())
	}

	return &attributes

}

// Returns the attributes as the options object that getContext() expects.
func (attributes *ContextAttributes) ToValue() js.Value {

	value := js.Global().Get("Object").New()

	value.Set("alpha", attributes.Alpha)
	value.Set("antialias", attributes.Antialias)
	value.Set("depth", attributes.Depth)
	value.Set("desynchronized", attributes.Desynchronized)
	value.Set("failIfMajorPerformanceCaveat", attributes.FailIfMajorPerformanceCaveat)
	value.Set("powerPreference", string(attributes.PowerPreference))
	value.Set("premultipliedAlpha", attributes.PremultipliedAlpha)
	value.Set("preserveDrawingBuffer", attributes.PreserveDrawingBuffer)
	value.Set("stencil", attributes.Stencil)

	return value

}
