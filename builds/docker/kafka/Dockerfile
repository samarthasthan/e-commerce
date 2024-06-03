FROM bitnami/kafka:3.6.1

WORKDIR /opt/bitnami/kafka/bin

USER root

RUN apt-get update && apt-get -y dist-upgrade
RUN apt-get -y install build-essential libssl-dev libffi-dev libblas3 libc6 liblapack3 gcc
RUN apt install -y netcat-traditional

ENV KAFKA_ENABLE_KRAFT=yes
ENV ALLOW_PLAINTEXT_LISTENER=yes

# Node-specific configurations
ENV KAFKA_CFG_NODE_ID=1
ENV KAFKA_CFG_PROCESS_ROLES=broker,controller

# Listener configurations
ENV KAFKA_CFG_CONTROLLER_LISTENER_NAMES=CONTROLLER
ENV KAFKA_CFG_INTER_BROKER_LISTENER_NAME=CLIENT
ENV KAFKA_CFG_LISTENERS=CLIENT://:29092,EXTERNAL://:9092,CONTROLLER://:9093
ENV KAFKA_CFG_LISTENER_SECURITY_PROTOCOL_MAP=CONTROLLER:PLAINTEXT,CLIENT:PLAINTEXT,EXTERNAL:PLAINTEXT
ENV KAFKA_CFG_ADVERTISED_LISTENERS=CLIENT://kafka:29092,EXTERNAL://localhost:9092

# ENV KAFKA_HEAP_OPTS="-Xmx256m -Xms265M"

# Controller quorum voters
ENV KAFKA_CFG_CONTROLLER_QUORUM_VOTERS=1@127.0.0.1:9093

COPY --chmod=755 kafka_setup.sh ./
COPY --chmod=755 wait-for.sh ./