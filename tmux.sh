#!/bin/bash

SESSION="mdpdf"

create_window() {
  local name=$1
  local dir=$2
  local cmd=$3

  tmux new-window -t "$SESSION" -n "$name"
  tmux send-keys -t "$SESSION":"$name" "cd $dir" C-m

  if [ -n "$cmd" ]; then
    tmux send-keys -t "$SESSION":"$name" "$cmd" C-m
  fi
}

# Create session with first window "home"
tmux new-session -d -s "$SESSION" -n "home"
tmux send-keys -t "$SESSION":home "cd $HOME" C-m

# Add more windows
create_window "nvim" "$HOME/mdpdf" "nvim"
create_window "cli"  "$HOME/mdpdf"
create_window "notes"       "$HOME/notes" "nvim"

# Attach
tmux attach -t "$SESSION"
