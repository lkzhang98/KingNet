# Copyright 2022 Innkeeper lkzhang98(张良康) <lkzhang98@163.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file. The original repo for
# this file is https://github.com/lkzhang98/kingnet.

FROM BASE_IMAGE
LABEL maintainer="<lkzhang98@163.com>"

WORKDIR /opt/kingnet

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    mkdir -p /var/log/kingnet

COPY kingnet /opt/kingnet/bin/

ENTRYPOINT ["/opt/kingnet/bin/kingnet"]
CMD ["-c", "/etc/kingnet/kingnet.yaml"]