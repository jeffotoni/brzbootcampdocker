from flask import Flask, request, jsonify

app = Flask(__name__)

# Simulando um banco de dados em memória
users = {}

@app.route('/v1/user', methods=['POST'])
def create_user():
    data = request.json
    if not data or 'id' not in data or 'name' not in data:
        return jsonify({"error": "Nao encontrado 'id' or 'name' field"}), 400

    user_id = data['id']
    if user_id in users:
        return jsonify({"error": "User já existe"}), 400

    users[user_id] = data['name']
    return jsonify({"message": "Usuario criado com sucesso", "user": data}), 201
    # return Response(json.dumps(response, ensure_ascii=False), mimetype='application/json'), 200

@app.route('/v1/user/<int:user_id>', methods=['GET'])
def get_user(user_id):
    if user_id not in users:
        return jsonify({"error": "User nao encontrado"}), 404

    return jsonify({"id": user_id, "name": users[user_id]}), 200

@app.route('/v1/users', methods=['GET'])
def list_users():
    return jsonify(users), 200

@app.route('/v1/user/<int:user_id>', methods=['PUT'])
def update_user(user_id):
    if user_id not in users:
        return jsonify({"error": "User nao encontrado"}), 404

    data = request.json
    if not data or 'name' not in data:
        return jsonify({"error": "Nao encontrado 'name' field"}), 400

    users[user_id] = data['name']
    return jsonify({"message": "User atualizado com sucesso", "id": user_id, "name": data['name']}), 200

@app.route('/v1/user/<int:user_id>', methods=['DELETE'])
def delete_user(user_id):
    if user_id not in users:
        return jsonify({"error": "User nao encontrado"}), 404

    del users[user_id]
    return jsonify({"message": "User deletado sucesso", "id": user_id}), 200

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)