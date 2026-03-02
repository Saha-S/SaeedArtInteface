package art

import "testing"

func TestDecode(t *testing.T) {
	svc := NewService()
	out, err := svc.Execute(ModeDecode, "[3 A][2 B]", false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "AAABB" {
		t.Fatalf("got %q want %q", out, "AAABB")
	}
}

func TestEncode(t *testing.T) {
	svc := NewService()
	out, err := svc.Execute(ModeEncode, "AAABB", false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != "[3 A][2 B]" {
		t.Fatalf("got %q want %q", out, "[3 A][2 B]")
	}
}

func TestDecodeInvalid(t *testing.T) {
	svc := NewService()
	_, err := svc.Execute(ModeDecode, "[x #]", false)
	if err == nil {
		t.Fatal("expected error for malformed input")
	}
}
