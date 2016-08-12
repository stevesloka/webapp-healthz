FROM scratch
MAINTAINER Steve Sloka <steve@stevesloka.com>
ADD webapp-healthz /webapp-healthz
ENTRYPOINT ["/webapp-healthz"]
