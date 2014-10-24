package core

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

type Textures map[string]image.Image

func NewTextures() Textures {
	return Textures{}
}

func (textures Textures) Fetch(path string) (texture image.Image, err error) {
	if texture, ok := textures[path]; ok {
		return texture, nil
	}

	var reader *os.File
	if reader, err = os.Open(path); err != nil {
		return nil, err
	}
	defer reader.Close()

	if texture, _, err = image.Decode(reader); err != nil {
		return nil, err
	}

	textures[path] = texture
	return texture, nil
}
