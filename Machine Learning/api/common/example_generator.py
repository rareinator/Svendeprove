import csv
import os
import random
from data import DataCache

data = DataCache()

with open('data/csv/example.csv', 'w', newline='') as f:
    csv_writer = csv.writer(f, delimiter=',', quotechar='"')
    csv_writer.writerow(["age", "gender", "chol", "bp", "symptoms", "diagnosis"])
    for d in data.diagnosis:
        for i in range(50):
            age_rand = random.randint(18, 85)
            gender_rand = random.choice(["m", "f"])
            chol_rand = random.randint(70, 280)
            bp_rand = random.choice([0, 1, 2]) # low, normal, high
            sym_len = len(d.symptoms)
            if not sym_len:
                continue
            rand_sym_num = random.randint(1, sym_len)
            rand_sym = random.sample([sym.id for sym in d.symptoms], rand_sym_num)
            csv_writer.writerow([age_rand, gender_rand, chol_rand, bp_rand, '[' + ','.join(rand_sym) + ']', d.id])
