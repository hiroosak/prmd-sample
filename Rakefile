#!/usr/bin/env rake

require 'prmd'
require './lib/prmd/commands/combine.rb'

task :clean do
  `rm -rf schema/schema.{json,md}`
end

task :combine do |t|
  data , errors = '', []
  options = {
    meta: "schema/meta.yaml"
  }

  data = Prmd.combine("schema/schemata", options).to_s
  puts "Success: Combined meta: #{options[:meta]} dir: schema/schemata"

  json = JSON.parse(data)
  Prmd.verify(json).each do |error|
    errors << error
  end
  if errors.empty?
    File.write("schema/schema.json", data)
    puts "Success: Verified schema/schema.json"
  else
    puts "Error: Verify Error"
    errors.each do |err|
      puts err
    end
  end
end

task :doc do |t|
  options = {
    prepend: ["schema/overview.md"]
  }
  data = File.read("schema/schema.json")

  schema = Prmd::Schema.new(JSON.parse(data))
  # TODO: Set up template
  options[:template] = File.expand_path(File.join(File.dirname(__FILE__), '..', 'lib', 'prmd', 'templates'))

  res = Prmd.render(schema, options)
  File.write("schema/schema.md", res)
end

task :default => :combine
