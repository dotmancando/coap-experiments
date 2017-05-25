require "coap"

puts "Hitting pulse endpoint"
client = CoAP::Client.new(host: "0.0.0.0")
resp = client.get("/pulse")
puts resp.payload

puts "Hitting echo endpoint"
resp = client.post("/echo", nil, nil, "banana")
puts resp.payload
