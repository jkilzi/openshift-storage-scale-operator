FROM registry.access.redhat.com/ubi9/nodejs-22:latest AS builder

USER root

WORKDIR /opt/app-root/src
COPY VERSION.txt .
COPY console/ .
# replace version in package.json
RUN VERSION=$(cat VERSION.txt) && sed -r -i "s|\"version\": \"0.0.1\"|\"version\": \"${VERSION}\"|;" ./package.json
RUN npm ci --ignore-scripts && npm run build

FROM registry.access.redhat.com/ubi9/nginx-120:latest
LABEL \
    com.redhat.component="Console plugin image for OpenShift Fusion Access Operator" \
    description="" \
    io.k8s.display-name="Console plugin image for OpenShift Fusion Access Operator" \
    io.k8s.description="" \
    io.openshift.tags="openshift,storage,scale" \
    distribution-scope="public" \
    name="openshift-fusion-access-console-plugin" \
    summary="Fusion Access Console Plugin" \
    release="v1.0" \
    version="v1.0" \
    maintainer="Red Hat jgil@redhat.com" \
    url="https://github.com/openshift-storage-scale/openshift-fusion-access-operator.git" \
    vendor="Red Hat, Inc." \
    License="Apache License 2.0"

COPY --from=builder /opt/app-root/src/docker/etc/default.conf "${NGINX_CONFIGURATION_PATH}"
COPY --from=builder /opt/app-root/src/dist .
USER 1001
CMD /usr/libexec/s2i/run
