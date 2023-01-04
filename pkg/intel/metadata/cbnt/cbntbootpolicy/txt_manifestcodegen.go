// Copyright 2017-2021 the LinuxBoot Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !manifestcodegen
// +build !manifestcodegen

// Code generated by "menifestcodegen". DO NOT EDIT.
// To reproduce: go run github.com/linuxboot/fiano/pkg/intel/metadata/common/manifestcodegen/cmd/manifestcodegen github.com/linuxboot/fiano/pkg/intel/metadata/cbnt/cbntbootpolicy

package cbntbootpolicy

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"

	"github.com/linuxboot/fiano/pkg/intel/metadata/cbnt"
	"github.com/linuxboot/fiano/pkg/intel/metadata/common/pretty"
)

var (
	// Just to avoid errors in "import" above in case if it wasn't used below
	_ = binary.LittleEndian
	_ = (fmt.Stringer)(nil)
	_ = (io.Reader)(nil)
	_ = pretty.Header
	_ = strings.Join
	_ = cbnt.StructInfo{}
)

// NewTXT returns a new instance of TXT with
// all default values set.
func NewTXT() *TXT {
	s := &TXT{}
	copy(s.StructInfo.ID[:], []byte(StructureIDTXT))
	s.StructInfo.Version = 0x21
	// Set through tag "default":
	s.SInitMinSVNAuth = 0
	// Set through tag "default":
	s.PTTCMOSOffset0 = 126
	// Set through tag "default":
	s.PTTCMOSOffset1 = 127
	// Set through tag "default":
	s.ACPIBaseOffset = 0x400
	// Set through tag "default":
	s.PwrMBaseOffset = 0xFE000000
	// Recursively initializing a child structure:
	s.DigestList = *cbnt.NewHashList()
	// Set through tag "required":
	s.SegmentCount = 0
	s.Rehash()
	return s
}

// Validate (recursively) checks the structure if there are any unexpected
// values. It returns an error if so.
func (s *TXT) Validate() error {
	// See tag "require"
	for idx := range s.Reserved0 {
		if s.Reserved0[idx] != 0 {
			return fmt.Errorf("'Reserved0[%d]' is expected to be 0, but it is %v", idx, s.Reserved0[idx])
		}
	}
	// See tag "require"
	for idx := range s.SetNumber {
		if s.SetNumber[idx] != 0 {
			return fmt.Errorf("'SetNumber[%d]' is expected to be 0, but it is %v", idx, s.SetNumber[idx])
		}
	}
	// See tag "require"
	for idx := range s.Reserved1 {
		if s.Reserved1[idx] != 0 {
			return fmt.Errorf("'Reserved1[%d]' is expected to be 0, but it is %v", idx, s.Reserved1[idx])
		}
	}
	// Recursively validating a child structure:
	if err := s.DigestList.Validate(); err != nil {
		return fmt.Errorf("error on field 'DigestList': %w", err)
	}
	// See tag "require"
	for idx := range s.Reserved3 {
		if s.Reserved3[idx] != 0 {
			return fmt.Errorf("'Reserved3[%d]' is expected to be 0, but it is %v", idx, s.Reserved3[idx])
		}
	}
	// See tag "require"
	if s.SegmentCount != 0 {
		return fmt.Errorf("field 'SegmentCount' expects value '0', but has %v", s.SegmentCount)
	}

	return nil
}

// StructureIDTXT is the StructureID (in terms of
// the document #575623) of element 'TXT'.
const StructureIDTXT = "__TXTS__"

// GetStructInfo returns current value of StructInfo of the structure.
//
// StructInfo is a set of standard fields with presented in any element
// ("element" in terms of document #575623).
func (s *TXT) GetStructInfo() cbnt.StructInfo {
	return s.StructInfo
}

// SetStructInfo sets new value of StructInfo to the structure.
//
// StructInfo is a set of standard fields with presented in any element
// ("element" in terms of document #575623).
func (s *TXT) SetStructInfo(newStructInfo cbnt.StructInfo) {
	s.StructInfo = newStructInfo
}

