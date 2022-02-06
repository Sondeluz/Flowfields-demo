package flowfields

import (
	"github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
    "image/color"
)

// Game implements ebiten.Game interface.
type Game struct{
    gopher  *ebiten.Image
    sg      *SharedGrid
    tick    chan interface{}
    barrier chan interface{}
    agents  int
}

func (g *Game) drawSharedGrid (sg *SharedGrid, screen *ebiten.Image) {
    screen.Clear()
    //screen.Fill(color.RGBA{102, 99, 169, 1})
    screen.Fill(color.RGBA{102, 99, 169, 0xff})
    op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.111, 0.111) // 77x52, approx.
    
    for y := range sg.grid {
		for x, pos := range sg.grid[y] {
            ebitenutil.DrawLine(screen, float64(x*77), float64(y*52), float64(x*77)+77, float64(y*52), color.White)
            ebitenutil.DrawLine(screen, float64(x*77), float64(y*52), float64(x*77), float64(y*52)+52, color.White)
            
            if pos.isObjective() && pos.isReached(){
                ebitenutil.DrawRect(screen, float64(x*77), float64(y*52), 77, 52, color.RGBA{102, 252, 169, 0xff})
            } else if pos.isObjective() && !pos.isReached(){
                ebitenutil.DrawRect(screen, float64(x*77), float64(y*52), 77, 52, color.RGBA{255, 60, 54, 0xff})
            }
            
            if pos.isOccupied() {
                op.GeoM.Translate(float64(x*77), float64(y*52)) // right, down
                screen.DrawImage(g.gopher, op)
            }
		}
	}
}


func (g *Game) Init(sg *SharedGrid, tick chan interface{}, barrier chan interface{}, agents int, tps int) {
    var err error
    g.gopher, _, err = ebitenutil.NewImageFromFile("gopher.png")
    if err != nil {
		log.Fatal(err)
	}
	
	g.sg = sg
	g.tick = tick
	g.barrier = barrier
	g.agents = agents
	
	ebiten.SetMaxTPS(tps)
}


// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {	
	for i := 0; i < g.agents; i++ {
        g.tick <- nil
    }
    
    for i := 0; i < g.agents; i++ {
        <- g.barrier
    }
    
    return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
    g.drawSharedGrid(g.sg, screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1540, 1040
}
