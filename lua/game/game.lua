local uv = vim.loop
local api = vim.api

local window = require("window.window")
local Pipes = require("game.pipes")
local bird = require("game.bird")
local tcp = require("tcp")

local Game = {}
Game.__index = Game

function Game:new(config)
    -- tcp client
    -- buffer and window
    -- game objects

    self.tcp_client = uv.new_tcp()
    local win_config = window.create_window(config.dim)
    self.buffer = win_config.buffer
    self.dim = win_config.dim
    self.pipes = Pipes:new(self.buffer, 4)
    return self
end

function Game:cmd(cmd, params)
    if cmd == "P" then
        self.pipes:place(params[1], params[2])
    else
        bird.update_bird(self.buffer, params[1], params[2])
    end
    self.buffer:render()

end

function Game:Start()

    self.tcp_client:connect("127.0.0.1", 42069, function (err)
        print("client created & connected to server")
    end)

    self.tcp_client:read_start(
        vim.schedule_wrap(
            function (err, chunk)

                print("DEBUG:CHUNK: ", chunk)
                for parse in tcp.process_tcp_packets(chunk) do
                    local cmd = parse.cmd
                    local params = parse.params
                    print("DEBUG:PARSE: ", vim.inspect({cmd, params}))

                    self:cmd(cmd, params)
                end
            end))
end

function Game:End()
    assert(self.tcp_client, "client already closed")
    self.tcp_client:read_stop()
end

return Game
