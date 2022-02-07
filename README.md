# Flowfields-demo
A basic demo for flowfield pathfinding

https://user-images.githubusercontent.com/56542714/152784713-702b53ec-58c4-4b47-b3b7-4b25addf4c62.mp4


This is a very limited demo for concurrent flowfield pathfinding I did in a couple of days using Go, loosely based on the guides from https://howtorts.github.io/

Its current limitations are:
- There are no physics (velocity, collisions, etc.): every agent moves exactly one cell per movement, based on its current cell neighbours and other agents which may be occupying these cells.

- There are no predictions and look-aheads, since there is a hard constraint that no two agents can occupy the same cell at any given time.

- The structure is designed for debugging rather than performance and usability: 
  - The solution is inherently concurrent but forces each agent to synchronize with a barrier in order to coordinate their movements and render them.
  - Each agent has one flowfield irregardless of other agents who may have the same objective, and also share a shared flowfield for tracking other agents. These two structures should be joined together in order to not duplicate data.


- The obstacle functionality is implemented, but not tested.


# Usage   

Compile it from the root directory with:
    
    go build main.go

Or simply run it with:
    
    go run main.go

There are some optional parameters:

    -agents int
        number of agents to start. (default 1)
        
    -debug 
        log movements of each agent in their respective files.
        
    -tps int 
        number of movements per second to do. (default 4)
        
Upon running it, a grid will show up. Each agent is represented by a Go's gopher, and the cells will be colour coded such as:
- If the cell is yellow: It has been crossed by an agent as part of its desired path
- If the cell is orange: A cell which was part of the path of an agent has not been crossed, in order to avoid colliding with another agent.
- If the cell is red: The cell is the objective for an agent
- If the cell is green(ish): The cell is an objective reached by an agent. Upon reaching it, the gopher will disappear, allowing other agents to cross it.
