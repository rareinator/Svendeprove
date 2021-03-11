import os
from tensorflow import keras

__location__ = os.path.realpath(
    os.path.join(os.getcwd(), os.path.dirname(__file__)))


class PredictionService():
    def __init__(self, model):
        self.model = keras.models.load_model(__location__ + "/data/models/" + model)

    def predict(self, data):
        return self.model.predict(data)[0]
