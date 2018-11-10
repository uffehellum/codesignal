package secretagents

import "testing"

func TestEmpty(t *testing.T) {
	result := secretAgentsMeetingProposal("", 1)
	if len(result) != 2 {
		t.Fatal("empty", result)
	}
}

func Test1(t *testing.T) {
	result := secretAgentsMeetingProposal("10.7.11.0.5.?.#.?._.15.-1.14", 1)
	if result[0] != "today-night-bar" || result[1] != "13.7" {
		t.Fatal("wrong", result)
	}
	a := []test{
		{
			"10.7.11.0.5.?.#.?._.15.-1.14",
			1,
			"today-night-bar",
			"13.7",
		},
		{
			"10.4.7.11.0.5._.10.?.*.?._.11.5.8.6.-4.2.11.-4.9.6",
			2,
			"twodays-morning-restaurant",
			"5.9.12",
		},
	}
	for _, tt := range a {
		result = secretAgentsMeetingProposal(tt.Message, tt.D)
		if result[0] != tt.Clear || result[1] != tt.Reply {
			t.Fatal(tt, result)
		}
	}
}

type test struct {
	Message string
	D       int
	Clear   string
	Reply   string
}
