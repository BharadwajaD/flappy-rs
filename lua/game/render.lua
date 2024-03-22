local api = vim.api

-- TODO: Build classes

local M = {}

local function set_row(width, x_loc)
    local line = ""
    for i = 0, width-1, 1 do
        if i == x_loc then
            line = line .. "x"
        else
            line = line .. " "
        end
    end
    return line
end

 --@parameter dim
 --@parameter buffer_id
 --@parameter x_loc
 --@parameter height
 --Draw pipe of given height at x_loc
 function M.draw_pipe(dim, buffer_id, x_loc, height)

    assert(height < math.floor(dim.height)/2, "too long pipe")

    local pipe = {}
    local line = set_row(dim.width, x_loc)
    for _ = 1, height, 1 do
        table.insert(pipe, line)
    end
    if api.nvim_buf_is_loaded(buffer_id) and api.nvim_buf_is_valid(buffer_id) then
        api.nvim_buf_set_lines(buffer_id, 0, height+1,0, pipe)
        api.nvim_buf_set_lines(buffer_id, -height, -1,0, pipe)
    end
 end

 --@parameter dim
 --@parameter buffer_id
 --@parameter x_loc
 --@parameter y_loc
 --Draw bird at the given (x, y) location
 function M.draw_bird(dim, buffer_id , x_loc, y_loc)
    local line = set_row(dim.width, x_loc)
    if api.nvim_buf_is_loaded(buffer_id) and api.nvim_buf_is_valid(buffer_id) then
        api.nvim_buf_set_lines(buffer_id,  y_loc, y_loc+1,0, {line})
    end
 end

return M
