local Game = require("game.game")

-- TODO: Don't know wat to do with this
local group = vim.api.nvim_create_augroup("flappy.window", {
    clear = true
})

local game
function START(host, port)
    -- assert(game == nil, "Game already started")

    local dim = {width = 80, height = 24}
    game = Game:New({dim = dim, host = host, port = port})
    game:Start()

end

function END()
    -- assert(game, "No Game to end")
    if game.running then
       game:End()
    end
    game = nil
end
