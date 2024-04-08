local api = vim.api


local M = {}

 --@param width
 --@param x_loc
 --@param char
 --@param height
-- creates a row
local function set_row(width, x_loc, char)
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
 function M.add_pipe(buffer, x_loc, height)

    local dim = buffer.dim
    assert(height < math.floor(dim.height)/2, "too long pipe")

    buffer:place_vline(x_loc, height, "x")
    buffer:place_vline(x_loc, height, "x", true)
 end

function M.remove_pipe(buffer, x_loc, height)
    buffer:remove_vline(x_loc, height)
    buffer:remove_vline(x_loc, height, true)
end

 --@param dim
 --@param buffer_id
 --@param x_loc
 --@param y_loc
 --Draw bird at the given (x, y) location
 function M.update_bird(buffer , x_loc, y_loc)
    if M.prev_bird == nil then
        M.prev_bird = {x_loc= x_loc, y_loc= y_loc}
    end
    local prev = M.prev_bird

    buffer:remove_point(prev.x_loc, prev.y_loc)
    buffer:place_point(x_loc, y_loc, "b")

    M.prev_bird = {x_loc= x_loc, y_loc= y_loc}

 end

return M
