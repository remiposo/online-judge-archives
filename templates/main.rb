#! /usr/bin/env ruby
# frozen_string_literal: true

def main
  puts gets.split.map(&:to_i).sum
end

main if __FILE__ == $PROGRAM_NAME
