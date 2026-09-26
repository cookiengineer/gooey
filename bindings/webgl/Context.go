//go:build wasm

package webgl

import "errors"
import "syscall/js"

type Context struct {
	DrawingBufferWidth     int                `json:"drawingBufferWidth"`
	DrawingBufferHeight    int                `json:"drawingBufferHeight"`
	Renderer               string             `json:"renderer"`
	ShadingLanguageVersion string             `json:"shadingLanguageVersion"`
	Vendor                 string             `json:"vendor"`
	Version                string             `json:"version"`
	Attributes             *ContextAttributes `json:"attributes"`
	Value                  *js.Value          `json:"value"`

	fnActiveTexture            js.Value
	fnAttachShader             js.Value
	fnBindAttribLocation       js.Value
	fnBindBuffer               js.Value
	fnBindFramebuffer          js.Value
	fnBindRenderbuffer         js.Value
	fnBindTexture              js.Value
	fnBindVertexArray          js.Value
	fnBlendEquation            js.Value
	fnBlendFunc                js.Value
	fnBlendFuncSeparate        js.Value
	fnBufferData               js.Value
	fnBufferSubData            js.Value
	fnCheckFramebufferStatus   js.Value
	fnClear                    js.Value
	fnClearColor               js.Value
	fnClearDepth               js.Value
	fnClearStencil             js.Value
	fnColorMask                js.Value
	fnCompileShader            js.Value
	fnCreateBuffer             js.Value
	fnCreateFramebuffer        js.Value
	fnCreateProgram            js.Value
	fnCreateRenderbuffer       js.Value
	fnCreateShader             js.Value
	fnCreateTexture            js.Value
	fnCreateVertexArray        js.Value
	fnCullFace                 js.Value
	fnDeleteBuffer             js.Value
	fnDeleteFramebuffer        js.Value
	fnDeleteProgram            js.Value
	fnDeleteRenderbuffer       js.Value
	fnDeleteShader             js.Value
	fnDeleteTexture            js.Value
	fnDeleteVertexArray        js.Value
	fnDepthFunc                js.Value
	fnDepthMask                js.Value
	fnDetachShader             js.Value
	fnDisable                  js.Value
	fnDisableVertexAttribArray js.Value
	fnDrawArrays               js.Value
	fnDrawArraysInstanced      js.Value
	fnDrawElements             js.Value
	fnDrawElementsInstanced    js.Value
	fnEnable                   js.Value
	fnEnableVertexAttribArray  js.Value
	fnFinish                   js.Value
	fnFlush                    js.Value
	fnFramebufferRenderbuffer  js.Value
	fnFramebufferTexture2D     js.Value
	fnFrontFace                js.Value
	fnGenerateMipmap           js.Value
	fnGetActiveAttrib          js.Value
	fnGetActiveUniform         js.Value
	fnGetAttribLocation        js.Value
	fnGetBufferSubData         js.Value
	fnGetError                 js.Value
	fnGetParameter             js.Value
	fnGetProgramInfoLog        js.Value
	fnGetProgramParameter      js.Value
	fnGetShaderInfoLog         js.Value
	fnGetShaderParameter       js.Value
	fnGetUniformLocation       js.Value
	fnIsContextLost            js.Value
	fnIsEnabled                js.Value
	fnLinkProgram              js.Value
	fnPixelStorei              js.Value
	fnReadPixels               js.Value
	fnRenderbufferStorage      js.Value
	fnScissor                  js.Value
	fnShaderSource             js.Value
	fnTexImage2D               js.Value
	fnTexParameteri            js.Value
	fnTexSubImage2D            js.Value
	fnUniform1f                js.Value
	fnUniform1fv               js.Value
	fnUniform1i                js.Value
	fnUniform1iv               js.Value
	fnUniform2f                js.Value
	fnUniform2i                js.Value
	fnUniform3f                js.Value
	fnUniform3i                js.Value
	fnUniform4f                js.Value
	fnUniform4i                js.Value
	fnUniformMatrix2fv         js.Value
	fnUniformMatrix3fv         js.Value
	fnUniformMatrix4fv         js.Value
	fnUseProgram               js.Value
	fnVertexAttribDivisor      js.Value
	fnVertexAttribIPointer     js.Value
	fnVertexAttribPointer      js.Value
	fnViewport                 js.Value

	staging typed_array_staging
}

func ToContext(value js.Value) *Context {

	if value.IsNull() || value.IsUndefined() {
		return nil
	}

	var context Context

	context.Value = &value

	context.bindFunctions(value)

	context.DrawingBufferWidth = value.Get("drawingBufferWidth").Int()
	context.DrawingBufferHeight = value.Get("drawingBufferHeight").Int()
	context.Renderer = context.GetParameterString(ParameterRenderer)
	context.ShadingLanguageVersion = context.GetParameterString(ParameterShadingLanguageVersion)
	context.Vendor = context.GetParameterString(ParameterVendor)
	context.Version = context.GetParameterString(ParameterVersion)

	attributes := value.Call("getContextAttributes")

	if !attributes.IsNull() && !attributes.IsUndefined() {
		context.Attributes = ToContextAttributes(attributes)
	}

	return &context

}

