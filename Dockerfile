FROM golang:onbuild
COPY main /
RUN mkdir /client
COPY client/* /client/
CMD ["/main"]