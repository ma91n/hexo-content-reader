# hexo-content-reader
Hexo content reader for go

## Usage

```sh
go get github.com/ma91n/hexoreader
```

```go
	file, err := os.ReadFile(path)

	post, err := hexoreader.New(bytes.NewReader(file)).ReadAll()
```
