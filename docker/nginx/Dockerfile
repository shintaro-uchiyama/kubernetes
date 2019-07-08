FROM nginx:1.15.8
LABEL maintainer="shintaro.0112@gmail.com"

RUN mkdir -p /etc/nginx/tls
RUN chown -R root:root /etc/nginx/tls
RUN chmod -R 600 /etc/nginx/tls
ADD ./index.html /usr/share/nginx/html
