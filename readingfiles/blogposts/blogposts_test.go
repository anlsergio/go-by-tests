package blogposts_test

import (
	"github.com/anlsergio/go-by-tests/readingfiles/blogposts"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: foo, bar
---
B
B
Q`
	)

	fs := fstest.MapFS{
		"hello world.md": {Data: []byte(firstBody)},
		"hello-world.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}

	if len(fs) != len(posts) {
		t.Errorf("want %d posts, got %d", len(fs), len(posts))
	}

	assertPost(t, blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
World`,
	}, posts[0])
}

func assertPost(t *testing.T, want blogposts.Post, got blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %+v, got %+v", want, got)
	}
}
