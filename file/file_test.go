package file

import (
    "testing"
    "os"
    "path/filepath"
    "github.com/stretchr/testify/assert"
)

func TestFindAsciiDocFilesInRootDir(t *testing.T) {
    tempDirPath := filepath.Join(os.TempDir(), "a")
    createDir(tempDirPath)
    adocPath1 := filepath.Join(tempDirPath, "1.adoc")
    adocPath2 := filepath.Join(tempDirPath, "abc.adoc")
    txtPath := filepath.Join(tempDirPath, "my.txt")
    binPath := filepath.Join(tempDirPath, "other.bin")
    docPath := filepath.Join(tempDirPath, "some.doc")
    createFile(adocPath1)
    createFile(adocPath2)
    createFile(txtPath)
    createFile(binPath)
    createFile(docPath)

    adocFiles := FindAsciiDocFiles(tempDirPath)
    assert.Equal(t, 2, len(adocFiles))
    assert.Equal(t, adocPath1, adocFiles[0])
    assert.Equal(t, adocPath2, adocFiles[1])

    deleteFile(adocPath1)
    deleteFile(adocPath2)
    deleteFile(txtPath)
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
    adocPath1 := filepath.Join(subDirPath, "1.adoc")
    adocPath2 := filepath.Join(subSubDirPath, "2.adoc")
    createFile(adocPath1)
    createFile(adocPath2)

    adocFiles := FindAsciiDocFiles(tempDirPath)
    assert.Equal(t, 2, len(adocFiles))
    assert.Equal(t, adocPath1, adocFiles[0])
    assert.Equal(t, adocPath2, adocFiles[1])

    deleteFile(adocPath1)
    deleteFile(adocPath2)
}

func TestFindAsciiDocFilesDifferentExtensions(t *testing.T) {
    tempDirPath := filepath.Join(os.TempDir(), "c")
    createDir(tempDirPath)
    adocPath1 := filepath.Join(tempDirPath, "1.adoc")
    adocPath2 := filepath.Join(tempDirPath, "2.asciidoc")
    adocPath3 := filepath.Join(tempDirPath, "3.asc")
    createFile(adocPath1)
    createFile(adocPath2)
    createFile(adocPath3)

    adocFiles := FindAsciiDocFiles(tempDirPath)
    assert.Equal(t, 3, len(adocFiles))
    assert.Equal(t, adocPath1, adocFiles[0])
    assert.Equal(t, adocPath2, adocFiles[1])
    assert.Equal(t, adocPath3, adocFiles[2])

    deleteFile(adocPath1)
    deleteFile(adocPath2)
    deleteFile(adocPath3)
}

func TestReadFile(t *testing.T) {
    expectedContent := "some text"
    tempDirPath := filepath.Join(os.TempDir(), "content")
    createDir(tempDirPath)
    adocPath1 := filepath.Join(tempDirPath, "1.adoc")
    createFile(adocPath1)
    writeFile(adocPath1, expectedContent)

    content := ReadFile(adocPath1)
    assert.Equal(t, expectedContent, content)

    deleteFile(adocPath1)
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
