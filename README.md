# Go2WebAssembly

Simple app that integrates Web Assembly to the JS enviroment. The Go code is built into a wasm file, the wasm file can read by the web app. The Go code can add functions into the global JS object, making Go functions runnable from the web app.

* To build the wasm file run `npm run build:wasm`
* To start the Go server run `npm run serve`
