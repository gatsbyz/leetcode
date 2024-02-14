type FileSystem struct {
    root *FileNode
}

type FileNode struct {
    children map[string]*FileNode
    value int
}


func Constructor() FileSystem {
    return FileSystem{
        root: &FileNode{
            children: make(map[string]*FileNode),
            value: 0,
        },
    }
}


func (fs *FileSystem) CreatePath(path string, value int) bool {
    pathSplit := strings.Split(path, "/")
    curr := fs.root
    for _, path := range pathSplit[1:len(pathSplit)-1] {
        if _, exists := curr.children[path] ; !exists {
            return false
        }
        fmt.Println(value)
        curr = curr.children[path]
    }
    if _, exists := curr.children[pathSplit[len(pathSplit)-1]] ; exists {
        return false
    }
    curr.children[pathSplit[len(pathSplit)-1]]  = &FileNode{
        children: make(map[string]*FileNode),
        value: value,
    }
    return true
}


func (fs *FileSystem) Get(path string) int {
    pathSplit := strings.Split(path, "/")
    curr := fs.root
    for _, path := range pathSplit[1:len(pathSplit)-1] {
        if _, exists := curr.children[path] ; !exists {
            return -1
        }
        curr = curr.children[path]
    }
    if _, exists := curr.children[pathSplit[len(pathSplit)-1]] ; !exists {
        return -1
    }

    return curr.children[pathSplit[len(pathSplit)-1]].value
}


/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.CreatePath(path,value);
 * param_2 := obj.Get(path);
 */