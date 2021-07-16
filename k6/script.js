import http from 'k6/http';
import { sleep } from 'k6';

const headers = { 'Content-Type': 'application/json' };

export let options = {
  vus: 10,
  duration: '1m',
};

export default function () {
  var url = `${__ENV.DOMAIN}/api/v1/growth`;

  http.get(`${__ENV.DOMAIN}/ping`,{ headers: headers });
  sleep(1);
	http.get(`${__ENV.DOMAIN}/api/v1/growth/brz/ngdp_r/2002`,{ headers: headers });
  sleep(1);
  const data = { value: 345.88 };
  http.put(`${__ENV.DOMAIN}/api/v1/growth/brz/ngdp_r/2002`,JSON.stringify(data), { headers: headers } );
}