// ReadFrom reads the TXT from 'r' in format defined in the document #575623.
func (s *TXT) ReadFrom(r io.Reader) (int64, error) {
	var totalN int64

	err := binary.Read(r, binary.LittleEndian, &s.StructInfo)
	if err != nil {
		return totalN, fmt.Errorf("unable to read structure info at %d: %w", totalN, err)
	}
	totalN += int64(binary.Size(s.StructInfo))

	n, err := s.ReadDataFrom(r)
	if err != nil {
		return totalN, fmt.Errorf("unable to read data: %w", err)
	}
	totalN += n

	return totalN, nil
}

// ReadDataFrom reads the TXT from 'r' excluding StructInfo,
// in format defined in the document #575623.
func (s *TXT) ReadDataFrom(r io.Reader) (int64, error) {
	totalN := int64(0)

	// StructInfo (ManifestFieldType: structInfo)
	{
		// ReadDataFrom does not read Struct, use ReadFrom for that.
	}

	// Reserved0 (ManifestFieldType: arrayStatic)
	{
		n, err := 1, binary.Read(r, binary.LittleEndian, s.Reserved0[:])
		if err != nil {
			return totalN, fmt.Errorf("unable to read field 'Reserved0': %w", err)
		}
		totalN += int64(n)
	}

	// SetNumber (ManifestFieldType: arrayStatic)
	{
		n, err := 1, binary.Read(r, binary.LittleEndian, s.SetNumber[:])
		if err != nil {
			return totalN, fmt.Errorf("unable to read field 'SetNumber': %w", err)
		}
		totalN += int64(n)
	}

	// SInitMinSVNAuth (ManifestFieldType: endValue)
	{
		n, err := 1, binary.Read(r, binary.LittleEndian, &s.SInitMinSVNAuth)
		if err != nil {
			return totalN, fmt.Errorf("unable to read field 'SInitMinSVNAuth': %w", err)
		}
		totalN += int64(n)
	}

	// Reserved1 (ManifestFieldType: arrayStatic)
	{
		n, err := 1, binary.Read(r, binary.LittleEndian, s.Reserved1[:])
		if err != nil {
			return totalN, fmt.Errorf("unable to read field 'Reserved1': %w", err)
		}
		totalN += int64(n)
	}

	// ControlFlags (ManifestFieldType: endValue)
	{
		n, err := 4, binary.Read(r, binary.LittleEndian, &s.ControlFlags)
		if err != nil {
			return totalN, fmt.Errorf("unable to read field 'ControlFlags': %w", err)
		}
		totalN += int64(n)
	}

	// PwrDownInterval (ManifestFieldType: endValue)
	{
		n, err := 2, binary.Read(r, binary.LittleEndian, &s.PwrDownInterval)
		if err != nil {
			return totalN, fmt.Errorf("unable to read field 'PwrDownInterval': %w", err)
		}
		totalN += int64(n)
	}

	// PTTCMOSOffset0 (ManifestFieldType: endValue)
	{
		n, err := 1, binary.Read(r, binary.LittleEndian, &s.PTTCMOSOffset0)
		if err != nil {
			return totalN, fmt.Errorf("unable to read field 'PTTCMOSOffset0': %w", err)
		}
		totalN += int64(n)
	}

	// PTTCMOSOffset1 (ManifestFieldType: endValue)
	{
		n, err := 1, binary.Read(r, binary.LittleEndian, &s.PTTCMOSOffset1)
		if err != nil {
			return totalN, fmt.Errorf("unable to read field 'PTTCMOSOffset1': %w", err)
		}
		totalN += int64(n)
	}

	// ACPIBaseOffset (ManifestFieldType: endValue)
	{
		n, err := 2, binary.Read(r, binary.LittleEndian, &s.ACPIBaseOffset)
		if err != nil {
			return totalN, fmt.Errorf("unable to read field 'ACPIBaseOffset': %w", err)
		}
		totalN += int64(n)
	}

	// Reserved2 (ManifestFieldType: arrayStatic)
	{
		n, err := 2, binary.Read(r, binary.LittleEndian, s.Reserved2[:])
		if err != nil {
			return totalN, fmt.Errorf("unable to read field 'Reserved2': %w", err)
		}
		totalN += int64(n)
	}

	// PwrMBaseOffset (ManifestFieldType: endValue)
	{
		n, err := 4, binary.Read(r, binary.LittleEndian, &s.PwrMBaseOffset)
		if err != nil {
			return totalN, fmt.Errorf("unable to read field 'PwrMBaseOffset': %w", err)
		}
		totalN += int64(n)
	}

	// DigestList (ManifestFieldType: subStruct)
	{
		n, err := s.DigestList.ReadFrom(r)
		if err != nil {
			return totalN, fmt.Errorf("unable to read field 'DigestList': %w", err)
		}
		totalN += int64(n)
	}

	// Reserved3 (ManifestFieldType: arrayStatic)
	{
		n, err := 3, binary.Read(r, binary.LittleEndian, s.Reserved3[:])
		if err != nil {
			return totalN, fmt.Errorf("unable to read field 'Reserved3': %w", err)
		}
		totalN += int64(n)
	}

	// SegmentCount (ManifestFieldType: endValue)
	{
		n, err := 1, binary.Read(r, binary.LittleEndian, &s.SegmentCount)
		if err != nil {
			return totalN, fmt.Errorf("unable to read field 'SegmentCount': %w", err)
		}
		totalN += int64(n)
	}

	return totalN, nil
}

