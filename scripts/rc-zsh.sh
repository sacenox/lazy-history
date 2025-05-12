

# Define the lazy-history widget function
lazy-history-widget() {
  BUFFER="lazy-history"
  zle accept-line
}

# Create the widget
zle -N lazy-history-widget

# Bind the widget to Ctrl+R
bindkey '^R' lazy-history-widget