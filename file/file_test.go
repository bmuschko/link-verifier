package file_test

import (
	"fmt"
	. "github.com/bmuschko/link-verifier/file"
	. "github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestFindAsciiDocFilesInRootDir(t *testing.T) {
	tempDirPath := filepath.Join(os.TempDir(), "a")
	createDir(tempDirPath)
	path1 := filepath.Join(tempDirPath, fmt.Sprintf("1.adoc"))
	path2 := filepath.Join(tempDirPath, fmt.Sprintf("abc.adoc"))
	jpgPath := filepath.Join(tempDirPath, "my.jpg")
	binPath := filepath.Join(tempDirPath, "other.bin")
	docPath := filepath.Join(tempDirPath, "some.doc")
	createFile(path1)
	createFile(path2)
	createFile(jpgPath)
	createFile(binPath)
	createFile(docPath)

	files := FindTextBasedFiles(tempDirPath, []string{"*.adoc"})
	Equal(t, 2, len(files))
	Equal(t, path1, files[0])
	Equal(t, path2, files[1])

	deleteFile(path1)
	deleteFile(path2)
	deleteFile(jpgPath)
	deleteFile(binPath)
	deleteFile(docPath)
}

func TestFindAsciiDocFilesInSubDirs(t *testing.T) {
	tempDirPath := filepath.Join(os.TempDir(), "b")
	createDir(tempDirPath)
	subDirPath := filepath.Join(tempDirPath, "sub")
	createDir(subDirPath)
	subSubDirPath := filepath.Join(subDirPath, "subsub")
	createDir(subSubDirPath)
	path1 := filepath.Join(subDirPath, fmt.Sprintf("1.adoc"))
	path2 := filepath.Join(subSubDirPath, fmt.Sprintf("2.adoc"))
	createFile(path1)
	createFile(path2)

	files := FindTextBasedFiles(tempDirPath, []string{"*.adoc"})
	Equal(t, 2, len(files))
	Equal(t, path1, files[0])
	Equal(t, path2, files[1])

	deleteFile(path1)
	deleteFile(path2)
}

func TestFindAsciiDocFilesDifferentExtensions(t *testing.T) {
	tempDirPath := filepath.Join(os.TempDir(), "c")
	createDir(tempDirPath)
	path1 := filepath.Join(tempDirPath, fmt.Sprintf("1.adoc"))
	path2 := filepath.Join(tempDirPath, fmt.Sprintf("2.asciidoc"))
	path3 := filepath.Join(tempDirPath, fmt.Sprintf("3.asc"))
	createFile(path1)
	createFile(path2)
	createFile(path3)

	files := FindTextBasedFiles(tempDirPath, []string{"*.adoc", "*.asciidoc", "*.asc"})
	Equal(t, 3, len(files))
	Equal(t, path1, files[0])
	Equal(t, path2, files[1])
	Equal(t, path3, files[2])

	deleteFile(path1)
	deleteFile(path2)
	deleteFile(path3)
}

func TestFilesForCustomIncludePatterns(t *testing.T) {
	tempDirPath := filepath.Join(os.TempDir(), "e")
	createDir(tempDirPath)
	path1 := filepath.Join(tempDirPath, fmt.Sprintf("1.html"))
	path2 := filepath.Join(tempDirPath, fmt.Sprintf("2.yml"))
	createFile(path1)
	createFile(path2)

	files := FindTextBasedFiles(tempDirPath, []string{"*.html", "*.yml"})
	Equal(t, 2, len(files))
	Equal(t, path1, files[0])
	Equal(t, path2, files[1])

	deleteFile(path1)
	deleteFile(path2)
}

func TestReadFile(t *testing.T) {
	expectedContent := "some text"
	tempDirPath := filepath.Join(os.TempDir(), "content")
	createDir(tempDirPath)
	path1 := filepath.Join(tempDirPath, fmt.Sprintf("1.adoc"))
	createFile(path1)
	writeFile(path1, expectedContent)

	content := ReadFile(path1)
	Equal(t, expectedContent, content)

	deleteFile(path1)
}

func TestPanicWhenReadingNonExistentFile(t *testing.T) {
	path := filepath.Join("1.adoc")
	Panics(t, func() {
		ReadFile(path)
	})
}

func createDir(path string) {
	err := os.MkdirAll(path, 0755)

	if err != nil {
		panic(err)
	}
}

func createFile(path string) {
	w, err := os.Create(path)

	if err != nil {
		panic(err)
	}

	defer w.Close()
}

func writeFile(path string, content string) {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, err = file.WriteString(content)

	if err != nil {
		panic(err)
	}

	err = file.Sync()

	if err != nil {
		panic(err)
	}
}

func deleteFile(path string) {
	var err = os.Remove(path)

	if err != nil {
		panic(err)
	}
}
