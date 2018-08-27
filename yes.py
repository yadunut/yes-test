#!/usr/bin/env python3

import sys

expletive = "y"
if len(sys.argv) != 1:
    expletive = sys.argv[1]
expletive += "\n"
while True:
    sys.stdout.write(expletive)

# [7.44MiB/s]