// Binds every method of the WebGL2RenderingContext once, so that a render loop
// pays for the property lookup once instead of once per call.
func (context *Context) bindFunctions(value js.Value) {

	context.fnActiveTexture = bind_function(value, "activeTexture")
	context.fnAttachShader = bind_function(value, "attachShader")
	context.fnBindAttribLocation = bind_function(value, "bindAttribLocation")
	context.fnBindBuffer = bind_function(value, "bindBuffer")
	context.fnBindFramebuffer = bind_function(value, "bindFramebuffer")
	context.fnBindRenderbuffer = bind_function(value, "bindRenderbuffer")
	context.fnBindTexture = bind_function(value, "bindTexture")
	context.fnBindVertexArray = bind_function(value, "bindVertexArray")
	context.fnBlendEquation = bind_function(value, "blendEquation")
	context.fnBlendFunc = bind_function(value, "blendFunc")
	context.fnBlendFuncSeparate = bind_function(value, "blendFuncSeparate")
	context.fnBufferData = bind_function(value, "bufferData")
	context.fnBufferSubData = bind_function(value, "bufferSubData")
	context.fnCheckFramebufferStatus = bind_function(value, "checkFramebufferStatus")
	context.fnClear = bind_function(value, "clear")
	context.fnClearColor = bind_function(value, "clearColor")
	context.fnClearDepth = bind_function(value, "clearDepth")
	context.fnClearStencil = bind_function(value, "clearStencil")
	context.fnColorMask = bind_function(value, "colorMask")
	context.fnCompileShader = bind_function(value, "compileShader")
	context.fnCreateBuffer = bind_function(value, "createBuffer")
	context.fnCreateFramebuffer = bind_function(value, "createFramebuffer")
	context.fnCreateProgram = bind_function(value, "createProgram")
	context.fnCreateRenderbuffer = bind_function(value, "createRenderbuffer")
	context.fnCreateShader = bind_function(value, "createShader")
	context.fnCreateTexture = bind_function(value, "createTexture")
	context.fnCreateVertexArray = bind_function(value, "createVertexArray")
	context.fnCullFace = bind_function(value, "cullFace")
	context.fnDeleteBuffer = bind_function(value, "deleteBuffer")
	context.fnDeleteFramebuffer = bind_function(value, "deleteFramebuffer")
	context.fnDeleteProgram = bind_function(value, "deleteProgram")
	context.fnDeleteRenderbuffer = bind_function(value, "deleteRenderbuffer")
	context.fnDeleteShader = bind_function(value, "deleteShader")
	context.fnDeleteTexture = bind_function(value, "deleteTexture")
	context.fnDeleteVertexArray = bind_function(value, "deleteVertexArray")
	context.fnDepthFunc = bind_function(value, "depthFunc")
	context.fnDepthMask = bind_function(value, "depthMask")
	context.fnDetachShader = bind_function(value, "detachShader")
	context.fnDisable = bind_function(value, "disable")
	context.fnDisableVertexAttribArray = bind_function(value, "disableVertexAttribArray")
	context.fnDrawArrays = bind_function(value, "drawArrays")
	context.fnDrawArraysInstanced = bind_function(value, "drawArraysInstanced")
	context.fnDrawElements = bind_function(value, "drawElements")
	context.fnDrawElementsInstanced = bind_function(value, "drawElementsInstanced")
	context.fnEnable = bind_function(value, "enable")
	context.fnEnableVertexAttribArray = bind_function(value, "enableVertexAttribArray")
	context.fnFinish = bind_function(value, "finish")
	context.fnFlush = bind_function(value, "flush")
	context.fnFramebufferRenderbuffer = bind_function(value, "framebufferRenderbuffer")
	context.fnFramebufferTexture2D = bind_function(value, "framebufferTexture2D")
	context.fnFrontFace = bind_function(value, "frontFace")
	context.fnGenerateMipmap = bind_function(value, "generateMipmap")
	context.fnGetActiveAttrib = bind_function(value, "getActiveAttrib")
	context.fnGetActiveUniform = bind_function(value, "getActiveUniform")
	context.fnGetAttribLocation = bind_function(value, "getAttribLocation")
	context.fnGetBufferSubData = bind_function(value, "getBufferSubData")
	context.fnGetError = bind_function(value, "getError")
	context.fnGetParameter = bind_function(value, "getParameter")
	context.fnGetProgramInfoLog = bind_function(value, "getProgramInfoLog")
	context.fnGetProgramParameter = bind_function(value, "getProgramParameter")
	context.fnGetShaderInfoLog = bind_function(value, "getShaderInfoLog")
	context.fnGetShaderParameter = bind_function(value, "getShaderParameter")
	context.fnGetUniformLocation = bind_function(value, "getUniformLocation")
	context.fnIsContextLost = bind_function(value, "isContextLost")
	context.fnIsEnabled = bind_function(value, "isEnabled")
	context.fnLinkProgram = bind_function(value, "linkProgram")
	context.fnPixelStorei = bind_function(value, "pixelStorei")
	context.fnReadPixels = bind_function(value, "readPixels")
	context.fnRenderbufferStorage = bind_function(value, "renderbufferStorage")
	context.fnScissor = bind_function(value, "scissor")
	context.fnShaderSource = bind_function(value, "shaderSource")
	context.fnTexImage2D = bind_function(value, "texImage2D")
	context.fnTexParameteri = bind_function(value, "texParameteri")
	context.fnTexSubImage2D = bind_function(value, "texSubImage2D")
	context.fnUniform1f = bind_function(value, "uniform1f")
	context.fnUniform1fv = bind_function(value, "uniform1fv")
	context.fnUniform1i = bind_function(value, "uniform1i")
	context.fnUniform1iv = bind_function(value, "uniform1iv")
	context.fnUniform2f = bind_function(value, "uniform2f")
	context.fnUniform2i = bind_function(value, "uniform2i")
	context.fnUniform3f = bind_function(value, "uniform3f")
	context.fnUniform3i = bind_function(value, "uniform3i")
	context.fnUniform4f = bind_function(value, "uniform4f")
	context.fnUniform4i = bind_function(value, "uniform4i")
	context.fnUniformMatrix2fv = bind_function(value, "uniformMatrix2fv")
	context.fnUniformMatrix3fv = bind_function(value, "uniformMatrix3fv")
	context.fnUniformMatrix4fv = bind_function(value, "uniformMatrix4fv")
	context.fnUseProgram = bind_function(value, "useProgram")
	context.fnVertexAttribDivisor = bind_function(value, "vertexAttribDivisor")
	context.fnVertexAttribIPointer = bind_function(value, "vertexAttribIPointer")
	context.fnVertexAttribPointer = bind_function(value, "vertexAttribPointer")
	context.fnViewport = bind_function(value, "viewport")

}

