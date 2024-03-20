local window = require("window")
local render = require("game.render")

local group = vim.api.nvim_create_augroup("flappy.window", {
    clear = true
})

 -- create a window
 -- b -> bird x -> pipe
 function START()
    local win_config = window.create_window()
    render.draw_bird(win_config.dim, win_config.buffer_id, 10, 16)
    render.draw_pipe(win_config.dim, win_config.buffer_id, 10, 5)
 end

 function END()
 end
