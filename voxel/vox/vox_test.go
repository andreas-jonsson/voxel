/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package vox

import (
	"image"
	"image/color"
	"image/color/palette"
	"os"
	"testing"

	"github.com/andreas-jonsson/voxel/voxel"
)

type voxelImage struct {
	data []*image.Paletted
}

func (img *voxelImage) SetBounds(b voxel.Box) {
	rect := image.Rect(0, 0, b.Max.X, b.Max.Y)
	img.data = make([]*image.Paletted, b.Max.Z)

	for i := 0; i < b.Max.Z; i++ {
		img.data[i] = image.NewPaletted(rect, palette.Plan9)
	}
}

func (img *voxelImage) SetPalette(pal color.Palette) {
	for _, layer := range img.data {
		layer.Palette = pal
	}
}

func (img *voxelImage) Set(x, y, z int, index uint8) {
	img.data[z].SetColorIndex(x, y, index)
}

func TestVox(t *testing.T) {
	if fp, err := os.Open("test.vox"); err == nil {
		defer fp.Close()

		var img voxelImage
		if err := Decode(fp, &img); err != nil {
			t.Error(err)
		}
	} else {
		t.Error(err)
	}
}
