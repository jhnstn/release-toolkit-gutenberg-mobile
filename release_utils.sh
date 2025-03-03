#!/bin/bash

# Utils adapted from https://github.com/Homebrew/install/blob/master/install.sh
# string formatters
if [[ -t 1 ]]; then
    tty_escape() { printf "\033[%sm" "$1"; }
else
    tty_escape() { :; }
fi
tty_mkbold() { tty_escape "1;$1"; }
tty_underline="$(tty_escape "4;39")"
tty_blue="$(tty_mkbold 34)"
tty_red="$(tty_mkbold 31)"
tty_reset="$(tty_escape 0)"


# Takes multiple arguments and prints them joined with single spaces in-between
# while escaping any spaces in arguments themselves
shell_join() {
    local arg
    printf "%s" "$1"
    shift
    for arg in "$@"; do
        printf " "
        printf "%s" "${arg// /\ }"
    done
}

# Takes multiple arguments, joins them and prints them in a colored format
ohai() {
  printf "${tty_blue}==> %s${tty_reset}\n" "$(shell_join "$@")"
}

# Takes a single argument and prints it in a colored format
warn() {
  printf "${tty_underline}${tty_red}Warning${tty_reset}: %s\n" "$1"
}

# Takes a single argument, prints it in a colored format and aborts the script
abort() {
    printf "\n${tty_red}%s${tty_reset}\n" "$1"
    exit 1
}

# Takes multiple arguments consisting a command and executes it. If the command
# is not successful aborts the script, printing the failed command and its
# arguments in a colored format.
#
# Returns the executed command's result if it's successful.
execute() {
    if ! "$@"; then
        abort "$(printf "Failed during: %s" "$(shell_join "$@")")"
    fi
}

# Takes multiple arguments consisting a command and executes it.
# If the command is not successful, it will be retried for a maximum of 10 times.
# In case it keeps failing, the script will be aborted, printing the failed command
# and its arguments in a colored format.
#
# Returns the executed command's result if it's successful.
execute_until_succeeds() {
    local max_retry=10
    local retry_count=0
    until "$@"
    do
        sleep 1
        [[ retry_count -eq "$max_retry" ]] && abort "$(printf "Failed after retrying $retry_count times: %s" "$(shell_join "$@")")"
        ((retry_count++))
        warn "Retrying '$(shell_join "$@")' command after $retry_count times..."
    done
}

#####
# Confirm to Proceed Prompt
#####

# Accepts a single argument: a yes/no question (ending with a ? most likely) to ask the user
function confirm_to_proceed() {
    read -r -p "$1 (y/n) " -n 1
    echo ""
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        abort "Aborting release..."
    fi
}

# Automation steps
#
# These could live in a dedicated file, because "utils" is a misleading name where to put them.
# However, they depend on functions such as execute and ohai, so for the moment it's easier to have them in here.
function update_ios_gutenberg_config() {
  ios_config_file=$1
  ref=$2
  github_org=$3

  ohai "Update GitHub organization and gutenberg-mobile ref"

  test -f "$ios_config_file" || abort "Error: Could not find $ios_config_file"

  yq eval ".github_org= \"$github_org\"" "$ios_config_file" -i || abort "Error: Failed updating GitHub organization in $ios_config_file"

  # Make iOS point to the given commit, removing the configured tag, if any
  yq eval ".ref.commit = \"$ref\"" "$ios_config_file" -i || abort "Error: Failed updating gutenberg ref in $ios_config_file (part 1 of 2, setting the commit)"
  yq eval 'del(.ref.tag)' "$ios_config_file" -i || abort "Error: Failed updating gutenberg ref in $ios_config_file (part 2 of 2, commenting the tag)"

  execute "bundle" "install"
  execute_until_succeeds "rake" "dependencies"

  execute "git" "add" "Podfile" "Podfile.lock" "$ios_config_file"
  execute "git" "commit" "-m" "Release script: Update gutenberg-mobile ref"
}
