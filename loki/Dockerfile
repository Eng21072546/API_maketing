FROM grafana/promtail

COPY promtail-config.yaml /etc/promtail/promtail-config.yaml

CMD ["promtail", "-config.file=/etc/promtail/promtail-config.yaml"]
