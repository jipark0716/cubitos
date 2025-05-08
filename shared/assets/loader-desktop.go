//go:build !js && !wasm

package assets

import "os"

type DesktopLoader struct{}

func (l DesktopLoader) Load(path string) ([]byte, error) {
	return os.ReadFile("./public/" + path)
}

func init() {
	LoaderInstance = &DesktopLoader{}
}
