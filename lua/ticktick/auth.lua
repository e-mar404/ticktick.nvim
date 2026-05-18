local utils = require 'ticktick.utils'

---@class Credentials
---@field client_id string
---@field client_secret string

local auth = {}

--- Function to provide credentials needed to authenticate with ticktick api. If there are any credentials saved then it will use those to re-authenticate with the api, if you want to provide new credentials then run `:TickLoginForce`
---
--- This function must be called from inside a coroutine 
---
--- If force parameter is present then even if there are credentials saved user will be prompted for new ones
---@param force boolean
auth.login = function (force)
  local co = coroutine.running()
  assert(co, "function must run inside a coroutine")

  --[@as Credentials]
  local creds
  local chan = utils._get_chan()
  local credentials_exist = vim.rpcrequest(chan, "Auth.CredentialsExist")

  if credentials_exist and not force then
    creds = vim.rpcrequest(chan, "Auth.RetrieveCredentials")
  else
    auth._prompt_for_creds(co)
    creds = coroutine.yield()
  end

  vim.cmd('stopinsert')

  print("Sign in to TickTick.com on your browser")

  vim.schedule(coroutine.wrap(function ()
    local reply = vim.rpcrequest(chan, "Auth.Login", creds)

    local msg = "could not sign in"
    if reply then
      msg = "successfully logged in"
    end

    vim.notify(msg, vim.log.levels.DEBUG)
  end))
end

auth._prompt_for_creds = function (co)
  local config = require('ticktick').config
  local instructions = {
    "How to create credentials to use for TickTick api.",
    "",
    "Go to https://developer.ticktick.com/manage and create a new app.",
    "Make sure to set the following fields:",
    "",
    "\t- Name of the app to `ticktick.nvim`",
    "\t- Redirect URI to `http://127.0.0.1" .. config.rpc_port .. "`",
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

  vim.keymap.set({ 'n', 'i' }, '<CR>', function ()
    vim.cmd('stopinsert')

    print("Sign in to TickTick.com on your browser")

    vim.schedule(coroutine.wrap(function ()
      local lines = vim.api.nvim_buf_get_lines(buf, 10, 12, false)
      local creds = utils._extract_creds(lines)
      coroutine.resume(co, creds)
    end))
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
