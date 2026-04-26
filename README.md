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

### Road map 

This will act as my checklist on the technical road map and how I am going to
implement the above.

1. **Access token**

There needs to be a way to input client_id and client_secret that you get from
the (ticktick development center)[https://developer.ticktick.com/manage]. This
is required to get the access token and do anything with the api so lets start
with that. This is tricky since I cant just add it as an opt in `setup()` since
that will usually get committed to user's dotfiles. It will have to be handled
on client and then stored on disk, outside the dotfiles, most likely
`XDG_DATA_DIR`.

2. **Start on list view user_command**

After Im able to get the access to the access token then I should get all the
lists (called projects on the api docs) and make a ui to show the different 
lists. Which I am thinking of doing one "tab" per list (a tab just being a
highlighted title with a list view), on a new window.

3. **Fetch tasks that belong to each list**

4. **Add check box  and actions on tasks**

5. **Repeat but for a habits view user_command**

## Requirements

- A TickTick account

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

Run the script [./link.sh](./link.sh) to get started with local development.
Make sure to run the script at the root of this repo, it uses `pwd`.

[^1]: https://github.com/neovim/neovim/issues/34765
