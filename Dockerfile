FROM localhost:5000/backend-comm-mongo
MAINTAINER Pål Karlsrud <paal@128.no>

ENV BASE_DIR "/var/publishing-templates"
ENV GOPATH "/root/.go"
ENV GOBIN ${GOPATH}/bin
ENV GIN_MODE "release"

RUN git clone https://github.com/microserv/publishing-templates ${BASE_DIR}
RUN apk add --update curl go bzr

RUN cp ${BASE_DIR}/publishing-templates.ini /etc/supervisor.d/
RUN curl -o /etc/supervisor.d/mongodb.ini https://128.no/f/mongodb.ini

WORKDIR ${BASE_DIR}
RUN go get -v
RUN go build

RUN rm -rf /run && mkdir -p /run

ENV SERVICE_NAME templates

EXPOSE 80
