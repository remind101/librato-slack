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
      post '/?account=test&token=1234', File.read(File.expand_path('../fixtures/alert.txt', __FILE__))
    end
  end
end
