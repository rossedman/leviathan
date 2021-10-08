package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/slok/go-http-metrics/middleware"
	"github.com/slok/go-http-metrics/middleware/std"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
)

var rootCmd = &cobra.Command{
	Use:   "leviathan",
	Short: "A big fish, lurking in the dark, hunted by the Kubernetes",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Infof("starting on %v...", viper.GetInt("port"))

		// Create our middleware for capturing HTTP metrics
		mdlw := middleware.New(middleware.Config{
			Recorder: metrics.NewRecorder(metrics.Config{}),
		})

		// setup server
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, `{"app": "leviathan", "version": "0.1.8"}`)
		})
		mux.HandleFunc("/error", func(w http.ResponseWriter, req *http.Request) {
			logrus.Infoln("received error request")
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, `{"error": "true"}`)
		})
		mux.HandleFunc("/health", func(w http.ResponseWriter, req *http.Request) {
			logrus.Infoln("received healthcheck request")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "OK\n")
		})
		mux.HandleFunc("/headers", func(w http.ResponseWriter, req *http.Request) {
			logrus.Infoln("received header request")
			w.WriteHeader(http.StatusOK)
			for name, headers := range req.Header {
				for _, h := range headers {
					fmt.Fprintf(w, "%v: %v\n", name, h)
				}
			}
		})
		mux.Handle("/metrics", promhttp.Handler())

		// wrap mux in prometheus middleware and serve
		logrus.Fatal(http.ListenAndServe(
			fmt.Sprintf(":%v", viper.GetInt("port")),
			std.Handler("", mdlw, mux),
		))
	},
}

func init() {
	cobra.OnInitialize(initConfig)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func initConfig() {
	viper.SetEnvPrefix("leviathan")
	viper.SetDefault("port", "8080")
	viper.AutomaticEnv()
}
