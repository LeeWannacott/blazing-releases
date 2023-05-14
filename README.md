## Installation
```
go install github.com/LeeWannacott/blazing-releases@latest
```

## Usage
```
blazing-releases -target=leewannacott/blazing-releases -tags=quick-lint/quick-lint-js -token=$(cat token.txt) -changelog=docs/CHANGELOG.md
```

## blazing-releases --help
```
Usage: blazing-releases [OPTIONS]
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
* [access token](https://github.com/settings/tokens) generate token with 'public_repo' access which is the least amount of permission required or select 'repo' to use with private repos in addition to pulic.

### Information:
* Go routines got performance from ~10 seconds to ~2 seconds.

* Aiming for compatibility with: https://keepachangelog.com/en/1.0.0/

* Developed for quick-lint-js: https://github.com/quick-lint/quick-lint-js/pull/669

### Projects using blazing-releases:
* https://github.com/quick-lint/quick-lint-js
