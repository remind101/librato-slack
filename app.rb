require 'bundler/setup'
Bundler.require :default, ENV['RACK_ENV']
require 'json'

class App < Grape::API
  helpers do
    def alert
      @alert ||= Hashie::Mash.new JSON.parse(params.payload)
    end
  end

  post do
    payload = {
      username: 'Librato',
      text: "Alert #{alert.alert.name} has triggered!"
    }

    payload[:attachments] = alert.violations.map do |source, violations|
      fields = []
      fields << { title: 'Source', value: source }

      violations.each do |violation|
        fields << { title: violation.metric, value: violation.value, short: true }
      end

      { fields: fields }
    end

    HTTParty.post "https://#{params.account}.slack.com/services/hooks/incoming-webhook?token=#{params.token}",
      body: "payload=#{JSON.dump(payload)}"
  end
end
