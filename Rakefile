#!/usr/bin/env rake

begin
  require 'rspec/core/rake_task'
  desc "Run specs"
  RSpec::Core::RakeTask.new do |t|
    t.pattern = 'spec/**/*_spec.rb'
  end

  task :default => [:spec]
rescue LoadError
  # NOOP
end
