#!/bin/bash

set -e

# Execute script commands from gutenberg-mobile root directory
GB_MOBILE_PATH="$1"
 cd "$GB_MOBILE_PATH"

#####
# PR Milestone check
#####

function check_num_milestone_prs() {
    MILESTONE_NAME="$1"
    curl -s "https://api.github.com/repos/wordpress-mobile/gutenberg-mobile/milestones" \
      | jq ".[] | select(.title == \"$MILESTONE_NAME\") | .open_issues"
}

function check_if_version_has_pending_prs_for_milestone() {
    number_milestone_prs=$(check_num_milestone_prs "$VERSION_NUMBER")
    # If we got no results and the version ends with ".0"
    if [[ -z "$number_milestone_prs" ]] && [[ "$VERSION_NUMBER" =~ [0-9]+\.[0-9]+\.0$ ]]; then
        # Remove the ending ".0" and check again because we usually drop ".0" from our
        # milestones. For example, the milestone for 1.34.0 was 1.34
        version_without_dot_zero=$(echo "$VERSION_NUMBER" | rev | cut -c 3- | rev)
        number_milestone_prs=$(check_num_milestone_prs "$version_without_dot_zero")
    fi
    echo "$number_milestone_prs"
}


#####
# Check for Aztec Release Versions
#####

SEMANTIC_VERSION_REGEX='\d+\.\d+\.\d+'

function check_android_aztec_is_release_version() {
    react_native_aztec_gradle="$GB_MOBILE_PATH/gutenberg/packages/react-native-aztec/android/build.gradle"
    release_version=$(grep aztecVersion "$react_native_aztec_gradle" | grep -oE "$SEMANTIC_VERSION_REGEX")
    if [[ -z "$release_version" ]]; then
        echo "A release version for AztecAndroid was not found in $react_native_aztec_gradle"
    fi
}

function check_ios_aztec_is_release_version() {
    result=''

    podspec_file="$GB_MOBILE_PATH/gutenberg/packages/react-native-aztec/RNTAztecView.podspec"
    aztec_version=$(grep WordPress-Aztec-iOS "$podspec_file" | grep -oE "$SEMANTIC_VERSION_REGEX")
    if [[ -z "$aztec_version" ]]; then
        result="A release version for WordPress-Aztec-iOS was not found in $podspec_file"
    fi

    echo "$result"
}

# If any problems, the problems are each printed on their own line
function check_android_and_ios_aztec_versions() {
    android_result=$(check_android_aztec_is_release_version)
    ios_result=$(check_ios_aztec_is_release_version)
    if [[ -n "$android_result" ]]; then
        printf "%s\n%s" "$android_result" "$ios_result"
    else
        echo "$ios_result"
    fi
}

printf "%s \n" "$(check_android_and_ios_aztec_versions)"
