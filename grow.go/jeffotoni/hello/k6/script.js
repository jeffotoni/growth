import http from 'k6/http';
import { sleep } from 'k6';

const headers = { 'Content-Type': 'application/json' };

// export let options = {
//   vus: 10,
//   duration: '1m',
// };

export default function() {
    http.get(`http://localhost:8080/hello`, { headers: headers });
}