func (context *Context) ActiveTexture(unit TextureUnit) {
	invoke_function(context.fnActiveTexture, uint(unit))
}

func (context *Context) AttachShader(program *Program, shader *Shader) {

	if program != nil && program.Value != nil && shader != nil && shader.Value != nil {
		invoke_function(context.fnAttachShader, *program.Value, *shader.Value)
	}

}

// Assigns an attribute location before the Program is linked, which is the only
// time the assignment takes effect.
func (context *Context) BindAttribLocation(program *Program, index uint, name string) {

	if program != nil && program.Value != nil {
		invoke_function(context.fnBindAttribLocation, *program.Value, index, name)
	}

}

// A nil Buffer binds the null object, which is how the specification says
// "unbind whatever is on this target".
func (context *Context) BindBuffer(target BufferTarget, buffer *Buffer) {

	if buffer != nil && buffer.Value != nil {
		invoke_function(context.fnBindBuffer, uint(target), *buffer.Value)
	} else {
		invoke_function(context.fnBindBuffer, uint(target), js.Null())
	}

}

func (context *Context) BindFramebuffer(target FramebufferTarget, framebuffer *Framebuffer) {

	if framebuffer != nil && framebuffer.Value != nil {
		invoke_function(context.fnBindFramebuffer, uint(target), *framebuffer.Value)
	} else {
		invoke_function(context.fnBindFramebuffer, uint(target), js.Null())
	}

}

// The specification only accepts RENDERBUFFER as the target of a Renderbuffer
// call, so the binding does not ask for it.
func (context *Context) BindRenderbuffer(renderbuffer *Renderbuffer) {

	if renderbuffer != nil && renderbuffer.Value != nil {
		invoke_function(context.fnBindRenderbuffer, renderbuffer_target, *renderbuffer.Value)
	} else {
		invoke_function(context.fnBindRenderbuffer, renderbuffer_target, js.Null())
	}

}

// A nil Texture binds the null object, see BindBuffer().
func (context *Context) BindTexture(target TextureTarget, texture *Texture) {

	if texture != nil && texture.Value != nil {
		invoke_function(context.fnBindTexture, uint(target), *texture.Value)
	} else {
		invoke_function(context.fnBindTexture, uint(target), js.Null())
	}

}

func (context *Context) BindVertexArray(array *VertexArray) {

	if array != nil && array.Value != nil {
		invoke_function(context.fnBindVertexArray, *array.Value)
	} else {
		invoke_function(context.fnBindVertexArray, js.Null())
	}

}

func (context *Context) BlendEquation(equation BlendEquation) {
	invoke_function(context.fnBlendEquation, uint(equation))
}

func (context *Context) BlendFunc(source BlendFactor, destination BlendFactor) {
	invoke_function(context.fnBlendFunc, uint(source), uint(destination))
}

func (context *Context) BlendFuncSeparate(source_rgb BlendFactor, destination_rgb BlendFactor, source_alpha BlendFactor, destination_alpha BlendFactor) {
	invoke_function(context.fnBlendFuncSeparate, uint(source_rgb), uint(destination_rgb), uint(source_alpha), uint(destination_alpha))
}

func (context *Context) BufferDataFloat32(target BufferTarget, data []float32, usage BufferUsage) {
	invoke_function(context.fnBufferData, uint(target), float32_to_typedarray(&context.staging, data), uint(usage))
}

