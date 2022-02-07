# Flowfields-demo
A basic demo for flowfield pathfinding

This is a very limited demo for concurrent flowfield pathfinding I did in a couple of days using Go, loosely based on the guides from https://howtorts.github.io/

Its current limitations are:
- There are no physics (velocity, collisions, etc.): every agent moves exactly one cell per movement, based on its current cell neighbours and other agents which may be occupying these cells.

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
