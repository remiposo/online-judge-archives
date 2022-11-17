#! /usr/bin/env ruby
# frozen_string_literal: true

def main
  d = gets.to_i
  ps = Array.new(d + 1, 0)
  n = gets.to_i
  n.times do
    l, r = gets.split.map(&:to_i)
    ps[l - 1] += 1
    ps[r] -= 1
  end
  ans = 0
  d.times do |i|
    ans += ps[i]
    puts ans
  end
end

main if __FILE__ == $PROGRAM_NAME
