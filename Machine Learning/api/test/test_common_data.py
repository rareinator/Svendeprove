import os
import tempfile

from unittest import TestCase
from api.common.data import DataCache

class TestDataCache(TestCase):
    def setUp(self):
        self.data = DataCache()

    def test_diagnosis_are_loaded(self):
        assert len(self.data.diagnosis) > 10

    # def test_symptoms_are_loaded(self):