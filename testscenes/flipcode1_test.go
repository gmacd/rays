package testscenes

import (
	"testing"

	"github.com/gmacd/rays/core"
	"github.com/gmacd/rays/renderer"
)

func BenchmarkFlipcode1(b *testing.B) {
	canvas := core.NewCanvas(800, 600)
	scene := CreateFlipcode1Scene()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		renderer.Render(scene, canvas)
	}
}
