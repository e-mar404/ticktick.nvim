local utils = require "ticktick.utils"
local server= require "ticktick.server"
---@module "ticktick.config"

---@class Api
---@field get_access_token function
---@field oauth_url string

---@class Credentials
---@field client_id string
---@field client_secret string

-- TODO: figure out how to get past the "missing x field" warning, maybe im doing type assertions wrong...
---@type Api
local api = {
  oauth_url = "https://ticktick.com/oauth/authorize",
}

---@param creds Credentials 
---@return string
api.get_access_token = function (creds)
  -- Using their dev docs on getting access token: https://developer.ticktick.com/api#/openapi?id=get-access-token
  local state = utils._generate_new_state()

  vim.ui.open(api.oauth_url ..
    "?scope=tasks:write tasks:read&client_id=" .. creds.client_id ..
    "&state=" .. state ..
    "&redirect_uri=" .. "http://127.0.0.1:8080" ..
    "&response_type=code")

  vim.notify("Sign in on browser tab", vim.log.levels.INFO)

  server.start(function (code)
    vim.notify(code, vim.log.levels.DEBUG)
  end)

  return "access_token"
end

return api
