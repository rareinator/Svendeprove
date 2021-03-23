import os
import base64
from flask_restful import Resource, reqparse
from common.ml import PredictionService

Prediction = PredictionService("scan")

class ScanResource(Resource):
    def post(self):
        parser = reqparse.RequestParser()
        parser.add_argument('scan')
        parser = parser.parse_args()
        scan_file = parser["scan"]

        if not scan_file:
            return {'code': 400,
                    'message': "No image received or image could not be read"}, 400
        # Save on disk - Unable to convert to numpy array while in memory
        with open("img.jpg", "wb") as fh:
            fh.write(base64.b64decode(scan_file))

        try:
            image = Prediction._format_image("img.jpg")
        except:
            os.remove("img.jpg")
            return {'code': 415,
                    'message': 'The image could not be processed'}, 415
        prediction = Prediction.predict(image)
        result = {'positive': "{:.2f}%".format(prediction[1] * 100.0),
                  'negative': "{:.2f}%".format(prediction[0] * 100.0)}
        return {'code': 200,
                'prediction': result}
