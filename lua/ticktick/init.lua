---@class Configuration 
---
--- Port used by the go rpc server that handles the TickTick api requests
---@field rpc_port string
---
--- Port used to create a server to handle the callback by the OAuth sign in flow from TickTick
---@field callback_port string
---
--- If set to true will not auto start the rpc server, this will have to be done separately ... for dev purposes
---@field dev boolean

local ticktick = {}

---@type Configuration
local defaults = {
  rpc_port = ":8080",
  callback_port = ":9090",
  dev = false
}

---@type Configuration
ticktick.config = vim.tbl_deep_extend('force', {}, defaults)

ticktick.setup = function (opts)
  ticktick.config = vim.tbl_deep_extend('force', {}, defaults, opts or {})

  if not ticktick.config.dev then
    vim.fn.jobstart({ 'tickgo' }, {})
  end
end

return ticktick
