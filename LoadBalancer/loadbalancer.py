from flask import Flask, jsonify, request
import requests

app = Flask(__name__)
global balance
balance = 0

@app.route('/getData/', methods=['GET'])
def getData():
    print('TESTE PASSEI AQUI')
    global balance
    value = None
    port = '8080'

    if balance == 0:
        value = requests.get('apiA:'+ port + '/getData/')
    else:
        value = requests.get('apiB:'+ port + '/getData/')

    balance = 1 - balance

    return value.json()

@app.route('/postData/', methods=['POST'])
def postData():
    print('TESTE PASSEI AQUI')
    global balance
    value = None
    x = request.get_json()
    port = '8080'

    inverse = {
        "x":x["y"],
        "y":x["x"]
    }

    if balance == 0:
        value = requests.post('apiA:'+ port + '/insert/', json=x)
    else:
        value = requests.post('apiB:'+ port + '/insert/', data=inverse)

    balance = 1 - balance

    return value.json()

if __name__ == '__main__':
    balance = 0
    app.run(host='localhost', port=8080)