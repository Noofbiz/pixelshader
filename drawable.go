package pixelshader

import "github.com/EngoEngine/gl"

// PixelRegion is a drawable to use with the pixel shader!
type PixelRegion struct{}

func (p PixelRegion) Texture() *gl.Texture {
	return nil
}

func (p PixelRegion) Width() float32 {
	return 0
}

func (p PixelRegion) Height() float32 {
	return 0
}

func (p PixelRegion) View() (float32, float32, float32, float32) {
	return 0, 0, 0, 0
}

func (p PixelRegion) Close() {}
