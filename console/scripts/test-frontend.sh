#!/usr/bin/env bash

set -euo pipefail
[[ -n "${DEBUGME+x}" ]] && set -x


# https://ci-operator-configresolver-ui-ci.apps.ci.l2s4.p1.openshiftapps.com/help#env
OPENSHIFT_CI=${OPENSHIFT_CI:=false}
ARTIFACT_DIR=${ARTIFACT_DIR:=/tmp/artifacts}

npm run i18n
GIT_STATUS="$(git status --short --untracked-files -- locales)"
if [ -n "$GIT_STATUS" ]; then
  echo "i18n files are not up to date. Run 'npm run i18n' then commit changes."
  git --no-pager diff
  exit 1
fi