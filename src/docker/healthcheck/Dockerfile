ARG PORT=80
FROM busybox
EXPOSE $PORT
#HEALTHCHECK CMD nc -z localhost $PORT
HEALTHCHECK CMD false
ENTRYPOINT ["nc", "-lk", "-p", "80", "-e", "hostname"]
