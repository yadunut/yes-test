#!/usr/bin/env ruby
expletive = ARGV[0]
expletive ||= "y"
expletive += "\n"

while true
    STDOUT << expletive
end
