type FileNode struct {
    children map[string]*FileNode // Children could be both files and directories
    content  string               // Non-empty for files
    isFile   bool
}

type FileSystem struct {
    root *FileNode
}

func Constructor() FileSystem {
    return FileSystem{root: &FileNode{children: make(map[string]*FileNode)}}
}

func (this *FileSystem) Ls(path string) []string {
    node := this.navigate(path)
    if node.isFile {
        // Extract file name from path
        splitPath := strings.Split(path, "/")
        return []string{splitPath[len(splitPath)-1]}
    }
    // Collect and return directory contents
    var contents []string
    for name := range node.children {
        contents = append(contents, name)
    }
    sort.Strings(contents) // Sort results lexicographically
    return contents
}

func (this *FileSystem) Mkdir(path string) {
    this.navigate(path) // Simply navigate to the path, nodes will be created if not exist
}

func (this *FileSystem) AddContentToFile(filePath string, content string) {
    node := this.navigate(filePath)
    node.content += content // Append content
    node.isFile = true      // Mark as file
}

func (this *FileSystem) ReadContentFromFile(filePath string) string {
    node := this.navigate(filePath)
    return node.content
}

// Helper function to navigate through the file system and optionally create nodes
func (this *FileSystem) navigate(path string) *FileNode {
    current := this.root
    if path == "/" {
        return current
    }
    parts := strings.Split(path, "/")[1:] // Ignore the first empty part due to leading '/'
    for _, part := range parts {
        if _, exists := current.children[part]; !exists {
            current.children[part] = &FileNode{children: make(map[string]*FileNode)}
        }
        current = current.children[part]
    }
    return current
}


/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ls(path);
 * obj.Mkdir(path);
 * obj.AddContentToFile(filePath,content);
 * param_4 := obj.ReadContentFromFile(filePath);
 */