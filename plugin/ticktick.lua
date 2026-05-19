local auth = require('ticktick.auth')
local ui = require('ticktick.ui')

vim.api.nvim_create_user_command('TickLogin', function (_)
  coroutine.wrap(auth.login)()
end, { nargs = 0 })

vim.api.nvim_create_user_command('TickLoginForce', function (_)
  coroutine.wrap(function ()
    auth.login(true)
  end)()
end, { nargs = 0 })

vim.api.nvim_create_user_command('TickOpenUI', function (_)
  ui.open()
end, { nargs = 0})
