import csv
import os

__location__ = os.path.realpath(
    os.path.join(os.getcwd(), os.path.dirname(__file__)))

class Symptom(object):
    def __init__(self, id, name, diagnosis=[]):
        self.id = id
        self.name = name
        self.diagnosis = diagnosis

class Diagnosis(object):
    def __init__(self, id, name, symptoms=[]):
        self.id = id
        self.name = name
        self.symptoms = symptoms

class DataCache(object):
    def __init__(self):
        self.diagnosis = self._load_diagnosis()
        self.symptoms = self._load_symptoms()
        self._set_relations()

    def _load_diagnosis(self):
        diagnosis = []
        with open(__location__ + "/data/csv/dis.csv", "r") as csvfile:
            reader = csv.reader(csvfile, delimiter=',', quotechar='"')
            next(reader, None) # skip headers
            for row in reader:
                diagnosis.append(Diagnosis(row[0], row[4]))
        return diagnosis

    def _load_symptoms(self):
        symptoms = []
        with open(__location__ + "/data/csv/sym.csv", "r") as csvfile:
            reader = csv.reader(csvfile, delimiter=',', quotechar='"')
            next(reader, None)
            for row in reader:
                symptoms.append(Symptom(row[2], row[3]))
        return symptoms

    def _set_relations(self):
        with open(__location__ + "/data/csv/sym_dis.csv", "r") as csvfile:
            reader = csv.reader(csvfile, delimiter=',', quotechar='"')
            next(reader, None)
            sym_dis = {}
            for row in reader:
                if row[2] and row[8]:
                    symptom = next(sym for sym in self.symptoms if sym.id == row[8])
                    if sym_dis.get(row[2]):
                        sym_dis[row[2]].append(symptom)
                    else:
                        sym_dis[row[2]] = [symptom]
            for key, value in sym_dis.items():
                diagnosis = next(dis for dis in self.diagnosis if dis.id == key)
                diagnosis.symptoms = value

    def parse_symptom_name(self, name):
        return [sym for sym in self.symptoms if name.lower() in sym.name.lower()]
    def parse_diagnosis_name(self, name):
        return [dis for dis in self.diagnosis if name.lower() in dis.name.lower()]
            
data = DataCache()