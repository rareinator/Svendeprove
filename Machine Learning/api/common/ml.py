import os
import ast
import cv2
import numpy as np
import pandas as pd
import tensorflow as tf
from tensorflow import keras

IMG_SIZE = 96
__location__ = os.path.realpath(
    os.path.join(os.getcwd(), os.path.dirname(__file__)))


class PredictionService():
    def __init__(self, model):
        self.model_name = model
        self.model_path = __location__ + "/data/models/" + model
        self.model = keras.models.load_model(self.model_path)

    def predict(self, data):
        return self.model.predict(data)[0]

    def _save_csv_data(self, data, name):
        with open(__location__ + "/data/csv/" + name, "w") as f:
            f.write(data)

    #Format image for machine learning predictions
    def _format_image(self, filepath):
        img_array = cv2.imread(filepath, cv2.IMREAD_COLOR)
        image = cv2.resize(img_array, (IMG_SIZE,IMG_SIZE))
        image = np.reshape(img_array, (-1, IMG_SIZE,IMG_SIZE,3))
        image = tf.image.convert_image_dtype(image, tf.float32)
        os.remove(filepath)
        return image

    ### BETA Feature ###
    def _train(self, data_file):
        training_data = pd.read_csv("data/csv/" + data_file)

        # Convert string list to actual list type
        training_data['symptoms'] = training_data['symptoms'].apply(lambda x: ast.literal_eval(x))
        # Pad sequences so all are same length
        training_data['symptoms'] = tf.keras.preprocessing.sequence.pad_sequences(training_data['symptoms'])
        y_train = training_data.pop("diagnosis")

        # "one-hot" encode category columns
        def one_hot_cat_column(feature_name, vocab):
            return tf.feature_column.indicator_column(
                tf.feature_column.categorical_column_with_vocabulary_list(feature_name, vocab))

        def build_feature_columns():
            feature_columns = []
            # Categorical columns
            for feature_name in ["gender"]:
                vocab = training_data[feature_name].unique()
                feature_columns.append(one_hot_cat_column(feature_name, vocab))
            # Numeric columns
            for feature_name in ["age", "chol", "bp"]:
                feature_columns.append(tf.feature_column.numeric_column(feature_name, dtype=tf.int32))
            # Sequence columns
            for feature_name in ["symptoms"]:
                symptoms = tf.feature_column.sequence_categorical_column_with_identity("symptoms", num_buckets=5000)
                feature_columns.append(symptoms)
            return feature_columns

        def make_input(x, y, n_epochs=None, shuffle=True):
            def input_fn():
                dataset = tf.data.Dataset.from_tensor_slices((dict(x), y))
                if shuffle:
                    dataset = dataset.shuffle(len(y_train))
                dataset = dataset.repeat(n_epochs)
                dataset = dataset.batch(len(y_train))
                return dataset
            return input_fn

        feature_columns = build_feature_columns()

        train_input = make_input(training_data, y_train)
        boosted_gradient = tf.estimator.BoostedTreesClassifier(feature_columns, n_classes=5000, n_batches_per_layer=1)
        model = boosted_gradient.train(train_input, max_steps=100)
        model.save(self.model_path)