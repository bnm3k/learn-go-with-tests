package c13

import (
	"reflect"
	"testing"
)

/*
golang challenge: write a function walk(x interface{}, fn func(string))
which takes a struct x and calls fn for all strings fields found inside.
difficulty level: recursively.
*/

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	t.Run("strings from slices, arrays, structs", func(t *testing.T) {
		cases := []struct {
			Name          string
			Input         interface{}
			ExpectedCalls []string
		}{
			{
				"Struct with one string field",
				struct {
					Name string
				}{"chris"},
				[]string{"chris"},
			},
			{
				"Struct with two string fields",
				struct {
					Name string
					City string
				}{"Chris", "London"},
				[]string{"Chris", "London"},
			},
			{
				"Struct with non string field",
				struct {
					Name string
					Age  int
				}{"Chris", 33},
				[]string{"Chris"},
			},
			{
				"Nested fields",
				Person{
					"Chris",
					Profile{33, "London"},
				},
				[]string{"Chris", "London"},
			},
			{
				"Pointers to things",
				&Person{
					"Chris",
					Profile{33, "London"},
				},
				[]string{"Chris", "London"},
			},
			{
				"Slices",
				[]Profile{
					{33, "London"},
					{34, "Reykjavík"},
				},
				[]string{"London", "Reykjavík"},
			},
			{
				"Arrays",
				[2]Profile{
					{33, "London"},
					{34, "Reykjavík"},
				},
				[]string{"London", "Reykjavík"},
			},
			{
				"Maps",
				map[string]string{
					"Foo": "Bar",
					"Baz": "Boz",
				},
				[]string{"Bar", "Boz"},
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
	})

	t.Run("strings from map values", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
