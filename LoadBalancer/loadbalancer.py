from flask import Flask, jsonify, request
import requests

app = Flask(__name__)
global balance
balance = 0

@app.route('/getData/', methods=['GET'])
def getData():
    global balance
    value = None
    port = '8080'

    if balance == 0:
        print('em A')
        value = requests.get('http://apiA:'+ port + '/getData/')
    else:
        print('em B')
        value = requests.get('http://apiB:'+ port + '/getData/')

    balance = 1 - balance

    return value.json()

@app.route('/postData/', methods=['POST'])
def postData():
    global balance
    value = None
    x = request.get_json()
    port = '8080'

    inverse = {
        "x":x["y"],
        "y":x["x"]
    }

    if balance == 0:
        value = requests.post('http://apiA:'+ port + '/insert/', json=x)
    else:
        value = requests.post('http://apiB:'+ port + '/insert/', json=inverse)

    balance = 1 - balance

    return value.json()

if __name__ == '__main__':
    balance = 0
    app.run(host='0.0.0.0', port=5000)