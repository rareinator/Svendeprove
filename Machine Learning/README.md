# Machine Learning Service

## API
This folder contains the API necessary to serve the machine learning models that have been developed.

To install and run this API server, you need to perform the following commands from the root of the API application (The api folder)

``` python
> pip install requirements.txt
> flask run
```

Running these commands starts a development server at localhost:5000. Now you can start sending requests to the API using either your browser, or software such as postman.

--------
### Availabe Endpoints

- GET   /diagnosis
- POST  /diagnosis
- POST  /scan

Please refer to the official documentation of the API server for more information

--------
## Training
This folder contains the Jupyter Notebook files that have been used during the development of the machine learning models.

They are not necessary for the operation of the machine learning service. They are present for documentation purposes only.

To run Jupyter and open these notebooks, you first need to install Jupyter.

``` python
> pip install jupyterlab
```
Next, you need to run a Jupyter notebook server

``` python
> jupyter notebook
```

Once the jupyter server is running, you need to open your browser at localhost:8888
Here you can navigate to the training folder containing the .ipynb files and open them.