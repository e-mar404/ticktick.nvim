local ticktick = {}

---@param args vim.api.keyset.user_command
ticktick.init = function (args)
  local instructions = {
    "How to create credentials to use for TickTick api.",
    "",
    "Go to https://developer.ticktick.com/manage and create a new app.",
    "Set the name of the app to `ticktick.nvim` and click next.",
    "Open the newly created app and add the Client ID and Secret bellow.",
    "",
    "Client ID: ",
    "Client Secret: ",
    "",
    "Press enter while either in insert or normal mode to submit.",
  }

  local width, height = 0, #instructions
  for _, line in pairs(instructions) do
    if #line > width then
      width = #line
    end
  end

  local buf = vim.api.nvim_create_buf(false, false)

  -- sensible options for a temporary floating window
  vim.api.nvim_set_option_value('buftype', 'nofile', {buf=buf})
  vim.api.nvim_set_option_value('bufhidden', 'wipe', {buf=buf})
  vim.api.nvim_set_option_value('swapfile', false, {buf=buf})

  local row = math.floor((vim.o.lines - height) / 2)
  local col = math.floor((vim.o.columns - width) / 2)

  ---@type vim.api.keyset.win_config
  local win_config = {
    relative='editor',
    row=row,
    col=col,
    width=width,
    height=height,
    style = 'minimal',
    border = 'rounded',
    title = 'TickTick Config',
  }

  local win = vim.api.nvim_open_win(buf, true, win_config)

  vim.api.nvim_buf_set_lines(buf, 0, -1, false, instructions)
  vim.api.nvim_win_set_cursor(win, {7, 10})

  vim.keymap.set({'i', 'n'}, '<CR>', function ()
    vim.cmd('stopinsert')

    vim.schedule(function ()
      local lines = vim.api.nvim_buf_get_lines(buf, 6, 8, false)
      vim.api.nvim_win_close(win, true)

      -- TODO: parse client id and secret then save them 
      print(vim.inspect(lines))
    end)
  end, { buffer = buf })

  vim.cmd('startinsert!')
end

return ticktick
