from flask import Flask, request, redirect
from flask_restful import Api

from api.resources.diagnosis import Diagnosis
from api.resources.scan import Scan

app = Flask(__name__)
api = Api(app)

api.add_resource(Diagnosis, '/diagnosis')
api.add_resource(Scan, '/scan')