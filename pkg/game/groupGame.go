package game

import (
	"github.com/bharadwajaD/flappy-go/pkg/spmc"
)

type GroupGame struct {
    group_id int  //TODO: create (group_id, GroupGame) map
    pipes *spmc.Spmc[Pipe]
    scores map[int]int //game_id, score of each game in the group
}

func NewGroupGame(opts *GameOpts) GroupGame {
    pipes := spmc.NewSpmc[Pipe]()

    pipes_gen := GenPipes(opts)
	go func() {
		for {
            pipe :=  pipes_gen()
            pipes.Broadcast(pipe)
		}
	}()

    return GroupGame{
        pipes: &pipes,
    }
}

