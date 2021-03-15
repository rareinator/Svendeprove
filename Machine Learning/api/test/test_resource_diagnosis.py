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
    assert len(data) == 6

def test_diagnosis_get_all(client):
    response = client.get("/diagnosis")
    data = response.json
    assert len(data) == 1347

def test_get_diagnosis_by_symptom(client):
    response = client.get("/diagnosis?symptoms=swell,pain")
    data = response.json
    assert len(data) > 0

def test_scan_positive_result(client):
    encoded_img = ""
    with open("test_pos.jpg", "rb") as img_file:
        encoded_img = base64.b64encode(img_file.read())
    response = client.post("/scan", data=dict(scan=encoded_img))
    data = response.json
    assert data[1] * 100.0 > 50.0
    assert data[0] * 100.0 < 50.0

def test_scan_negative_result(client):
    encoded_img = ""
    with open("test_neg.jpg", "rb") as img_file:
        encoded_img = base64.b64encode(img_file.read())
    response = client.post("/scan", data=dict(scan=encoded_img))
    data = response.json
    assert data[0] * 100.0 > 50.0
    assert data[1] * 100.0 < 50.0
