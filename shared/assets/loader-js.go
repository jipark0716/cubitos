//go:build js && wasm

package assets

import (
	"io"
	"net/http"
)

type WasmLoader struct{}

func init() {
	LoaderInstance = &WasmLoader{}
}

func (l WasmLoader) Load(path string) ([]byte, error) {
	resp, err := http.Get("/" + path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