// RehashRecursive calls Rehash (see below) recursively.
func (s *TXT) RehashRecursive() {
	s.StructInfo.Rehash()
	s.DigestList.Rehash()
	s.Rehash()
}

// Rehash sets values which are calculated automatically depending on the rest
// data. It is usually about the total size field of an element.
func (s *TXT) Rehash() {
	s.Variable0 = 0
	s.ElementSize = uint16(s.TotalSize())
}

// WriteTo writes the TXT into 'w' in format defined in
// the document #575623.
func (s *TXT) WriteTo(w io.Writer) (int64, error) {
	totalN := int64(0)
	s.Rehash()

	// StructInfo (ManifestFieldType: structInfo)
	{
		n, err := s.StructInfo.WriteTo(w)
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'StructInfo': %w", err)
		}
		totalN += int64(n)
	}

	// Reserved0 (ManifestFieldType: arrayStatic)
	{
		n, err := 1, binary.Write(w, binary.LittleEndian, s.Reserved0[:])
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'Reserved0': %w", err)
		}
		totalN += int64(n)
	}

	// SetNumber (ManifestFieldType: arrayStatic)
	{
		n, err := 1, binary.Write(w, binary.LittleEndian, s.SetNumber[:])
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'SetNumber': %w", err)
		}
		totalN += int64(n)
	}

	// SInitMinSVNAuth (ManifestFieldType: endValue)
	{
		n, err := 1, binary.Write(w, binary.LittleEndian, &s.SInitMinSVNAuth)
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'SInitMinSVNAuth': %w", err)
		}
		totalN += int64(n)
	}

	// Reserved1 (ManifestFieldType: arrayStatic)
	{
		n, err := 1, binary.Write(w, binary.LittleEndian, s.Reserved1[:])
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'Reserved1': %w", err)
		}
		totalN += int64(n)
	}

	// ControlFlags (ManifestFieldType: endValue)
	{
		n, err := 4, binary.Write(w, binary.LittleEndian, &s.ControlFlags)
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'ControlFlags': %w", err)
		}
		totalN += int64(n)
	}

	// PwrDownInterval (ManifestFieldType: endValue)
	{
		n, err := 2, binary.Write(w, binary.LittleEndian, &s.PwrDownInterval)
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'PwrDownInterval': %w", err)
		}
		totalN += int64(n)
	}

	// PTTCMOSOffset0 (ManifestFieldType: endValue)
	{
		n, err := 1, binary.Write(w, binary.LittleEndian, &s.PTTCMOSOffset0)
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'PTTCMOSOffset0': %w", err)
		}
		totalN += int64(n)
	}

	// PTTCMOSOffset1 (ManifestFieldType: endValue)
	{
		n, err := 1, binary.Write(w, binary.LittleEndian, &s.PTTCMOSOffset1)
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'PTTCMOSOffset1': %w", err)
		}
		totalN += int64(n)
	}

	// ACPIBaseOffset (ManifestFieldType: endValue)
	{
		n, err := 2, binary.Write(w, binary.LittleEndian, &s.ACPIBaseOffset)
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'ACPIBaseOffset': %w", err)
		}
		totalN += int64(n)
	}

	// Reserved2 (ManifestFieldType: arrayStatic)
	{
		n, err := 2, binary.Write(w, binary.LittleEndian, s.Reserved2[:])
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'Reserved2': %w", err)
		}
		totalN += int64(n)
	}

	// PwrMBaseOffset (ManifestFieldType: endValue)
	{
		n, err := 4, binary.Write(w, binary.LittleEndian, &s.PwrMBaseOffset)
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'PwrMBaseOffset': %w", err)
		}
		totalN += int64(n)
	}

	// DigestList (ManifestFieldType: subStruct)
	{
		n, err := s.DigestList.WriteTo(w)
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'DigestList': %w", err)
		}
		totalN += int64(n)
	}

	// Reserved3 (ManifestFieldType: arrayStatic)
	{
		n, err := 3, binary.Write(w, binary.LittleEndian, s.Reserved3[:])
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'Reserved3': %w", err)
		}
		totalN += int64(n)
	}

	// SegmentCount (ManifestFieldType: endValue)
	{
		n, err := 1, binary.Write(w, binary.LittleEndian, &s.SegmentCount)
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'SegmentCount': %w", err)
		}
		totalN += int64(n)
	}

	return totalN, nil
}

