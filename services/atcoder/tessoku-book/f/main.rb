#! /usr/bin/env ruby
# frozen_string_literal: true

def main
  _, q = gets.split.map(&:to_i)
  as = gets.split.map(&:to_i)
  sas = as.each_with_object([0]) { |a, obj| obj << obj[-1] + a }
  q.times do
    l, r = gets.split.map(&:to_i)
    puts sas[r] - sas[l-1]
  end
end

main if __FILE__ == $PROGRAM_NAME
