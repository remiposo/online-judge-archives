#! /usr/bin/env ruby
# frozen_string_literal: true

def main
  puts gets.split(' ').map(&:to_i).sum
end

__FILE__ == $PROGRAM_NAME && main
