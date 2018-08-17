##########################################################
# Dockerfile to build Omigad container images
# Based on scratch
##########################################################
#!/bin/bash
FROM scratch

ADD ./omigad /wst/omigad
ADD ./conf /wst/conf

WORKDIR /wst

EXPOSE 18012

CMD ["/wst/omigad"]