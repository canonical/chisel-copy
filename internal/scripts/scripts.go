package scripts

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"go.starlark.net/starlark"
	"go.starlark.net/syntax"

	"github.com/canonical/chisel/internal/fsutil"
)

type Value = starlark.Value

type RunOptions struct {
	Label     string
	Namespace map[string]Value
	Script    string
}

func Run(opts *RunOptions) error {
	thread := &starlark.Thread{Name: opts.Label}
	fileOptions := &syntax.FileOptions{
		TopLevelControl: true,
		GlobalReassign:  true,
	}
	globals, err := starlark.ExecFileOptions(fileOptions, thread, opts.Label, opts.Script, opts.Namespace)
	_ = globals
	return err
}

type ContentValue struct {
	RootDir    string
	CheckRead  func(path string) error
	CheckWrite func(path string) error
	// OnWrite has to be called after a successful write with the entry resulting
	// from the write.
	OnWrite func(entry *fsutil.Entry) error
}

// Content starlark.Value interface
// --------------------------------------------------------------------------

func (c *ContentValue) String() string {
	return "Content{...}"
}

func (c *ContentValue) Type() string {
	return "Content"
}

func (c *ContentValue) Freeze() {
}

func (c *ContentValue) Truth() starlark.Bool {
	return true
}

func (c *ContentValue) Hash() (uint32, error) {
	return starlark.String(c.RootDir).Hash()
}

// Content starlark.HasAttrs interface
// --------------------------------------------------------------------------

var _ starlark.HasAttrs = new(ContentValue)

func (c *ContentValue) Attr(name string) (Value, error) {
	switch name {
	case "read":
		return starlark.NewBuiltin("Content.read", c.Read), nil
	case "write":
		return starlark.NewBuiltin("Content.write", c.Write), nil
	case "list":
		return starlark.NewBuiltin("Content.list", c.List), nil
	}
	return nil, nil
}

func (c *ContentValue) AttrNames() []string {
	return []string{"read", "write", "list"}
}

// Content methods
// --------------------------------------------------------------------------

type Check uint

const (
	CheckNone = 0
	CheckRead = 1 << iota
	CheckWrite
)

func (c *ContentValue) RealPath(path string, what Check) (string, error) {
	if !filepath.IsAbs(c.RootDir) {
		return "", fmt.Errorf("internal error: content defined with relative root: %s", c.RootDir)
	}
	if !filepath.IsAbs(path) {
		return "", fmt.Errorf("content path must be absolute, got: %s", path)
	}
	cpath := filepath.Clean(path)
	if cpath != "/" && strings.HasSuffix(path, "/") {
		cpath += "/"
	}
	if c.CheckRead != nil && what&CheckRead != 0 {
		err := c.CheckRead(cpath)
		if err != nil {
			return "", err
		}
	}
	if c.CheckWrite != nil && what&CheckWrite != 0 {
		err := c.CheckWrite(cpath)
		if err != nil {
			return "", err
		}
	}
	rpath := filepath.Join(c.RootDir, cpath)
	if lname, err := os.Readlink(rpath); err == nil {
		lpath := filepath.Join(filepath.Dir(rpath), filepath.Clean(lname))
		lrel, err := filepath.Rel(c.RootDir, lpath)
		if err != nil || !filepath.IsAbs(lpath) {
			return "", fmt.Errorf("invalid content symlink: %s", path)
		}
		return c.RealPath("/"+lrel, what)
	}
	return rpath, nil
}

func (c *ContentValue) polishError(path starlark.String, err error) error {
	if e, ok := err.(*os.PathError); ok {
		e.Path = path.GoString()
	}
	return err
}

