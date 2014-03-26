require 'spec_helper'

describe App do
  include Rack::Test::Methods

  def app
    App
  end

  before do
    header 'Content-Type', 'application/x-www-form-urlencoded'
  end

  describe 'librato' do
    it 'sends a message', :vcr do
      post '/librato?account=test&token=1234', File.read(File.expand_path('../fixtures/librato/alert.txt', __FILE__))
    end
  end

  describe 'papertrail' do
    it 'sends a message', :vcr do
      post '/papertrail?account=test&token=1234', File.read(File.expand_path('../fixtures/papertrail/event.txt', __FILE__))
    end
  end
end
