#!/bin/bash
set -euo pipefail

set +u; #relax unbound variable check for this part
if [ ! -z "${DEBUG}" ]; then
  echo "DEBUG switch specified, printing current gcloud config"
  gcloud config list --all
fi
set -u;

#Check that the vulnerability check policy passes (no signature )
#This is done for ALL branches.
/kritis/signer \
    -v=10 \
    -alsologtostderr \
    -mode=check-only \
    -image=$(/bin/cat $DIGEST_FILENAME) \
    -policy=${VULNZ_CHECK_POLICY} 

# On master or a hotfix branch, add a signature that indicates that the build is vetted for promotion in our environments (coming from master or a hotfix version)
HOTFIX_PATTERN='^hotfix/([0-9]+_[0-9]+_[0-9]+)$'
if [[ "${BRANCH_NAME}" = "master" ]] || [[ "${BRANCH_NAME}" =~ ${HOTFIX_PATTERN} ]]; then
    echo "Signing image with the '${ATTESTOR_NAME}' binauthz attestor.";

    #In theory, this should not be required when in Cloud Build. Only way to test is to get there.
    set +u;
    if [ ! -z "${GOOGLE_APPLICATION_CREDENTIALS}" ]; then
        echo "Activating service account from ${GOOGLE_APPLICATION_CREDENTIALS}";
        gcloud auth activate-service-account --key-file=${GOOGLE_APPLICATION_CREDENTIALS};
    fi
    set -u;

    gcloud beta container binauthz attestations sign-and-create \
    --artifact-url="$(/bin/cat $DIGEST_FILENAME)" \
    --attestor="${ATTESTOR_NAME}" \
    --keyversion="${KMS_KEY_NAME}" \
    --project=${PROJECT_ID}
else
    echo "Not signing the image with binauthz attestor, wrong branch(current branch: ${BRANCH_NAME}).";
fi
