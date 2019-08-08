package main

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/puppetlabs/horsehead/instrumentation/metrics"
	"github.com/puppetlabs/horsehead/instrumentation/metrics/collectors"
	"github.com/puppetlabs/horsehead/instrumentation/metrics/delegates"
	"github.com/puppetlabs/horsehead/instrumentation/metrics/server"
)

const (
	namespace           = "example"
	storageBackendCall  = "storage_backend_call"
	taskCount           = "task_count"
	requestCount        = "request_count"
	httpRequestDuration = "http_handler_duration"
)

type testApp struct {
	m *metrics.Metrics
}

func (t testApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := t.m.MustCounter(requestCount, metrics.NewLabel("method", r.Method))
	c.Add(1)

	gcsTimer := t.m.MustTimer(storageBackendCall,
		metrics.NewLabel("backend", "gcs"),
		metrics.NewLabel("action", "put"),
	)

	t.m.OnTimer(gcsTimer, func() {
		random := rand.Intn(10)
		<-time.After(time.Millisecond * time.Duration(random))
	})

	s3Timer := t.m.MustTimer(storageBackendCall,
		metrics.NewLabel("backend", "s3"),
		metrics.NewLabel("action", "get"),
	)

	t.m.OnTimer(s3Timer, func() {
		random := rand.Intn(10)
		<-time.After(time.Millisecond * time.Duration(random))
	})
}

func main() {
	m, err := metrics.NewNamespace(namespace, metrics.Options{
		DelegateType:  delegates.PrometheusDelegate,
		ErrorBehavior: metrics.ErrorBehaviorLog,
	})
	if err != nil {
		log.Fatal(err)
	}

	copts := collectors.CounterOptions{
		Description: "a count of requests executed",
		Labels:      []string{"method"},
	}
	if err := m.RegisterCounter(requestCount, copts); err != nil {
		log.Fatal(err)
	}

	m.MustRegisterCounter(taskCount, collectors.CounterOptions{
		Description: "a count of tasks executed",
		Labels:      []string{"status"},
	})

	topts := collectors.TimerOptions{
		Description: "duration of requests to google cloud storage",
		Labels:      []string{"backend", "action"},
	}
	if err := m.RegisterTimer(storageBackendCall, topts); err != nil {
		log.Fatal(err)
	}

	hopts := collectors.DurationMiddlewareOptions{
		Description: "http request duration middleware",
		Labels:      []string{"handler", "code", "method"},
	}
	if err := m.RegisterDurationMiddleware(httpRequestDuration, hopts); err != nil {
		log.Fatal(err)
	}

	failedTasks := m.MustCounter(taskCount, metrics.NewLabel("status", "failed"))
	failedTasks.Add(5)

	succeededTasks := m.MustCounter(taskCount, metrics.NewLabel("status", "succeeded"))
	succeededTasks.Add(1)
	succeededTasks.Inc()

	mw := m.MustDurationMiddleware(httpRequestDuration, metrics.NewLabel("handler", "test_handler"))
	mux := http.NewServeMux()
	mux.Handle("/", mw.Wrap(testApp{m}))

	go http.ListenAndServe("localhost:2399", mux)

	server := server.New(m, server.Options{
		BindAddr: "localhost:2398",
		Path:     "/metrics",
	})

	if err := server.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
}