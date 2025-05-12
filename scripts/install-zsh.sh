#! /bin/zsh

# set shell options to print commands as they are executed
set -e

# Install the lazy-history command using go install
echo "Installing lazy-history"
go install github.com/sacenox/lazy-history@latest

# Add the lazy-history command to the PATH if it's not already there
echo "Adding lazy-history to PATH"
if ! echo "$PATH" | grep -q "$HOME/go/bin"; then
    export PATH="$PATH:$HOME/go/bin"
fi

# TODO: Uncomment this when we have a way to add query with the tui
# Get the directory where this script is located
#SCRIPT_DIR="$(cd "$(dirname "${(%):-%x}")" && pwd)"

# Source the rc.sh file from the same directory as this script
# into the zsh rc file
#echo "adding rc-zsh.sh to ~/.zshrc"
#cat "$SCRIPT_DIR/rc-zsh.sh" >> ~/.zshrc

echo "Done! Please restart your shell.\nOr source the ~/.zshrc file directly."
