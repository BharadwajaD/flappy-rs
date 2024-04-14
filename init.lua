local uv = vim.loop
local api = vim.api

local keypress = require("window.keypress")
local Game = require("game.game")

-- TODO: Don't know wat to do with this
local group = vim.api.nvim_create_augroup("flappy.window", {
    clear = true
})

local M = {}

function KeyPressStart()
    local win_config = window.create_window({width = 80, height = 24})
    local buffer_id = win_config.buffer.buffer_id
    keypress.KeyPressEvent(buffer_id)
end

function START()

    local dim = {width = 80, height = 24}
    local game = Game:new({dim = dim})
    M.game = game
    game:Start()

end

function END()
    M.game:End()
end
