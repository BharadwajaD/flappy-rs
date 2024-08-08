import { Game } from "./game.js"

const url = "ws://127.0.0.1:42069/game-ws"

const ws = new WebSocket(url)

/**
 *@type {Document}*/
const doc = $(document)

const game_config = {buffer_config: {rows: 24, cols: 80}, pipes_count: 5,
    win_config: { height: doc.height(), width: doc.width() }}

const game = new Game(game_config)

ws.addEventListener("open", (_event) => {
    console.log(`ws connected to ${url}`)
})

ws.addEventListener("message", async (event) => {

    const msg = event.data
    const sep = ':'
    const splits = msg.split(sep)
    const cmd = splits[0]
    const params = splits.slice(1)

    await game.Cmd(cmd, ... params)

})

$("#canvas").addEventListener("keydown", (event) => {
    console.log("Sent: ", event.key)
    ws.send(event.key)
})
