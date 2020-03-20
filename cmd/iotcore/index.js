const iotCore = require("aws-iot-device-sdk");
const services = require("./api/driver/driver_grpc_pb");
const messages = require("./api/driver/driver_pb");
const grpc = require("grpc");

let client;

const DRIVER_HOST = process.env.DRIVER_HOST;
const IOT_KEY_PATH = process.env.IOT_KEY_PATH;
const IOT_CERT_PATH = process.env.IOT_CERT_PATH;
const IOT_CLIENT_ID = process.env.IOT_CLIENT_ID;
const IOT_HOST = process.env.IOT_HOST;

const IOT_CA_PATH = "root-CA.crt";

const iotDevice = new iotCore.device({
  keyPath: IOT_KEY_PATH,
  certPath: IOT_CERT_PATH,
  caPath: IOT_CA_PATH,
  clientId: IOT_CLIENT_ID,
  host: IOT_HOST
});

iotDevice.on("connect", function() {
  client = new services.DriverClient(
    DRIVER_HOST,
    grpc.credentials.createInsecure()
  );

  iotDevice.subscribe("topic_1");
});

iotDevice.on("message", function(topic, payload) {
  const action = payload.toString();

  request = new messages.CommandRequest();

  client.up(request, function(err, response) {
    if (err != nil) {
      console.log(err);
    } else {
      console.log(response);
    }
  });

  switch (action) {
    case "left":
      client.left(request, (err, res) => {
        if (err != nil) {
          //TODO: Add logging
          console.log(err);
        } else {
          //TODO: Add logging
          console.log(res);
        }
      });
      break;
    case "up":
      client.up(request, (err, res) => {
        if (err != nil) {
          //TODO: Add logging
          console.log(err);
        } else {
          //TODO: Add logging
          console.log(res);
        }
      });
      break;
    case "right":
      client.right(request, (err, res) => {
        if (err != nil) {
          //TODO: Add logging
          console.log(err);
        } else {
          //TODO: Add logging
          console.log(res);
        }
      });
      break;
    case "down":
      client.down(request, (err, res) => {
        if (err != nil) {
          //TODO: Add logging
          console.log(err);
        } else {
          //TODO: Add logging
          console.log(res);
        }
      });
      break;
    case "upLeft":
      client.upLeft(request, (err, res) => {
        if (err != nil) {
          //TODO: Add logging
          console.log(err);
        } else {
          //TODO: Add logging
          console.log(res);
        }
      });
      break;
    case "downLeft":
      client.downLeft(request, (err, res) => {
        if (err != nil) {
          //TODO: Add logging
          console.log(err);
        } else {
          //TODO: Add logging
          console.log(res);
        }
      });
      break;
    case "upRight":
      client.upRight(request, (err, res) => {
        if (err != nil) {
          //TODO: Add logging
          console.log(err);
        } else {
          //TODO: Add logging
          console.log(res);
        }
      });
      break;
    case "downRight":
      client.downRight(request, (err, res) => {
        if (err != nil) {
          //TODO: Add logging
          console.log(err);
        } else {
          //TODO: Add logging
          console.log(res);
        }
      });
      break;
    case "stop":
      client.stop(request, (err, res) => {
        if (err != nil) {
          //TODO: Add logging
          console.log(err);
        } else {
          //TODO: Add logging
          console.log(res);
        }
      });
      break;
    case "fire":
      client.fire(request, (err, res) => {
        if (err != nil) {
          //TODO: Add logging
          console.log(err);
        } else {
          //TODO: Add logging
          console.log(res);
        }
      });
      break;
  }

  console.log(topic, payload.toString());
});