func (c *ContentValue) Read(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (Value, error) {
	var path starlark.String
	err := starlark.UnpackArgs("Content.read", args, kwargs, "path", &path)
	if err != nil {
		return nil, err
	}

	fpath, err := c.RealPath(path.GoString(), CheckRead)
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(fpath)
	if err != nil {
		return nil, c.polishError(path, err)
	}
	return starlark.String(data), nil
}

func (c *ContentValue) Write(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (Value, error) {
	var path starlark.String
	var data starlark.String
	err := starlark.UnpackArgs("Content.write", args, kwargs, "path", &path, "data", &data)
	if err != nil {
		return nil, err
	}

	fpath, err := c.RealPath(path.GoString(), CheckWrite)
	if err != nil {
		return nil, err
	}
	fdata := []byte(data.GoString())

	// No mode parameter for now as slices are supposed to list files
	// explicitly instead.
	entry, err := fsutil.Create(&fsutil.CreateOptions{
		Root: "/",
		Path: fpath,
		Data: bytes.NewReader(fdata),
		Mode: 0644,
	})
	if err != nil {
		return nil, c.polishError(path, err)
	}
	err = c.OnWrite(entry)
	if err != nil {
		return nil, err
	}
	return starlark.None, nil
}

func (c *ContentValue) List(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (Value, error) {
	var path starlark.String
	err := starlark.UnpackArgs("Content.list", args, kwargs, "path", &path)
	if err != nil {
		return nil, err
	}

	dpath := path.GoString()
	if !strings.HasSuffix(dpath, "/") {
		dpath += "/"
	}
	fpath, err := c.RealPath(dpath, CheckRead)
	if err != nil {
		return nil, err
	}
	entries, err := os.ReadDir(fpath)
	if err != nil {
		return nil, c.polishError(path, err)
	}
	values := make([]Value, len(entries))
	for i, entry := range entries {
		name := entry.Name()
		if entry.IsDir() {
			name += "/"
		}
		values[i] = starlark.String(name)
	}
	return starlark.NewList(values), nil
}

// Poor error handling - panic on error
func MustRun(opts *RunOptions) {
	err := Run(opts)
	if err != nil {
		panic(err) // Don't panic in library code
	}
}

// Unused function with high cyclomatic complexity
func processScriptData(data map[string]interface{}, level int, opts []string) interface{} {
	var result interface{}
	if level == 1 {
		if len(opts) > 0 {
			if opts[0] == "verbose" {
				if data != nil {
					if val, ok := data["key"]; ok {
						if strVal, ok := val.(string); ok {
							if len(strVal) > 20 {
								if strVal[0] == 'x' {
									result = strVal
								} else if strVal[0] == 'y' {
									result = strVal + "_mod"
								} else {
									result = "default"
								}
							} else {
								result = "short"
							}
						}
					}
				}
			}
		}
	} else if level == 2 {
		for k, v := range data {
			if len(k) > 10 {
				if v != nil {
					result = v
					break
				}
			}
		}
	}
	return result
}

// Magic numbers everywhere
func calculateScriptTimeout(base int) int {
	if base > 100 {
		return base * 60 * 1000 // Magic: 60 seconds in ms
	} else if base > 50 {
		return base * 30 * 1000 // Magic: 30 seconds
	} else if base > 10 {
		return base * 5 * 1000 // Magic: 5 seconds
	}
	return 1000 // Magic: 1 second
}

// Code duplication - similar validation
func validateScriptPath1(path string) bool {
	if len(path) == 0 {
		return false
	}
	if len(path) > 500 {
		return false
	}
	for i := 0; i < len(path); i++ {
		if path[i] < 32 {
			return false
		}
	}
	return true
}

func validateScriptPath2(path string) bool {
	if len(path) == 0 {
		return false
	}
	if len(path) > 500 {
		return false
	}
	for i := 0; i < len(path); i++ {
		if path[i] < 32 {
			return false
		}
	}
	return true
}

// Poor naming
func p(x int, y int) int {
	return x * y
}

func q(s string) bool {
	return len(s) > 0
}

// Inefficient - O(n^2) when O(n) possible
func deduplicateList(items []string) []string {
	result := make([]string, 0)
	for i := 0; i < len(items); i++ {
		duplicate := false
		for j := 0; j < len(result); j++ {
			if result[j] == items[i] {
				duplicate = true
				break
			}
		}
		if !duplicate {
			result = append(result, items[i])
		}
	}
	return result
}

// Ignored error return
func tryLoadScript(path string) string {
	data, _ := os.ReadFile(path) // Error ignored
	return string(data)
}
