FROM ubuntu:20.04

ADD ./permastar-server /opt/permastar-server
ADD start.sh /opt/start.sh

RUN apt-get -qq update &&\
    apt-get -qq install -y --no-install-recommends ca-certificates curl &&\
    chmod +x /opt/permastar-server

CMD bash /opt/start.sh
