# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: tn.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder

# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2

DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x08tn.proto\x12\x04grpc\x1a\x1bgoogle/protobuf/empty.proto\"6\n\x08\x43\x61llBack\x12\x0c\n\x04Name\x18\x01 \x01(\t\x12\x0c\n\x04\x44\x61te\x18\x02 \x01(\t\x12\x0e\n\x06Number\x18\x03 \x01(\t\"9\n\x15RequestCreateCallback\x12 \n\x08\x63\x61llback\x18\x01 \x01(\x0b\x32\x0e.grpc.CallBack\" \n\x12RequestGetCallBack\x12\n\n\x02id\x18\x01 \x01(\x05\"E\n\x15RequestUpdateCallBack\x12\n\n\x02id\x18\x01 \x01(\x05\x12 \n\x08\x63\x61llback\x18\x02 \x01(\x0b\x32\x0e.grpc.CallBack\"#\n\x15RequestDeleteCallBack\x12\n\n\x02id\x18\x01 \x01(\x05\"$\n\x16ResponseDeleteCallBack\x12\n\n\x02id\x18\x01 \x01(\x05\x32\x90\x02\n\x02TN\x12\x45\n\x0e\x43reateCallBack\x12\x1b.grpc.RequestCreateCallback\x1a\x16.google.protobuf.Empty\x12\x37\n\x0bGetCallBack\x12\x18.grpc.RequestGetCallBack\x1a\x0e.grpc.CallBack\x12=\n\x0eUpdateCallBack\x12\x1b.grpc.RequestUpdateCallBack\x1a\x0e.grpc.CallBack\x12K\n\x0e\x44\x65leteCallBack\x12\x1b.grpc.RequestDeleteCallBack\x1a\x1c.grpc.ResponseDeleteCallBackB\x07Z\x05/grpcb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'tn_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\005/grpc'
  _globals['_CALLBACK']._serialized_start=47
  _globals['_CALLBACK']._serialized_end=101
  _globals['_REQUESTCREATECALLBACK']._serialized_start=103
  _globals['_REQUESTCREATECALLBACK']._serialized_end=160
  _globals['_REQUESTGETCALLBACK']._serialized_start=162
  _globals['_REQUESTGETCALLBACK']._serialized_end=194
  _globals['_REQUESTUPDATECALLBACK']._serialized_start=196
  _globals['_REQUESTUPDATECALLBACK']._serialized_end=265
  _globals['_REQUESTDELETECALLBACK']._serialized_start=267
  _globals['_REQUESTDELETECALLBACK']._serialized_end=302
  _globals['_RESPONSEDELETECALLBACK']._serialized_start=304
  _globals['_RESPONSEDELETECALLBACK']._serialized_end=340
  _globals['_TN']._serialized_start=343
  _globals['_TN']._serialized_end=615
# @@protoc_insertion_point(module_scope)
