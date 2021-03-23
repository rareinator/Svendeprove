from flask import Flask, request, redirect
from flask_restful import Api

from resources.diagnosis import DiagnosisResource
from resources.scan import ScanResource

app = Flask(__name__)
api = Api(app)

api.add_resource(DiagnosisResource, '/diagnosis')
api.add_resource(ScanResource, '/scan')

if __name__ == "__main__":
    app.run()