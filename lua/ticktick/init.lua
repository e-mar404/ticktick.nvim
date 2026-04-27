local config = require "ticktick.config"
local ticktick = {}

---@param args vim.api.keyset.user_command
ticktick.init = function (args)
  config.load_access_token()
end

return ticktick
