// Copyright © 2021 Alibaba Group Holding Ltd.
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

// This file is extracted from "github.com/docker/docker/pkg/system"

package utils

import "golang.org/x/sys/unix"

// Lgetxattr retrieves the value of the extended attribute identified by attr
// and associated with the given path in the file system.
// It will return a nil slice and nil error if the xattr is not set.
func Lgetxattr(path string, attr string) ([]byte, error) {
	// Start with a 128 length byte array
	dest := make([]byte, 128)
	sz, errno := unix.Lgetxattr(path, attr, dest)

	switch {
	case errno == unix.ENODATA:
		return nil, nil
	case errno == unix.ERANGE:
		// 128 byte array might just not be good enough. A dummy buffer is used
		// to get the real size of the xattrs on disk
		sz, errno = unix.Lgetxattr(path, attr, []byte{})
		if errno != nil {
			return nil, errno
		}
		dest = make([]byte, sz)
		sz, errno = unix.Lgetxattr(path, attr, dest)
		if errno != nil {
			return nil, errno
		}
	case errno != nil:
		return nil, errno
	}
	return dest[:sz], nil
}

// Lsetxattr sets the value of the extended attribute identified by attr
// and associated with the given path in the file system.
func Lsetxattr(path string, attr string, data []byte, flags int) error {
	return unix.Lsetxattr(path, attr, data, flags)
}
