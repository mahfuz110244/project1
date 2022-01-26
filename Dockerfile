FROM scratch
COPY project1 ./
COPY config.env ./
EXPOSE 4003/tcp
EXPOSE 5003/tcp
ENTRYPOINT ["/project1"]