FROM prom/prometheus:latest

# copy the config file
COPY prometheus.yml /etc/prometheus/prometheus.yml

## use a non-root user
RUN adduser --disabled-password -u 1000 user
USER user
