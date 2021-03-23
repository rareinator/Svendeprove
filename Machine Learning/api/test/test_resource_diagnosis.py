import os
import tempfile
import base64

import pytest

from app import app

@pytest.fixture
def client():
    return app.test_client()


def test_diagnosis_partial_name(client):
    response = client.get("/diagnosis?name=abdominal")
    data = response.json
    assert len(data['diagnosis']) == 6

def test_diagnosis_get_all(client):
    response = client.get("/diagnosis")
    data = response.json
    assert len(data['diagnosis']) == 1347

def test_diagnosis_by_symptom_post(client):
    response = client.post("/diagnosis", data=dict(age=25, gender='m', symptoms=['swell']))
    data = response.json
    assert len(data['diagnosis']) > 0

def test_scan_positive_result(client):
    encoded_img = ""
    with open("test_pos.jpg", "rb") as img_file:
        encoded_img = base64.b64encode(img_file.read())

    response = client.post("/scan", data=dict(scan=encoded_img))
    data = response.json

    positive = float(data['prediction']['positive'].replace('%', ''))
    negative = float(data['prediction']['negative'].replace('%', ''))
    
    assert positive > 50.0
    assert negative < 50.0

def test_scan_negative_result(client):
    encoded_img = ""
    with open("test_neg.jpg", "rb") as img_file:
        encoded_img = base64.b64encode(img_file.read())

    response = client.post("/scan", data=dict(scan=encoded_img))
    data = response.json

    positive = float(data['prediction']['positive'].replace('%', ''))
    negative = float(data['prediction']['negative'].replace('%', ''))

    assert negative > 50.0
    assert positive < 50.0
