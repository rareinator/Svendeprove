import os
import tempfile

from unittest import TestCase
from api import app

class TestResourceDiagnosis(TestCase):
    def setUp(self):
        self.app = app.test_client()

    def test_diagnosis_count(self):
        response = self.app.get("/diagnosis")
        data = response.json
        assert len(data) > 10

    # def test_diagnosis_name_filter(self):
    #     assert

    # def test_diagnosis_prediction_by_symptom(self):
    #     assert

    # def test_diagnosis_prediction_by_partial_symptom(self):
    #     assert