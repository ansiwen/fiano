package bgkey

import (
	"fmt"

	"github.com/linuxboot/fiano/pkg/intel/metadata/common/pretty"
)

// Print prints the Key Manifest.
func (m *Manifest) Print() {
	if m.KeyAndSignature.Signature.DataTotalSize() < 1 {
		fmt.Printf("%v\n", m.PrettyString(1, true, pretty.OptionOmitKeySignature(true)))
		fmt.Printf("  --KeyAndSignature--\n\tKey Manifest not signed!\n\n")
	} else {
		fmt.Printf("%v\n", m.PrettyString(1, true, pretty.OptionOmitKeySignature(false)))
	}
}