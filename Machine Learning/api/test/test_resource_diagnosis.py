import os
import tempfile

import pytest

from unittest import TestCase
from app import app

@pytest.fixture
def client():
    return app.test_client()


def test_diagnosis_count(client):
    response = client.get("/diagnosis")
    data = response.json
    assert len(data) > 10

def test_diagnosis_partial_name(client):
    response = client.get("/diagnosis?name=abdominal")
    data = response.json
    assert len(data) == 6
    # def test_diagnosis_name_filter(self):
    #     assert

    # def test_diagnosis_prediction_by_symptom(self):
    #     assert

    # def test_diagnosis_prediction_by_partial_symptom(self):
    #     assert