FROM handsonsecurity/seed-ubuntu:dev AS builder

# Install netcat
RUN apt-get update && apt-get install -y netcat

# Copy the source code to the image
COPY . /oracle
WORKDIR /oracle

RUN make


FROM handsonsecurity/seed-ubuntu:small

RUN apt-get update && apt-get install -y netcat

# Copy the executable binaries to the image
COPY --from=builder /oracle/build /oracle
WORKDIR /oracle

CMD ./server
