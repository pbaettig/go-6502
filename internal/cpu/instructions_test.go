package cpu

import (
	"fmt"
	"testing"
)

func TestCPU_FetchDecodeExecute(t *testing.T) {
	proc := New()
	proc.Mem[ResetVector] = 0x00
	proc.Mem[ResetVector+1] = 0x00

	ldaTests := []struct {
		value       uint8
		wantStatusZ bool
		wantStatusN bool
	}{
		{
			value:       42,
			wantStatusN: false,
			wantStatusZ: false,
		},
		{
			value:       0,
			wantStatusN: false,
			wantStatusZ: true,
		},
		{
			value:       127,
			wantStatusN: false,
			wantStatusZ: false,
		},
		{
			value:       128,
			wantStatusN: true,
			wantStatusZ: false,
		},
	}
	for i, tt := range ldaTests {
		t.Run(fmt.Sprintf(
			"LDA_Immediate_%d", i),
			func(t *testing.T) {
				proc.Mem[0x0000] = 0xA9
				proc.Mem[0x0001] = tt.value
				proc.Reset()

				proc.FetchDecodeExecute()
				if proc.AC != tt.value {
					t.Errorf("want AC=%d, got AC=%d", tt.value, proc.AC)
				}
				if z := proc.SR.Z(); z != tt.wantStatusZ {
					t.Errorf("want SR.Z=%v, got SR.Z=%v", tt.wantStatusZ, z)
				}
				if n := proc.SR.N(); n != tt.wantStatusN {
					t.Errorf("want SR.N=%v, got SR.N=%v", tt.wantStatusN, n)
				}
			})
	}
}
