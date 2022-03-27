package netstat

import (
	"reflect"
	"testing"

	mp "github.com/mackerelio/go-mackerel-plugin"
)

func TestNetstatPlugin_GraphDefinition(t *testing.T) {
	type fields struct {
		Prefix string
		Extend bool
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]mp.Graphs
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NetstatPlugin{
				Prefix: tt.fields.Prefix,
				Extend: tt.fields.Extend,
			}
			if got := n.GraphDefinition(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NetstatPlugin.GraphDefinition() = %v, want %v", got, tt.want)
			}
		})
	}
}
