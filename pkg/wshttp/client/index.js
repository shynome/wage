const path = require("path");
const wasmUrl = path.join(module.path, "wshttp-go.wasm");
require("./wasm_exec_node");
require("websocket-polyfill");
const b = fs.readFileSync(wasmUrl);
const go = new Go();

exports.GoFetchInit = Promise.resolve(1)
  .then(() => WebAssembly.instantiate(b, go.importObject))
  .then(({ instance }) => instance)
  .then((inst) => {
    go.run(inst);
  });
