### Install blazingly-fast-releases
$```go install github.com/LeeWannacott/blazingly-fast-releases@latest```

### $ blazingly-fast-releases --help
```
Usage: /tmp/go-build2653530073/b001/exe/main [OPTIONS]
Usage: blazingly-fast-releases [OPTIONS]
  --changelog string
    	Relative folder path to changelog.md file (default "docs/CHANGELOG.md")
  --help
    	Example: $ blazingly-fast-releases --target=leewannacott/blazingly-fast-releases --tags=quick-lint/quick-lint-js --changelog=docs/CHANGELOG.md --token=$(cat token.txt)
  --tags string
    	Repo where you want to get the tags from. (default "leewannacott/quick-lint-js")
  --target string
    	Repo where you want to update the releases. (default "leewannacott/blazingly-fast-releases")
  --token string
    	(https://github.com/settings/tokens) generate a token with 'public_repo' or 'repo' permissions. Store token in a file (token.txt). Example: --token=$(cat token.txt)
```

### Example of running in the terminal (CLI)
$ ```blazingly-fast-releases -target=leewannacott/blazingly-fast-releases -tags=quick-lint/quick-lint-js -token=$(cat token.txt) -changelog=docs/CHANGELOG.md```

### Prerequisites: Install golang 
* `https://go.dev/doc/install`

### Info:
* Go routines got performance from ~10 seconds to ~2 seconds.

* Aiming for compatibility with: https://keepachangelog.com/en/1.0.0/

* Developed for quick-lint-js: https://github.com/quick-lint/quick-lint-js/pull/669

### Projects using blazingly-fast-releases:
* https://github.com/quick-lint/quick-lint-js
