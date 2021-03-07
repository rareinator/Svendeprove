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
        self.diagnosis = self.load_diagnosis()
        self.symptoms = self.load_symptoms()
        self.set_relations()

    def load_diagnosis(self):
        diagnosis = []
        with open(__location__ + "/data/dis.csv", "r") as csvfile:
            reader = csv.reader(csvfile, delimiter=',', quotechar='"')
            next(reader, None) # skip headers
            for row in reader:
                diagnosis.append(Diagnosis(row[0], row[4]))
        return diagnosis

    def load_symptoms(self):
        symptoms = []
        with open(__location__ + "/data/sym.csv", "r") as csvfile:
            reader = csv.reader(csvfile, delimiter=',', quotechar='"')
            next(reader, None)
            for row in reader:
                symptoms.append(Symptom(row[2], row[3]))
        return symptoms

    def set_relations(self):
        with open(__location__ + "/data/sym_dis.csv", "r") as csvfile:
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
                    
data = DataCache()