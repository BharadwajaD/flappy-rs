local uv = vim.loop
local api = vim.api

local window = require("window.window")
local ascii = require("window.ascii")
local Pipes = require("game.pipes")
local bird = require("game.bird")
local tcp = require("tcp")

local Game = {}
Game.__index = Game

function Game:New(config)
    -- tcp client
    -- buffer and window
    -- game objects

    self.tcp_client = uv.new_tcp()
    local win_config = window.create_window(config.dim)
    self.window = win_config.window
    self.buffer = win_config.buffer
    self.dim = win_config.dim
    self.pipes = Pipes:new(self.buffer, 4)
    self.running = false
    return self
end

function Game:cmd(cmd, params)
    print("DEBUG: ", cmd, vim.inspect(params))
    if cmd == "P" then
        self.pipes:place(params[1], params[2])
    elseif cmd == "B" then
        bird.update_bird(self.buffer, params[1], params[2])
    elseif cmd == "E" then
        ascii.draw_score(self.buffer, self.dim, params[1])
        self:End()
    else
        print("Not yet coded ... ")
    end

    self.buffer:render()
end

function Game:Start()

    self.running = true
    self.tcp_client:connect("127.0.0.1", 42069, function (err)
        print("client created & connected to server")
    end)

    self.tcp_client:read_start(
        vim.schedule_wrap(
            function (err, chunk)
                for parse in tcp.process_tcp_packets(chunk) do
                    local cmd = parse.cmd
                    local params = parse.params

                    self:cmd(cmd, params)
                end
            end))

    local function write(key)
        self.tcp_client:write("K:"..key.."?")
    end

    vim.keymap.set('n', 'k', function () write("k") end , { silent = true, buffer = self.buffer.buffer_id })
    vim.keymap.set('n', 'q', function ()
        -- close window and buffer
        if self.running then
            self:End()
        end
        window.close_window(self.window)
        self.buffer:close()
    end , { silent = true, buffer = self.buffer.buffer_id })
end

function Game:End()
    assert(self.tcp_client, "client already closed")
    self.tcp_client:read_stop()
    self.tcp_client:close()
    self.running = false
end

return Game
