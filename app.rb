require 'bundler/setup'
Bundler.require :default, ENV['RACK_ENV']
require 'json'

require_relative './services/librato'

class App < Grape::API
  mount Librato
end
