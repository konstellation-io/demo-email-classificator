# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: public_input.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x12public_input.proto\x12\x04main\"(\n\x14\x43lassificatorRequest\x12\x10\n\x08\x66ilename\x18\x01 \x01(\t\"A\n\x19\x42\x61tchClassificatorRequest\x12\x10\n\x08\x66ilename\x18\x01 \x01(\t\x12\x12\n\nbatch_size\x18\x02 \x01(\x05\"\x1b\n\x08Response\x12\x0f\n\x07message\x18\x01 \x01(\t2\x94\x01\n\nEntrypoint\x12=\n\rClassificator\x12\x1a.main.ClassificatorRequest\x1a\x0e.main.Response\"\x00\x12G\n\x12\x42\x61tchClassificator\x12\x1f.main.BatchClassificatorRequest\x1a\x0e.main.Response\"\x00\x42\tZ\x07./protob\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'public_input_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\007./proto'
  _CLASSIFICATORREQUEST._serialized_start=28
  _CLASSIFICATORREQUEST._serialized_end=68
  _BATCHCLASSIFICATORREQUEST._serialized_start=70
  _BATCHCLASSIFICATORREQUEST._serialized_end=135
  _RESPONSE._serialized_start=137
  _RESPONSE._serialized_end=164
  _ENTRYPOINT._serialized_start=167
  _ENTRYPOINT._serialized_end=315
# @@protoc_insertion_point(module_scope)
