local uv = vim.uv

local server = {}

server.start = function ()
  local s = uv.new_tcp()
  assert(s, "unable to start tcp server")

  s:bind("127.0.0.1", 8080)

  s:listen(128, function (err)
    assert(not err, err)

    local client = uv.new_tcp()
    assert(client, "unable to start tcp client")

    s:accept(client)
    -- TODO: this executes twice, once for the actual redirect and another for the favicon, dont think its a big deal but I end up with the correct state and code at the end so prob fine
    client:read_start(function(read_err, chunk)
      assert(not read_err, read_err)

      local url = chunk:match("%/%?%S+")

      local _, code_end = assert(url:find("code="))
      local sep = url:find("&")
      local _, state_end = assert(url:find("state="))

      local code = url:sub(code_end + 1, sep - 1)
      local state = url:sub(state_end + 1, #url)

      print("code: " .. code .. "\nstate: " .. state)

      -- TODO: hard coded values, will need to account for a few errors (like the asserts above)
      local body = "successful login, you may close the browser"

      local response = table.concat({
        "HTTP/1.1 200 OK",
        "Content-Type: text/html",
        "Content-Length: " .. #body,
        "",
        body,
      }, "\r\n")

      client:write(response, function ()
        client:shutdown(function ()
          client:close()

        end)
      end)

      client:read_stop()
    end)
  end)

  print("TCP server listening at 127.0.0.1:8080")
end

return server
