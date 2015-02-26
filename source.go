/*
Package source allows you to access the files in your source tree.
*/
package source

import (
	"path"
	"runtime"
)

// Path returns the path you passed as relPath as an absolute path relative to
// your source folder, assuming you built the source on the computer you use the
// binary on.
func Path(relPath ...string) string {
	_, file, _, _ := runtime.Caller(1)

	pathSlice := append([]string{path.Dir(file)}, relPath...)

	return path.Join(pathSlice...)
}
