require 'spec_helper'

describe App do
  use_vcr_cassette
  include Rack::Test::Methods

  def app
    App
  end

  before do
    header 'Content-Type', 'application/x-www-form-urlencoded'
  end

  it 'creates a slack attachment' do
    post '/', File.read(File.expand_path('../fixtures/alert.txt', __FILE__))
  end
end
