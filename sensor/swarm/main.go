package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	// This registers all registries and scanners.
	_ "bitbucket.org/stack-rox/apollo/pkg/registries/all"
	_ "bitbucket.org/stack-rox/apollo/pkg/scanners/all"

	"bitbucket.org/stack-rox/apollo/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/benchmarks"
	"bitbucket.org/stack-rox/apollo/pkg/clientconn"
	"bitbucket.org/stack-rox/apollo/pkg/enforcers"
	"bitbucket.org/stack-rox/apollo/pkg/env"
	"bitbucket.org/stack-rox/apollo/pkg/grpc"
	"bitbucket.org/stack-rox/apollo/pkg/grpc/routes"
	"bitbucket.org/stack-rox/apollo/pkg/listeners"
	"bitbucket.org/stack-rox/apollo/pkg/logging"
	"bitbucket.org/stack-rox/apollo/pkg/mtls/verifier"
	"bitbucket.org/stack-rox/apollo/pkg/orchestrators"
	"bitbucket.org/stack-rox/apollo/pkg/sensor"
	"bitbucket.org/stack-rox/apollo/pkg/sources"
	"bitbucket.org/stack-rox/apollo/sensor/swarm/enforcer"
	"bitbucket.org/stack-rox/apollo/sensor/swarm/listener"
	"bitbucket.org/stack-rox/apollo/sensor/swarm/orchestrator"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	grpcLib "google.golang.org/grpc"
)

var (
	logger = logging.LoggerForModule()

	clusterID          string
	centralEndpoint    string
	advertisedEndpoint string
	image              string

	server                 grpc.API
	listenerInstance       listeners.Listener
	enforcerInstance       enforcers.Enforcer
	benchScheduler         *benchmarks.SchedulerClient
	orchestratorInstance   orchestrators.Orchestrator
	imageIntegrationPoller *sources.Client

	conn *grpcLib.ClientConn

	sensorInstance sensor.Sensor
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	initialize()

	start()

	for {
		select {
		case sig := <-sigs:
			logger.Infof("Caught %s signal", sig)
			stop()
			logger.Info("Swarm Sensor terminated")
			return
		}
	}
}

// Fetch all needed environment information and initialize all needed objects.
func initialize() {
	// Read environment.
	clusterID = env.ClusterID.Setting()
	centralEndpoint = env.CentralEndpoint.Setting()
	advertisedEndpoint = env.AdvertisedEndpoint.Setting()
	image = env.Image.Setting()

	// Start up connections.
	var err error
	conn, err = clientconn.GRPCConnection(centralEndpoint)
	if err != nil {
		logger.Fatalf("Error connecting to central: %s", err)
	}

	listenerInstance, err = listener.New()
	if err != nil {
		panic(err)
	}

	enforcerInstance, err = enforcer.New()
	if err != nil {
		panic(err)
	}

	orchestratorInstance, err = orchestrator.New()
	if err != nil {
		panic(err)
	}

	benchScheduler, err = benchmarks.NewSchedulerClient(orchestratorInstance, advertisedEndpoint, image, conn, clusterID)
	if err != nil {
		panic(err)
	}

	imageIntegrationPoller = sources.NewImageIntegrationsClient(conn, clusterID)

	logger.Info("Swarm Sensor Initialized")
}

// Start all necessary side processes then start sensor.
func start() {
	// Create grpc server with custom routes
	config := grpc.Config{
		TLS:          verifier.NonCA{},
		CustomRoutes: customRoutes(),
	}
	server = grpc.NewAPI(config)

	logger.Infof("Connecting to Central server %s", centralEndpoint)
	registerAPIServices()
	server.Start()

	// Start all of our channels and listeners
	if listenerInstance != nil {
		go listenerInstance.Start()
	}
	if enforcerInstance != nil {
		go enforcerInstance.Start()
	}
	if benchScheduler != nil {
		go benchScheduler.Start()
	}
	if imageIntegrationPoller != nil {
		go imageIntegrationPoller.Start()
	}

	// Wait for central so we can initiate our GRPC connection to send sensor events.
	waitUntilCentralIsReady(conn)

	// If everything is brought up correctly, start the sensor.
	if listenerInstance != nil && enforcerInstance != nil {
		sensorInstance = sensor.NewSensor(imageIntegrationPoller, conn)
		sensorInstance.Start(listenerInstance.Events(), enforcerInstance.Actions())
	}

	logger.Info("Swarm Sensor Started")
}

// Stop stops the sensor and all necessary side processes.
func stop() {
	// Stop the sensor.
	sensorInstance.Stop()

	// Stop all of our listeners.
	if listenerInstance != nil {
		listenerInstance.Stop()
	}
	if enforcerInstance != nil {
		enforcerInstance.Stop()
	}
	if benchScheduler != nil {
		benchScheduler.Stop()
	}
	if imageIntegrationPoller != nil {
		imageIntegrationPoller.Stop()
	}

	logger.Info("Swarm Sensor Stopped")
}

// Helper functions.
////////////////////

// Provides the custom routes to provide.
func customRoutes() map[string]routes.CustomRoute {
	routeMap := map[string]routes.CustomRoute{
		"/metrics": {
			ServerHandler: promhttp.Handler(),
			Compression:   false,
		},
	}
	return routeMap
}

// Registers our connection for benchmarking.
func registerAPIServices() {
	server.Register(benchmarks.NewBenchmarkResultsService(benchmarks.NewLRURelayer(conn)))
	logger.Info("API services registered")
}

// Function does not complete until central is pingable.
func waitUntilCentralIsReady(conn *grpcLib.ClientConn) {
	pingService := v1.NewPingServiceClient(conn)
	err := pingWithTimeout(pingService)
	for err != nil {
		logger.Infof("Ping to Central failed: %s. Retrying...", err)
		time.Sleep(2 * time.Second)
		err = pingWithTimeout(pingService)
	}
}

// Ping a service with a timeout of 10 seconds.
func pingWithTimeout(svc v1.PingServiceClient) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = svc.Ping(ctx, &empty.Empty{})
	return
}
