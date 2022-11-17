#! /usr/bin/env ruby
# frozen_string_literal: true

def main
  n, k = gets.split.map(&:to_i)
  result = 0
  (1..n).each do |s|
    (1..n).each do |t|
      u = k - s - t
      result += 1 if u >= 1 && u <= n
    end
  end
  puts result
end

main if __FILE__ == $PROGRAM_NAME
