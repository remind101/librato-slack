require 'bundler/setup'
Bundler.require :default, ENV['RACK_ENV']
require 'json'

require_relative './services/librato'
require_relative './services/papertrail'

class App < Grape::API
  helpers do
    def post(payload)
      payload[:channel] = params.channel if params.channel?
      HTTParty.post "https://#{params.account}.slack.com/services/hooks/incoming-webhook?token=#{params.token}",
        body: "payload=#{JSON.dump(payload)}"
    end
  end

  params do
    requires :account
    requires :token
    optional :channel
  end

  mount Librato
  mount Papertrail
end
