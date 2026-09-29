// Package virtual provides a DOM reconciler that patches a parent element's
// children in place instead of replacing them wholesale.
//
// The algorithm operates on the virtual.Element interface so that it can be
// unit tested natively with the pure-Go virtual/memory.Element double, while
// virtual/dom.Element adapts the real bindings/dom.Element for wasm builds.
package virtual
