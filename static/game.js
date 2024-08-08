import { BufferWindow } from "./buffer.js"
import { Pipe } from "./pipe.js"

export class Game{
    constructor(game_config){

        const buffer_config = game_config.buffer_config

        this.bwindow = new BufferWindow(buffer_config)
        this.pipes = new Pipe(game_config.pipes_count, this.bwindow)

        this.prevb = []
    }

    async Cmd(_cmd, ... params){

        if (_cmd == "B"){
            if (this.prevb.length != 0) {
                this.bwindow.RemovePoint(this.prevb[0], this.prevb[1])
            }
            this.bwindow.PlacePoint(params[0], params[1], 'b')
            this.prevb = [params[0], params[1]]

        }else if(_cmd == "P"){
            this.pipes.Place(params[0], params[1])
        }

        await this.bwindow.Render()

    }
}
