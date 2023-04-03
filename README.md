
## go run main.go --help

```
Usage: /tmp/go-build1589409234/b001/exe/main [OPTIONS]
  -AuthToken string
    	Visit: (https://github.com/settings/tokens) generate a token with 'public_repo' or 'repo' permissions. Store access token in a file (token.txt). Example usage: -AuthToken=$(cat token.txt)
  -Repo string
    	GitHub repo where release notes to be released. (default "leewannacott/quick-lint-js")
  -TagsRepo string
    	GitHub repo to get release tags from. (default "quick-lint/quick-lint-js")
  -help
    	Print usage information. Example: $ go run main.go -Repo=leewannacott/quick-lint-js -TagsRepo=quick-lint/quick-lint-js -AuthToken=$(cat token.txt)
```

`$ go run main.go -Repo=leewannacott/quick-lint-js -TagsRepo=quick-lint/quick-lint-js -AuthToken=$(cat token.txt)`

Go routines got performance from ~10 seconds to ~2 seconds.

Aiming for compatibility with: https://keepachangelog.com/en/1.0.0/

Developed for: https://github.com/quick-lint/quick-lint-js/pull/669
