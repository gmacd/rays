package geom

import (
	"github.com/gmacd/rays/core"
)

// TODO Lights as separate entities?
// TODO Light details as part of material?

type Light interface {
	IsLight() bool
	SetIsLight(isLight bool)
	LightCentre() core.Vec3
}

type LightData struct {
	isLight bool
	pos     core.Vec3
}

func NewLightData(pos core.Vec3) *LightData {
	return &LightData{false, pos}
}

func NewLightDataNone() *LightData {
	return &LightData{false, core.NewVec3Zero()}
}

func (light *LightData) IsLight() bool {
	return light.isLight
}

func (light *LightData) SetIsLight(isLight bool) {
	light.isLight = isLight
}

func (light *LightData) LightCentre() core.Vec3 {
	return light.pos
}
