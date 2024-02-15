type inode uint16

type FileSystem struct {
	root         *FileNode
	inodeCounter inode
	inodeDataMap map[inode][]byte
}

func (fs *FileSystem) incrementInode() {
	fs.inodeCounter = fs.inodeCounter + 1
}

type FileNode struct {
	name        string
	isDirectory bool
	inode       inode
	children    map[string]*FileNode
}

func newFileNode(root *FileSystem, name string, isDirectory bool) *FileNode {
	defer root.incrementInode()

	return &FileNode{
		name:        name,
		isDirectory: isDirectory,
		inode:       root.inodeCounter,
		children:    make(map[string]*FileNode),
	}
}

func Constructor() FileSystem {
    return FileSystem{
        root: &FileNode{
            name: "/",
            isDirectory: true,
            inode: 0,
            children: make(map[string]*FileNode),
        },
        inodeCounter: 1,
        inodeDataMap: make(map[inode][]byte),
    }
}

// If path is a file path, returns a list that only contains this file's name.
// If path is a directory path, returns the list of file and directory names in this directory.
func (fs *FileSystem) Ls(path string) []string {
    var allFiles []string
    node := fs.root
    for _, pathStr := range splitFilePath(path) {
        node, _ = node.children[pathStr]
    }
    if !node.isDirectory {
        return []string{node.name}
    } else {
        for childNodeName := range node.children {
            allFiles = append(allFiles, childNodeName)
        }
    }
    sort.Strings(allFiles)
    return allFiles
}

// Makes a new directory according to the given path. The given directory path does not exist. If the middle directories in the path do not exist, you should create them as well.
func (fs *FileSystem) Mkdir(path string)  {
    createFileIfNotExists(path, fs, true)
}

// If filePath does not exist, creates that file containing given content.
// If filePath already exists, appends the given content to original content.
func (fs *FileSystem) AddContentToFile(filePath string, content string)  {
    node := createFileIfNotExists(filePath, fs, false)
    fs.inodeDataMap[node.inode] = append(fs.inodeDataMap[node.inode], []byte(content)...)
}

// Returns the content in the file at filePath.
func (fs *FileSystem) ReadContentFromFile(filePath string) string {
    node := createFileIfNotExists(filePath, fs, false)
    return string(fs.inodeDataMap[node.inode])
}

func splitFilePath(path string) []string {
	if path == "/" {
		return []string{}
	}
	split := strings.Split(path, "/")
	return split[1:]
}

func createFileIfNotExists(path string, root *FileSystem, isDirectory bool) *FileNode {
	node := root.root
	for _, pathStr := range splitFilePath(path) {
		//if child is not found, then create a new file node
		if childNode, ok := node.children[pathStr]; !ok {
			//we need to create the file node here
			fileNode := newFileNode(root, pathStr, isDirectory)
			//tag this fileNode to the filesystem
			node.children[pathStr] = fileNode
			//the newly created node is used to recur
			node = fileNode
		} else {
			//recur for the already created child
			node = childNode
		}
	}
	return node
}


/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ls(path);
 * obj.Mkdir(path);
 * obj.AddContentToFile(filePath,content);
 * param_4 := obj.ReadContentFromFile(filePath);
 */