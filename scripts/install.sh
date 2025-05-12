#! /bin/sh

# set shell options to print commands as they are executed
set -xe

# Install the lazy-history command using go install
go install github.com/sacenox/lazy-history@latest

# Add the lazy-history command to the PATH if it's not already there
if ! echo "$PATH" | grep -q "$HOME/go/bin"; then
    export PATH="$PATH:$HOME/go/bin"
fi

# bind the command to ctrl+r
bind -x '"\C-r": lazy-history'