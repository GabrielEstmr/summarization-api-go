FROM golang:1.20.1


# RUN apk add --no-cache git

# ELASTIC - APM
RUN curl -L -O https://artifacts.elastic.co/downloads/apm-server/apm-server-7.13.0-amd64.deb && \
    dpkg -i apm-server-7.13.0-amd64.deb
RUN chown root /etc/apm-server/apm-server.yml
COPY apm/apm-server.yml /etc/apm-server/apm-server.yml

# ELASTIC - Metrics
RUN curl -L -O https://artifacts.elastic.co/downloads/beats/metricbeat/metricbeat-7.13.0-amd64.deb && \
    dpkg -i metricbeat-7.13.0-amd64.deb
RUN chown root /etc/metricbeat/metricbeat.yml
COPY apm/metricbeat.yml /etc/metricbeat/metricbeat.yml

# ELASTIC - Logs
RUN curl -L -O https://artifacts.elastic.co/downloads/beats/filebeat/filebeat-7.13.0-amd64.deb && \
    dpkg -i filebeat-7.13.0-amd64.deb
RUN chown root /etc/filebeat/filebeat.yml
COPY apm/filebeat.yml /etc/filebeat/filebeat.yml











# Set the Current Working Directory inside the container
WORKDIR /app/mpindicatorgo

COPY apm/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# We want to populate the module cache based on the go.{mod,sum} files.
COPY src/go.mod .
COPY src/go.sum .

RUN go mod download

COPY src/ .


# Build the Go app
RUN go build -o ./out/mpindicatorgo .

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
# CMD ["./out/mpindicatorgo"]


ENTRYPOINT [ "/entrypoint.sh" ]