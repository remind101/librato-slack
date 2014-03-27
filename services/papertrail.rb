class Papertrail < Grape::API
  namespace :papertrail do
    helpers do
      def event
        @event ||= Hashie::Mash.new JSON.parse(params.payload)
      end
    end

    post do
      title = "Search <#{event.saved_search.html_search_url}|#{event.saved_search.name}> has triggered!\n"

      payload = {
        username: 'Papertrail',
        text: title + event.events.map { |event| "#{event.hostname}: #{event.message}" }.join("\n")
      }

      post payload
    end
  end
end
