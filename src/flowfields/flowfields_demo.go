package flowfields

import (
	"log"
    "github.com/hajimehoshi/ebiten/v2"
    "os"
    "strconv"
)

const (
    MAX_MOVES_DEMO = 1000
)


func InitDemo(agents int, tps int, debug bool) {
    sg := newSharedGrid()
	//f := newRandomFlowFieldWithoutObstacles()
	//obj := f.getObjective()
	tick := make(chan interface{}, 2)
    barrier := make(chan interface{}, 2)
    
    // Run each agent in a Goroutine
	for y := 0; y < agents; y++ {
        go func(offset int) {
            var logger *log.Logger
            
            if debug {
                f, err := os.OpenFile(strconv.Itoa(offset)+"_log", os.O_RDWR | os.O_CREATE, 0666)
                if err != nil {
                    log.Fatalf("error opening file for debugging: %v", err)
                }
                defer f.Close()
                
                logger = log.New(f, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
            }
            
            
            flowfield := newRandomFlowFieldWithoutObstacles()
            a := NewAgent(offset, XYPosition{X: 0, Y: offset}, flowfield, sg)
            moves := 0
            
            if debug {
                logger.Println("Objective:", a.flowfield.objective)
            }
            
            for !a.isInObjective() && moves < MAX_MOVES_DEMO {
                if debug {
                    logger.Println("move of agent",offset,"to",a.position)
                }
                <- tick 
                a.move()
                barrier <- nil
            }
            
            if debug {
                logger.Println("FINAL move of agent", offset, "to", a.position)
            }
                
            a.sharedGrid.free(a.position)                   // Free the objective
            a.sharedGrid.setReached(a.flowfield.objective)  // And mark it as reached
            
            for {
                <- tick 
                barrier <- nil
            }
        }(y) 
    }
	
	
	game := &Game{}
    // Specify the window size as you like. Here, a doubled size is specified.
    ebiten.SetWindowSize(1541, 1059)
    //ebiten.SetWindowSize(1000, 1000)
    ebiten.SetWindowTitle("Flowfields demo")
    game.Init(sg, tick, barrier, agents, tps)
    // Call ebiten.RunGame to start your game loop.
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
} 
