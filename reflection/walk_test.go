package reflection__test

import (
	reflection_ "hello/reflection"
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	type Profile struct {
		Age  int
		City string
	}

	type Person struct {
		Name    string
		Profile Profile
	}

	cases := []struct {
		Name          string
		Input         any
		ExpectedCalls []string
	}{
		{
			Name: "struct with one string field",
			Input: struct {
				Name string
			}{
				Name: "Chris",
			},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "struct with two string fields",
			Input: struct {
				Name string
				City string
			}{
				Name: "Chris",
				City: "London",
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "struct with non string field",
			Input: struct {
				Name string
				Age  int
			}{
				Name: "Chris",
				Age:  33,
			},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "nested fields",
			Input: Person{
				Name: "Chris",
				Profile: Profile{
					Age:  33,
					City: "London",
				},
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "pointers to things",
			Input: &Person{
				Name: "Chris",
				Profile: Profile{
					Age:  33,
					City: "London",
				},
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "slices",
			Input: []Profile{
				{
					Age:  33,
					City: "London",
				},
				{
					Age:  34,
					City: "Reykjavík",
				},
			},
			ExpectedCalls: []string{"London", "Reykjavík"},
		},
		{
			Name: "arrays",
			Input: [2]Profile{
				{
					Age:  33,
					City: "London",
				},
				{
					Age:  34,
					City: "Reykjavík",
				},
			},
			ExpectedCalls: []string{"London", "Reykjavík"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			var got []string
			reflection_.Walk(tt.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, tt.ExpectedCalls) {
				t.Errorf("want %v, got %v", tt.ExpectedCalls, got)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		reflection_.Walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, "Bar", got)
		assertContains(t, "Boz", got)
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{Age: 33, City: "Berlin"}
			aChannel <- Profile{Age: 34, City: "Katowice"}
			close(aChannel)
		}()

		want := []string{"Berlin", "Katowice"}
		var got []string

		reflection_.Walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}

func assertContains(t *testing.T, needle string, haystack []string) {
	t.Helper()

	var contains bool
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
