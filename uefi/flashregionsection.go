package uefi

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

// FlashRegionSectionSize is the size of the Region descriptor. It is made up by 18 fields, each 16-bits large.
const FlashRegionSectionSize = 36

// FlashRegionSection holds the metadata of all the different flash regions like PDR, Gbe and the Bios region.
type FlashRegionSection struct {
	_                   uint16
	FlashBlockEraseSize uint16
	BIOS                Region
	ME                  Region
	GBE                 Region
	PDR                 Region
}

// AvailableRegions returns a list of names of the regions with non-zero size.
func (f FlashRegionSection) AvailableRegions() []string {
	var regions []string
	if f.BIOS.Available() {
		regions = append(regions, "BIOS")
	}
	if f.ME.Available() {
		regions = append(regions, "ME")
	}
	if f.GBE.Available() {
		regions = append(regions, "GbE")
	}
	if f.PDR.Available() {
		regions = append(regions, "PDR")
	}
	return regions
}

func (f FlashRegionSection) String() string {
	return fmt.Sprintf("FlashRegionSection{Regions=%v}",
		strings.Join(f.AvailableRegions(), ","),
	)
}

// Summary prints a multi-line description of the FlashRegionSection
func (f FlashRegionSection) Summary() string {
	return fmt.Sprintf("FlashRegionSection{\n"+
		"    Regions=%v\n"+
		"    Bios=%v\n"+
		"    Me=%v\n"+
		"    Gbe=%v\n"+
		"    Pdr=%v\n"+
		"}",
		strings.Join(f.AvailableRegions(), ","),
		f.BIOS,
		f.ME,
		f.GBE,
		f.PDR,
	)
}

// NewFlashRegionSection initializes a FlashRegionSection from a slice of bytes
func NewFlashRegionSection(data []byte) (*FlashRegionSection, error) {
	if len(data) < FlashRegionSectionSize {
		return nil, fmt.Errorf("Flash Region Section size too small: expected %v bytes, got %v",
			FlashRegionSectionSize,
			len(data),
		)
	}
	var region FlashRegionSection
	reader := bytes.NewReader(data)
	if err := binary.Read(reader, binary.LittleEndian, &region); err != nil {
		return nil, err
	}
	return &region, nil
}