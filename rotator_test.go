package rotator

import (
	"fmt"
	"testing"
)

func TestRotator_Add(t *testing.T) {
	r := New(func(v string) string {
		return v
	})

	_ = r.Add("a", "b", "c", "d")
	if r.Len() != 4 {
		t.Error("expected 4 got ", r.Len())
	}

}

func TestRotator_Rotate(t *testing.T) {
	t.Run("with values", func(t *testing.T) {
		r := New(func(v string) string { return v })

		_ = r.Add("a", "b", "c")
		v, _ := r.Rotate()
		if v != "a" {
			t.Error("expected a got ", v)
		}

		v, _ = r.Rotate()
		if v != "b" {
			t.Error("expected b got ", v)
		}

		v, _ = r.Rotate()
		if v != "c" {
			t.Error("expected c got ", v)
		}

		v, _ = r.Rotate()
		if v != "a" {
			t.Error("expected a got ", v)
		}

		_ = r.Add("d")

		v, _ = r.Rotate()
		if v != "b" {
			t.Error("expected b got ", v)
		}

	})

	t.Run("without values", func(t *testing.T) {
		r := New(func(v string) string { return v })
		v, _ := r.Rotate()
		if v != "" {
			t.Error("expected empty string got ", v)
		}
	})

	t.Run("with struct values", func(t *testing.T) {
		type test struct {
			name string
		}

		r := New(func(v test) string { return v.name })
		_ = r.Add(test{name: "a"}, test{name: "b"}, test{name: "c"})

		v, _ := r.Rotate()
		if v.name != "a" {
			t.Error("expected a got ", v)
		}

		v, _ = r.Rotate()
		if v.name != "b" {
			t.Error("expected b got ", v)
		}

		v, _ = r.Rotate()
		if v.name != "c" {
			t.Error("expected c got ", v)
		}

		v, _ = r.Rotate()
		if v.name != "a" {
			t.Error("expected a got ", v)
		}

		_ = r.Add(test{name: "d"})

		v, _ = r.Rotate()
		if v.name != "b" {
			t.Error("expected b got ", v)
		}
	})

	t.Run("with concurrent access", func(t *testing.T) {
		r := New(func(v string) string { return v })
		for i := 0; i < 100; i++ {
			i := i
			go func() { _ = r.Add(fmt.Sprint(i)) }()
			go func() {
				_, _ = r.Rotate()
			}()
			go r.Get(fmt.Sprint(i))
		}
	})
}

func TestRotator_Get(t *testing.T) {
	r := New(func(v string) string { return v })
	_ = r.Add("a", "b", "c")
	v, _ := r.Get("a")
	if v != "a" {
		t.Error("expected a got ", v)
	}

	v, _ = r.Get("b")
	if v != "b" {
		t.Error("expected b got ", v)
	}

	v, _ = r.Get("c")
	if v != "c" {
		t.Error("expected c got ", v)
	}

	v, _ = r.Get("d")
	if v != "" {
		t.Error("expected empty string got ", v)
	}
}

func TestRotator_Delete(t *testing.T) {
	r := New(func(v string) string { return v })
	_ = r.Add("a", "b", "c")

	_ = r.Delete("a")
	if r.Len() != 2 {
		t.Error("expected 2 got ", r.Len())
	}

	_ = r.Delete("b")
	if r.Len() != 1 {
		t.Error("expected 1 got ", r.Len())
	}

	_ = r.Delete("c")
	if r.Len() != 0 {
		t.Error("expected 0 got ", r.Len())
	}

	_ = r.Delete("d")
	if r.Len() != 0 {
		t.Error("expected 0 got ", r.Len())
	}

}
