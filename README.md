## Installation
```
go install github.com/LeeWannacott/blazingly-fast-releases@latest
```

## Usage
```
blazingly-fast-releases -target=leewannacott/blazingly-fast-releases -tags=quick-lint/quick-lint-js -token=$(cat token.txt) -changelog=docs/CHANGELOG.md
```

## blazingly-fast-releases --help
```
Usage: blazingly-fast-releases [OPTIONS]
  --changelog string
    	Relative path to CHANGELOG.md. (default "docs/CHANGELOG.md")
  --help
    	Display help
  --tags string
    	Repo where you want tags from. (default "owner/repo")
  --target string
    	Repo where you want to update the releases. (default "owner/repo")
  --token string
    	(https://github.com/settings/tokens) generate token with 'public_repo' or 'repo' (private repos) permission.
```

### Prerequisites:
* Go: `https://go.dev/doc/install`

### Information:
* Go routines got performance from ~10 seconds to ~2 seconds.

* Aiming for compatibility with: https://keepachangelog.com/en/1.0.0/

* Developed for quick-lint-js: https://github.com/quick-lint/quick-lint-js/pull/669

### Projects using blazingly-fast-releases:
* https://github.com/quick-lint/quick-lint-js
