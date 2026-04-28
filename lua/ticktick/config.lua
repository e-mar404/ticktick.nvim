---@class TickTickConfig
---@field load_access_token function

local utils = require "ticktick.utils"
local api   = require "ticktick.api"

---@type TickTickConfig
local config = {}

config.load_access_token = function ()
  -- TODO: for now it assumes that no access token is saved locally, so it will always ask for client id and secret
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

  -- TODO: maybe extract this function out, something like a on_submit() and move the keymap somewhere else
  -- This will need to take into account that it needs to be after creating the buf
  -- I can potentially do a prep_buf(buf) that will create the keymap and add the sensible settings
  vim.keymap.set({'i', 'n'}, '<CR>', function ()
    vim.cmd('stopinsert')

    vim.schedule(function ()
      local lines = vim.api.nvim_buf_get_lines(buf, 10, 12, false)

      ---@type Credentials
      local creds = utils._parse_credentials(lines)

      if creds.client_id == "" then
        error("No Client ID was set, please add it")
      elseif creds.client_secret == "" then
        error("No Client Secret was set, please add it")
      else
        vim.api.nvim_win_close(0, true)

        api.get_access_token(creds)
      end
    end)
  end, { buffer = buf })

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

return config
