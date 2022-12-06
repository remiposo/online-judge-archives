#! /usr/bin/env ruby
# frozen_string_literal: true

def main
  _, x = gets.chomp.split.map(&:to_i)
  as = gets.chomp.split.map(&:to_i)
  idx = as.bsearch_index { |a| a >= x }
  puts idx + 1
end

main if __FILE__ == $PROGRAM_NAME
