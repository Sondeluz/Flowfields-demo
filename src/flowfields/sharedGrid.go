package flowfields

import (
	"sync"
    "log"
)

const (
	UNOCCUPIED = -1
)

// A position on the shared map, with
// an optional agent ID who may be going
// to occupy it in its next move, and 
// auxiliary flags for drawing it
type SharedGridPosition struct {
	agentID                int
	mtx                    sync.Mutex // mutex for this position
	objective              bool
	reached                bool
	desired                bool 
}

// Shared grid between all agents
type SharedGrid struct {
	grid [][]SharedGridPosition
}

func newSharedGrid() *SharedGrid {
	grid := make([][]SharedGridPosition, GRID_HEIGHT)
    
	for y := range grid {
		grid[y] = make([]SharedGridPosition, GRID_WIDTH)
        
		for col := range grid[y] {
			grid[y][col].agentID = UNOCCUPIED
		}
	}
	
	log.Println(grid)

	return &SharedGrid{grid}
}

// Attempts to occupy a position in the shared grid, returning true if success,
// false otherwise
func (s *SharedGrid) attemptToOccupy(pos XYPosition, agent Agent) bool {
	if pos.X < 0 || pos.Y < 0 || pos.X > GRID_WIDTH || pos.Y > GRID_HEIGHT {
		panic("Invalid shared grid position")
	}

	s.grid[pos.Y][pos.X].mtx.Lock()
	defer s.grid[pos.Y][pos.X].mtx.Unlock()
    
	if s.grid[pos.Y][pos.X].agentID == UNOCCUPIED {
		s.grid[pos.Y][pos.X].agentID = agent.id // Occupy it
		return true
	} else {
		return false
	}
}

// Sets a position in the shared grid as unoccupied
func (s *SharedGrid) free(pos XYPosition) {
	if pos.X < 0 || pos.Y < 0 || pos.X > GRID_WIDTH || pos.Y > GRID_HEIGHT {
		panic("Invalid shared grid position")
	}

	s.grid[pos.Y][pos.X].mtx.Lock()
	defer s.grid[pos.Y][pos.X].mtx.Unlock()

	s.grid[pos.Y][pos.X].agentID = UNOCCUPIED
}


// Returns true if an agent is in the position
func (s *SharedGridPosition) isOccupied() bool {
	s.mtx.Lock()
	defer s.mtx.Unlock()

    return s.agentID != UNOCCUPIED
}

// Returns true if an agent has this position as an objective
func (s *SharedGridPosition) isObjective() bool {
	s.mtx.Lock()
	defer s.mtx.Unlock()

    return s.objective
}

// Returns true if an agent which had this position
// as an objective has reached it
func (s *SharedGridPosition) isReached() bool {
	s.mtx.Lock()
	defer s.mtx.Unlock()

    return s.reached
}

// Returns true if an agent had this position
// as part of its desired part
func (s *SharedGridPosition) isDesired() bool {
	s.mtx.Lock()
	defer s.mtx.Unlock()

    return s.desired
}


func (s *SharedGrid) setObjective(pos XYPosition) {
	if pos.X < 0 || pos.Y < 0 || pos.X > GRID_WIDTH || pos.Y > GRID_HEIGHT {
		panic("Invalid shared grid position")
	}

	s.grid[pos.Y][pos.X].mtx.Lock()
	defer s.grid[pos.Y][pos.X].mtx.Unlock()

	s.grid[pos.Y][pos.X].objective = true
}

func (s *SharedGrid) setReached(pos XYPosition) {
	if pos.X < 0 || pos.Y < 0 || pos.X > GRID_WIDTH || pos.Y > GRID_HEIGHT {
		panic("Invalid shared grid position")
	}

	s.grid[pos.Y][pos.X].mtx.Lock()
	defer s.grid[pos.Y][pos.X].mtx.Unlock()

	s.grid[pos.Y][pos.X].reached = true
}

func (s *SharedGrid) setDesired(pos XYPosition) {
	if pos.X < 0 || pos.Y < 0 || pos.X > GRID_WIDTH || pos.Y > GRID_HEIGHT {
		panic("Invalid shared grid position")
	}

	s.grid[pos.Y][pos.X].mtx.Lock()
	defer s.grid[pos.Y][pos.X].mtx.Unlock()

	s.grid[pos.Y][pos.X].desired = true
}
