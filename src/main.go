package main

import (
	main_configurations_cache "mpindicator/main/configurations/cache"
	main_configurations_mongodb "mpindicator/main/configurations/mongodb"
	main_configurations_rabbitmq "mpindicator/main/configurations/rabbitmq"
	main_configurations_router "mpindicator/main/configurations/router"
	main_configurations_tracer "mpindicator/main/configurations/tracer"
	main_configurations_yml "mpindicator/main/configurations/yml"
	main_gateways_http_httphandler "mpindicator/main/gateways/http/httphandler"
	main_gateways_rabbitmq_listeners "mpindicator/main/gateways/rabbitmq/listeners"
	main_usecases_beans "mpindicator/main/usecases/beans"
	main_utils "mpindicator/main/utils"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmgorilla"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
)

const MSG_INITIALIZING_APPLICATION = "Initializing application."
const MSG_APPLICATION_FAILED = "Application has failed to start"
const MSG_LISTENER = "Listener on port: %s"

const IDX_APPLICATION_PORT = "Application.Port"
const IDX_TRACING_SERVER_NAME = "Tracing.server.name"

// Estudar reflect para criação de beans
// Estudar Create Bean via redis lock: usando sync é threadSave mas nao é pod safe

// TODO:
// [  ] - Persistência DB collection Error
// [OK] - Controller Pegar Indicador
// [  ] - Controller Pegar Indicador newest
// [OK] - Implementation Redis Cache
// [  ] - Use OpenTelemetry
// [OK] - Redis lock
// [  ] - Monitoria aplicação (Prometheus)
// [  ] - HealthCheck endpoint e integração com o kubernetes

func init() {
	log.Println(MSG_INITIALIZING_APPLICATION)

	main_configurations_yml.GetYmlConfigBean()
	main_configurations_mongodb.GetDatabaseBean()
	main_configurations_router.GetRouterBean()
	main_configurations_cache.GetRedisClusterBean()
	main_usecases_beans.GetBeans()

	main_configurations_rabbitmq.SetAmqpConfig()
	go main_gateways_rabbitmq_listeners.Listen()
	go main_gateways_rabbitmq_listeners.IndicatorProcessorListener()
}

var log = &logrus.Logger{
	Out:   os.Stderr,
	Hooks: make(logrus.LevelHooks),
	Level: logrus.DebugLevel,
	Formatter: &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "log.level",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyFunc:  "function.name", // non-ECS
		},
	},
}

func main() {

	log.SetLevel(logrus.TraceLevel)

	file, errLogFile := os.OpenFile("out.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if errLogFile == nil {
		log.SetOutput(file)
	}
	defer file.Close()

	fields := logrus.Fields{"userId": 12}
	log.WithFields(fields).Info("User logged in!")

	apm.DefaultTracer.SetLogger(log)

	main_configurations_tracer.InitTracer()

	applicationPort := main_configurations_yml.GetBeanPropertyByName(IDX_APPLICATION_PORT)
	tracingServerName := main_configurations_yml.GetBeanPropertyByName(IDX_TRACING_SERVER_NAME)
	routes := main_gateways_http_httphandler.GetRoutesBean()

	router := main_gateways_http_httphandler.ConfigRoutes(main_configurations_router.GetRouterBean(), *routes)
	router.Use(otelmux.Middleware(tracingServerName))
	router.Use(apmgorilla.Middleware())

	err := http.ListenAndServe(":"+applicationPort, router)
	if err != nil {
		main_utils.FailOnError(err, MSG_APPLICATION_FAILED)
	}
	log.Printf(MSG_LISTENER, applicationPort)
}
