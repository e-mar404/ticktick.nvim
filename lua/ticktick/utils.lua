---@module  "ticktick.config" 

local utils = {}

---@param lines string[]
---@return TickTickConfig
utils._parse_credentials = function (lines)
  ---@type TickTickConfig
  local creds = {}

  local start, _ = lines[1]:find(":")
  creds.client_id = lines[1]:sub(start + 1):gsub(' ', '')

  start, _ = lines[2]:find(":")
  creds.client_secret = lines[2]:sub(start + 1):gsub(' ', '')

  return creds
end

return utils
