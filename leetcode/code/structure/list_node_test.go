package structure

import (
	"fmt"
	"testing"
)

func Test_buildListNode(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{nums: []int{1, 2, 3, 4, 5, 6}},
			//want: buildListNode(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := BuildListNode(tt.args.nums)
			for node != nil {
				if node.Next != nil {
					fmt.Printf("  %d  ---", node.Val)
				} else {
					fmt.Printf("  %d", node.Val)
				}
				node = node.Next
			}
			//if got := buildListNode(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("buildListNode() = %v, want %v", got, tt.want)
			//}
		})
	}
}
