# ticktick.nvim

A Neovim lua plugin using a go server to interact with the TickTick API directly
from neovim.

> ⚠️ **Disclaimer**  
> This is an unofficial plugin.  
> TickTick is a trademark of its respective owner(s).  
> This project is not affiliated with, endorsed by, or sponsored by TickTick.

## Features

What features do I want?

- [ ] manage tasks (crud operations)
- [ ] ability to see different lists in their own view 
- [ ] habit check in
- [ ] filtering / grep of tasks

If I do the above that will be a good enough proof of concept for me.

### Road map 

This will act as my checklist on the technical road map and how I am going to
implement the above.

1. **Access token** ✅

Access token is able to be retrieved and saved to disk by adding the client ID
and secret to a pop up menu in neovim.

2. **fetch all available tasks and display them on a new buffer**

3. **Fetch tasks that belong to a specific list**

4. **Add check box  and actions on tasks**

5. **Repeat  2-4 but for a habits**

## Requirements

- A TickTick account
- go
- just 

### Notes on local development

Note: I will put this in here just. Not expecting public contributions on this 
so mainly just for me.

Recently I moved to the built-in package manager (`nvim.pack`), and right now
is the first time I have needed to use local plugins since moving to nvim.pack.
Well it turns out that local plugins are not very well supported[^1]. The
explanation given is fine and I don't mind the thoughts in there. However using
`file:///` to develop plugins is not very nice, specially since it will only
pull committed code and I will have to do `vim.pack.update()` every time I want
to test things. Second however, a good point was brought up, symlinks. So that
is what I am using. I made a symlink from this dir to
`~/.local/share/nvim/site/pack/personal/start/ticktick.nvim`. Now it gets loaded
automatically.

Run `just link` to get started with local development. Make sure to run the
script at the root of this repo, it uses `pwd`.

[^1]: https://github.com/neovim/neovim/issues/34765
