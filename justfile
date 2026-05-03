# https://just.systems

dev:
  air

lua-todo:
  grep "TODO" -rn ./lua ./plugin --color=always

go-todo:
  grep "TODO" -rn ./tickgo --color=always

lint:
  gofmt -w ./tickgo
