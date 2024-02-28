import logging
import math
import time
from concurrent import futures

import grpc
from proto import recommender_pb2, recommender_pb2_grpc

_ONE_DAY_IN_SECONDS = 60 * 60 * 24


def GetRecommendations(request, context):
    print("Received request: ", request)
    
    
    return recommender_pb2.RecommendationResponse(
        recommended_items = ["item1", "item2", "item3"]
    )


class RecommenderServicer(recommender_pb2_grpc.RecommenderServicer):
    def GetRecommendations(self, request, context):
        return GetRecommendations(request, context)


def serve():
    print("Server starting...")
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    recommender_pb2_grpc.add_RecommenderServicer_to_server(
        RecommenderServicer(), server
    )
    server.add_insecure_port("[::]:50051")
    server.start()
    print("Server started, listening on 50051")
    server.wait_for_termination()
    # try:
    #     while True:
    #         time.sleep(_ONE_DAY_IN_SECONDS)
    # except KeyboardInterrupt:
    #     server.stop(0)


if __name__ == "__main__":
    logging.basicConfig()
    serve()
