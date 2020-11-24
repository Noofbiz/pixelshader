package pixelshader

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/EngoEngine/gl"
)

const vert = `
attribute vec2 position;
uniform mat3 matrixProjection;
uniform mat3 matrixView;
uniform mat3 matrixModel;
void main()
{
	vec3 matr = matrixProjection * matrixView * matrixModel * vec3(position, 1.0);
  gl_Position = vec4(matr.xy, 0, matr.z);
}
`

type PixelShader struct {
	FragShader string

	program                                         *gl.Program
	vertices                                        []float32
	buffer                                          *gl.Buffer
	coordinates                                     int
	resolutionLocation, mouseLocation, timeLocation *gl.UniformLocation

	matrixProjection *gl.UniformLocation
	matrixView       *gl.UniformLocation
	matrixModel      *gl.UniformLocation

	projectionMatrix *engo.Matrix
	viewMatrix       *engo.Matrix
	modelMatrix      *engo.Matrix

	camera        *common.CameraSystem
	cameraEnabled bool
}

func (s *PixelShader) Setup(w *ecs.World) error {
	s.vertices = []float32{-1, -1, -1, 1, 1, -1, 1, 1, -1, 1, 1, -1}
	s.buffer = engo.Gl.CreateBuffer()
	var err error
	s.program, err = common.LoadShader(vert, s.FragShader)
	if err != nil {
		return err
	}
	s.resolutionLocation = engo.Gl.GetUniformLocation(s.program, "u_resolution")
	s.mouseLocation = engo.Gl.GetUniformLocation(s.program, "u_mouse")
	s.timeLocation = engo.Gl.GetUniformLocation(s.program, "u_time")
	s.coordinates = engo.Gl.GetAttribLocation(s.program, "position")

	s.matrixProjection = engo.Gl.GetUniformLocation(s.program, "matrixProjection")
	s.matrixView = engo.Gl.GetUniformLocation(s.program, "matrixView")
	s.matrixModel = engo.Gl.GetUniformLocation(s.program, "matrixModel")

	s.projectionMatrix = engo.IdentityMatrix()
	s.viewMatrix = engo.IdentityMatrix()
	s.modelMatrix = engo.IdentityMatrix()
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

	s.projectionMatrix.Identity()
	if engo.ScaleOnResize() {
		s.projectionMatrix.Scale(1/(engo.GameWidth()/2), 1/(-engo.GameHeight()/2))
	} else {
		s.projectionMatrix.Scale(1/(engo.CanvasWidth()/(2*engo.CanvasScale())), 1/(-engo.CanvasHeight()/(2*engo.CanvasScale())))
	}

	s.viewMatrix.Identity()
	if s.cameraEnabled {
		s.viewMatrix.Scale(1/s.camera.Z(), 1/s.camera.Z())
		s.viewMatrix.Translate(-s.camera.X(), -s.camera.Y()).Rotate(s.camera.Angle())
	} else {
		scaleX, scaleY := s.projectionMatrix.ScaleComponent()
		s.viewMatrix.Translate(-1/scaleX, 1/scaleY)
	}

	engo.Gl.UniformMatrix3fv(s.matrixProjection, false, s.projectionMatrix.Val[:])
	engo.Gl.UniformMatrix3fv(s.matrixView, false, s.viewMatrix.Val[:])
}

func (s *PixelShader) Draw(render *common.RenderComponent, space *common.SpaceComponent) {
	engo.Gl.Uniform2f(s.resolutionLocation, space.Width, space.Height)
	engo.Gl.Uniform2f(s.mouseLocation, engo.Input.Mouse.X, engo.Input.Mouse.Y)
	engo.Gl.Uniform1f(s.timeLocation, engo.Time.Time())

	s.modelMatrix.Identity().Scale(engo.GetGlobalScale().X, engo.GetGlobalScale().Y).Translate(space.Position.X, space.Position.Y)
	if space.Rotation != 0 {
		s.modelMatrix.Rotate(space.Rotation)
	}
	s.modelMatrix.Scale(render.Scale.X, render.Scale.Y)
	s.modelMatrix.Scale(space.Width, space.Height)

	engo.Gl.UniformMatrix3fv(s.matrixModel, false, s.modelMatrix.Val[:])

	engo.Gl.DrawArrays(engo.Gl.TRIANGLES, 0, 6)
}

func (s *PixelShader) Post() {
	// Cleanup
	engo.Gl.DisableVertexAttribArray(s.coordinates)

	engo.Gl.BindBuffer(engo.Gl.ARRAY_BUFFER, nil)

	engo.Gl.Disable(engo.Gl.BLEND)
}

func (s *PixelShader) SetCamera(c *common.CameraSystem) {
	if s.cameraEnabled {
		s.camera = c
	}
}
