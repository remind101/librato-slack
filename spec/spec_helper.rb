ENV['RACK_ENV'] = 'test'

require_relative '../app.rb'
Bundler.require :default, :test

# Requires supporting ruby files with custom matchers and macros, etc,
# in spec/support/ and its subdirectories.
Dir[File.join(File.dirname(__FILE__), "support/**/*.rb")].each {|f| require f}

VCR.config do |config|
  config.cassette_library_dir = 'spec/cassettes'
  config.default_cassette_options = {
    match_requests_on: [:method, :uri, :body],
    record: :once,
    allow_playback_repeats: false,
    allow_unused_http_interactions: false,
  }
  config.stub_with :webmock
  config.configure_rspec_metadata!
end

RSpec.configure do |config|
  config.order = 'random'
  config.filter_run :focus => true
  config.run_all_when_everything_filtered = true
  config.extend VCR::RSpec::Macros
  config.treat_symbols_as_metadata_keys_with_true_values = true
end
