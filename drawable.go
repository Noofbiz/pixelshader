package pixelshader

import (
	"github.com/EngoEngine/engo/common"
	"github.com/EngoEngine/gl"
)

// PixelRegion is a drawable to use with the pixel shader!
type PixelRegion struct {
	// Tex0, Tex1 and Tex2 are textures available to the shader
	// They're available as u_tex0, u_tex1, and u_tex2
	// Note: The position for the textures is opposite the frag coordinates
	// Due to OpenGL storing the textures flipped on the y-axis.
	// so if
	// vec2 p = gl_FragCoord.xy/u_resolution.xy;
	// the coordinate positions are
	// vec2 texPos = vec2(p.x, -p.y);
	Tex0, Tex1, Tex2 common.Drawable
}

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
