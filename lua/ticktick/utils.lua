---@module  "ticktick.api" 

local utils = {}

---@param lines string[]
---@return Credentials 
utils._parse_credentials = function (lines)
  ---@type Credentials 
  local creds = {}

  local start, _ = lines[1]:find(":")
  creds.client_id = lines[1]:sub(start + 1):gsub(' ', '')

  start, _ = lines[2]:find(":")
  creds.client_secret = lines[2]:sub(start + 1):gsub(' ', '')

  return creds
end

---@return string
utils._generate_new_state = function ()
  -- TODO: temp, make sure to create a unique state and save it to disk
  return "state"
end

return utils
