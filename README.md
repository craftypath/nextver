# VerIt
```
$ git init && echo "hello" > README.md && git add . && git commit -m "initial commit"
$ verit 0.x.0
0.1.0
# end of (successful) build
$ git push --tags
$ verit 0.x.0
0.2.0
# manually bump major version 0 -> 1, don't move the x for doing just that!
$ verit 1.x.0
1.0.0
$ verit 1.x.0
1.1.0
```
