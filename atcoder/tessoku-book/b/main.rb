#! /usr/bin/env ruby
# frozen_string_literal: true

def main
  _, x = gets.split(' ')
  puts gets.split(' ').include?(x) ? 'Yes' : 'No'
end

__FILE__ == $PROGRAM_NAME && main
