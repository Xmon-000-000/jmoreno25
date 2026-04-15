FROM ubuntu:latest

ENV SCRIPTS=/var/scripts

RUN apt-get update && \
    apt-get install -y netcat-openbsd && \
    rm -rf /var/lib/apt/lists/* && \
    mkdir -p $SCRIPTS

COPY backend.sh $SCRIPTS/backend.sh
RUN chmod +x $SCRIPTS/backend.sh

EXPOSE 8080

CMD ["/var/scripts/backend.sh"]
