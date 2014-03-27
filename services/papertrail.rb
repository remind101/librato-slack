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
        text: "Search <#{event.saved_search.html_search_url}|#{event.saved_search.name}> has triggered!"
      }

      payload[:attachments] = event.events.map do |event|
        { fields: [{ title: event.hostname, value: event.message }] }
      end

      post payload
    end
  end
end
