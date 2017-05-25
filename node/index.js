var coap = require("coap");

console.log("Preparing COAP request");

var req = coap.request({
  method: "POST",
  pathname: "/echo",
})

// write a payload into the request
req.write("hello coap");

console.log("Sending COAP request");

req.end();

// set up a callback to output any response we get back from the server
req.on("response", function(res) {
  console.log("Received response:");

  res.pipe(process.stdout)
  res.on("end", function() {
    process.exit(0)
  })
});
