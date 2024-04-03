local M = {}

function M.parse(str)
    local splits = {}
    local sep = ":"
    for split in string.gmatch(str, "([^"..sep.."]+)") do
        table.insert(splits, split)
    end
    return splits
end

return M
