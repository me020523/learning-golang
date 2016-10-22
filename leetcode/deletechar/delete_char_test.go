package deletechar

import "testing"

func TestDeleteChar(t *testing.T) {
	runes := []rune{
		'a', 'b', 'b', 'd', 'b',
	}
	result := []rune{
		'a', 'd',
	}
	runes = DeleteChar(runes, 'b')
	t.Logf("%v\n", runes)
	if len(runes) != len(result) {
		t.FailNow()
	}
	for i := 0; i < len(result); i++ {
		if runes[i] != result[i] {
			t.FailNow()
		}

	}
}
