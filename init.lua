local uv = vim.loop
local api = vim.api

local window = require("window.window")
local render = require("game.render")
local tcp = require("tcp")
local queue = require("queue")

-- TODO: Don't know wat to do with this
local group = vim.api.nvim_create_augroup("flappy.window", {
    clear = true
})

local M = {}

function START()

    local client = uv.new_tcp()
    M.client = client
    local pqueue = queue:new(4)
    local buffer


    client:connect("127.0.0.1", 42069, function (err)
        print("client created & connected to server")
    end)
    client:read_start(
        vim.schedule_wrap(
            function (err, chunk)

                print("DEBUG:CHUNK: ", chunk)
                for _, cmd in pairs(tcp.process_packets(chunk)) do
                    local obj = cmd[1]
                    local param1 = tonumber(cmd[2])
                    local param2 = tonumber(cmd[3])

                    if obj == "S" and M.win_config ~= nil then
                        M.win_config = window.create_window({width = param1, height = param2})
                        buffer = M.win_config.buffer
                    end

                    if M.win_config == nil then
                        M.win_config = window.create_window({width = 80, height = 24})
                        buffer = M.win_config.buffer
                    end

                    if obj == "P" then
                        local rpipe = pqueue:enqueue({param1 = param1, param2 = param2})
                        if rpipe ~= nil then
                            render.remove_pipe(M.win_config.buffer, rpipe.param1, rpipe.param2)
                        end
                        render.add_pipe(M.win_config.buffer,  param1, param2)
                    else
                        render.update_bird(M.win_config.buffer, param1, param2)
                    end

                    buffer:render()
                end
            end))
end

function END()
    assert(M.client, "client already closed")
    M.client:read_stop()
end
