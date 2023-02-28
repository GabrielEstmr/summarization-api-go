#!/bin/bash

chmod go-w /etc/apm-server/apm-server.yml
service apm-server start

chmod go-w /etc/metricbeat/metricbeat.yml
metricbeat modules enable golang
metricbeat setup
service metricbeat start

chmod go-w /etc/filebeat/filebeat.yml
service filebeat start

./out/mpindicatorgo