# WebGL Example

This example uses [webgl](/bindings/webgl) to put a triangle on a canvas without a single line of
JavaScript. It asks the canvas for a `WebGL2RenderingContext` with its own context attributes,
compiles a `#version 300 es` vertex and fragment shader, links them into a program, uploads one
interleaved vertex buffer and one index buffer through a vertex array, and draws the result with
`DrawElements()` over a cleared frame. It also registers the `webglcontextlost` and
`webglcontextrestored` listeners on the canvas, so the page survives a context that goes away.

Build and run it with `./build.sh`, which compiles `main.go` to `public/main.wasm`, copies Go's
`wasm_exec.js` next to it and serves `public/` on <http://localhost:3000>. The console logs the
version string the engine reports, the renderer name, and one line per frame drawn. This is what the
page looks like:

![WebGL Example](/examples/bindings/webgl/screenshot.png)

## Used in practice

The bindings were written for a real renderer before they were tidied into a package. A turn-based
strategy game written in Go draws its isometric voxel maps through them, both in the browser and,
through the same call surface, natively: one instanced draw call per distinct model, 78 calls for a
40×40 map, and the browser frame matches the native frame to within 1.04 % of pixels. This is one of
its maps, rendered in the browser:

![A 40×40 isometric voxel map rendered through the webgl bindings](/examples/bindings/webgl/tof3-map.png)

And every unit kind it ships, each rendered on its own tile:

![Nineteen voxel unit kinds rendered through the webgl bindings](/examples/bindings/webgl/tof3-units.png)
