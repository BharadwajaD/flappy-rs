local api = vim.api


local M = {}

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

    -- pcall to handle errors
    buffer:RemovePoint(prev.x_loc, prev.y_loc)
    buffer:PlacePoint(x_loc, y_loc, "b")

    M.prev_bird = {x_loc= x_loc, y_loc= y_loc}

 end

return M
