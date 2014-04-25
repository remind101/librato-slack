class Librato < Grape::API
  namespace :librato do
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
          fields << { title: violation.metric, value: violation.value.to_s, short: true }
        end

        { fields: fields }
      end

      post payload
    end
  end
end
