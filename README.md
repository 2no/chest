# chest

The Easy Metafile Manager

## Example
```
$ ls
index.js package.json metafile_a metafile_b

$ chest put metafile_a metafile_b

$ ls -la
.chest index.js package.json

$ chest list
metafile_a metafile_b

$ chest open
.chest index.js package.json metafile_a metafile_b
```

## Usage
```
chest usage       Show this message
chest put <file>  Put specified config file in the chest
chest take <file> Take specified config file from the chest into the project root
chest list        Show list of config files in the chest
chest open        Place config files from the chest into the project root as symbolic links
chest close       Remove symbolic links created by "open"
```

## Install
```
$ go get github.com/wakuworks/chest
```

## License
MIT
