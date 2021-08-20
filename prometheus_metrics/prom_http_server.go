package main

import (
	"math/rand"
	"net/http"

	//"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	counterRegistry *prometheus.Registry
	registry        *prometheus.Registry
)

func init() {
	counterRegistry = prometheus.NewRegistry()
	registry = prometheus.NewRegistry()
}

func generateTestGauge() {
	testGauge := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "test_gauge",
			Help: "A gauge generated by appplication",
		},
	)
	var gaugeTestValue float64 = 1
	testGauge.Set(gaugeTestValue)
	registry.MustRegister(testGauge)
	for {
		time.Sleep(3 * time.Second)
		val := rand.Intn(1001)
		if float64(val) > gaugeTestValue {
			gaugeTestValue = float64(val)
		}
		testGauge.Set(gaugeTestValue)
	}
}

func generateTestCounter() {
	testCounter := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "test_counter",
			Help: "A counter generated by application",
		},
	)
	counterRegistry.MustRegister(testCounter)
	for {
		time.Sleep(3 * time.Second)
		testCounter.Inc()
	}
}

func generateTestHistogram() {
	testHistogram := prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "test_histogram",
			Help:    "A histogram generated by application",
			Buckets: []float64{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000},
		},
	)
	registry.MustRegister(testHistogram)
	for {
		val := rand.Intn(1001)
		testHistogram.Observe(float64(val))
		time.Sleep(3 * time.Second)
	}
}

func main() {
	gatherers := prometheus.Gatherers{
		counterRegistry,
		registry,
	}
	router := mux.NewRouter()
	//promhttp.Handler()
	router.Handle("/metrics", promhttp.HandlerFor(gatherers, promhttp.HandlerOpts{}))
	log.Info().Msg("Listening on 8888 for metrics")
	go func() {
		log.Fatal().Msg(http.ListenAndServe(":8888", router).Error())
	}()
	//var wg sync.WaitGroup
	//wg.Add(3)
	log.Info().Msg("Started generating test_gauge")
	go generateTestGauge()
	log.Info().Msg("Started generating test_counter")
	go generateTestCounter()
	log.Info().Msg("Started observing test_histogram")
	generateTestHistogram()
	//wg.Wait()
}
