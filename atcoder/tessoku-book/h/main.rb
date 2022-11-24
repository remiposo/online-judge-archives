#! /usr/bin/env ruby
# frozen_string_literal: true

def main
  h, w = gets.split.map(&:to_i)
  x = [Array.new(w + 1, 0)]
  h.times do
    x << [0] + gets.split.map(&:to_i)
  end
  (1..h).each do |i|
    (1..w).each do |j|
      x[i][j] = x[i][j - 1] + x[i][j]
    end
  end
  (1..w).each do |i|
    (1..h).each do |j|
      x[j][i] = x[j - 1][i] + x[j][i]
    end
  end
  gets.to_i.times do
    a, b, c, d = gets.split.map(&:to_i)
    puts x[c][d] + x[a - 1][b - 1] - x[c][b - 1] - x[a - 1][d]
  end
end

main if __FILE__ == $PROGRAM_NAME
