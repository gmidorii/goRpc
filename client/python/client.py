import gorpc_pb2
import gorpc_pb2_grpc
import grpc

def run():
    channel = grpc.insecure_channel('localhost:8888')
    stub = gorpc_pb2.GoRpcStub(channel)
    res = stub.GetSample(gorpc_pb2.SampleReq(id=1, name='test'))
    print(res.mes)

if __name__ == '__main__':
    run()