#! /usr/bin/env ruby
# frozen_string_literal: true

def main
  _, k = gets.split.map(&:to_i)
  ps = gets.split.map(&:to_i)
  qs = gets.split.map(&:to_i)

  f = ps.product(qs).any? do |a|
    a.sum == k
  end

  puts f ? 'Yes' : 'No'
end

main if __FILE__ == $PROGRAM_NAME
