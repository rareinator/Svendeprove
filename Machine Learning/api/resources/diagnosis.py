from flask import request
from flask_restful import Resource, reqparse

from common.data import data
from common.ml import PredictionService

Prediction = PredictionService('diagnosis')

class Diagnosis(Resource):
    def get(self):
        name = request.args.get('name', '')
        symptoms = request.args.get('symptoms', '')
        diagnosis = data.diagnosis
        s_dis = []
        n_dis = []
        if name:
            n_dis = data.parse_diagnosis_name(name)
        if symptoms:
            s_dis = list(filter(lambda dis: any([dis_sym for dis_sym in dis.symptoms if sym.lower() in dis_sym.name.lower()] for sym in symptoms), data.diagnosis))
        if name or symptoms:
            diagnosis = n_dis.extend(s_dis)
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
        parser.add_argument('symptoms', action='append')
        parser.add_argument('age', type=int)
        parser.add_argument('gender', choices=('m', 'f'))
        parser.add_argument('bp', type=int, choices=(0, 1, 2)) #Low, normal, high
        parser.add_argument('chol', type=int)
        parser = parser.parse_args()
        
        symptoms = parser['symptoms'].split(",")
        age = parser['age']
        gender = parser['gender']
        cholesterol = parser['chol']
        b_pressure = parser['bp']

        symptoms = [symptoms] if type(symptoms) is not list else symptoms
        symptoms = [data.parse_symptom_name(sym).id for sym in symptoms]

        prediction = PredictionService.predict([age, gender, cholesterol, b_pressure, symptoms])
        diagnosis = list(filter(lambda dis: dis.id in prediction, data.diagnosis))
        return {"code": 200,
                "diagnosis": [{'name': dis.name} for dis in diagnosis]}

        # diagnosis = filter(lambda dis: any([dis_sym for dis_sym in dis.symptoms if sym.lower() in dis_sym.name.lower()] for sym in symptoms), data.diagnosis)
        # if not diagnosis:
        #     return {'code': 404,
        #             'message': "Not found"}, 404

        # return {'code': 200,
        #     'diagnosis': [{
        #     'name': dis.name,
        #     'symptoms': [sym.name for sym in dis.symptoms]
        # } for dis in diagnosis]}, 200