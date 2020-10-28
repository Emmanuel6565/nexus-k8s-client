FROM debian
COPY ./nexus-client /nexus-client
ENTRYPOINT /nexus-client