// Allocates the currently bound Buffer without uploading anything into it.
func (context *Context) BufferDataSize(target BufferTarget, size int, usage BufferUsage) {
	invoke_function(context.fnBufferData, uint(target), size, uint(usage))
}

func (context *Context) BufferDataUint16(target BufferTarget, data []uint16, usage BufferUsage) {
	invoke_function(context.fnBufferData, uint(target), uint16_to_typedarray(&context.staging, data), uint(usage))
}

func (context *Context) BufferDataUint32(target BufferTarget, data []uint32, usage BufferUsage) {
	invoke_function(context.fnBufferData, uint(target), uint32_to_typedarray(&context.staging, data), uint(usage))
}

func (context *Context) BufferDataUint8(target BufferTarget, data []byte, usage BufferUsage) {
	invoke_function(context.fnBufferData, uint(target), uint8_to_typedarray(&context.staging, data), uint(usage))
}

// Uploads into a Buffer that BufferDataSize() allocated earlier. The offset is
// counted in bytes, as the specification counts it.
func (context *Context) BufferSubDataFloat32(target BufferTarget, offset int, data []float32) {
	invoke_function(context.fnBufferSubData, uint(target), offset, float32_to_typedarray(&context.staging, data))
}

func (context *Context) BufferSubDataUint8(target BufferTarget, offset int, data []byte) {
	invoke_function(context.fnBufferSubData, uint(target), offset, uint8_to_typedarray(&context.staging, data))
}

func (context *Context) CheckFramebufferStatus(target FramebufferTarget) FramebufferStatus {

	value := invoke_function(context.fnCheckFramebufferStatus, uint(target))

	if value.IsNull() || value.IsUndefined() {
		return FramebufferStatusUnsupported
	}

	return FramebufferStatus(value.Int())

}

// Clears the given buffers. Combine the BufferBit values with a bitwise or.
func (context *Context) Clear(mask BufferBit) {
	invoke_function(context.fnClear, uint(mask))
}

func (context *Context) ClearColor(red float32, green float32, blue float32, alpha float32) {
	invoke_function(context.fnClearColor, float64(red), float64(green), float64(blue), float64(alpha))
}

func (context *Context) ClearDepth(depth float32) {
	invoke_function(context.fnClearDepth, float64(depth))
}

func (context *Context) ClearStencil(stencil int) {
	invoke_function(context.fnClearStencil, stencil)
}

func (context *Context) ColorMask(red bool, green bool, blue bool, alpha bool) {
	invoke_function(context.fnColorMask, red, green, blue, alpha)
}

// Compiles a Shader and updates its Compiled and InfoLog properties. The
// returned error carries the info log, because a Shader that does not compile
// is a programming error that should not be silent.
func (context *Context) CompileShader(shader *Shader) error {

	if shader == nil || shader.Value == nil {
		return errors.New("Shader: Cannot compile a Shader that does not exist")
	}

	invoke_function(context.fnCompileShader, *shader.Value)

	shader.Compiled = context.GetShaderParameterBool(shader, ShaderParameterCompileStatus)
	shader.InfoLog = context.GetShaderInfoLog(shader)

	if shader.Compiled == false {
		return errors.New("Shader: Failed to compile\n" + shader.InfoLog)
	}

	return nil

}

func (context *Context) CreateBuffer() *Buffer {
	return ToBuffer(invoke_function(context.fnCreateBuffer))
}

func (context *Context) CreateFramebuffer() *Framebuffer {
	return ToFramebuffer(invoke_function(context.fnCreateFramebuffer))
}

func (context *Context) CreateProgram() *Program {
	return ToProgram(invoke_function(context.fnCreateProgram))
}

func (context *Context) CreateRenderbuffer() *Renderbuffer {
	return ToRenderbuffer(invoke_function(context.fnCreateRenderbuffer))
}

func (context *Context) CreateShader(kind ShaderType) *Shader {

	shader := ToShader(invoke_function(context.fnCreateShader, uint(kind)))

	if shader != nil {
		shader.Type = kind
	}

	return shader

}

func (context *Context) CreateTexture() *Texture {
	return ToTexture(invoke_function(context.fnCreateTexture))
}

func (context *Context) CreateVertexArray() *VertexArray {
	return ToVertexArray(invoke_function(context.fnCreateVertexArray))
}

func (context *Context) CullFace(face CullFace) {
	invoke_function(context.fnCullFace, uint(face))
}

func (context *Context) DeleteBuffer(buffer *Buffer) {

	if buffer != nil && buffer.Value != nil {

		invoke_function(context.fnDeleteBuffer, *buffer.Value)

		buffer.Value = nil

	}

}

func (context *Context) DeleteFramebuffer(framebuffer *Framebuffer) {

	if framebuffer != nil && framebuffer.Value != nil {

		invoke_function(context.fnDeleteFramebuffer, *framebuffer.Value)

		framebuffer.Value = nil

	}

}

func (context *Context) DeleteProgram(program *Program) {

	if program != nil && program.Value != nil {

		invoke_function(context.fnDeleteProgram, *program.Value)

		program.Linked = false
		program.Value = nil

	}

}

