local M = {}

-- Commands are seperated by ?
-- Example: P:12:24?B:11:12?...
local stream = ""
local cmd_sep = "?"
local sep = ":"

function M.parse(chunk)

    stream = stream .. chunk
    local idx = string.find(stream, cmd_sep)

    assert(idx, "no ? in stream")

    local cmd = string.sub(stream, 1, idx-1)
    stream = string.sub(stream, idx+1, string.len(stream))

    local splits = {}
    for split in string.gmatch(cmd, "([^"..sep.."]+)") do
        table.insert(splits, split)
    end
    return splits
end


return M
