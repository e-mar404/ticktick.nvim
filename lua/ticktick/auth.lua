local utils = require 'ticktick.utils'

---@class Credentials
---@field client_id string
---@field client_secret string

local auth = {}

auth.login = function (chan)
  -- get credentials from user
  local instructions = {
    "How to create credentials to use for TickTick api.",
    "",
    "Go to https://developer.ticktick.com/manage and create a new app.",
    "Make sure to set the following fields:",
    "",
    "\t- Name of the app to `ticktick.nvim`",
    "\t- Redirect URI to `http://127.0.0.1:8080`",
    "",
    "Open the newly created app and add the Client ID and Secret bellow.",
    "",
    "Client ID: ",
    "Client Secret: ",
    "",
    "Press enter while either in insert or normal mode to submit.",
  }

  local buf = vim.api.nvim_create_buf(false, false)

  -- sensible options for a temporary floating window
  vim.api.nvim_set_option_value('buftype', 'nofile', {buf=buf})
  vim.api.nvim_set_option_value('bufhidden', 'wipe', {buf=buf})
  vim.api.nvim_set_option_value('swapfile', false, {buf=buf})
  vim.api.nvim_buf_set_lines(buf, 0, -1, false, instructions)

  -- send request to go server to get access token in a keymap
  vim.keymap.set({ 'n', 'i' }, '<CR>', function ()
    vim.cmd('stopinsert')

    vim.schedule(function ()
      print("Sign in to TickTick.com on your browser")

      local lines = vim.api.nvim_buf_get_lines(buf, 10, 12, false)
      local creds = utils._extract_creds(lines)
      local reply = vim.rpcrequest(chan, "Auth.Login", creds)

      local msg = "request did not succeed"
      if reply then
        msg = "request succeeded"
      end

      vim.notify(msg, vim.log.levels.DEBUG)
    end)
  end, { buf=buf })

  local width, height = 0, #instructions
  for _, line in pairs(instructions) do
    if #line > width then
      width = #line
    end
  end

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
  vim.api.nvim_win_set_cursor(win, {11, 10})

  vim.cmd('startinsert!')
end

return auth