// StructInfoSize returns the size in bytes of the value of field StructInfo
func (s *TXT) StructInfoTotalSize() uint64 {
	return s.StructInfo.TotalSize()
}

// Reserved0Size returns the size in bytes of the value of field Reserved0
func (s *TXT) Reserved0TotalSize() uint64 {
	return 1
}

// SetNumberSize returns the size in bytes of the value of field SetNumber
func (s *TXT) SetNumberTotalSize() uint64 {
	return 1
}

// SInitMinSVNAuthSize returns the size in bytes of the value of field SInitMinSVNAuth
func (s *TXT) SInitMinSVNAuthTotalSize() uint64 {
	return 1
}

// Reserved1Size returns the size in bytes of the value of field Reserved1
func (s *TXT) Reserved1TotalSize() uint64 {
	return 1
}

// ControlFlagsSize returns the size in bytes of the value of field ControlFlags
func (s *TXT) ControlFlagsTotalSize() uint64 {
	return 4
}

// PwrDownIntervalSize returns the size in bytes of the value of field PwrDownInterval
func (s *TXT) PwrDownIntervalTotalSize() uint64 {
	return 2
}

// PTTCMOSOffset0Size returns the size in bytes of the value of field PTTCMOSOffset0
func (s *TXT) PTTCMOSOffset0TotalSize() uint64 {
	return 1
}

// PTTCMOSOffset1Size returns the size in bytes of the value of field PTTCMOSOffset1
func (s *TXT) PTTCMOSOffset1TotalSize() uint64 {
	return 1
}

// ACPIBaseOffsetSize returns the size in bytes of the value of field ACPIBaseOffset
func (s *TXT) ACPIBaseOffsetTotalSize() uint64 {
	return 2
}

// Reserved2Size returns the size in bytes of the value of field Reserved2
func (s *TXT) Reserved2TotalSize() uint64 {
	return 2
}

// PwrMBaseOffsetSize returns the size in bytes of the value of field PwrMBaseOffset
func (s *TXT) PwrMBaseOffsetTotalSize() uint64 {
	return 4
}

// DigestListSize returns the size in bytes of the value of field DigestList
func (s *TXT) DigestListTotalSize() uint64 {
	return s.DigestList.TotalSize()
}

// Reserved3Size returns the size in bytes of the value of field Reserved3
func (s *TXT) Reserved3TotalSize() uint64 {
	return 3
}

// SegmentCountSize returns the size in bytes of the value of field SegmentCount
func (s *TXT) SegmentCountTotalSize() uint64 {
	return 1
}

