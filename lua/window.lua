local api = vim.api

local M = {}

local function create_win_config()
    return {
        relative = "editor",
        title = "flappy-bird",
        style = "minimal",
        border = "double",
        anchor = "NW",
        row = 0,-- (row,col) pos relative to the external window
        col = 0,
        width = 80,
        height = 24
    }
end

function M.create_win_center_config()

    local win_config = create_win_config()
    local ui = api.nvim_list_uis()[1]

    win_config.row = math.floor((ui.height - win_config.height)/2)
    win_config.col = math.floor((ui.width - win_config.width)/2)

    return win_config
end

function M.create_window()
    local win_config = M.create_win_center_config()
    local buf = api.nvim_create_buf(false, true)
    local window = api.nvim_open_win(buf, true, win_config)

    return window
end


return M
