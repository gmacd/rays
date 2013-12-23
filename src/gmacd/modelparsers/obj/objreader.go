package obj

import (
	"fmt"
	"gmacd/core"
	"gmacd/geom"
	"gmacd/modelreaders/obj"
	"path/filepath"
)

// TODO rename?
type ObjReader struct {
	textures core.Textures
	filename string

	model        *geom.Model
	currMaterial *core.Material
}

func NewObjReader(textures core.Textures, filename string) *ObjReader {
	return &ObjReader{
		textures, filename,
		geom.NewModel(), nil}
}

func (reader *ObjReader) Filename() string { return reader.filename }

func (reader *ObjReader) Comment(s string) {}

func (reader *ObjReader) Vertex(components []float64) {
	if len(components) == 3 {
		reader.model.Vertices = append(reader.model.Vertices, core.NewVec3(components[0], components[1], components[2]))
	}
}

func (reader *ObjReader) TextureVertex(components []float64) {
	if len(components) == 2 {
		reader.model.TextureVertices = append(reader.model.TextureVertices, core.NewVec2(components[0], components[1]))
	}
}

func (reader *ObjReader) Normal(components []float64) {
	if len(components) == 3 {
		reader.model.Normals = append(reader.model.Normals, core.NewVec3(components[0], components[1], components[2]))
	}
}

func (reader *ObjReader) Group(names []string) {}

func (reader *ObjReader) UseMaterial(materialName string) {
	reader.currMaterial = reader.model.Materials[materialName]
}

func (reader *ObjReader) Face(vertexIds, textureVertexIds, normalIds []int) {
	hasTextureVertices := len(textureVertexIds) > 0
	hasNormals := len(normalIds) > 0

	numVerts := len(reader.model.Vertices)
	numTexVerts := len(reader.model.TextureVertices)
	numNormals := len(reader.model.Normals)

	t := geom.NewTri()
	for i := 0; i < len(vertexIds); i++ {

		vertId := vertexIds[i]
		if vertId < 0 {
			vertId += numVerts // Offset from last
		}
		t.Vertices = append(t.Vertices, &reader.model.Vertices[vertId])

		if hasTextureVertices {
			texVertId := textureVertexIds[i]
			if texVertId < 0 {
				texVertId += numTexVerts // Offset from last
			}
			t.TextureVertices = append(t.TextureVertices, &reader.model.TextureVertices[texVertId])
		}

		if hasNormals {
			normalId := textureVertexIds[i]
			if normalId < 0 {
				normalId += numNormals // Offset from last
			}
			t.Normals = append(t.Normals, &reader.model.Normals[normalId])
		}
	}

	reader.model.Triangles = append(reader.model.Triangles, *t)
}

func (reader *ObjReader) MaterialLibrary(filename string) {
	actualFilename := filepath.Join(filepath.Dir(reader.Filename()), filename)

	mtlReader := NewObjMtlReader(reader.textures, reader.model, actualFilename)
	if err := obj.ReadMaterial(mtlReader); err != nil {
		fmt.Printf("**  Error loading %v: %v\n", actualFilename, err.Error())
		return
	}
}
