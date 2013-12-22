package obj

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ObjParser interface {
	Filename() string

	Comment(s string)
	Vertex(components []float64)
	TextureVertex(components []float64)
	Normal(components []float64)
	Group(names []string)
	Face(vertexIds, textureVertexIds, normalIds []int)
	MaterialLibrary(filename string)
	UseMaterial(materialName string)
}

func Read(objParser ObjParser) error {
	file, err := os.Open(objParser.Filename())
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineTokens := strings.Fields(scanner.Text())
		if len(lineTokens) > 0 {
			parseObjLine(objParser, lineTokens)
		}
	}

	fmt.Printf(" Loaded: %v\n", objParser.Filename())
	return nil
}

// Parse a line by dispatching to the appropriate function in parser
// Assumes at least one token.
func parseObjLine(parser ObjParser, tokens []string) {
	switch tokens[0] {
	case "#":
		parser.Comment(strings.Join(tokens[1:], " "))
	case "v":
		parser.Vertex(parseFloats(tokens[1:]))
	case "vt":
		parser.TextureVertex(parseFloats(tokens[1:]))
	case "vn":
		parser.Normal(parseFloats(tokens[1:]))
	case "g":
		parser.Group(tokens[1:])
	case "f":
		parser.Face(parseFaces(tokens[1:]))
	case "mtllib":
		parser.MaterialLibrary(tokens[1])
	case "usemtl":
		parser.UseMaterial(tokens[1])

	default:
		fmt.Printf("Unknown obj parameter: %v\n", tokens[0])
	}
}

func parseFaces(tokens []string) (vertexIds, textureVertexIds, normalIds []int) {
	vertexIds = make([]int, len(tokens))
	textureVertexIds = make([]int, len(tokens))
	normalIds = make([]int, len(tokens))

	for i := 0; i < len(tokens); i++ {
		// Each token is split by '/'
		elementTokens := strings.Split(tokens[i], "/")
		// Ugh
		for j := 0; j < len(elementTokens); j++ {
			id, _ := strconv.ParseInt(elementTokens[j], 10, 32)
			switch j {
			case 0:
				vertexIds[i] = int(id)
			case 1:
				textureVertexIds[i] = int(id)
			case 2:
				normalIds[i] = int(id)
			}
		}
	}

	return vertexIds, textureVertexIds, normalIds
}

func parseFloat(token string) float64 {
	value, _ := strconv.ParseFloat(token, 64)
	return value
}

func parseFloats(tokens []string) []float64 {
	floats := make([]float64, len(tokens))
	for i := 0; i < len(tokens); i++ {
		floats[i], _ = strconv.ParseFloat(tokens[i], 64)
	}
	return floats
}

type ObjMaterialParser interface {
	Filename() string

	Comment(s string)
	NewMaterial(materialName string)
	Specular(value float64)
	IndexOfRefraction(value float64)
	Transparency(value float64)
	AmbientColour(values []float64)
	DiffuseColour(values []float64)
	SpecularColour(values []float64)
	EmissiveColour(values []float64)
	DiffuseTexture(filename string)
}

func ReadMaterial(parser ObjMaterialParser) error {
	file, err := os.Open(parser.Filename())
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineTokens := strings.Fields(strings.TrimSpace(scanner.Text()))
		if len(lineTokens) > 0 {
			parseMaterialLine(parser, lineTokens)
		}
	}

	fmt.Printf(" Loaded: %v\n", parser.Filename())
	return nil
}

// Parse a line by dispatching to the appropriate function in parser
// Assumes at least one token.
func parseMaterialLine(parser ObjMaterialParser, tokens []string) {
	switch tokens[0] {
	case "#":
		parser.Comment(strings.Join(tokens[1:], " "))
	case "newmtl":
		parser.NewMaterial(tokens[1])
	case "Ns":
		parser.Specular(parseFloat(tokens[1]))
	case "Ni":
		parser.IndexOfRefraction(parseFloat(tokens[1]))
	case "Tr":
		parser.Transparency(parseFloat(tokens[1]))
	case "Ka":
		parser.AmbientColour(parseFloats(tokens[1:]))
	case "Kd":
		parser.DiffuseColour(parseFloats(tokens[1:]))
	case "Ks":
		parser.SpecularColour(parseFloats(tokens[1:]))
	case "Ke":
		parser.EmissiveColour(parseFloats(tokens[1:]))
	case "map_Kd":
		parser.DiffuseTexture(tokens[1])

	default:
		fmt.Printf("Unknown material parameter: %v\n", tokens[0])
	}
}
