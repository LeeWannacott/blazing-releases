package main

import (
	"os"
	"os/user"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestGetChangeLogPath(t *testing.T) {
	currentUser, err := user.Current()
	if err != nil {
		t.Fatal(err)
	}
	homeDir := currentUser.HomeDir
	expected := filepath.Join(homeDir, "Desktop/quick-release-notes/docs/CHANGELOG.md")
	currentDir, _ := os.Getwd()
	path := filepath.Join(currentDir, "docs/CHANGELOG.md")
	if !strings.HasSuffix(path, expected) {
		t.Errorf("getChangeLogPath() = %v; want %v", path, expected)
	}
}

func TestGetTagsFromGitHub(t *testing.T) {
	tags, err := getTagsFromGitHub("quick-lint/quick-lint-js")
	if err != nil {
		t.Errorf("getTagsFromGitHub() = %v; want nil", err)
	}
	expectedTags := []string{"2.12.0", "2.11.0", "2.10.0", "2.9.0", "2.8.0", "2.7.0", "2.6.0", "2.5.0", "2.4.2", "2.4.1", "2.4.0", "2.3.1", "2.3.0", "2.2.0", "2.1.0", "2.0.0", "1.0.0", "0.7.1", "0.7.0", "0.6.0", "0.5.0", "0.4.0", "0.3.0", "0.2.0"}
	gotTags := make([]string, len(tags))
	for i, tag := range tags {
		gotTags[i] = tag.Name
	}
	if !reflect.DeepEqual(gotTags, expectedTags) {
		t.Errorf("unexpected tags: got %v, want %v", gotTags, expectedTags)
	}
}
