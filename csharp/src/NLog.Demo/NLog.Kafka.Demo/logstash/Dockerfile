FROM logstash:7.2.0 as base
WORKDIR /opt/logstash

FROM BASE AS final
WORKDIR /opt/logstash
ENTRYPOINT ["bin/logstash","-f", "/opt/logstash/conf.d/"]