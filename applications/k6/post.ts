import http from 'k6/http';
import { sleep } from 'k6';

const port = __ENV.PORT || '8080';

// gerar aleatorio ids e names
function generateData() {
    return {
        id: Math.floor(Math.random() * 1000000), // Gera IDs únicos
        name: `User-${__ITER}`,
    };
}

export const options = {
    insecureSkipTLSVerify: true,
};

export default function () {
    const index = __ITER + 1; // Iteração atual (começa em 0, incrementa para começar em 1)
    const payload = JSON.stringify(generateData(index));

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    //url 
    const url = `http://localhost:${port}/v1/user`;
    
    // Enviar o POST
    const res = http.post(url, payload, params);
    console.log(`Status: ${res.status}, Payload enviado: ${payload}`);
}