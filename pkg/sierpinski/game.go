package sierpinski

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// speedFactor is how many pixels to set per tick.
var speedFactor = 10

var colorMappings = map[int]color.Color{
	0: color.RGBA{0x00, 0x00, 0xFF, 0xFF},
	1: color.RGBA{0x00, 0xFF, 0x00, 0xFF},
	2: color.RGBA{0xFF, 0x00, 0x00, 0xFF},
}

// pt represents a fixed point determined by x,y co-ordinates on a
// cartesian plane.
type pt struct {
	x, y int
}

// existsInTriangle returns true if p is found within the triangle formed by
// v1, v2, and v3 else false.
// https://stackoverflow.com/questions/2049582/how-to-determine-if-a-point-is-in-a-2d-triangle
func existsInTriangle(p, v1, v2, v3 pt) bool {
	sign := func(p1, p2, p3 pt) int { return (p1.x-p3.x)*(p2.y-p3.y) - (p2.x-p3.x)*(p1.y-p3.y) }

	d1 := sign(p, v1, v2)
	d2 := sign(p, v2, v3)
	d3 := sign(p, v3, v1)

	has_neg := (d1 < 0) || (d2 < 0) || (d3 < 0)
	has_pos := (d1 > 0) || (d2 > 0) || (d3 > 0)

	return !(has_neg && has_pos)
}

// midwayPoint returns a pt that exists halfway between p1, and p2
func midwayPoint(p1, p2 pt) pt {
	return pt{(p1.x + p2.x) / 2, (p1.y + p2.y) / 2}
}

type Game struct {
	width, height   int
	lastPosition    pt
	initialTriangle [3]pt
	image           *ebiten.Image
}

func NewGame(width, height int) *Game {
	// Set the initial triangle to start with.
	pts := [3]pt{
		{x: width / 2, y: 0},
		{x: 1, y: height - 1},
		{x: width - 1, y: height - 1},
	}

	// Find an initial point within the triangle to start with.
	var position pt
	found := false
	for !found {
		p := pt{x: rand.Intn(width), y: rand.Intn(height)}
		found = existsInTriangle(p, pts[0], pts[1], pts[2])
		position = p
	}
	return &Game{
		width:           width,
		height:          height,
		image:           ebiten.NewImage(width, height),
		initialTriangle: pts,
		lastPosition:    position,
	}
}

func (g *Game) findNearestPoint(p1 pt) pt {
	var smallestDistance float64
	var smallestPoint pt
	for _, p := range g.initialTriangle {
		x := float64(p1.x - p.x)
		y := float64(p1.y - p.y)
		d := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
		if d < smallestDistance {
			smallestDistance = d
			smallestPoint = p
		}
	}
	return smallestPoint
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		if speedFactor <= 0 {
			speedFactor = 0
		} else {
			speedFactor--
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		speedFactor++
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.image = ebiten.NewImage(g.width, g.height)
	}
	for i := 0; i < speedFactor; i++ {
		r := rand.Intn(3)
		m := midwayPoint(g.lastPosition, g.initialTriangle[r])
		g.image.Set(m.x, m.y, colorMappings[r])
		g.lastPosition = m
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.image, &ebiten.DrawImageOptions{})
	t := fmt.Sprintf("TPS: %f\nSpeed: %d (press <- or -> arrow to change)\n", ebiten.CurrentTPS(), speedFactor)
	ebitenutil.DebugPrint(screen, t)
}

func (g *Game) Layout(_, _ int) (int, int) { return g.width, g.height }
