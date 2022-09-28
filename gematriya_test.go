package gematriya

import (
	"fmt"
	"testing"
)

func assertEqual(t *testing.T, expected string, actual string) {
	if actual != expected {
		t.Errorf(`Expected "%s", got "%s"`, expected, actual)
	}
}

func TestGematriya(t *testing.T) {
	assertEqual(t, "תשמ״ט", Gematriya(5749))
	assertEqual(t, "תשע״ד", Gematriya(5774))
	assertEqual(t, "תש״פ", Gematriya(5780))
	assertEqual(t, "ג׳", Gematriya(3))
	assertEqual(t, "י״ד", Gematriya(14))
	assertEqual(t, "ט״ו", Gematriya(15))
	assertEqual(t, "ט״ז", Gematriya(16))
	assertEqual(t, "י״ז", Gematriya(17))
	assertEqual(t, "כ׳", Gematriya(20))
	assertEqual(t, "כ״ה", Gematriya(25))
	assertEqual(t, "ס׳", Gematriya(60))
	assertEqual(t, "קכ״ג", Gematriya(123))
	assertEqual(t, "תרי״ג", Gematriya(613))
	assertEqual(t, "תשמ״ט", Gematriya(5749))
	assertEqual(t, "ג׳תשס״א", Gematriya(3761))
	assertEqual(t, "ו׳תשמ״ט", Gematriya(6749))
	assertEqual(t, "ח׳תשס״ה", Gematriya(8765))
	assertEqual(t, "כב׳ת״ש", Gematriya(22700))
	assertEqual(t, "טז׳קכ״ג", Gematriya(16123))
	assertEqual(t, "א׳קכ״ג", Gematriya(1123))
	assertEqual(t, "ו׳", Gematriya(6000))
	assertEqual(t, "ז׳ז׳", Gematriya(7007))
}

func ExampleGematriya() {
	gematriya := Gematriya(5749)
	fmt.Println(gematriya)
	// Output: תשמ״ט
}
