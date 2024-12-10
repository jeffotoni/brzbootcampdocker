//@jeffotoni
import http from 'k6/http';
import { sleep } from 'k6';

const port = __ENV.PORT || '8080';

const headers = { 'Content-Type': 'application/json' };

//url 
const url = `http://localhost:${port}/v1/user`;

export default function() {
    http.get(url, { headers: headers });
}