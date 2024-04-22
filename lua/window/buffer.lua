local api = vim.api

local Buffer = {}
Buffer.__index = Buffer

function Buffer:new(height, width)
    local buffer = {}
    for _ = 1, height, 1 do
        local row = {}
        for _ = 1, width, 1 do
            table.insert(row, " ")
        end
        table.insert(buffer, row)
    end

    self.buffer = buffer
    self.dim = {height = height, width = width}
    self.buffer_id = api.nvim_create_buf(false, true)

    return self
end

function Buffer:render()
    local tbl = {}
    for _, row in ipairs(self.buffer) do
        local line = ""
        for _, col in ipairs(row) do
            line = line .. col
        end
        table.insert(tbl, line)
    end
    api.nvim_buf_set_lines(self.buffer_id, 0, -1, 0, tbl)
end

function Buffer:out_of_bounds(x, y)
    local height = self.dim.height
    if y < 1 or y > height then
        return true
    end
    return false
end

function Buffer:PlacePoint(x, y, char)
    self.buffer[y][x] = char
end

function Buffer:RemovePoint(x, y)
    self.buffer[y][x] = " "
end

function Buffer:PlaceVline(x, h, char, fromEnd)
    for y = 1, h, 1 do
        self.buffer[y][x] = char
    end

    if fromEnd then
        local sze = #self.buffer
        for y = 0, h-1, 1 do
            self.buffer[sze-y][x] = char
        end
    end
end

function Buffer:RemoveVline(x, h, fromEnd)
    for y = 1, h, 1 do
        self.buffer[y][x] = " "
    end

    if fromEnd then
        local sze = #self.buffer
        for y = 0, h-1, 1 do
            self.buffer[sze-y][x] = " "
        end
    end
end

function Buffer:close()
    api.nvim_buf_delete(self.buffer_id, {force= true})
end

return Buffer
