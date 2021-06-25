FROM busybox
RUN {                  \
    echo "echo hello"; \
    echo "echo world"; \
} > /bin/svc

RUN echo "\
xxxx      \
" > /tmp/README


RUN echo $'#!/bin/sh\n\        
echo best\n\
echo wish\n\
' > /tmp/README2

RUN echo """\ 
$PID\
kafka\
""" > /tmp/xxx
ADD ./svc /bin/
RUN chmod +x /bin/svc /tmp/README2
CMD ["nc", "-lk", "-p", "80", "-e", "/bin/svc"]
EXPOSE 80
