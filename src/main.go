package main

import (
    "flowfieldsDemo/flowfields"
    "flag"
)

func main() {
    var agents int
    var tps int
    var debug bool
    
    flag.IntVar(&agents, "agents", 1, "number of agents to start.")
    flag.IntVar(&tps, "tps", 4, "number of movements per second to do.")
    flag.BoolVar(&debug, "debug", false, "log movements of each agent in their respective files.")
    flag.Parse()
    
	flowfields.InitDemo(agents, tps, debug)
}
