FROM scratch

COPY out/* /
COPY web/frontend/dist  /web/frontend/dist

CMD ["/kubez-linux-amd64"]