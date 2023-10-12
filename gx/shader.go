package gx

import (
	"github.com/hajimehoshi/ebiten/v2"
	//"fmt"
)

type Shader = ebiten.Shader
type ShaderOptions struct {
	Shader *Shader
	Uniforms map[string] any
	Images [4]*Image
}

var (
	// The shader is for example only.
	SolidWhiteColorShader = MustNewShader([]byte(`
		package main
		
		var Random float
		
		func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
			//ts := imageSrcTextureSize()
			
			//_, size := imageSrcRegionOnTexture()
			/*return vec4(
				position.y/size.y,
				position.y/size.y,
				position.y/size.y,
				position.y/size.y,
			)*/
			/*py := int(position.y / size.y) % 5
			px := int(position.x / size.x) % 5
			if py >= 1 && px >= 1 {
				return vec4(
					1,
					0,
					0,
					1,
				)
			}*/
			
			/*ret := vec4(
				0,
				sin(position.x),
				sin(position.y),
				1,
			)
			
			return imageSrc0UnsafeAt(texCoord) * (ret)*/
			return imageSrc0UnsafeAt(texCoord) * 
				vec4(1, 1, Random, Random)
		}
	`))
)

func MustNewShader(src []byte) (*Shader) {
	shader, err := NewShader(src)
	if err != nil {
		panic(err)
	}
	
	return shader
}

func NewShader(src []byte) (*Shader, error) {
	return ebiten.NewShader(src)
}