// StructInfoOffset returns the offset in bytes of field StructInfo
func (s *TXT) StructInfoOffset() uint64 {
	return 0
}

// Reserved0Offset returns the offset in bytes of field Reserved0
func (s *TXT) Reserved0Offset() uint64 {
	return s.StructInfoOffset() + s.StructInfoTotalSize()
}

// SetNumberOffset returns the offset in bytes of field SetNumber
func (s *TXT) SetNumberOffset() uint64 {
	return s.Reserved0Offset() + s.Reserved0TotalSize()
}

// SInitMinSVNAuthOffset returns the offset in bytes of field SInitMinSVNAuth
func (s *TXT) SInitMinSVNAuthOffset() uint64 {
	return s.SetNumberOffset() + s.SetNumberTotalSize()
}

// Reserved1Offset returns the offset in bytes of field Reserved1
func (s *TXT) Reserved1Offset() uint64 {
	return s.SInitMinSVNAuthOffset() + s.SInitMinSVNAuthTotalSize()
}

// ControlFlagsOffset returns the offset in bytes of field ControlFlags
func (s *TXT) ControlFlagsOffset() uint64 {
	return s.Reserved1Offset() + s.Reserved1TotalSize()
}

// PwrDownIntervalOffset returns the offset in bytes of field PwrDownInterval
func (s *TXT) PwrDownIntervalOffset() uint64 {
	return s.ControlFlagsOffset() + s.ControlFlagsTotalSize()
}

// PTTCMOSOffset0Offset returns the offset in bytes of field PTTCMOSOffset0
func (s *TXT) PTTCMOSOffset0Offset() uint64 {
	return s.PwrDownIntervalOffset() + s.PwrDownIntervalTotalSize()
}

// PTTCMOSOffset1Offset returns the offset in bytes of field PTTCMOSOffset1
func (s *TXT) PTTCMOSOffset1Offset() uint64 {
	return s.PTTCMOSOffset0Offset() + s.PTTCMOSOffset0TotalSize()
}

// ACPIBaseOffsetOffset returns the offset in bytes of field ACPIBaseOffset
func (s *TXT) ACPIBaseOffsetOffset() uint64 {
	return s.PTTCMOSOffset1Offset() + s.PTTCMOSOffset1TotalSize()
}

// Reserved2Offset returns the offset in bytes of field Reserved2
func (s *TXT) Reserved2Offset() uint64 {
	return s.ACPIBaseOffsetOffset() + s.ACPIBaseOffsetTotalSize()
}

// PwrMBaseOffsetOffset returns the offset in bytes of field PwrMBaseOffset
func (s *TXT) PwrMBaseOffsetOffset() uint64 {
	return s.Reserved2Offset() + s.Reserved2TotalSize()
}

// DigestListOffset returns the offset in bytes of field DigestList
func (s *TXT) DigestListOffset() uint64 {
	return s.PwrMBaseOffsetOffset() + s.PwrMBaseOffsetTotalSize()
}

// Reserved3Offset returns the offset in bytes of field Reserved3
func (s *TXT) Reserved3Offset() uint64 {
	return s.DigestListOffset() + s.DigestListTotalSize()
}

// SegmentCountOffset returns the offset in bytes of field SegmentCount
func (s *TXT) SegmentCountOffset() uint64 {
	return s.Reserved3Offset() + s.Reserved3TotalSize()
}

// Size returns the total size of the TXT.
func (s *TXT) TotalSize() uint64 {
	if s == nil {
		return 0
	}

	var size uint64
	size += s.StructInfoTotalSize()
	size += s.Reserved0TotalSize()
	size += s.SetNumberTotalSize()
	size += s.SInitMinSVNAuthTotalSize()
	size += s.Reserved1TotalSize()
	size += s.ControlFlagsTotalSize()
	size += s.PwrDownIntervalTotalSize()
	size += s.PTTCMOSOffset0TotalSize()
	size += s.PTTCMOSOffset1TotalSize()
	size += s.ACPIBaseOffsetTotalSize()
	size += s.Reserved2TotalSize()
	size += s.PwrMBaseOffsetTotalSize()
	size += s.DigestListTotalSize()
	size += s.Reserved3TotalSize()
	size += s.SegmentCountTotalSize()
	return size
}

