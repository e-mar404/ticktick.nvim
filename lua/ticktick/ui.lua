local utils = require('ticktick.utils')
local ui = {}

ui.open = function ()
  -- TODO: open a split window
  local buf = vim.api.nvim_create_buf(false, false)

  -- Uses 80 as the arbitrary value that is "used up" because I should technically not have many things that are past 80 chars, this is personal preference and will have to change if other people use it lol
  local width = math.floor(vim.o.columns - 80)

  ---@type vim.api.keyset.win_config
  local win_config = {
    vertical = true,
    split = 'right',
    style = 'minimal',
    width = width,
  }

  local _ = vim.api.nvim_open_win(buf, true, win_config)

  -- TODO: fetch all lists/projects & make 'tabs' for those
  local chan = utils._get_chan()
  local lists = vim.rpcrequest(chan, "ProjectService.GetAll")

  print('got ' .. lists)

  -- TODO: fetch the tasks for those lists/projects asynchronously in the background 
end

return ui
