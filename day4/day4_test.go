package main

import "testing"

func TestExec(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name      string
		args      args
		wantAnsP1 int
		wantAnsP2 int
	}{
		// TODO: Add test cases.
		{"Example", args{[]string{"[1518-11-01 00:00] Guard #10 begins shift",
			"[1518-11-01 00:05] falls asleep",
			"[1518-11-01 00:25] wakes up",
			"[1518-11-01 00:30] falls asleep",
			"[1518-11-01 00:55] wakes up",
			"[1518-11-01 23:58] Guard #99 begins shift",
			"[1518-11-02 00:40] falls asleep",
			"[1518-11-02 00:50] wakes up",
			"[1518-11-03 00:05] Guard #10 begins shift",
			"[1518-11-03 00:24] falls asleep",
			"[1518-11-03 00:29] wakes up",
			"[1518-11-04 00:02] Guard #99 begins shift",
			"[1518-11-04 00:36] falls asleep",
			"[1518-11-04 00:46] wakes up",
			"[1518-11-05 00:03] Guard #99 begins shift",
			"[1518-11-05 00:45] falls asleep",
			"[1518-11-05 00:55] wakes up"}}, 240, 4455},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAnsP1, gotAnsP2 := Exec(tt.args.lines)
			if gotAnsP1 != tt.wantAnsP1 {
				t.Errorf("Exec() gotAnsP1 = %v, want %v", gotAnsP1, tt.wantAnsP1)
			}
			if gotAnsP2 != tt.wantAnsP2 {
				t.Errorf("Exec() gotAnsP2 = %v, want %v", gotAnsP2, tt.wantAnsP2)
			}
		})
	}
}
