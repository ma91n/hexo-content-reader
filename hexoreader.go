package hexoreader

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/goccy/go-yaml"
)

func New(r io.Reader) HexoReader {
	return HexoReader{
		r: r,
	}
}

type HexoReader struct {
	r io.Reader
}

func (h HexoReader) ReadAll() (Post, error) {
	all, err := ioutil.ReadAll(h.r)
	if err != nil {
		return Post{}, err
	}

	splits := strings.SplitN(string(all), "---", 2)
	if len(splits) < 2 {
		return Post{}, errors.New("not found front matter")
	}

	frontMatterBody := splits[0]
	contentBody := strings.TrimLeft(strings.TrimLeft(splits[1], "\r"), "\n")

	var fm FrontMatter
	if err := yaml.Unmarshal([]byte(frontMatterBody), &fm); err != nil {
		return Post{}, fmt.Errorf("front-matter unmarshal: %w", err)
	}

	return Post{
		FrontMatter: fm,
		Content:     contentBody,
	}, nil

}