func (context *Context) DeleteRenderbuffer(renderbuffer *Renderbuffer) {

	if renderbuffer != nil && renderbuffer.Value != nil {

		invoke_function(context.fnDeleteRenderbuffer, *renderbuffer.Value)

		renderbuffer.Value = nil

	}

}

func (context *Context) DeleteShader(shader *Shader) {

	if shader != nil && shader.Value != nil {

		invoke_function(context.fnDeleteShader, *shader.Value)

		shader.Compiled = false
		shader.Value = nil

	}

}

func (context *Context) DeleteTexture(texture *Texture) {

	if texture != nil && texture.Value != nil {

		invoke_function(context.fnDeleteTexture, *texture.Value)

		texture.Value = nil

	}

}

func (context *Context) DeleteVertexArray(array *VertexArray) {

	if array != nil && array.Value != nil {

		invoke_function(context.fnDeleteVertexArray, *array.Value)

		array.Value = nil

	}

}

func (context *Context) DepthFunc(function DepthFunc) {
	invoke_function(context.fnDepthFunc, uint(function))
}

func (context *Context) DepthMask(writable bool) {
	invoke_function(context.fnDepthMask, writable)
}

func (context *Context) DetachShader(program *Program, shader *Shader) {

	if program != nil && program.Value != nil && shader != nil && shader.Value != nil {
		invoke_function(context.fnDetachShader, *program.Value, *shader.Value)
	}

}

func (context *Context) Disable(capability Capability) {
	invoke_function(context.fnDisable, uint(capability))
}

func (context *Context) DisableVertexAttribArray(index uint) {
	invoke_function(context.fnDisableVertexAttribArray, index)
}

func (context *Context) DrawArrays(mode DrawMode, first int, count int) {
	invoke_function(context.fnDrawArrays, uint(mode), first, count)
}

func (context *Context) DrawArraysInstanced(mode DrawMode, first int, count int, instances int) {
	invoke_function(context.fnDrawArraysInstanced, uint(mode), first, count, instances)
}

// Draws indexed geometry. The offset is counted in bytes into the Buffer that
// is bound to BufferTargetElementArray, as the specification counts it.
func (context *Context) DrawElements(mode DrawMode, count int, kind DataType, offset int) {
	invoke_function(context.fnDrawElements, uint(mode), count, uint(kind), offset)
}

func (context *Context) DrawElementsInstanced(mode DrawMode, count int, kind DataType, offset int, instances int) {
	invoke_function(context.fnDrawElementsInstanced, uint(mode), count, uint(kind), offset, instances)
}

func (context *Context) Enable(capability Capability) {
	invoke_function(context.fnEnable, uint(capability))
}

func (context *Context) EnableVertexAttribArray(index uint) {
	invoke_function(context.fnEnableVertexAttribArray, index)
}

func (context *Context) Finish() {
	invoke_function(context.fnFinish)
}

func (context *Context) Flush() {
	invoke_function(context.fnFlush)
}

func (context *Context) FramebufferRenderbuffer(target FramebufferTarget, attachment FramebufferAttachment, renderbuffer *Renderbuffer) {

	if renderbuffer != nil && renderbuffer.Value != nil {
		invoke_function(context.fnFramebufferRenderbuffer, uint(target), uint(attachment), renderbuffer_target, *renderbuffer.Value)
	}

}

func (context *Context) FramebufferTexture2D(target FramebufferTarget, attachment FramebufferAttachment, texture_target TextureTarget, texture *Texture, level int) {

	if texture != nil && texture.Value != nil {
		invoke_function(context.fnFramebufferTexture2D, uint(target), uint(attachment), uint(texture_target), *texture.Value, level)
	}

}

func (context *Context) FrontFace(winding FrontFace) {
	invoke_function(context.fnFrontFace, uint(winding))
}

func (context *Context) GenerateMipmap(target TextureTarget) {
	invoke_function(context.fnGenerateMipmap, uint(target))
}

func (context *Context) GetActiveAttrib(program *Program, index uint) *ActiveInfo {

	if program != nil && program.Value != nil {
		return ToActiveInfo(invoke_function(context.fnGetActiveAttrib, *program.Value, index))
	}

	return nil

}

func (context *Context) GetActiveUniform(program *Program, index uint) *ActiveInfo {

	if program != nil && program.Value != nil {
		return ToActiveInfo(invoke_function(context.fnGetActiveUniform, *program.Value, index))
	}

	return nil

}

// Returns the location of a vertex attribute, or -1 when the linked Program
// does not use one by that name.
func (context *Context) GetAttribLocation(program *Program, name string) int {

	if program != nil && program.Value != nil {

		value := invoke_function(context.fnGetAttribLocation, *program.Value, name)

		if !value.IsNull() && !value.IsUndefined() {
			return value.Int()
		}

	}

	return -1

}

// Reads a Buffer back into Go memory. This is a synchronous read of data the
// GPU owns, so it stalls the pipeline; it belongs in a debugging path, not in a
// render loop.
func (context *Context) GetBufferSubDataFloat32(target BufferTarget, offset int, length int) []float32 {

	array := js.Global().Get("Float32Array").New(length)

	invoke_function(context.fnGetBufferSubData, uint(target), offset, array)

	return typedarray_to_float32(array)

}

