local utils = require('ticktick.utils')
local ui = {}

ui.open = function ()
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

  local chan = utils._get_chan()
  local projects = vim.rpcrequest(chan, "Project.GetAll")

  if not projects then
    print("no projects received")
    return
  end

  print('projects: \n\n')
  local str = ''
  for _, project in ipairs(projects) do
    str = str .. 'id: ' .. project.id .. '\nname: ' .. project.name .. '\n'
  end
  print(str)

  -- TODO: make tabs for the projects received

  -- TODO: fetch the tasks for those projects asynchronously in the background
end

return ui
