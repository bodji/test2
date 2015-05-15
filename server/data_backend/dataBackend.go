/**

    Plik upload server

The MIT License (MIT)

Copyright (c) <2015>
	- Mathieu Bodjikian <mathieu@bodjikian.fr>
	- Charles-Antoine Mathieu <skatkatt@root.gg>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
**/

package data_backend

import (
	"io"

	"github.com/root-gg/plik/server/common"
	"github.com/root-gg/plik/server/data_backend/file"
	"github.com/root-gg/plik/server/data_backend/swift"
	"github.com/root-gg/plik/server/data_backend/weedfs"
)

var dataBackend DataBackend

// DataBackend interface describes methods that data backends
// must implements to be compatible with plik.
type DataBackend interface {
	GetFile(ctx *common.PlikContext, u *common.Upload, id string) (rc io.ReadCloser, err error)
	AddFile(ctx *common.PlikContext, u *common.Upload, file *common.File, fileReader io.Reader) (backendDetails map[string]interface{}, err error)
	RemoveFile(ctx *common.PlikContext, u *common.Upload, id string) (err error)
	RemoveUpload(ctx *common.PlikContext, u *common.Upload) (err error)
}

// GetDataBackend is a singleton pattern.
// Init static backend if not already and return it
func GetDataBackend() DataBackend {
	if dataBackend == nil {
		Initialize()
	}
	return dataBackend
}

// Initialize backend from type found in configuration
func Initialize() {
	if dataBackend == nil {
		switch common.Config.DataBackend {
		case "file":
			dataBackend = file.NewFileBackend(common.Config.DataBackendConfig)
		case "swift":
			dataBackend = swift.NewSwiftBackend(common.Config.DataBackendConfig)
		case "weedfs":
			dataBackend = weedfs.NewWeedFsBackend(common.Config.DataBackendConfig)
		default:
			common.Log().Fatalf("Invalid data backend %s", common.Config.DataBackend)
		}
	}
}
