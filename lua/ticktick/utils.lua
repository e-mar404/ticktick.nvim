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

utils._get_chan = function ()
  if utils.chan then
    return utils.chan
  end

  local addrs = '127.0.0.1' .. require('ticktick').config.rpc_port
  utils.chan = vim.fn.sockconnect('tcp', addrs, { rpc = true })

  return utils.chan
end

return utils
