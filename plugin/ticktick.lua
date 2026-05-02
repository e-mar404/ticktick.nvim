local auth = require('ticktick.auth')

local chan

local get_chan = function ()
  if chan then
    return chan
  end

  -- TODO: have the host and port be pulled from env vars or something similar
  chan = vim.fn.sockconnect('tcp', '127.0.0.1:8080', { rpc = true })

  return chan
end

vim.api.nvim_create_user_command('TickLogin', function (_)
  auth.login(get_chan())
end, { nargs = 0 })
