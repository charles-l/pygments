package pygments

import (
	"testing"
)

func TestBinary(t *testing.T) {
	Binary("/usr/bin/pygmentize")
}

func TestWhich(t *testing.T) {
	v := Which()
	e := "/usr/bin/pygmentize"
	if v != e {
		t.Error("Expected", e, ", got", v)
	}
}

func TestHighlighting(t *testing.T) {
	v, _ := Highlight([]byte("print \"Hello World!\""), "python", "html", "utf-8")
	e := "<span class=\"k\">print</span> <span class=\"s2\">&quot;Hello World!&quot;</span>\n"
	if v != e {
		t.Error("Expected", e, ", got", v)
	}
}

func TestUnknownFileType(t *testing.T) {
	_, err := Highlight([]byte("# markdown"), "md", "html", "utf-8")
	if err == nil {
		t.Error("Expected an error")
	}
}
