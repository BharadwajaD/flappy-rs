local M = {}

local stream = ""
function M.process_tcp_packets(chunk)
    stream = stream .. chunk
    local sep = ":"
    local idx
    return function ()
        if stream == "" then
            return
        end

        idx = string.find(stream, "?")
        local cmd_str = string.sub(stream, 1, idx)
        local cmd_idx = string.find(cmd_str, sep)
        local cmd = string.sub(cmd_str, 1, cmd_idx-1)
        cmd_str = string.sub(cmd_str, cmd_idx+1, #cmd_str - 1)
        local params = {}

        for param in string.gmatch(cmd_str, "([^" .. sep .. "]+)") do
            table.insert(params, tonumber(param))
        end

        stream = string.sub(stream, idx+1)
        return { cmd = cmd,  params = params}
    end
end

return M
