import os
import tempfile
import pytest

from api.common.data import DataCache

@pytest.fixture
def data():
    return DataCache()

def test_diagnosis_are_loaded(data):
    assert len(data.diagnosis) > 10

def test_symptoms_are_loaded(data):
    assert len(data.symptoms) > 10

def test_diagnosis_symptom_relation(data):
    diag = next(filter(lambda dis: dis.name == "Abdominal trauma", data.diagnosis))
    assert len(diag.symptoms) == 2
    
    # def test_symptoms_are_loaded(self):