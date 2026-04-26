local ticktick = require('ticktick')

vim.api.nvim_create_user_command('TickInit', ticktick.init, {})
