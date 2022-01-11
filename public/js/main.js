var wasm;

function loadWasm() {
    console.log("Loading wasm file");

    const go = new Go(); // Defined in wasm_exec.js
    const WASM_URL = '/js/main.wasm';
    
    if ('instantiateStreaming' in WebAssembly) {
        WebAssembly.instantiateStreaming(fetch(WASM_URL), go.importObject).then(function (obj) {
            wasm = obj.instance;
            go.run(wasm);
        })
    } else {
        fetch(WASM_URL).then(resp =>
            resp.arrayBuffer()
        ).then(bytes =>
            WebAssembly.instantiate(bytes, go.importObject).then(function (obj) {
                wasm = obj.instance;
                go.run(wasm);
            })
        )
    }    
}

window.addEventListener("DOMContentLoaded", () => {
    loadWasm();

    document.querySelector("#load-image").addEventListener("click", (ev) => {
        ev.preventDefault();

        console.log("Loading image");
        wasm.exports.loadImage();
    })
})