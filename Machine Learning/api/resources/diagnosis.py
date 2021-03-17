from flask import request
from flask_restful import Resource, reqparse

from common.data import data
from common.ml import PredictionService

#Prediction = PredictionService('diagnosis')

class DiagnosisResource(Resource):
    def get(self):
        name = request.args.get('name', '')
        diagnosis = data.diagnosis
        if name:
            diagnosis = list(filter(lambda dis: name.lower() in dis.name.lower(), data.diagnosis))
            if not diagnosis:
                return {
                    'code': 404,
                    'message': 'Not Found'
                }, 404
        return {'code': 200,
                'diagnosis': [{
                    'name': dis.name,
                    'symptoms': [s.name for s in dis.symptoms]
            } for dis in diagnosis]}, 200
        
    def post(self):
        parser = reqparse.RequestParser()
        parser.add_argument('symptoms', action='append', required=True)
        parser.add_argument('age', type=int)
        parser.add_argument('gender', choices=('m', 'f'))
        parser = parser.parse_args()
        
        symptoms = parser['symptoms']
        symptoms = [s.id for s in symptoms for s in data.parse_symptom_name(s)]

        diagnosis = [{'name': diag.name,
                      'symptoms': [sym.name for sym in diag.symptoms],
                      'hits': len([sym for sym in diag.symptoms if sym.id in symptoms])}
                      for diag in data.diagnosis if any(diag_sym.id in symptoms for diag_sym in diag.symptoms)]

        # Prepared for Machine Learning prediction
        #age = parser['age']
        #gender = parser['gender']
        #diagnosis = PredictionService.predict({'age': age, 'gender': gender, 'symptoms': symptoms})

        if not diagnosis:
            return {'code': 404,
                    'message': "Not found"}, 404

        return {'code': 200,
                'diagnosis': sorted(diagnosis, key=lambda x: x['hits'], reverse=True)}, 200