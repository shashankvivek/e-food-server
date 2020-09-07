FROM swaggerapi/swagger-ui@sha256:5a65209869b6ca41e1382ff2ea260a06d95dccde2ed7e09ec61ff0e8d37932ce

COPY api/index.html /usr/share/nginx/html/index.html
