import http from 'k6/http';
import { SharedArray } from "k6/data";
import { sleep } from 'k6';

var payload = new SharedArray("growth", function () {
  var f = JSON.parse(open("./data/growth.json"));
  //console.log(JSON.stringify(f));
  return f;
});

export let options = {
  vus: 10,
  duration: '1m',
};

export default function () {
  var url = `${__ENV.DOMAIN}/api/v1/growth`;
  var params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  http.post(url,  JSON.stringify(payload), params);
}