func (context *Context) GetBufferSubDataUint8(target BufferTarget, offset int, length int) []byte {

	array := js.Global().Get("Uint8Array").New(length)

	invoke_function(context.fnGetBufferSubData, uint(target), offset, array)

	return typedarray_to_uint8(array)

}

// Returns the first error that was flagged since the last call, and resets the
// error state. ErrorCodeNoError means the Context is clean.
func (context *Context) GetError() ErrorCode {

	value := invoke_function(context.fnGetError)

	if value.IsNull() || value.IsUndefined() {
		return ErrorCodeNoError
	}

	return ErrorCode(value.Int())

}

func (context *Context) GetParameterBool(parameter Parameter) bool {

	var result bool

	value := invoke_function(context.fnGetParameter, uint(parameter))

	if !value.IsNull() && !value.IsUndefined() {
		result = value.Bool()
	}

	return result

}

func (context *Context) GetParameterFloat(parameter Parameter) float64 {

	var result float64

	value := invoke_function(context.fnGetParameter, uint(parameter))

	if !value.IsNull() && !value.IsUndefined() {
		result = value.Float()
	}

	return result

}

func (context *Context) GetParameterInt(parameter Parameter) int {

	var result int

	value := invoke_function(context.fnGetParameter, uint(parameter))

	if !value.IsNull() && !value.IsUndefined() {
		result = value.Int()
	}

	return result

}

// Returns a parameter that the specification defines as an array of numbers,
// such as ParameterViewport or ParameterMaxViewportDims.
func (context *Context) GetParameterInts(parameter Parameter) []int {

	result := make([]int, 0)

	value := invoke_function(context.fnGetParameter, uint(parameter))

	if !value.IsNull() && !value.IsUndefined() {

		for v := 0; v < value.Length(); v++ {
			result = append(result, value.Index(v).Int())
		}

	}

	return result

}

func (context *Context) GetParameterString(parameter Parameter) string {

	var result string

	value := invoke_function(context.fnGetParameter, uint(parameter))

	if !value.IsNull() && !value.IsUndefined() {
		result = value.String()
	}

	return result

}

func (context *Context) GetProgramInfoLog(program *Program) string {

	var result string

	if program != nil && program.Value != nil {

		value := invoke_function(context.fnGetProgramInfoLog, *program.Value)

		if !value.IsNull() && !value.IsUndefined() {
			result = value.String()
		}

	}

	return result

}

func (context *Context) GetProgramParameterBool(program *Program, parameter ProgramParameter) bool {

	var result bool

	if program != nil && program.Value != nil {

		value := invoke_function(context.fnGetProgramParameter, *program.Value, uint(parameter))

		if !value.IsNull() && !value.IsUndefined() {
			result = value.Bool()
		}

	}

	return result

}

func (context *Context) GetProgramParameterInt(program *Program, parameter ProgramParameter) int {

	var result int

	if program != nil && program.Value != nil {

		value := invoke_function(context.fnGetProgramParameter, *program.Value, uint(parameter))

		if !value.IsNull() && !value.IsUndefined() {
			result = value.Int()
		}

	}

	return result

}

func (context *Context) GetShaderInfoLog(shader *Shader) string {

	var result string

	if shader != nil && shader.Value != nil {

		value := invoke_function(context.fnGetShaderInfoLog, *shader.Value)

		if !value.IsNull() && !value.IsUndefined() {
			result = value.String()
		}

	}

	return result

}

func (context *Context) GetShaderParameterBool(shader *Shader, parameter ShaderParameter) bool {

	var result bool

	if shader != nil && shader.Value != nil {

		value := invoke_function(context.fnGetShaderParameter, *shader.Value, uint(parameter))

		if !value.IsNull() && !value.IsUndefined() {
			result = value.Bool()
		}

	}

	return result

}

func (context *Context) GetShaderParameterInt(shader *Shader, parameter ShaderParameter) int {

	var result int

	if shader != nil && shader.Value != nil {

		value := invoke_function(context.fnGetShaderParameter, *shader.Value, uint(parameter))

		if !value.IsNull() && !value.IsUndefined() {
			result = value.Int()
		}

	}

	return result

}

// Returns the location of a uniform, or nil when the linked Program does not
// use one by that name.
func (context *Context) GetUniformLocation(program *Program, name string) *UniformLocation {

	if program != nil && program.Value != nil {

		location := ToUniformLocation(invoke_function(context.fnGetUniformLocation, *program.Value, name))

		if location != nil {
			location.Name = name
		}

		return location

	}

	return nil

}

func (context *Context) IsContextLost() bool {

	var result bool

	value := invoke_function(context.fnIsContextLost)

	if !value.IsNull() && !value.IsUndefined() {
		result = value.Bool()
	}

	return result

}

func (context *Context) IsEnabled(capability Capability) bool {

	var result bool

	value := invoke_function(context.fnIsEnabled, uint(capability))

	if !value.IsNull() && !value.IsUndefined() {
		result = value.Bool()
	}

	return result

}

