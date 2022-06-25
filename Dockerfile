FROM golang:1.18

# Build arguments
ARG UID

# Install base dependencies
RUN apt-get update && \
    apt-get install build-essential librdkafka-dev -y

# Create user user 
RUN useradd -ms /bin/bash -u ${UID} code
RUN usermod -G root code
USER code

# Configure environment
WORKDIR $GOPATH/src
ENV PATH="/go/bin:${PATH}"

CMD [ "tail", "-f", "/dev/null" ]