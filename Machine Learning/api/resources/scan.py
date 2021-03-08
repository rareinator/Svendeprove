import werkzeug
from flask import request
from flask_restful import Resource, reqparse

class Scan(Resource):

    def post(self):
        parser = reqparse.RequestParser()
        parser.add_argument('scan', type=werkzeug.FileStorage, location="files")
        parser = parser.parse_args()
        scan_file = parser["scan"]
        return {'200': 'ok'}