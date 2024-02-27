FROM centos:7

WORKDIR /

RUN mkdir log

COPY main-server main-server
COPY conf/config_*.toml /conf

EXPOSE  8080

ENTRYPOINT ["/main-server"]