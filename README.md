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
```
```
| File                   | Mime             | Orig ext. | Real ext. | Notes         |
| --------------------------------------------------------------------------------- |
| /png-picture2.png      | image/png        | .png      | .png      |               |
| /archive.tar.gz        | application/gzip | .gz       | .gz       |               |
| /chat                  | unknown          |           |           |               |
| /empty.pdf             | unknown          | .pdf      |           | File is empty |
| /pdf-doc.pdf           | application/pdf  | .pdf      | .pdf      |               |
| /pdf-doc.txt           | application/pdf  | .txt      | .pdf      | Fix required  |
| /png-picture1.jpg      | image/png        | .jpg      | .png      | Fix required  |

Process has been completed in 300.514µs.
- 7 file(s) processed
- 0 file(s) ignored
- 2 file(s) require fix
- 0 file(s) fixed
- 1 error(s)
```

```sh
# Run the app in fix mode
gofixt --dir=./tmp fix
```
```
 | File                   | Mime             | Orig ext. | Real ext. | Notes                                  |
 | ---------------------------------------------------------------------------------------------------------- |
 | /png-picture2.png      | image/png        | .png      | .png      |                                        |
 | /archive.tar.gz        | application/gzip | .gz       | .gz       |                                        |
 | /chat                  | unknown          |           |           |                                        |
 | /empty.pdf             | unknown          | .pdf      |           | File is empty                          |
 | /pdf-doc.pdf           | application/pdf  | .pdf      | .pdf      |                                        |
 | /pdf-doc.txt           | application/pdf  | .txt      | .pdf      | File with the same name already exists |
 | /png-picture1.jpg      | image/png        | .jpg      | .png      | Fixed                                  |
 
 Process has been completed in 250.123µs.
 - 7 file(s) processed
 - 0 file(s) ignored
 - 0 file(s) require fix
 - 1 file(s) fixed
 - 2 error(s)
```

# Contents
- [Contents](#contents)
- [Install](#install)
- [Usage](#usage)
- [License](#license)

# Install
```sh
# Install the gofixt binary in your $GOPATH/bin directory
go install github.com/alexandermac/gofixt/cmd/gofixt
```

# Usage
## CLI
```
Usage: gofixt [flags] command

Flags:
  --dir    Scanning directory (absolute or relative path)
  --exts   Comma separated list of file extensions, files with other extensions will be ignored, default: empty
  --print  Print mode: all,important,report,none, default: important

Commands:
  scan     Scans files in the provided directory recursively. Prints files info in a table format
  fix      Scans files in the provided directory recursively and fixes their extensions (when needed). Prints files info in a table format
  help     Shows this help
  version  Prints app version

Examples:
  gofixt --dir=~/images scan --exts=jpeg,png,webp
  gofixt --dir=~/files --silent fix
```

# License
Licensed under the MIT license.

# Author
Alexander Mac
