<div align="center">
  <h1>Go-Fit</h1>
  <p>Fix binary files extensions using magic numbers header signature</p>
  <p>
    <a href="https://github.com/alexandermac/gofit/actions/workflows/ci.yml?query=branch%3Amaster"><img src="https://github.com/alexandermac/gofit/actions/workflows/ci.yml/badge.svg" alt="Build Status"></a>
    <a href="LICENSE"><img src="https://img.shields.io/github/license/alexandermac/gofit.svg" alt="License"></a>
  </p>
</div>


```sh
# Run the app in scan mode
gofit --dir=./tmp scan
# Output
| File                | Mime              | Orig ext.  | Real ext.  | 
| ----                | ----              | ---------  | ---------  | 
| tmp/archive.tar.gz  | application/gzip  | .gz        | .gz        |
| tmp/chat.svg        |                   | .svg       | .unknown   |
| tmp/pdf file.txt    | application/pdf   | .txt       | .pdf       |
| tmp/png file1.jpg   | image/png         | .jpg       | .png       |
| tmp/png file2.png   | image/png         | .png       | .png       |
Done!

# Run the app in fix mode
gofit --dir=./tmp fix
# Output
| File                | Mime              | Orig ext.  | Real ext.  | Fixed  | 
| ----                | ----              | ---------  | ---------  | -----  | 
| tmp/archive.tar.gz  | application/gzip  | .gz        | .gz        |        |
| tmp/chat.svg        |                   | .svg       | .unknown   |        |
| tmp/pdf file.txt    | application/pdf   | .txt       | .pdf       | yes    |
| tmp/png file1.jpg   | image/png         | .jpg       | .png       | yes    |
| tmp/png file2.png   | image/png         | .png       | .png       |        |
Done!
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
  --dir Scanning directory (absolute or relative path)

Commands:
  scan     Scans files in the provided directory recursively. Prints files info in a table format
  fix      Scans files in the provided directory recursively and fixes their extensions (when needed). Prints files info in a table format
  help     Shows this help
  version  Prints app version

Examples:
  gofit --dir=~/images scan
  gofit --dir=~/files fix
```

# License
Licensed under the MIT license.

# Author
Alexander Mac
