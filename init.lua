local uv = vim.loop

local window = require("window")
local render = require("game.render")
local tcp = require("tcp")

local group = vim.api.nvim_create_augroup("flappy.window", {
    clear = true
})

local M = {}

function START()

    local client = uv.new_tcp()
    M.client = client

    client:connect("127.0.0.1", 42069, function (err)
        print("client created & connected to server")
    end)

    client:read_start(vim.schedule_wrap(function (err, chunk)

        client:write(chunk)

        local obj = tcp.parse(chunk)
        assert(obj, "string parsing went wrong")

        local param1 = tonumber(obj[2])
        local param2 = tonumber(obj[3])

        if obj[1] == "S" then
            if  M.win_config ~= nil then
                M.win_config = window.create_window({width = param1, height = param2})
            end
        elseif obj[1] == "P" then
            if M.win_config == nil then
                M.win_config = window.create_window({width = 80, height = 24})
            end
            render.draw_pipe(M.win_config.dim, M.win_config.buffer_id,  param1, param2)
        else
            if M.win_config == nil then
                M.win_config = window.create_window({width = 80, height = 24})
            end
            render.draw_bird(M.win_config.dim, M.win_config.buffer_id, param1, param2)
        end
    end))
end

function END()
    assert(M.client, "client already closed")
    M.client:read_stop()
end

