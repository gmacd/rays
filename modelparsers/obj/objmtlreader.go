package obj

import (
	"gmacd/core"
	"gmacd/geom"
)

type ObjMtlReader struct {
	textures core.Textures
	filename string

	model        *geom.Model
	currMaterial *core.Material
}

func NewObjMtlReader(textures core.Textures, model *geom.Model, filename string) *ObjMtlReader {
	return &ObjMtlReader{
		textures, filename, model, nil}
}

func (reader *ObjMtlReader) Filename() string { return reader.filename }

func (reader *ObjMtlReader) Comment(s string) {}

func (reader *ObjMtlReader) NewMaterial(materialName string) {
	material := core.NewMaterialBlank()
	material.Name = materialName

	reader.currMaterial = material
	reader.model.Materials[materialName] = material
}

func (reader *ObjMtlReader) Specular(value float64) {
	reader.currMaterial.Specular = value
}

func (reader *ObjMtlReader) IndexOfRefraction(value float64) {
	reader.currMaterial.RefractiveIndex = value
}

func (reader *ObjMtlReader) DiffuseColour(values []float64) {
	reader.currMaterial.Colour = core.NewColourRGB(values[0], values[1], values[2])
}

func (reader *ObjMtlReader) Transparency(value float64)      {}
func (reader *ObjMtlReader) AmbientColour(values []float64)  {}
func (reader *ObjMtlReader) SpecularColour(values []float64) {}
func (reader *ObjMtlReader) EmissiveColour(values []float64) {}
func (reader *ObjMtlReader) DiffuseTexture(filename string)  {}
