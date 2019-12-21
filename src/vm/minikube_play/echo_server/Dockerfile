FROM busybox
EXPOSE 8888
COPY ./xecho /bin/xecho
CMD ["nc","-ll","-p","8888","-e","/bin/xecho"]
