// Copyright 2017-2023 the LinuxBoot Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bgkey

import (
	"testing"

	"github.com/linuxboot/fiano/pkg/intel/metadata/common/unittest"
)

func TestReadWrite(t *testing.T) {
	unittest.BGManifestReadWrite(t, &Manifest{}, "testdata/km.bin")
}
