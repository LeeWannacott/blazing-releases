
## go run main.go --help

```
Usage: /tmp/go-build2653530073/b001/exe/main [OPTIONS]
  --help
      Example: $ go run main.go --target=leewannacott/quick-release-notes --tags=quick-lint/quick-lint-js --token=$(cat token.txt)
  --changelog string
    	Relative folder path to changelog.md file (default "docs/CHANGELOG.md")
  --tags string
    	Repo where you want to get the tags from. (default "leewannacott/quick-lint-js")
  --target string
    	Repo where you want to update the releases. (default "leewannacott/quick-release-notes")
  --token string
    	(https://github.com/settings/tokens) generate a token with 'public_repo' or 'repo' permissions. Store token in a file (token.txt). Example: --token=$(cat token.txt)
```

```
$ go run main.go -target=leewannacott/quick-release-notes -tags=quick-lint/quick-lint-js -token=$(cat token.txt) -changelog=docs/CHANGELOG.md
```

Go routines got performance from ~10 seconds to ~2 seconds.

Aiming for compatibility with: https://keepachangelog.com/en/1.0.0/

Developed for quick-lint-js: https://github.com/quick-lint/quick-lint-js/pull/669
