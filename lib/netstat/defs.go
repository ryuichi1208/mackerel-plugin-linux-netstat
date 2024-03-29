package netstat

import (
	"strings"

	mp "github.com/mackerelio/go-mackerel-plugin"
)

func (n NetstatPlugin) GraphDefinition() map[string]mp.Graphs {
	labelPrefix := strings.Title(n.MetricKeyPrefix())
	graphs := map[string]mp.Graphs{
		"tcp.segments": {
			Label: labelPrefix + " segments",
			Unit:  mp.UnitBytes,
			Metrics: []mp.Metrics{
				{Name: "TcpInSegs", Label: "", Diff: true},
				{Name: "TcpOutSegs", Label: "", Diff: true},
				{Name: "TcpRetransSegs", Label: "", Diff: true},
				{Name: "TcpInErrs", Label: "", Diff: true},
				{Name: "TcpOutRsts", Label: "", Diff: true},
			},
		},
		"tcp.memory": {
			Label: labelPrefix + " memory",
			Unit:  mp.UnitBytes,
			Metrics: []mp.Metrics{
				{Name: "tcp_memory_size", Label: "", Diff: false},
			},
		},
		"tcp.misc_errors": {
			Label: labelPrefix + " misc_errors",
			Unit:  mp.UnitBytes,
			Metrics: []mp.Metrics{
				{Name: "TcpExtEmbryonicRsts", Label: "", Diff: true},
				{Name: "TcpExtPruneCalled", Label: "", Diff: true},
				{Name: "TcpExtRcvPruned", Label: "", Diff: true},
				{Name: "TcpExtOfoPruned", Label: "", Diff: true},
				{Name: "TcpExtListenOverflows", Label: "", Diff: true},
				{Name: "TcpExtListenDrops", Label: "", Diff: true},
				{Name: "TcpExtTCPSchedulerFailed", Label: "", Diff: true},
				{Name: "TcpExtTCPRcvCollapsed", Label: "", Diff: true},
			},
		},
	}
	if n.Extend {
		extendGraphs := map[string]mp.Graphs{
			"tcp.abort": {
				Label: labelPrefix + " abort",
				Unit:  mp.UnitBytes,
				Metrics: []mp.Metrics{
					{Name: "TcpExtTCPAbortOnData", Label: "", Diff: true},
					{Name: "TcpExtTCPAbortOnClose", Label: "", Diff: true},
					{Name: "TcpExtTCPAbortOnMemory", Label: "", Diff: true},
					{Name: "TcpExtTCPAbortOnTimeout", Label: "", Diff: true},
					{Name: "TcpExtTCPAbortOnLinger", Label: "", Diff: true},
					{Name: "TcpExtTCPAbortFailed", Label: "", Diff: true},
				},
			},
			"tcp.delayed_acks": {
				Label: labelPrefix + " delayed_acks",
				Unit:  mp.UnitBytes,
				Metrics: []mp.Metrics{
					{Name: "TcpExtDelayedACKs", Label: "", Diff: true},
					{Name: "TcpExtDelayedACKLocked", Label: "", Diff: true},
					{Name: "TcpExtDelayedACKLost", Label: "", Diff: true},
				},
			},
			"tcp.timewait.timewait": {
				Label: labelPrefix + " timewait",
				Unit:  mp.UnitBytes,
				Metrics: []mp.Metrics{
					{Name: "TcpExtTW", Label: "", Diff: true},
					{Name: "TcpExtTWRecycled", Label: "", Diff: true},
					{Name: "TcpExtTWKilled", Label: "", Diff: true},
				},
			},
			"tcp.syncookies": {
				Label: labelPrefix + " syncookies",
				Unit:  mp.UnitBytes,
				Metrics: []mp.Metrics{
					{Name: "TcpExtSyncookiesSent", Label: "", Diff: true},
					{Name: "TcpExtSyncookiesRecv", Label: "", Diff: true},
					{Name: "TcpExtSyncookiesFailed", Label: "", Diff: true},
				},
			},
			"tcp.timeout": {
				Label: labelPrefix + " timeout",
				Unit:  mp.UnitBytes,
				Metrics: []mp.Metrics{
					{Name: "TcpExtTCPTimeouts", Label: "", Diff: true},
				},
			},
		}
		graphs = mapMerge(graphs, extendGraphs)
	}
	return graphs
}
