local utils = require('ticktick.utils')

local ns_id = vim.api.nvim_create_namespace("TickTickUI")
vim.api.nvim_set_hl(0, "TickTickActiveTab", {
  fg = "#1a1b26",
  bg = "#7aa2f7",
  bold = true,
})

vim.api.nvim_set_hl(0, "TickTickInactiveTab", {
  bg = "#1a1b26",
  fg = "#7aa2f7",
  bold = true,
})

local ui = {
  projects = {},
  tasks = {},
  taskbar = {
    active = 1,
  },
  canvas = {}
}

ui.tasks._write_to_canvas = function ()
  local height = vim.o.lines - 2
  for _ = 1, height, 1 do
    table.insert(ui.canvas, "")
  end
end

ui.taskbar.tabs = function ()
  local tabs = {}
  for idx, project in ipairs(ui.projects) do
    local hl = idx == ui.taskbar.active and 'TickTickActiveTab' or 'TickTickInactiveTab'
    table.insert(tabs, { " " .. project.name .. " ", hl })
  end
  return tabs
end

ui.open = function ()
  local chan = utils._get_chan()
  ui.projects = vim.rpcrequest(chan, "Project.GetAll")

  if not ui.projects then
    print("no projects received")
    return
  end

  local buf = vim.api.nvim_create_buf(false, false)

  vim.api.nvim_set_option_value('buftype', 'nofile', {buf=buf})
  vim.api.nvim_set_option_value('bufhidden', 'wipe', {buf=buf})
  vim.api.nvim_set_option_value('swapfile', false, {buf=buf})

  ui.tasks._write_to_canvas()
  vim.api.nvim_buf_set_lines(buf, 0, -1, false, ui.canvas)

  vim.api.nvim_buf_set_extmark(buf, ns_id, (#ui.canvas - 1), 0, {
    virt_text = ui.taskbar.tabs(),
    virt_text_pos = 'overlay',
  })

  -- Uses 80 as the arbitrary value that is "used up" because I should technically not have many things that are past 80 chars, this is personal preference and will have to change if other people use it lol
  local height = vim.o.lines - 1
  local width = math.floor(vim.o.columns - 80)

  ---@type vim.api.keyset.win_config
  local win_config = {
    vertical = true,
    split = 'right',
    style = 'minimal',
    height = height,
    width = width,
  }

  local _ = vim.api.nvim_open_win(buf, true, win_config)

  -- TODO: make tabs for the projects received

  -- TODO: fetch the tasks for those projects asynchronously in the background
end

-- TODO: make a function for a floating prompt

return ui
