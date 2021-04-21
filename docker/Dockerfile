FROM alpine

ENV service_name=meta-server
ENV service_group=web

COPY dist/* /opt/${service_group}/${service_name}/

WORKDIR /opt/${service_group}/${service_name}

RUN chmod +x ${service_name}

ENTRYPOINT ["sh","-c","./${service_name}"]