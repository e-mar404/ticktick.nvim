local config = require "ticktick.config"
local ticktick = {}

---@param args vim.api.keyset.user_command
ticktick.init = function (args)
  -- find out if there are saved creds
  if config.load() then
    return
  end

  config.get_creds()
end

return ticktick
