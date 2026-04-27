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
    client:read_start(function(read_err, chunk)
      assert(not read_err, read_err)

      -- TODO: hard coded values, will need to account for a few errors
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

        -- TODO: extract the code and find a way to get it back to the api.get_access_token function
        print("received chunk: " .. chunk)
      end)

      client:read_stop()
    end)
  end)

  print("TCP server listening at 127.0.0.1:8080")
end

return server
