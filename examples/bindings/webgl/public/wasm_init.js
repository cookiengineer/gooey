(async () => {

	const wasm_buffer = await fetch("main.wasm").then((response) => response.arrayBuffer());

	const go = new Go();
	const module = await WebAssembly.instantiate(wasm_buffer, go.importObject);

	go.run(module.instance);

})();

