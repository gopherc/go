// +build js,wasm,goc

package debug

func GOID() uintptr {
	return getDebugG()
}