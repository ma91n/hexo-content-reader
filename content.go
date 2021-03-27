package hexoreader

type Post struct {
	FrontMatter
	Content string
}

type FrontMatter struct {
	Title      string   `yaml:"title"`
	Date       string   `yaml:"date"`
	Tags       []string `yaml:"tags"`
	Categories []string `yaml:"category"`
}
