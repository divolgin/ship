package filetree

import (
	"os"
	"path"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

// A Loader returns a struct representation
// of a filesystem directory tree
type Loader interface {
	LoadTree(root string) (*Node, error)
}

// NewLoader builds an aferoLoader, used with dig
func NewLoader(
	fs afero.Afero,
	logger log.Logger,
) Loader {
	return &aferoLoader{
		FS:     fs,
		Logger: logger,
	}
}

type aferoLoader struct {
	Logger log.Logger
	FS     afero.Afero
}

func (a *aferoLoader) LoadTree(root string) (*Node, error) {

	fs := afero.Afero{Fs: afero.NewBasePathFs(a.FS, root)}

	files, err := fs.ReadDir("/")
	if err != nil {
		return nil, errors.Wrapf(err, "read dir %q", root)
	}

	rootNode := Node{
		Path:     "/",
		Name:     "/",
		Children: []Node{},
	}
	populated, err := a.loadTree(rootNode, files)

	return &populated, errors.Wrap(err, "load tree")
}

func (a *aferoLoader) loadTree(current Node, files []os.FileInfo) (Node, error) {
	if len(files) == 0 {
		return current, nil
	}

	file, rest := files[0], files[1:]
	filePath := path.Join(current.Path, file.Name())

	// no thanks
	if isSymlink(file) {
		level.Debug(a.Logger).Log("event", "symlink.skip", "file", filePath)
		return a.loadTree(current, rest)
	}

	if !file.IsDir() {
		return a.loadTree(current.withChild(Node{
			Name: file.Name(),
			Path: filePath,
		}), rest)
	}

	subFiles, err := a.FS.ReadDir(filePath)
	if err != nil {
		return current, errors.Wrapf(err, "read dir %q", file.Name())
	}

	subTree := Node{
		Name:     file.Name(),
		Path:     filePath,
		Children: []Node{},
	}

	subTreeLoaded, err := a.loadTree(subTree, subFiles)
	if err != nil {
		return current, errors.Wrapf(err, "load tree %q", file.Name())
	}

	return a.loadTree(current.withChild(subTreeLoaded), rest)
}

func isSymlink(file os.FileInfo) bool {
	return file.Mode()&os.ModeSymlink != 0
}

func (n Node) withChild(child Node) Node {
	return Node{
		Name:     n.Name,
		Path:     n.Path,
		Children: append(n.Children, child),
	}
}