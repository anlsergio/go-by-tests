package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

const (
	titleTag       = "Title: "
	descriptionTag = "Description: "
	tagsTag        = "Tags: "
	tagsSeparator  = ", "
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	return Post{
		Title:       readMetaLine(titleTag),
		Description: readMetaLine(descriptionTag),
		Tags:        strings.Split(readMetaLine(tagsTag), tagsSeparator),
		Body:        readBody(scanner),
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() //ignore a line (---)
	buffer := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buffer, scanner.Text())
	}
	body := strings.TrimSuffix(buffer.String(), "\n")

	return body
}
