
## go run main.go --help

```
Usage: /tmp/go-build81954260/b001/exe/main [OPTIONS]
  -help
    	Example: $ go run main.go --target-repo=leewannacott/quick-release-notes --tags-repo=quick-lint/quick-lint-js --token=$(cat token.txt)
  -tags-repo string
    	GitHub repo to get release tags from. (default "leewannacott/quick-lint-js")
  -target-repo string
    	GitHub repo where you want the releases posted. (default "leewannacott/quick-release-notes")
  -token string
    	(https://github.com/settings/tokens) generate a token with 'public_repo' or 'repo' permissions. Store token in a file (token.txt). Example: --token=$(cat token.txt)
```

```
$ go run main.go --target-repo=leewannacott/quick-release-notes --tags-repo=quick-lint/quick-lint-js --token=$(cat token.txt)
```

Go routines got performance from ~10 seconds to ~2 seconds.

Aiming for compatibility with: https://keepachangelog.com/en/1.0.0/

Developed for: https://github.com/quick-lint/quick-lint-js/pull/669