// Links a Program and updates its Linked and InfoLog properties. The returned
// error carries the info log, for the same reason CompileShader() does.
func (context *Context) LinkProgram(program *Program) error {

	if program == nil || program.Value == nil {
		return errors.New("Program: Cannot link a Program that does not exist")
	}

	invoke_function(context.fnLinkProgram, *program.Value)

	program.Linked = context.GetProgramParameterBool(program, ProgramParameterLinkStatus)
	program.InfoLog = context.GetProgramInfoLog(program)

	if program.Linked == false {
		return errors.New("Program: Failed to link\n" + program.InfoLog)
	}

	return nil

}

// The specification types every pixelStorei value as a GLint, including the
// three WebGL flags, so a bool is converted rather than handed over to
// JavaScript's own coercion.
func (context *Context) PixelStoreBool(parameter PixelStoreParameter, enabled bool) {

	if enabled == true {
		context.PixelStoreInt(parameter, 1)
	} else {
		context.PixelStoreInt(parameter, 0)
	}

}

func (context *Context) PixelStoreInt(parameter PixelStoreParameter, value int) {
	invoke_function(context.fnPixelStorei, uint(parameter), value)
}

// Reads the colour buffer of whatever is bound to FramebufferTargetRead back
// into Go memory. This waits for the GPU, so it belongs in a screenshot path
// rather than in a render loop.
//
// The destination view is built from kind and sized from format, because the
// specification requires the two to match and answers a mismatch with
// INVALID_OPERATION and an untouched view — zeroes, with nothing to read the
// failure from. A kind that is not a pixel type, or a format whose component
// count this package does not know, returns nil instead of a wrong answer.
//
// The bytes come back exactly as the browser wrote them, so a caller reading a
// 16-bit or floating point buffer decodes them itself, and rows are padded to
// the current ParameterPackAlignment. Set PixelStoreParameterPackAlignment to 1
// for tightly packed rows.
func (context *Context) ReadPixels(x int, y int, width int, height int, format TextureFormat, kind DataType) []byte {

	components := format_components(format)
	_, size := pixel_typedarray(kind)

	if components == 0 || size == 0 || width <= 0 || height <= 0 {
		return nil
	}

	alignment := context.GetParameterInt(ParameterPackAlignment)

	if alignment < 1 {
		alignment = 4
	}

	row := width * components * size
	stride := ((row + alignment - 1) / alignment) * alignment
	total := stride*(height-1) + row

	array := new_pixel_typedarray(kind, (total+size-1)/size)

	if array.IsUndefined() {
		return nil
	}

	invoke_function(context.fnReadPixels, x, y, width, height, uint(format), uint(kind), array)

	return typedarray_to_uint8(array)

}

// See BindRenderbuffer() on the missing target parameter.
func (context *Context) RenderbufferStorage(format RenderbufferFormat, width int, height int) {
	invoke_function(context.fnRenderbufferStorage, renderbuffer_target, uint(format), width, height)
}

func (context *Context) Scissor(x int, y int, width int, height int) {
	invoke_function(context.fnScissor, x, y, width, height)
}

func (context *Context) ShaderSource(shader *Shader, source string) {

	if shader != nil && shader.Value != nil {

		shader.Source = source

		invoke_function(context.fnShaderSource, *shader.Value, source)

	}

}

// Uploads a level of a Texture. An empty pixels slice allocates the storage
// without filling it, which is what an attachment of a Framebuffer wants.
//
// pixels is raw bytes, and the view the browser reads them through is built
// from kind — a Uint8Array for DataTypeUnsignedByte, a Uint16Array for the
// 16-bit types, a Float32Array for DataTypeFloat — because the specification
// answers a view that does not match the type with INVALID_OPERATION. A kind
// that is not a pixel type, or a byte count that is not a whole number of
// elements of it, is an error rather than a call the driver will refuse.
func (context *Context) TexImage2D(target TextureTarget, level int, internal_format TextureFormat, width int, height int, format TextureFormat, kind DataType, pixels []byte) error {

	_, size := pixel_typedarray(kind)

	if size == 0 {
		return errors.New("Texture: TexImage2D was given a type that is not a pixel type")
	}

	if len(pixels) == 0 {

		invoke_function(context.fnTexImage2D, uint(target), level, uint(internal_format), width, height, 0, uint(format), uint(kind), js.Null())

		return nil

	}

	array := pixels_to_typedarray(&context.staging, kind, pixels)

	if array.IsUndefined() {
		return errors.New("Texture: TexImage2D was given a byte count that is not a whole number of pixels")
	}

	invoke_function(context.fnTexImage2D, uint(target), level, uint(internal_format), width, height, 0, uint(format), uint(kind), array)

	return nil

}

func (context *Context) TexParameterFilter(target TextureTarget, parameter TextureParameter, filter TextureFilter) {
	invoke_function(context.fnTexParameteri, uint(target), uint(parameter), int(filter))
}

func (context *Context) TexParameterInt(target TextureTarget, parameter TextureParameter, value int) {
	invoke_function(context.fnTexParameteri, uint(target), uint(parameter), value)
}

func (context *Context) TexParameterWrap(target TextureTarget, parameter TextureParameter, wrap TextureWrap) {
	invoke_function(context.fnTexParameteri, uint(target), uint(parameter), int(wrap))
}

