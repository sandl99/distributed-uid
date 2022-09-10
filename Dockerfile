FROM nginx:latest

ENV EPOCH=1288834974657
ENV NODE_BITS=10
ENV STEP_BITS=12
ENV LOG_FILE_PATH=/data/src/error_log.txt

RUN mkdir -p /data/src/
WORKDIR /data/src

COPY . /data/src/
RUN cp deploy/nginx.conf /etc/nginx/nginx.conf

RUN cp deploy/uid.service /etc/init.d/uid
RUN chmod 755 /etc/init.d/uid
RUN chmod 755 /data/src/deploy/start_service.sh

CMD /etc/init.d/nginx start && /etc/init.d/uid start
