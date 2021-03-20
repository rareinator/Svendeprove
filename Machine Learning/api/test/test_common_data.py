import os
import cv2
import numpy as np
import pandas as pd
import tensorflow as tf
import tempfile
import pytest

from common.data import DataCache
from common.ml import PredictionService

IMG_SIZE = 96

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

def test_diagnosis_name_parse(data):
    diagnosis = data.parse_diagnosis_name("inal tra")
    assert diagnosis[0].name == "Abdominal trauma"

def test_machine_learning_model_loaded(data):
    Prediction = PredictionService("scan")
    assert Prediction.model

def test_scan_model_prediction_positive(data):
    Prediction = PredictionService("scan")

    img = _format_image("test_pos.jpg")
    result = Prediction.predict(img)

    assert result[1] * 100.0 > 50.0
    assert result[0] * 100.0 < 50.0

def test_scan_model_prediction_negative(data):
    Prediction = PredictionService("scan")

    img = _format_image("test_neg.jpg")
    result = Prediction.predict(img)
    
    assert result[0] * 100.0 > 50.0
    assert result[1] * 100.0 < 50.0

def _format_image(filepath):
    img_array = cv2.imread(filepath, cv2.IMREAD_COLOR)
    image = cv2.resize(img_array, (IMG_SIZE,IMG_SIZE))
    image = np.reshape(img_array, (-1, IMG_SIZE,IMG_SIZE,3))
    image = tf.image.convert_image_dtype(image, tf.float32)
    return image