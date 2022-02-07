package flowfields

import (
	"github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
    "image/color"
    "embed"
)

const (
    GOPHER_SCALE = 0.111 // Brings it down to 77x52, approx.
    GOPHER_X_OFFSET = 77
    GOPHER_Y_OFFSET = 52
) 

// Game implements ebiten.Game interface.
type Game struct{
    gopher  *ebiten.Image
    sg      *SharedGrid
    tick    chan interface{}
    barrier chan interface{}
    agents  int
}

//go:embed gopher.png
var gopher embed.FS

func (g *Game) drawSharedGrid (sg *SharedGrid, screen *ebiten.Image) {
    screen.Clear()
    screen.Fill(color.RGBA{102, 99, 169, 0xff})
    
    
    for y := range sg.grid {
		for x, pos := range sg.grid[y] {
            // Draw grid lines
            ebitenutil.DrawLine(screen, float64(x*GOPHER_X_OFFSET), float64(y*GOPHER_Y_OFFSET), 
                                float64(x*GOPHER_X_OFFSET)+GOPHER_X_OFFSET, float64(y*GOPHER_Y_OFFSET), color.White)
            ebitenutil.DrawLine(screen, float64(x*GOPHER_X_OFFSET), float64(y*GOPHER_Y_OFFSET), 
                                float64(x*GOPHER_X_OFFSET), float64(y*GOPHER_Y_OFFSET)+GOPHER_Y_OFFSET, color.White)
            
            // Draw positions with different colour depending on whether they are an objective or desired position, and whether they have been reached or not
            if pos.isObjective() && pos.isReached(){
                ebitenutil.DrawRect(screen, float64(x*GOPHER_X_OFFSET), float64(y*GOPHER_Y_OFFSET), 
                                    GOPHER_X_OFFSET, GOPHER_Y_OFFSET, color.RGBA{102, 52, 169, 0xff})
            } else if pos.isObjective() && !pos.isReached(){
                ebitenutil.DrawRect(screen, float64(x*GOPHER_X_OFFSET), float64(y*GOPHER_Y_OFFSET), 
                                    GOPHER_X_OFFSET, GOPHER_Y_OFFSET, color.RGBA{255, 60, 54, 0xff})
            } else if pos.isDesired() && pos.isReached() {
                ebitenutil.DrawRect(screen, float64(x*GOPHER_X_OFFSET), float64(y*GOPHER_Y_OFFSET), 
                                    GOPHER_X_OFFSET, GOPHER_Y_OFFSET, color.RGBA{255, 255, 45, 0xff})
            } else if pos.isDesired() && !pos.isReached() {
                ebitenutil.DrawRect(screen, float64(x*GOPHER_X_OFFSET), float64(y*GOPHER_Y_OFFSET), 
                                    GOPHER_X_OFFSET, GOPHER_Y_OFFSET, color.RGBA{60, 94, 66, 0xff})
            } else if !pos.isDesired() && pos.isReached() {
                ebitenutil.DrawRect(screen, float64(x*GOPHER_X_OFFSET), float64(y*GOPHER_Y_OFFSET), 
                                    GOPHER_X_OFFSET, GOPHER_Y_OFFSET, color.RGBA{255, 150, 45, 0xff})
            } // else: uncoloured
            
            // Draw a gopher in occupied positions
            if pos.isOccupied() {
                op := &ebiten.DrawImageOptions{}
                op.GeoM.Scale(GOPHER_SCALE, GOPHER_SCALE) 
                op.GeoM.Translate(float64(x*GOPHER_X_OFFSET), float64(y*GOPHER_Y_OFFSET)) // right, down
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
