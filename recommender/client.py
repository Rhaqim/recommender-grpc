from __future__ import print_function

import logging

import grpc
from proto import recommender_pb2, recommender_pb2_grpc


def ger_recommendations(stub, user_id, num_recommendations):
    response = stub.GetRecommendations(
        recommender_pb2.RecommendationRequest(
            user_id=user_id, num_recommendations=num_recommendations
        )
    )
    print("Recommendations for user %s: %s" % (user_id, response.recommended_items))


def run():
    # NOTE(gRPC Python Team): .close() is possible on a channel and should be
    # used in circumstances in which the with statement does not fit the needs
    # of the code.
    with grpc.insecure_channel("localhost:50051") as channel:
        stub = recommender_pb2_grpc.RecommenderStub(channel)
        print("-------------- GetRecommendations --------------")
        ger_recommendations(stub, "user1", 5)


if __name__ == "__main__":
    logging.basicConfig()
    run()
