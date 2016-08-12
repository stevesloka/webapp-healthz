FROM scratch
MAINTAINER Steve Sloka <steve@stevesloka.com>
ADD mysql-healthz /mysql-healthz
ENTRYPOINT ["/webapp-healthz"]
