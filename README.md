# ticktick.nvim

A Neovim plugin for interacting with the TickTick API directly from your editor.

> ⚠️ **Disclaimer**  
> This is an unofficial plugin.  
> TickTick is a trademark of its respective owner(s).  
> This project is not affiliated with, endorsed by, or sponsored by TickTick.

## Features

What features do I want?

- [ ] habit check in
- [ ] manage tasks (create, read, update, delete)
- [ ] each task list will have a tab of its own
- [ ] handle api key and storage through neovim, and not deal with .env files

If I do the above that will be a good enough proof of concept for me.

## Requirements

- A TickTick account

### Local development

Note: I will put this in here just in case there I ever need it again.

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

Run the script [./link.sh](./link.sh) to get started with local development.
Make sure to run the script at the root of this repo, it uses `pwd`.

[^1]: https://github.com/neovim/neovim/issues/34765
