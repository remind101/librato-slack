class Papertrail < Grape::API
  namespace :papertrail do
    helpers do
      def event
        @event ||= Hashie::Mash.new JSON.parse(params.payload)
      end
    end

    post do
      payload = {
        username: 'Papertrail',
        text: event.events.map { |event| "#{event.hostname}: #{event.message}" }.join("\n")
      }

      post payload
    end
  end
end
