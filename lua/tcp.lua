local M = {}

-- Define buffer as a global variable
local buffer = ""

-- Commands are separated by ?
-- Example: P:12:24?B:11:12?...
function M.process_packets(chunk)
    local stream = buffer .. chunk
    local cmd_sep = "?"
    local sep = ":"
    local splits = {}

    local idx = string.find(stream, cmd_sep)
    while idx do
        local cmd = string.sub(stream, 1, idx - 1)
        stream = string.sub(stream, idx + 1)

        local cmd_splits = {}
        for split in string.gmatch(cmd, "([^" .. sep .. "]+)") do
            table.insert(cmd_splits, split)
        end

        table.insert(splits, cmd_splits)

        idx = string.find(stream, cmd_sep)
    end

    -- Update the global buffer
    buffer = stream

    return splits
end


return M
