#!/usr/bin/env bash
set -e

OPENAPI_GENERATOR_VERSION=7.13.0
OPENAPI_JSON=openapi.json
OPENAPI_CLIENT_DIR=v2

# Detect flavor
sed --version >/dev/null 2>&1 && SED_FLAVOR="GNU" || SED_FLAVOR="BSD"

# Set flags based on flavor
if [[ "$SED_FLAVOR" == "BSD" ]]; then
  SED_FLAGS=(-i '')
else
  SED_FLAGS=(-i)
fi

# Pre-Generation Steps ########################################################

clean_sdk() {
  [ -d "$OPENAPI_CLIENT_DIR" ] && chmod -R u+w "$OPENAPI_CLIENT_DIR" || true
  rm -rf \
    $OPENAPI_CLIENT_DIR/docs \
    $OPENAPI_CLIENT_DIR/test \
    $OPENAPI_CLIENT_DIR/go.mod \
    $OPENAPI_CLIENT_DIR/go.sum \
    $OPENAPI_CLIENT_DIR/*.go || true
}

bump_patch() {
  GO_GENERATOR_YAML="./go-generator-config.yaml"

  current_version=$(grep "packageVersion" "$GO_GENERATOR_YAML" | awk -F': ' '{print $2}' | tr -d '"')
  IFS='.' read -r major minor patch <<<"$current_version"
  patch=$((patch + 1))
  new_version="${major}.${minor}.${patch}"

  sed "${SED_FLAGS[@]}" "s/packageVersion: \"$current_version\"/packageVersion: \"$new_version\"/" "$GO_GENERATOR_YAML"

  echo "SDK Version bumped to $new_version"
}

# Generation ##################################################################

generate_sdk() {
  docker run \
    --rm \
    --user "$(id -u):$(id -g)" \
    -v "${PWD}":/local \
    openapitools/openapi-generator-cli:v$OPENAPI_GENERATOR_VERSION generate \
    -i /local/"$OPENAPI_JSON" \
    -g go \
    --git-user-id "vulncheck-oss" \
    --git-repo-id "sdk-go-v2" \
    -c /local/go-generator-config.yaml \
    -o /local/"$OPENAPI_CLIENT_DIR"
}

generate_readme() {
  echo "Hydrating README.template"
  cog -d -o README.md ./README.template
}

# Helpers #####################################################################

check_git_status() {
  if [ -n "$(git status --porcelain)" ]; then
    echo "New API changes!"
  else
    echo "No changes detected."
    exit 0
  fi
}

get_openapi_spec() {
  # use the combined output
  curl --silent --fail \
   --url https://api.vulncheck.com/openapi/combined \
   --header "Accept: application/json" \
   >"$OPENAPI_JSON"
}

clean_openapi_spec() {
  rm -rf "$OPENAPI_JSON" || true
}

# Main ########################################################################

main() {
  # First we'll check if there are new changes, by skipping the bump_patch step
  # (otherwise check_git_status would always see a new change, i.e. the version)
  clean_sdk
  clean_openapi_spec
  get_openapi_spec
  generate_sdk
  generate_readme
  check_git_status # If there's no new changes, we exit here
  # If there's new changes, let's bump the version and prepare for a new PR
  clean_sdk
  bump_patch
  generate_sdk
  generate_readme
}

main
