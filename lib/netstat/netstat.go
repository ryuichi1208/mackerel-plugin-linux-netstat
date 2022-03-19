package netstat

import (
	"flag"

	mackerelplugin "github.com/mackerelio/go-mackerel-plugin"
	mp "github.com/mackerelio/go-mackerel-plugin"
)

type NetstatPlugin struct {
	Prefix string
	Extend bool
}

func mapMerge(m ...map[string]mackerelplugin.Graphs) map[string]mackerelplugin.Graphs {
	ans := make(map[string]mackerelplugin.Graphs, 0)

	for _, c := range m {
		for k, v := range c {
			ans[k] = v
		}
	}
	return ans
}

func (n NetstatPlugin) MetricKeyPrefix() string {
	if n.Prefix == "" {
		n.Prefix = "netstat"
	}
	return n.Prefix
}

func (v NetstatPlugin) FetchMetrics() (map[string]float64, error) {
	m := get()
	return m, nil

}

func Do() {
	optPrefix := flag.String("metric-key-prefix", "", "Metric key prefix")
	optExtend := flag.Bool("extend", false, "extend flag")
	flag.Parse()
	v := &NetstatPlugin{
		Prefix: *optPrefix,
		Extend: *optExtend,
	}
	plugin := mp.NewMackerelPlugin(v)
	plugin.Run()
}
