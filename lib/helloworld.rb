# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: helloworld.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "helloworld.Person" do
    optional :name, :string, 1
  end
  add_message "helloworld.Greeting" do
    optional :message, :string, 1
  end
end

module Helloworld
  Person = Google::Protobuf::DescriptorPool.generated_pool.lookup("helloworld.Person").msgclass
  Greeting = Google::Protobuf::DescriptorPool.generated_pool.lookup("helloworld.Greeting").msgclass
end
