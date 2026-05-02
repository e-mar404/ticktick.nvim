local utils = {}

---@param lines string[]
---@return Credentials
utils._extract_creds = function (lines)
  --[@as Credentials]
  local creds = {}

  creds.client_id = lines[1]
    :gsub('Client ID:', '')
    :gsub(' ' , '')

  creds.client_secret = lines[2]
    :gsub('Client Secret:', '')
    :gsub(' ' , '')

  return creds
end

return utils
