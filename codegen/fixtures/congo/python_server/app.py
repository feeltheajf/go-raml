from flask import Flask, send_from_directory, send_file
from deliveries_api import deliveries_api
from drones_api import drones_api

import os
dir_path = os.path.dirname(os.path.realpath(__file__))

app = Flask(__name__)

app.register_blueprint(deliveries_api)
app.register_blueprint(drones_api)



@app.route('/apidocs/<path:path>')
def send_js(path):
    return send_from_directory(dir_path + '/' + 'apidocs', path)


@app.route('/', methods=['GET'])
def home():
    return send_file('index.html')

if __name__ == "__main__":
    app.run(debug=True)
