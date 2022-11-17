#! /usr/bin/env ruby
# frozen_string_literal: true

def main
  gets
  f = gets.split.map(&:to_i).combination(3).any? { |a| a.sum == 1000 }
  puts f ? 'Yes' : 'No'
end

__FILE__ == $PROGRAM_NAME && main
