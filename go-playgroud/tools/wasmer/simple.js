const { readFileSync } = require('fs');
const { WASI } = require('wasi');
const { argv, env } = require('process');

(async function () {
    const wasi = new WASI({ args: argv, env });
    const importObject = { wasi_snapshot_preview1: wasi.wasiImport, /** or: wasi_unstable: wasi.wasiImport **/ };
    const wasm = await WebAssembly.compile(readFileSync("./sum.wasm"));
    const instance = await WebAssembly.instantiate(wasm, importObject);
    wasi.start(instance);

    const { Sum } = instance.exports;
    console.log(Sum(3, 15));
}());