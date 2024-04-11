local api = vim.api
local Buffer = require("window.buffer")

local M = {}

local function create_win_config(dim)
    return {
        relative = "editor",
        title = "flappy-bird",
        style = "minimal",
        border = "double",
        anchor = "NW",
        row = 0,-- (row,col) pos relative to the external window
        col = 0,
        width = dim.width or 80,
        height = dim.height or 24,
    }
end

local function create_win_center_config(dim)

    local win_config = create_win_config(dim)
    local ui = api.nvim_list_uis()[1]

    win_config.row = math.floor((ui.height - win_config.height)/2)
    win_config.col = math.floor((ui.width - win_config.width)/2)

    return win_config
end

function M.create_window(dim)
    local win_config = create_win_center_config(dim)
    local width = win_config.width
    local height = win_config.height

    local buffer = Buffer:new(height, width)
    local window = api.nvim_open_win(buffer.buffer_id, true, win_config)
    api.nvim_win_set_buf(window, buffer.buffer_id)


    return {window = window, dim = {width = width, height = height}, buffer = buffer}
end

return M
