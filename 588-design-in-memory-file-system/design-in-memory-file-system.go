type File struct {
    children map[string]*File
    contents string
    isDirectory bool
}

type FileSystem struct {
    root *File
}


func Constructor() FileSystem {
    return FileSystem{ &File{children: make(map[string]*File), isDirectory: true} }
}


func (fs *FileSystem) Ls(path string) []string {
    file := fs.discover(path, false, false)
    if file == nil {
        return []string{}
    }
    if file.isDirectory {
        keys := make([]string, 0)
        for name, _ := range file.children {
            keys = append(keys,name)
        }
        slices.Sort(keys)
        return keys
    }
    split := strings.Split(path, "/")
    return split[len(split)-1:]
}


func (fs *FileSystem) Mkdir(path string)  {
    fs.discover(path, false, true)
}


func (fs *FileSystem) AddContentToFile(filePath string, content string)  {
    file := fs.discover(filePath, true, true)
    // add conditional - when returned is a directory
    file.isDirectory = false 
    file.contents += content
}


func (fs *FileSystem) ReadContentFromFile(filePath string) string {
    file := fs.discover(filePath, true, false)
    if file.isDirectory {
        // panic
    }
    return file.contents
}

func (fs *FileSystem) discover(filePath string, isFile, isCreate bool) *File {
    filePathSplits := strings.Split(filePath, "/")

    curr := fs.root
    for index, path := range filePathSplits {
        if path == "" {
            continue
        }
        if _, exists := curr.children[path] ; !exists {
            if isCreate {
                curr.children[path] = &File{children: make(map[string]*File), isDirectory: true}
            } else {
                return nil
            }
        }
        curr = curr.children[path]
        if index == len(filePathSplits)-1 && isFile {
            curr.isDirectory = false
        }
    }
    return curr
}
