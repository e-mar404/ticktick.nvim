local auth = require('ticktick.auth')

vim.api.nvim_create_user_command('TickLogin', function (_)
  auth.login()
end, { nargs = 0 })
