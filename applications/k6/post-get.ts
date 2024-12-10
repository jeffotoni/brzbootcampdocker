import http from 'k6/http';
import { sleep } from 'k6';

const port = __ENV.PORT || '8080';

const ids = [];

function generateData() {
    return {
        id: Math.floor(Math.random() * 1000000), // Gera IDs únicos
        name: `User-${__ITER}`,
    };
}

export default function () {
    const index = __ITER + 1; // Iteração atual (começa em 0, incrementa para começar em 1)
    const payload = JSON.stringify(generateData(index));

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    // URL do POST
    const postUrl = `http://localhost:${port}/v1/user`;

    // Enviar o POST
    const postRes = http.post(postUrl, payload, params);
    console.log(`POST Status: ${postRes.status}, Payload enviado: ${payload}`);

    // Se o POST foi bem-sucedido, armazene o ID para consultas GET
    if (postRes.status === 201) {
        const responseData = JSON.parse(postRes.body);
        ids.push(responseData.user.id); // Adiciona o ID ao array de IDs
    }

    //console.log(`ids.length: ${ids.length}`);
    // Realizar GET para todos os IDs disponíveis
    if (ids.length > 50) {
        for (let i = 0; i < ids.length; i++) {
            const idToGet = ids[i];
            const getUrl = `http://localhost:${port}/v1/user/${idToGet}`;
            const getRes = http.get(getUrl);
            console.log(`GET Status: ${getRes.status}, ID consultado: ${idToGet}`);
        }
        // limpar ids
        ids.length = 0;
    }
}