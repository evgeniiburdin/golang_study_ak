package main

import "testing"

func Test_defineListLength(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
							},
						},
					},
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := defineListLength(tt.args.head); got != tt.want {
				t.Errorf("defineListLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateTwinSum(t *testing.T) {
	type args struct {
		twin       *ListNode
		twinIndex  int
		listLength int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				twin: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 3,
							},
						},
					},
				},
				twinIndex:  1,
				listLength: 4,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateTwinSum(tt.args.twin, tt.args.twinIndex, tt.args.listLength); got != tt.want {
				t.Errorf("calculateTwinSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
