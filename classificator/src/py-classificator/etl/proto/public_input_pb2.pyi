"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import google.protobuf.descriptor
import google.protobuf.message
import sys

if sys.version_info >= (3, 8):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

@typing_extensions.final
class Request(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    FILENAME_FIELD_NUMBER: builtins.int
    BATCH_SIZE_FIELD_NUMBER: builtins.int
    filename: builtins.str
    batch_size: builtins.int
    def __init__(
        self,
        *,
        filename: builtins.str = ...,
        batch_size: builtins.int = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["batch_size", b"batch_size", "filename", b"filename"]) -> None: ...

global___Request = Request

@typing_extensions.final
class Response(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    MESSAGE_FIELD_NUMBER: builtins.int
    message: builtins.str
    def __init__(
        self,
        *,
        message: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["message", b"message"]) -> None: ...

global___Response = Response
