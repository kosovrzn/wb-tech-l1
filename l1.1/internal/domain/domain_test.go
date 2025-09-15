package domain

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() { os.Stdout = old }()

	f()
	_ = w.Close()
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	return buf.String()
}

func TestMethodPromotionAndPointerReceiver(t *testing.T) {
	a := Action{Human: Human{Name: "Test", Age: 20}, Role: "DevOps"}

	out := captureStdout(func() {
		a.Greet()
		a.BirthDay()
		a.Greet()
	})

	if !strings.Contains(out, "Test (20 y.o.)") || !strings.Contains(out, "Test (21 y.o.)") {
		t.Fatalf("unexpected output: %s", out)
	}
}

func TestOverrideAndParentCall(t *testing.T) {
	a := Action{Human: Human{Name: "X", Age: 1}, Role: "Role"}
	if got := a.Who(); got != "Action/Role" {
		t.Fatalf("Who() override failed: got %q", got)
	}
	if got := a.Human.Who(); got != "Human" {
		t.Fatalf("parent Who() call failed: got %q", got)
	}
}

func TestInterfaceSatisfaction(t *testing.T) {
	var _ WhoGreeter = Human{}
	var _ WhoGreeter = Action{}
	var _ WhoGreeter = &Action{}
	_ = fmt.Sprintf
}
