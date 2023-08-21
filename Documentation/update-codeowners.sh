#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source_dir="$(cd "${script_dir}/.." && pwd)"

echo ".." > "${script_dir}/codeowners.rst"
echo "    This file was autogenerated via Documentation/update-codeowners.sh, do not edit manually" \
    >> "${script_dir}/codeowners.rst"
echo >> "${script_dir}/codeowners.rst"
sed '/END_CODEOWNERS_DOCS/q' "${source_dir}/CODEOWNERS" \
    | head -n-2 \
    | cut -c 3- \
    | sed 's|@cilium/\([-_a-zA-Z0-9]*\)\(:\?\)|`@cilium/\1 <https://github.com/orgs/cilium/teams/\1>`__\2|g' \
    >> "${script_dir}/codeowners.rst"
