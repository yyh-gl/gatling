package main

import "github.com/prometheus/client_golang/prometheus"

var (
	// embed at build time
	version string

	runningVersion = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "running_version",
		Help: "A gauge of running version.",
	}, []string{"version"})
)

func init() {
	prometheus.MustRegister(runningVersion)
}

func countUpRunningVersion(version string) {
	runningVersion.With(prometheus.Labels{"version": version}).Inc()
}

func countDownRunningVersion(version string) {
	runningVersion.With(prometheus.Labels{"version": version}).Dec()
}
