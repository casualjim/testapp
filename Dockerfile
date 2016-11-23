FROM casualjim/base:latest
MAINTAINER Ivan Porto Carrero <ivan@flanders.co.nz> (@casualjim)

ADD testapp /testapp
EXPOSE 8080
ENTRYPOINT ["/testapp"]
