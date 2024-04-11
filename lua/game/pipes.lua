local queue = require("queue")

local Pipes = {}
Pipes.__index = Pipes

function Pipes:new(buffer, capacity)
    self.pqueue = queue:new(capacity)
    self.buffer = buffer
    return self
end

 --@param dim
 --@param buffer_id
 --@param x_loc
 --@param height
 --Draw pipe of given height at x_loc
function Pipes:add_pipe(x_loc, height)

    local dim = self.buffer.dim
    assert(height < math.floor(dim.height)/2, "too long pipe")

    self.buffer:place_vline(x_loc, height, "x")
    self.buffer:place_vline(x_loc, height, "x", true)
 end

function Pipes:remove_pipe( x_loc, height)
    self.buffer:remove_vline(x_loc, height)
    self.buffer:remove_vline(x_loc, height, true)
end

function Pipes:place(param1, param2)
    local rpipe = self.pqueue:enqueue({param1 = param1, param2 = param2})
    if rpipe ~= nil then
        self:remove_pipe(rpipe.param1, rpipe.param2)
    end
    self:add_pipe(param1, param2)
end

return Pipes
