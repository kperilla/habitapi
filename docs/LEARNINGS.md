# Learnings

- 2024/09/03 - Tried to make Fastify work with Typescript, and boy was it annoying,
- The test cases for the examples that come with it, just don't work. I think that Express is the way to go
- 2024/09/05 - Modules in Go is surprisingly complicated. Especially if you have multiple packages
but haven't deployed anything
 - Apparently it all depends on a go mod replace command, that was NOWHERE except one
very specific doc. Idk if it's just me or if everyone else already knows that. Very frustrating
- I'm using these resources
 - https://www.gobeyond.dev/standard-package-layout/
 - https://www.gobeyond.dev/packages-as-layers/
 - https://go.dev/doc/tutorial/call-module-code
- Testing was a big headache too. Took forever to realize that it depends 100% on
the directory structure. main.go HAD to be at the root directory for it to recognize
the http directory
