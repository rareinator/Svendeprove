from flask import request
from flask_restful import Resource, reqparse

from api.common.data import data

class Diagnosis(Resource):
    def get(self):
        name = request.args.get('name', '')
        if name:
            diagnosis = list(filter(lambda dis: name.lower() in dis.name.lower(), data.diagnosis))
            if diagnosis:
                return [{
                    'name': dis.name,
                    'symptoms': [s.name for s in dis.symptoms]
                } for dis in diagnosis]
            else:
                return {
                    '404': 'Not Found'
                }, 404
        else:
            return [{
                'name': dis.name,
                'symptoms': [s.name for s in dis.symptoms]
            } for dis in data.diagnosis]
        
    def post(self):
        parser = reqparse.RequestParser()
        parser.add_argument('symptoms', action='append')
        parser.add_argument('age', type=int)
        parser.add_argument('gender', choices=('m', 'f'))
        parser = parser.parse_args()
        
        symptoms = parser['symptoms']
        age = parser['age']
        gender = parser['gender']


        diagnosis = filter(lambda dis: any([dis_sym for dis_sym in dis.symptoms if sym.lower() in dis_sym.name.lower()] for sym in symptoms), data.diagnosis)
        return {'diagnosis': [{
            'name': dis.name,
            'symptoms': [sym.name for sym in dis.symptoms]
        } for dis in diagnosis]}