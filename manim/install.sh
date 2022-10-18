#!/bin/bash

# This is slightly different from the official OSX instructions:
# https://docs.manim.community/en/stable/installation/macos.html
# Had to install Cython to get manimpango going on an M1, was advised here:
# https://github.com/ManimCommunity/ManimPango/issues/94

brew install py3cairo ffmpeg
brew install pango scipy
pip3 install Cython
pip3 install manimpango
pip3 install manim
