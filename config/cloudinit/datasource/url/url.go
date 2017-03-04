// Copyright 2015 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package url

import (
	"github.com/rancher/os/config/cloudinit/datasource"
	"github.com/rancher/os/config/cloudinit/pkg"
)

type RemoteFile struct {
	url string
}

func NewDatasource(url string) *RemoteFile {
	return &RemoteFile{url}
}

func (f *RemoteFile) IsAvailable() bool {
	client := pkg.NewHTTPClient()
	_, err := client.Get(f.url)
	return (err == nil)
}

func (f *RemoteFile) AvailabilityChanges() bool {
	return true
}

func (f *RemoteFile) ConfigRoot() string {
	return ""
}

func (f *RemoteFile) FetchMetadata() (datasource.Metadata, error) {
	return datasource.Metadata{}, nil
}

func (f *RemoteFile) FetchUserdata() ([]byte, error) {
	client := pkg.NewHTTPClient()
	return client.GetRetry(f.url)
}

func (f *RemoteFile) Type() string {
	return "url"
}