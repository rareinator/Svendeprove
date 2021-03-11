import os
import cv2
import numpy as np
import tensorflow as tf
from flask import request
from flask_restful import Resource
from api.common.ml import PredictionService

IMG_SIZE = 96
Prediction = PredictionService("scan")

class Scan(Resource):

    def post(self):
        scan_file = request.files["scan"]
        if not scan_file or scan_file.filename == '':
            return {'code': 400,
                    'message': "No image received or image could not be read"}, 400
        scan_file.save(scan_file.filename)
        try:
            image = self._format_image(scan_file.filename)
        except:
            os.remove(scan_file.filename)
            return {'code': 415,
                    'message': 'The image could not be processed'}, 415
        prediction = Prediction.predict(image)
        result = {'positive': "{:.2f}%".format(prediction[1] * 100.0),
                  'negative': "{:.2f}%".format(prediction[0] * 100.0)}
        return {'code': 200,
                'prediction': result}

    #Format image for machine learning predictions
    def _format_image(self, filepath):
        img_array = cv2.imread(filepath, cv2.IMREAD_COLOR)
        image = cv2.resize(img_array, (IMG_SIZE,IMG_SIZE))
        image = np.reshape(img_array, (-1, IMG_SIZE,IMG_SIZE,3))
        image = tf.image.convert_image_dtype(image, tf.float32)
        os.remove(filepath)
        return image