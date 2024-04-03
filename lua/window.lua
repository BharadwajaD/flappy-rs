local api = vim.api

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

function M.create_win_center_config(dim)

    local win_config = create_win_config(dim)
    local ui = api.nvim_list_uis()[1]

    win_config.row = math.floor((ui.height - win_config.height)/2)
    win_config.col = math.floor((ui.width - win_config.width)/2)

    return win_config
end

function M.create_window(dim)
    local win_config = M.create_win_center_config(dim)
    local width = win_config.width
    local height = win_config.height

    local empty = {}
    for _ = 1, height, 1 do
        table.insert(empty, "")
    end

    local buffer_id = api.nvim_create_buf(false, true)
    local window = api.nvim_open_win(buffer_id, true, win_config)
    api.nvim_win_set_buf(window, buffer_id)

    -- Fill buffer with empty lines
    api.nvim_buf_set_lines(buffer_id, 0, 1, 0, empty)

    return {dim = {width = width, height = height}, buffer_id = buffer_id}
end


return M
