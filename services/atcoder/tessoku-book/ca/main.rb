#! /usr/bin/env ruby
# frozen_string_literal: true

def main
  a, b = gets.split(' ').map(&:to_i)
  f = false
  (a..b).each do |i|
    if (100 % i).zero?
      f = true
      break
    end
  end
  puts f ? 'Yes' : 'No'
end

__FILE__ == $PROGRAM_NAME && main
