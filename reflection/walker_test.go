package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Oscar"},
			[]string{"Oscar"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Oscar", "Ocotlan"},
			[]string{"Oscar", "Ocotlan"},
		},
		{
			"struct with non string fields",
			struct {
				Name string
				Age  int
			}{"Oscar", 24},
			[]string{"Oscar"},
		},
		{
			"nested fields",
			Person{
				"Oscar",
				Profile{24, "Ocotlan"},
			},
			[]string{"Oscar", "Ocotlan"},
		},
		{
			"pointers to things",
			&Person{
				"Oscar",
				Profile{24, "Ocotlan"},
			},
			[]string{"Oscar", "Ocotlan"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}
