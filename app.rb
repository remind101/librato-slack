require 'bundler/setup'
Bundler.require :default, ENV['RACK_ENV']
require 'json'

SLACK_ACCOUNT = ENV['SLACK_ACCOUNT']
SLACK_TOKEN   = ENV['SLACK_TOKEN']

class App < Grape::API
  helpers do
    def payload
      @payload ||= Hashie::Mash.new JSON.parse(params.payload)
    end
  end

  post do
    p payload

    attachment = {
      color: '#f00'
    }

    HTTParty.post "https://#{SLACK_ACCOUNT}.slack.com",
      body: "payload=#{JSON.dump(attachments: [attachment])}"
  end
end
