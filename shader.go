package pixelshader

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/EngoEngine/gl"
)

const vert = `
attribute vec2 position;
void main()
{
  gl_Position = vec4(position,0.0,1.0);
}
`

type PixelShader struct {
	FragShader string

	program                                         *gl.Program
	vertices                                        []float32
	buffer                                          *gl.Buffer
	coordinates                                     int
	resolutionLocation, mouseLocation, timeLocation *gl.UniformLocation
}

func (s *PixelShader) Setup(w *ecs.World) error {
	s.vertices = []float32{-1, -1, -1, 1, 1, -1, 1, 1, -1, 1, 1, -1}
	s.buffer = engo.Gl.CreateBuffer()
	var err error
	s.program, err = common.LoadShader(vert, frag)
	if err != nil {
		return err
	}
	s.resolutionLocation = engo.Gl.GetUniformLocation(s.program, "u_resolution")
	s.mouseLocation = engo.Gl.GetUniformLocation(s.program, "u_mouse")
	s.timeLocation = engo.Gl.GetUniformLocation(s.program, "u_time")
	s.coordinates = engo.Gl.GetAttribLocation(s.program, "position")
	return nil
}

func (s *PixelShader) Pre() {
	engo.Gl.Enable(engo.Gl.BLEND)
	engo.Gl.BlendFunc(engo.Gl.SRC_ALPHA, engo.Gl.ONE_MINUS_SRC_ALPHA)
	// Enable shader and buffer, enable attributes in shader
	engo.Gl.UseProgram(s.program)
	engo.Gl.BindBuffer(engo.Gl.ARRAY_BUFFER, s.buffer)
	engo.Gl.BufferData(engo.Gl.ARRAY_BUFFER, s.vertices, engo.Gl.STATIC_DRAW)
	engo.Gl.VertexAttribPointer(s.coordinates, 2, engo.Gl.FLOAT, false, 0, 0)
	engo.Gl.EnableVertexAttribArray(s.coordinates)
}

func (s *PixelShader) Draw(render *common.RenderComponent, space *common.SpaceComponent) {
	engo.Gl.Uniform2f(s.resolutionLocation, engo.CanvasWidth(), engo.CanvasHeight())
	engo.Gl.Uniform2f(s.mouseLocation, engo.Input.Mouse.X, engo.Input.Mouse.Y)
	engo.Gl.Uniform1f(s.timeLocation, engo.Time.Time())
	engo.Gl.DrawArrays(engo.Gl.TRIANGLES, 0, 6)
}

func (s *PixelShader) Post() {
	// Cleanup
	engo.Gl.DisableVertexAttribArray(s.coordinates)

	engo.Gl.BindBuffer(engo.Gl.ARRAY_BUFFER, nil)

	engo.Gl.Disable(engo.Gl.BLEND)
}

func (s *PixelShader) SetCamera(c *common.CameraSystem) {}
