package main

import (
	"image/color"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"

	"github.com/Noofbiz/pixelshader"
)

type sprite struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

type demoScene struct{}

func (*demoScene) Type() string { return "demo scene" }

func (*demoScene) Preload() {
	common.AddShader(pShader)

	engo.Files.Load("icon.png")
	engo.Files.Load("moon.png")
	engo.Files.Load("rainbow.png")
	engo.Files.Load("sun.png")
}

func (*demoScene) Setup(u engo.Updater) {
	w := u.(*ecs.World)

	var renderable *common.Renderable
	var notrenderable *common.NotRenderable
	w.AddSystemInterface(&common.RenderSystem{}, renderable, notrenderable)

	w.AddSystem(&common.FPSSystem{Display: true})

	bg := sprite{BasicEntity: ecs.NewBasic()}
	tex0, _ := common.LoadedSprite("moon.png")
	tex1, _ := common.LoadedSprite("rainbow.png")
	tex2, _ := common.LoadedSprite("sun.png")
	bg.Drawable = pixelshader.PixelRegion{
		Tex0: tex0,
		Tex1: tex1,
		Tex2: tex2,
	}
	bg.SetShader(pShader)
	bg.SetZIndex(0)
	w.AddEntity(&bg)

	rect := sprite{BasicEntity: ecs.NewBasic()}
	rect.Drawable = common.Rectangle{}
	rect.Width = 50
	rect.Height = 50
	rect.Color = color.RGBA{0, 0, 100, 255}
	rect.SetCenter(engo.Point{X: 100, Y: 100})
	rect.SetZIndex(1)
	w.AddEntity(&rect)

	hero := sprite{BasicEntity: ecs.NewBasic()}
	hero.Drawable, _ = common.LoadedSprite("icon.png")
	hero.SetZIndex(1)
	w.AddEntity(&hero)
}

var pShader = &pixelshader.PixelShader{FragShader: `
	// Created by inigo quilez - iq/2014
	// License Creative Commons Attribution-NonCommercial-ShareAlike 3.0 Unported License.
  #ifdef GL_ES
  #define LOWP lowp
  precision mediump float;
  #else
  #define LOWP
  #endif
  uniform vec2 u_resolution;  // Canvas size (width,height)
  uniform vec2 u_mouse;       // mouse position in screen pixels
  uniform float u_time;       // Time in seconds since load
	uniform sampler2D u_tex0;   // Drawable Tex0
	uniform sampler2D u_tex1;   // Drawable Tex1
	uniform sampler2D u_tex2;   // Drawable Tex2
  void main()
  {
    vec2 p = gl_FragCoord.xy/u_resolution.xy;
		vec2 texPos = vec2(p.x, -p.y);
		vec2 q = p - vec2(0.33,0.7);

		float pct = abs(sin(u_time));
		float pct2 = abs(cos(u_time));

		vec3 col = mix( vec3(1.0,0.3,0.0), vec3(1.0,0.8,0.3), sqrt(p.y) );

		float r = 0.2 + 0.1*cos( atan(q.y,q.x)*10.0 + 20.0*q.x + 1.0);
		col *= smoothstep( r, r+0.01, length( q ) );

		r = 0.015;
		r += 0.002*sin(120.0*q.y);
		r += exp(-40.0*p.y);
    col *= 1.0 - (1.0-smoothstep(r,r+0.002, abs(q.x-0.25*sin(2.0*q.y))))*(1.0-smoothstep(0.0,0.1,q.y));
		if (any(greaterThan(col, vec3(0.0, 0.0, 0.0))))
  		gl_FragColor = vec4(col, 1.0);
		else
			gl_FragColor = mix(mix(texture2D(u_tex2, texPos),texture2D(u_tex0, texPos),pct2), texture2D(u_tex1, texPos), pct);
  }
  `}

func main() {
	engo.Run(engo.RunOptions{
		Title:  "Pixel Shader Demo!",
		Width:  512, //16
		Height: 288, //9
	}, &demoScene{})
}
