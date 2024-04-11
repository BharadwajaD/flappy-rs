local uv = vim.loop
local api = vim.api

local window = require("window.window")
local bird = require("game.bird")
local tcp = require("tcp")
local Pipes = require("game.pipes")

-- TODO: Don't know wat to do with this
local group = vim.api.nvim_create_augroup("flappy.window", {
    clear = true
})

local M = {}

function START()

    local client = uv.new_tcp()
    M.client = client
    M.win_config = window.create_window({width = 80, height = 24})
    local buffer = M.win_config.buffer

    local pipes = Pipes:new(buffer, 4)

    client:connect("127.0.0.1", 42069, function (err)
        print("client created & connected to server")
    end)
    client:read_start(
        vim.schedule_wrap(
            function (err, chunk)

                print("DEBUG:CHUNK: ", chunk)
                for _, cmd in pairs(tcp.process_packets(chunk)) do
                    local obj = cmd[1]
                    local param1 = tonumber(cmd[2]) + 1 -- fuck 1 indexing
                    local param2 = tonumber(cmd[3]) + 1

                    if obj == "S" and M.win_config ~= nil then
                        M.win_config = window.create_window({width = param1, height = param2})
                        buffer = M.win_config.buffer
                    end

                    if obj == "P" then
                        pipes:place(param1, param2)
                    else
                        bird.update_bird(buffer, param1, param2)
                    end

                    buffer:render()
                end
            end))
end

function END()
    assert(M.client, "client already closed")
    M.client:read_stop()
end
