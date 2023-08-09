FROM BASE_IMAGE
LABEL maintainer="<lkzhang98@163.com>"

WORKDIR /opt/kingnet

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    mkdir -p /var/log/kingnet

COPY kingnet /opt/kingnet/bin/

ENTRYPOINT ["/opt/kingnet/bin/kingnet"]
CMD ["-c", "/etc/kingnet/kingnet.yaml"]