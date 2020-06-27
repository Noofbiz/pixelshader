# pixelShader
A pixel shader for engo

Pretty straightforward. Creates a pixel shader using the provided FragShader.
Play with neat designs and such from places like [Shader Toy](https://www.shadertoy.com)
right inside your engo games! The demo shows off how to use it in engo.

The fragment shader has access to three uniforms. If you think of more, submit a PR or
create an issue and I can add it. The uniforms are the following:

```
uniform vec2 u_resolution;  // Canvas size (width,height)
uniform vec2 u_mouse;       // mouse position in screen pixels
uniform float u_time;       // Time in seconds since load
```
