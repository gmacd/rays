package core

import ()

type Material struct {
	Colour          ColourRGB
	Reflection      float64
	Refraction      float64
	RefractiveIndex float64
	Diffuse         float64
	Specular        float64
}

func NewMaterialBlank() *Material {
	return &Material{NewColourRGB(0.2, 0.2, 0.2), 0.0, 1.0, 1.5, 0.2, 0.8}
}

func NewMaterial(colour ColourRGB, reflection, refraction, refractiveIndex, diffuse, specular float64) *Material {
	return &Material{colour, reflection, refraction, refractiveIndex, diffuse, specular}
}