// Overwrites part of a level of a Texture. The view is built from kind, see
// TexImage2D().
func (context *Context) TexSubImage2D(target TextureTarget, level int, x int, y int, width int, height int, format TextureFormat, kind DataType, pixels []byte) error {

	array := pixels_to_typedarray(&context.staging, kind, pixels)

	if array.IsUndefined() {
		return errors.New("Texture: TexSubImage2D was given a type or a byte count it cannot upload")
	}

	invoke_function(context.fnTexSubImage2D, uint(target), level, x, y, width, height, uint(format), uint(kind), array)

	return nil

}

func (context *Context) Uniform1f(location *UniformLocation, x float32) {

	if location != nil && location.Value != nil {
		invoke_function(context.fnUniform1f, *location.Value, float64(x))
	}

}

func (context *Context) Uniform1fv(location *UniformLocation, data []float32) {

	if location != nil && location.Value != nil {
		invoke_function(context.fnUniform1fv, *location.Value, float32_to_typedarray(&context.staging, data))
	}

}

func (context *Context) Uniform1i(location *UniformLocation, x int32) {

	if location != nil && location.Value != nil {
		invoke_function(context.fnUniform1i, *location.Value, int(x))
	}

}

func (context *Context) Uniform1iv(location *UniformLocation, data []int32) {

	if location != nil && location.Value != nil {
		invoke_function(context.fnUniform1iv, *location.Value, int32_to_typedarray(&context.staging, data))
	}

}

func (context *Context) Uniform2f(location *UniformLocation, x float32, y float32) {

	if location != nil && location.Value != nil {
		invoke_function(context.fnUniform2f, *location.Value, float64(x), float64(y))
	}

}

func (context *Context) Uniform2i(location *UniformLocation, x int32, y int32) {

	if location != nil && location.Value != nil {
		invoke_function(context.fnUniform2i, *location.Value, int(x), int(y))
	}

}

func (context *Context) Uniform3f(location *UniformLocation, x float32, y float32, z float32) {

	if location != nil && location.Value != nil {
		invoke_function(context.fnUniform3f, *location.Value, float64(x), float64(y), float64(z))
	}

}

func (context *Context) Uniform3i(location *UniformLocation, x int32, y int32, z int32) {

	if location != nil && location.Value != nil {
		invoke_function(context.fnUniform3i, *location.Value, int(x), int(y), int(z))
	}

}

func (context *Context) Uniform4f(location *UniformLocation, x float32, y float32, z float32, w float32) {

	if location != nil && location.Value != nil {
		invoke_function(context.fnUniform4f, *location.Value, float64(x), float64(y), float64(z), float64(w))
	}

}

func (context *Context) Uniform4i(location *UniformLocation, x int32, y int32, z int32, w int32) {

	if location != nil && location.Value != nil {
		invoke_function(context.fnUniform4i, *location.Value, int(x), int(y), int(z), int(w))
	}

}

// Uploads a 2x2 matrix in column major order, which is the order the
// specification requires and the order GLSL reads.
func (context *Context) UniformMatrix2fv(location *UniformLocation, data []float32) {

	if location != nil && location.Value != nil {
		invoke_function(context.fnUniformMatrix2fv, *location.Value, false, float32_to_typedarray(&context.staging, data))
	}

}

// Uploads a 3x3 matrix in column major order, see UniformMatrix2fv().
func (context *Context) UniformMatrix3fv(location *UniformLocation, data []float32) {

	if location != nil && location.Value != nil {
		invoke_function(context.fnUniformMatrix3fv, *location.Value, false, float32_to_typedarray(&context.staging, data))
	}

}

// Uploads a 4x4 matrix in column major order, see UniformMatrix2fv().
func (context *Context) UniformMatrix4fv(location *UniformLocation, data []float32) {

	if location != nil && location.Value != nil {
		invoke_function(context.fnUniformMatrix4fv, *location.Value, false, float32_to_typedarray(&context.staging, data))
	}

}

func (context *Context) UseProgram(program *Program) {

	if program != nil && program.Value != nil {
		invoke_function(context.fnUseProgram, *program.Value)
	} else {
		invoke_function(context.fnUseProgram, js.Null())
	}

}

func (context *Context) VertexAttribDivisor(index uint, divisor uint) {
	invoke_function(context.fnVertexAttribDivisor, index, divisor)
}

// Describes an integer vertex attribute, which a WebGL2 shader reads as an int
// or a uint rather than as a float.
func (context *Context) VertexAttribIPointer(index uint, size int, kind DataType, stride int, offset int) {
	invoke_function(context.fnVertexAttribIPointer, index, size, uint(kind), stride, offset)
}

// Describes a floating point vertex attribute. The stride and the offset are
// counted in bytes into the Buffer that is bound to BufferTargetArray.
func (context *Context) VertexAttribPointer(index uint, size int, kind DataType, normalized bool, stride int, offset int) {
	invoke_function(context.fnVertexAttribPointer, index, size, uint(kind), normalized, stride, offset)
}

func (context *Context) Viewport(x int, y int, width int, height int) {
	invoke_function(context.fnViewport, x, y, width, height)
}
