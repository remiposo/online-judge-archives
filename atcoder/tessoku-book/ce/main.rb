#! /usr/bin/env ruby
# frozen_string_literal: true

def main
  gets
  as = gets.split.map(&:to_i)
  sy = [0]
  sn = [0]
  as.each do |a|
    sy << (a == 1 ? sy[-1] + 1 : sy[-1])
    sn << (a == 0 ? sn[-1] + 1 : sn[-1])
  end
  q = gets.to_i
  q.times do
    l, r = gets.split.map(&:to_i)
    y = sy[r] - sy[l - 1]
    n = sn[r] - sn[l - 1]
    if y > n
      puts 'win'
    elsif y < n
      puts 'lose'
    else
      puts 'draw'
    end
  end
end

main if __FILE__ == $PROGRAM_NAME
