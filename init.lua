local window = require("window")

local group = vim.api.nvim_create_augroup("flappy.window", {
    clear = true
})

 -- create a window
 -- b -> bird x -> pipe
 function START()
    window.create_window()
 end

 function END()
 end
