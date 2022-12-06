#! /usr/bin/env ruby
# frozen_string_literal: true

def main
  _, k = gets.chomp.split.map(&:to_i)
  as = gets.chomp.split.map(&:to_i)
  result = (1..10**9).bsearch do |i|
    s = as.inject(0) { |res, a| res + i / a }
    s >= k
  end
  puts result
end

main if __FILE__ == $PROGRAM_NAME
