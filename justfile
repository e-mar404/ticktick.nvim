# https://just.systems

dev:
  air

lua-todo:
  -grep "TODO" -rn ./lua ./plugin --color=always

go-todo:
  -grep "TODO" -rn ./tickgo --color=always

lint:
  gofmt -w ./tickgo

link:
  # will create a symlink and place the directory in a special neovim dir (pack/*/start) that will auto start the plugin. This is purely for development and will not be how the plugin will be installed when it is rolled out
  mkdir -p ~/.local/share/nvim/site/pack/personal/start
  ln -s `pwd` ~/.local/share/nvim/site/pack/personal/start/ticktick.nvim

