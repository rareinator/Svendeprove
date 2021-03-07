from flask import Flask, request, redirect
from flask_restful import Api

from resources.diagnosis import Diagnosis
from common.data import DataCache

app = Flask(__name__)
api = Api(app)

api.add_resource(Diagnosis, '/diagnosis')