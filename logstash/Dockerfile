FROM elastic/logstash:7.15.0

COPY config/logstash-sample.conf /usr/share/logstash/config/logstash.conf

COPY config/logstash.yml /usr/share/logstash/config/logstash.yml

COPY config/jvm.options /usr/share/logstash/config/jvm.options

COPY config/log4j2.properties  /usr/share/logstash/config/log4j2.properties

COPY config/pipelines.yml  /usr/share/logstash/config/pipelines.yml

COPY config/startup.options /usr/share/logstash/config/startup.options