// PrettyString returns the content of the structure in an easy-to-read format.
func (s *TXT) PrettyString(depth uint, withHeader bool, opts ...pretty.Option) string {
	var lines []string
	if withHeader {
		lines = append(lines, pretty.Header(depth, "TXT", s))
	}
	if s == nil {
		return strings.Join(lines, "\n")
	}
	// ManifestFieldType is structInfo
	lines = append(lines, pretty.SubValue(depth+1, "Struct Info", "", &s.StructInfo, opts...)...)
	// ManifestFieldType is arrayStatic
	lines = append(lines, pretty.SubValue(depth+1, "Reserved 0", "", &s.Reserved0, opts...)...)
	// ManifestFieldType is arrayStatic
	lines = append(lines, pretty.SubValue(depth+1, "Set Number", "", &s.SetNumber, opts...)...)
	// ManifestFieldType is endValue
	lines = append(lines, pretty.SubValue(depth+1, "S Init Min SVN Auth", "", &s.SInitMinSVNAuth, opts...)...)
	// ManifestFieldType is arrayStatic
	lines = append(lines, pretty.SubValue(depth+1, "Reserved 1", "", &s.Reserved1, opts...)...)
	// ManifestFieldType is endValue
	lines = append(lines, pretty.SubValue(depth+1, "Control Flags", "", &s.ControlFlags, opts...)...)
	// ManifestFieldType is endValue
	lines = append(lines, pretty.SubValue(depth+1, "Pwr Down Interval", "", &s.PwrDownInterval, opts...)...)
	// ManifestFieldType is endValue
	lines = append(lines, pretty.SubValue(depth+1, "PTT CMOS Offset 0", "", &s.PTTCMOSOffset0, opts...)...)
	// ManifestFieldType is endValue
	lines = append(lines, pretty.SubValue(depth+1, "PTT CMOS Offset 1", "", &s.PTTCMOSOffset1, opts...)...)
	// ManifestFieldType is endValue
	lines = append(lines, pretty.SubValue(depth+1, "ACPI Base Offset", "", &s.ACPIBaseOffset, opts...)...)
	// ManifestFieldType is arrayStatic
	lines = append(lines, pretty.SubValue(depth+1, "Reserved 2", "", &s.Reserved2, opts...)...)
	// ManifestFieldType is endValue
	lines = append(lines, pretty.SubValue(depth+1, "ACPI MMIO Offset", "", &s.PwrMBaseOffset, opts...)...)
	// ManifestFieldType is subStruct
	lines = append(lines, pretty.SubValue(depth+1, "Digest List", "", &s.DigestList, opts...)...)
	// ManifestFieldType is arrayStatic
	lines = append(lines, pretty.SubValue(depth+1, "Reserved 3", "", &s.Reserved3, opts...)...)
	// ManifestFieldType is endValue
	lines = append(lines, pretty.SubValue(depth+1, "Segment Count", "", &s.SegmentCount, opts...)...)
	if depth < 2 {
		lines = append(lines, "")
	}
	return strings.Join(lines, "\n")
}

// PrettyString returns the bits of the flags in an easy-to-read format.
func (v Duration16In5Sec) PrettyString(depth uint, withHeader bool, opts ...pretty.Option) string {
	return v.String()
}

// TotalSize returns the total size measured through binary.Size.
func (v Duration16In5Sec) TotalSize() uint64 {
	return uint64(binary.Size(v))
}

// WriteTo writes the Duration16In5Sec into 'w' in binary format.
func (v Duration16In5Sec) WriteTo(w io.Writer) (int64, error) {
	return int64(v.TotalSize()), binary.Write(w, binary.LittleEndian, v)
}

// ReadFrom reads the Duration16In5Sec from 'r' in binary format.
func (v Duration16In5Sec) ReadFrom(r io.Reader) (int64, error) {
	return int64(v.TotalSize()), binary.Read(r, binary.LittleEndian, v)
}