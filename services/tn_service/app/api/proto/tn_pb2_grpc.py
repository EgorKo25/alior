# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from . import tn_pb2 as tn__pb2

GRPC_GENERATED_VERSION = '1.64.1'
GRPC_VERSION = grpc.__version__
EXPECTED_ERROR_RELEASE = '1.65.0'
SCHEDULED_RELEASE_DATE = 'June 25, 2024'
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    warnings.warn(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in tn_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
        + f' This warning will become an error in {EXPECTED_ERROR_RELEASE},'
        + f' scheduled for release on {SCHEDULED_RELEASE_DATE}.',
        RuntimeWarning
    )


class TNStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateCallBack = channel.unary_unary(
                '/grpc.TN/CreateCallBack',
                request_serializer=tn__pb2.RequestCreateCallback.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)
        self.GetCallBack = channel.unary_unary(
                '/grpc.TN/GetCallBack',
                request_serializer=tn__pb2.RequestGetCallBack.SerializeToString,
                response_deserializer=tn__pb2.CallBack.FromString,
                _registered_method=True)
        self.GetAllCallBacks = channel.unary_stream(
                '/grpc.TN/GetAllCallBacks',
                request_serializer=tn__pb2.RequestGetAllCallBacks.SerializeToString,
                response_deserializer=tn__pb2.CallBack.FromString,
                _registered_method=True)
        self.DeleteCallBack = channel.unary_unary(
                '/grpc.TN/DeleteCallBack',
                request_serializer=tn__pb2.RequestDeleteCallBack.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)


class TNServicer(object):
    """Missing associated documentation comment in .proto file."""

    def CreateCallBack(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetCallBack(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetAllCallBacks(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteCallBack(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_TNServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateCallBack': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateCallBack,
                    request_deserializer=tn__pb2.RequestCreateCallback.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'GetCallBack': grpc.unary_unary_rpc_method_handler(
                    servicer.GetCallBack,
                    request_deserializer=tn__pb2.RequestGetCallBack.FromString,
                    response_serializer=tn__pb2.CallBack.SerializeToString,
            ),
            'GetAllCallBacks': grpc.unary_stream_rpc_method_handler(
                    servicer.GetAllCallBacks,
                    request_deserializer=tn__pb2.RequestGetAllCallBacks.FromString,
                    response_serializer=tn__pb2.CallBack.SerializeToString,
            ),
            'DeleteCallBack': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteCallBack,
                    request_deserializer=tn__pb2.RequestDeleteCallBack.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'grpc.TN', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('grpc.TN', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class TN(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def CreateCallBack(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/grpc.TN/CreateCallBack',
            tn__pb2.RequestCreateCallback.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def GetCallBack(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/grpc.TN/GetCallBack',
            tn__pb2.RequestGetCallBack.SerializeToString,
            tn__pb2.CallBack.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def GetAllCallBacks(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(
            request,
            target,
            '/grpc.TN/GetAllCallBacks',
            tn__pb2.RequestGetAllCallBacks.SerializeToString,
            tn__pb2.CallBack.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def DeleteCallBack(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/grpc.TN/DeleteCallBack',
            tn__pb2.RequestDeleteCallBack.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
