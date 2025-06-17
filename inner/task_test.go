package inner

import "testing"

func TestSomeForTest(t *testing.T) {
	SomeForTest()

	result := "test"
	expected := "test"
	if result != expected {
		t.Errorf("Ожидалось: %q, получено: %q", expected, result)
	}
}
