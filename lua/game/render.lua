local api = vim.api


local M = {}

 --@param width
 --@param x_loc
 --@param char
 --@param height
-- creates a row
local function set_row(width, x_loc, char)
    print("DEBUG:SET_ROW:", width, x_loc)
    local line = ""
    for i = 0, width-1, 1 do
        if i == x_loc then
            line = line .. char
        else
            line = line .. " "
        end
    end
    return line
end

 --@param dim
 --@param buffer_id
 --@param x_loc
 --@param height
 --Draw pipe of given height at x_loc
 function M.draw_pipe(dim, buffer_id, x_loc, height)

    print(height, math.floor(dim.height)/2)
    assert(height < math.floor(dim.height)/2, "too long pipe")

    local pipe = {}
    local line = set_row(dim.width, x_loc, "x")
    for _ = 1, height, 1 do
        table.insert(pipe, line)
    end

    api.nvim_buf_set_lines(buffer_id, 0, height,1, pipe)
    api.nvim_buf_set_lines(buffer_id, dim.height-height, dim.height,1, pipe)
 end

 --@param dim
 --@param buffer_id
 --@param x_loc
 --@param y_loc
 --Draw bird at the given (x, y) location
 function M.draw_bird(dim, buffer_id , x_loc, y_loc)
    local line = set_row(dim.width, x_loc, "b")
    api.nvim_buf_set_lines(buffer_id,  y_loc, y_loc+1,0, {line})
 end

return M
