package component

import (
	"fmt"
	"testing"
)

func intP(i int) *int {
	return &i
}

func TestBaseComponent_DeepCopy(t *testing.T) {
	baseOptions := BaseComponentOptions{
		Label:        "label",
		Placeholder:  "placeholder",
		Required:     true,
		State:        "asdf",
		DefaultValue: 23232,
		Messages: []Message{
			{
				State: "asdf",
				Text:  "23232",
			},
		},
		Width: "11010em",
		Grid: &Grid{
			Sm:       intP(1),
			Md:       intP(2),
			Lg:       intP(3),
			SmOffset: intP(4),
			MdOffset: intP(5),
			LgOffset: intP(6),
			Order:    intP(7),
		},
	}

	base := NewBaseComponent("vn-asdf", baseOptions)
	baseP := &base

	copiedBase := baseP.DeepCopy()
	if copiedBase == baseP {
		t.Fatalf("copiedBase == baseP")
	} else {
		fmt.Println("Success!")
	}
}
