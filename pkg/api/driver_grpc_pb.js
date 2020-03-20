// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var driver_pb = require('./driver_pb.js');

function serialize_driver_CommandReply(arg) {
  if (!(arg instanceof driver_pb.CommandReply)) {
    throw new Error('Expected argument of type driver.CommandReply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_driver_CommandReply(buffer_arg) {
  return driver_pb.CommandReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_driver_CommandRequest(arg) {
  if (!(arg instanceof driver_pb.CommandRequest)) {
    throw new Error('Expected argument of type driver.CommandRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_driver_CommandRequest(buffer_arg) {
  return driver_pb.CommandRequest.deserializeBinary(new Uint8Array(buffer_arg));
}


var DriverService = exports.DriverService = {
  left: {
    path: '/driver.Driver/Left',
    requestStream: false,
    responseStream: false,
    requestType: driver_pb.CommandRequest,
    responseType: driver_pb.CommandReply,
    requestSerialize: serialize_driver_CommandRequest,
    requestDeserialize: deserialize_driver_CommandRequest,
    responseSerialize: serialize_driver_CommandReply,
    responseDeserialize: deserialize_driver_CommandReply,
  },
  up: {
    path: '/driver.Driver/Up',
    requestStream: false,
    responseStream: false,
    requestType: driver_pb.CommandRequest,
    responseType: driver_pb.CommandReply,
    requestSerialize: serialize_driver_CommandRequest,
    requestDeserialize: deserialize_driver_CommandRequest,
    responseSerialize: serialize_driver_CommandReply,
    responseDeserialize: deserialize_driver_CommandReply,
  },
  right: {
    path: '/driver.Driver/Right',
    requestStream: false,
    responseStream: false,
    requestType: driver_pb.CommandRequest,
    responseType: driver_pb.CommandReply,
    requestSerialize: serialize_driver_CommandRequest,
    requestDeserialize: deserialize_driver_CommandRequest,
    responseSerialize: serialize_driver_CommandReply,
    responseDeserialize: deserialize_driver_CommandReply,
  },
  down: {
    path: '/driver.Driver/Down',
    requestStream: false,
    responseStream: false,
    requestType: driver_pb.CommandRequest,
    responseType: driver_pb.CommandReply,
    requestSerialize: serialize_driver_CommandRequest,
    requestDeserialize: deserialize_driver_CommandRequest,
    responseSerialize: serialize_driver_CommandReply,
    responseDeserialize: deserialize_driver_CommandReply,
  },
  upLeft: {
    path: '/driver.Driver/UpLeft',
    requestStream: false,
    responseStream: false,
    requestType: driver_pb.CommandRequest,
    responseType: driver_pb.CommandReply,
    requestSerialize: serialize_driver_CommandRequest,
    requestDeserialize: deserialize_driver_CommandRequest,
    responseSerialize: serialize_driver_CommandReply,
    responseDeserialize: deserialize_driver_CommandReply,
  },
  downLeft: {
    path: '/driver.Driver/DownLeft',
    requestStream: false,
    responseStream: false,
    requestType: driver_pb.CommandRequest,
    responseType: driver_pb.CommandReply,
    requestSerialize: serialize_driver_CommandRequest,
    requestDeserialize: deserialize_driver_CommandRequest,
    responseSerialize: serialize_driver_CommandReply,
    responseDeserialize: deserialize_driver_CommandReply,
  },
  upRight: {
    path: '/driver.Driver/UpRight',
    requestStream: false,
    responseStream: false,
    requestType: driver_pb.CommandRequest,
    responseType: driver_pb.CommandReply,
    requestSerialize: serialize_driver_CommandRequest,
    requestDeserialize: deserialize_driver_CommandRequest,
    responseSerialize: serialize_driver_CommandReply,
    responseDeserialize: deserialize_driver_CommandReply,
  },
  downRight: {
    path: '/driver.Driver/DownRight',
    requestStream: false,
    responseStream: false,
    requestType: driver_pb.CommandRequest,
    responseType: driver_pb.CommandReply,
    requestSerialize: serialize_driver_CommandRequest,
    requestDeserialize: deserialize_driver_CommandRequest,
    responseSerialize: serialize_driver_CommandReply,
    responseDeserialize: deserialize_driver_CommandReply,
  },
  stop: {
    path: '/driver.Driver/Stop',
    requestStream: false,
    responseStream: false,
    requestType: driver_pb.CommandRequest,
    responseType: driver_pb.CommandReply,
    requestSerialize: serialize_driver_CommandRequest,
    requestDeserialize: deserialize_driver_CommandRequest,
    responseSerialize: serialize_driver_CommandReply,
    responseDeserialize: deserialize_driver_CommandReply,
  },
  fire: {
    path: '/driver.Driver/Fire',
    requestStream: false,
    responseStream: false,
    requestType: driver_pb.CommandRequest,
    responseType: driver_pb.CommandReply,
    requestSerialize: serialize_driver_CommandRequest,
    requestDeserialize: deserialize_driver_CommandRequest,
    responseSerialize: serialize_driver_CommandReply,
    responseDeserialize: deserialize_driver_CommandReply,
  },
};

exports.DriverClient = grpc.makeGenericClientConstructor(DriverService);
