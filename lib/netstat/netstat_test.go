package netstat

import (
	"reflect"
	"testing"

	mp "github.com/mackerelio/go-mackerel-plugin"
)

func Test_mapMerge(t *testing.T) {
	type args struct {
		m []map[string]mp.Graphs
	}
	tests := []struct {
		name string
		args args
		want map[string]mp.Graphs
	}{
		{
			name: "default",
			args: args{
				m: []map[string]mp.Graphs{
					{
						"netstat": mp.Graphs{
							Label: "netstat",
							Unit:  mp.UnitBytes,
							Metrics: []mp.Metrics{
								{Name: "TcpInSegs", Label: "", Diff: true},
							},
						},
					},
					{
						"netstat2": mp.Graphs{
							Label: "netstat",
							Unit:  mp.UnitBytes,
							Metrics: []mp.Metrics{
								{Name: "TcpInRets", Label: "", Diff: true},
							},
						},
					},
				},
			},
			want: map[string]mp.Graphs{
				"netstat": mp.Graphs{
					Label: "netstat",
					Unit:  mp.UnitBytes,
					Metrics: []mp.Metrics{
						{Name: "TcpInSegs", Label: "", Diff: true},
					},
				},
				"netstat2": mp.Graphs{
					Label: "netstat",
					Unit:  mp.UnitBytes,
					Metrics: []mp.Metrics{
						{Name: "TcpInRets", Label: "", Diff: true},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapMerge(tt.args.m...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNetstatPlugin_MetricKeyPrefix(t *testing.T) {
	type fields struct {
		Prefix string
		Extend bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NetstatPlugin{
				Prefix: tt.fields.Prefix,
				Extend: tt.fields.Extend,
			}
			if got := n.MetricKeyPrefix(); got != tt.want {
				t.Errorf("NetstatPlugin.MetricKeyPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNetstatPlugin_FetchMetrics(t *testing.T) {
	type fields struct {
		Prefix string
		Extend bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[string]float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NetstatPlugin{
				Prefix: tt.fields.Prefix,
				Extend: tt.fields.Extend,
			}
			got, err := v.FetchMetrics()
			if (err != nil) != tt.wantErr {
				t.Errorf("NetstatPlugin.FetchMetrics() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NetstatPlugin.FetchMetrics() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Do()
		})
	}
}
