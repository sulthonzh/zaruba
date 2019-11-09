package organizer

import (
	"io/ioutil"
	"path"
	"strings"
	"testing"
	"time"
)

func TestOrganize(t *testing.T) {
	project := path.Join("..", "playground", "projects", "test-organize")
	// start watcher and wait for a while
	Organize(project)
	time.Sleep(1 * time.Second)

	// inspect repos/classifiers/pre.txt
	expectedPreContent := "pre"
	preContent, err := readGeneratedFile(project, path.Join("repos", "classifiers", "pre.txt"))
	if err != nil {
		t.Error(err)
	} else if strings.Trim(expectedPreContent, "\n") != strings.Trim(preContent, "\n") {
		t.Errorf("Expected:\n%s\nActual:\n%s", expectedPreContent, preContent)
	}

	// inspect repos/classifiers/post.txt
	expectedPostContent := "post"
	postContent, err := readGeneratedFile(project, path.Join("repos", "classifiers", "post.txt"))
	if err != nil {
		t.Error(err)
	} else if strings.Trim(expectedPostContent, "\n") != strings.Trim(postContent, "\n") {
		t.Errorf("Expected:\n%s\nActual:\n%s", expectedPostContent, postContent)
	}

	// inspect services/sentiment/classifiers/classifier.py
	_, err = readGeneratedFile(project, path.Join("services", "sentiment", "classifiers", "classifier.py"))
	if err != nil {
		t.Error(err)
	}

	// inspect services/sentiment/post.txt
	_, err = readGeneratedFile(project, path.Join("services", "sentiment", "post.txt"))
	if err != nil {
		t.Error(err)
	}

	// inspect services/image/classifiers/classifier.py
	_, err = readGeneratedFile(project, path.Join("services", "image", "classifiers", "classifier.py"))
	if err != nil {
		t.Error(err)
	}

	// inspect services/image/post.txt
	_, err = readGeneratedFile(project, path.Join("services", "image", "post.txt"))
	if err != nil {
		t.Error(err)
	}

}

func readGeneratedFile(project, filepath string) (string, error) {
	data, err := ioutil.ReadFile(path.Join(project, filepath))
	return string(data), err
}