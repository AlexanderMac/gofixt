<div align="center">
  <h1>Go-Fix-Ext</h1>
  <p>Fix binary file extensions using magic numbers header signature</p>
  <p>
    <a href="https://github.com/alexandermac/gofixt/actions/workflows/ci.yml?query=branch%3Amaster"><img src="https://github.com/alexandermac/gofixt/actions/workflows/ci.yml/badge.svg" alt="Build Status"></a>
    <a href="LICENSE"><img src="https://img.shields.io/github/license/alexandermac/gofixt.svg" alt="License"></a>
  </p>
</div>

```sh
# Run the app in scan mode
gofixt --dir=./tmp scan
# Output
| File                   | Mime             | Orig ext. | Real ext. | Notes         |
| ----                   | ----             | --------- | --------- | -----         |
| <dir>/png-picture2.png | image/png        | .png      | .png      |               |
| <dir>/archive.tar.gz   | application/gzip | .gz       | .gz       |               |
| <dir>/chat             | unknown          |           |           |               |
| <dir>/empty.pdf        | unknown          | .pdf      |           | File is empty |
| <dir>/pdf-doc.pdf      | application/pdf  | .pdf      | .pdf      |               |
| <dir>/pdf-doc.txt      | application/pdf  | .txt      | .pdf      | Fix required  |
| <dir>/png-picture1.jpg | image/png        | .jpg      | .png      | Fix required  |
7 file(s) processed and 0 file(s) fixed in 211.589µs

# Run the app in fix mode
gofixt --dir=./tmp fix
| File                   | Mime             | Orig ext. | Real ext. | Notes                                     |
| ----                   | ----             | --------- | --------- | -----                                     |
| <dir>/png-picture2.png | image/png        | .png      | .png      |                                           |
| <dir>/archive.tar.gz   | application/gzip | .gz       | .gz       |                                           |
| <dir>/chat             | unknown          |           |           |                                           |
| <dir>/empty.pdf        | unknown          | .pdf      |           | File is empty                             |
| <dir>/pdf-doc.pdf      | application/pdf  | .pdf      | .pdf      |                                           |
| <dir>/pdf-doc.txt      | application/pdf  | .txt      | .pdf      | File with the same name is already exists |
| <dir>/png-picture1.jpg | image/png        | .jpg      | .png      | Fixed                                     |
7 file(s) processed and 1 file(s) fixed in 246.946µs
```

# Contents
- [Contents](#contents)
- [Install](#install)
- [Usage](#usage)
- [License](#license)

# Install
```sh
# Install the gofixt binary in your $GOPATH/bin directory
go install github.com/alexandermac/gom/cmd/gom
```

# Usage
## CLI
```
Usage: gofixt [flags] command

Flags:
  --dir    Scanning directory (absolute or relative path)
  --silent Don't print report

Commands:
  scan     Scans files in the provided directory recursively. Prints files info in a table format
  fix      Scans files in the provided directory recursively and fixes their extensions (when needed). Prints files info in a table format
  help     Shows this help
  version  Prints app version

Examples:
  gofixt --dir=~/images scan
  gofixt --dir=~/files --silent fix
```

# License
Licensed under the MIT license.

# Author
Alexander Mac
