package demo

import (
  "github.com/EngoEngine/engo"
  "github.com/EngoEngine/engo/common"

  "github.com/Noofbiz/pixelshader"
)

type demoScene struct{}

func (*demoScene) Type() string { return "demo scene" }

func (*demoScene) Preload() {
	common.AddShader(pShader)
}

func (*demoScene) Setup(u engo.Updater) {
	w := u.(*ecs.World)

	var renderable *common.Renderable
	var notrenderable *common.NotRenderable
	w.AddSystemInterface(&common.RenderSystem{}, renderable, notrenderable)

	bg := sprite{BasicEntity: ecs.NewBasic()}
	bg.Drawable = pixelshader.PixelRegion{}
	bg.SetShader(pShader)
	w.AddEntity(&bg)
}

pShader := &pixelshader.PixelShader{FragShader: `
  #ifdef GL_ES
  #define LOWP lowp
  precision mediump float;
  #else
  #define LOWP
  #endif
  uniform vec2 u_resolution;  // Canvas size (width,height)
  uniform vec2 u_mouse;       // mouse position in screen pixels
  uniform float u_time;       // Time in seconds since load
  void main()
  {
    vec2 st = gl_FragCoord.xy/u_resolution;
  	gl_FragColor = vec4(st.x,st.y,0.0,1.0);
  }
  `}

func main() {
	engo.Run(engo.RunOptions{
		Title:                      "Pixel Shader Demo!",
		Width:                      512, //16
		Height:                     288, //9
	}, &demoScene{})
}
