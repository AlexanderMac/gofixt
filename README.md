<div align="center">
  <h1>Go-Fit</h1>
  <p>Fix binary file extensions using magic numbers header signature</p>
  <p>
    <a href="https://github.com/alexandermac/gofit/actions/workflows/ci.yml?query=branch%3Amaster"><img src="https://github.com/alexandermac/gofit/actions/workflows/ci.yml/badge.svg" alt="Build Status"></a>
    <a href="LICENSE"><img src="https://img.shields.io/github/license/alexandermac/gofit.svg" alt="License"></a>
  </p>
</div>


```sh
# Run the app in scan mode
gofit --dir=./tmp scan
# Output
| File                 | Mime             | Orig ext. | Real ext. | Fixed | Error         |
| ----                 | ----             | --------- | --------- | ----- | -----         |
| tmp/report           | unknown          |           |           |       |               |
| tmp/empty.pdf        | unknown          | .pdf      |           |       | File is empty |
| tmp/pdf-doc.txt      | application/pdf  | .txt      | .pdf      |       |               |
| tmp/png-picture2.png | image/png        | .png      | .png      |       |               |
| tmp/archive.tar.gz   | application/gzip | .gz       | .gz       |       |               |
| tmp/chat             | unknown          |           |           |       |               |
| tmp/png-picture1.jpg | image/png        | .jpg      | .png      |       |               |

7 file(s) processed in 495.507µs

# Run the app in fix mode
gofit --dir=./tmp fix
| File                 | Mime             | Orig ext. | Real ext. | Fixed | Error         |
| ----                 | ----             | --------- | --------- | ----- | -----         |
| tmp/report           | unknown          |           |           |       |               |
| tmp/empty.pdf        | unknown          | .pdf      |           |       | File is empty |
| tmp/png-picture2.png | image/png        | .png      | .png      |       |               |
| tmp/chat             | unknown          |           |           |       |               |
| tmp/archive.tar.gz   | application/gzip | .gz       | .gz       |       |               |
| tmp/png-picture1.jpg | image/png        | .jpg      | .png      | Yes   |               |
| tmp/pdf-doc.txt      | application/pdf  | .txt      | .pdf      | Yes   |               |

7 file(s) processed in 582.46µs
```

# Contents
- [Contents](#contents)
- [Install](#install)
- [Usage](#usage)
- [License](#license)

# Install
```sh
# Install the gofit binary in your $GOPATH/bin directory
go install github.com/alexandermac/gom/cmd/gom
```

# Usage
## CLI
```
Usage: gofit [flags] command

Flags:
  --dir    Scanning directory (absolute or relative path)
  --silent Don't print report

Commands:
  scan     Scans files in the provided directory recursively. Prints files info in a table format
  fix      Scans files in the provided directory recursively and fixes their extensions (when needed). Prints files info in a table format
  help     Shows this help
  version  Prints app version

Examples:
  gofit --dir=~/images scan
  gofit --dir=~/files --silent fix
```

# License
Licensed under the MIT license.

# Author
Alexander Mac
