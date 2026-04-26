local utils = require "ticktick.utils"
local config= require "ticktick.config"
local ticktick = {}

---@param args vim.api.keyset.user_command
ticktick.init = function (args)
  config.editor()
end

return ticktick
