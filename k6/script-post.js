import http from 'k6/http';
import { SharedArray } from "k6/data";
import { sleep } from 'k6';

var payload = new SharedArray("growth", function () {
  var f = JSON.parse(open("../3mb-growth_json.json"));
  //console.log(JSON.stringify(f));
  return f;
});

export let options = {
  vus: 10,
  duration: '1m',
  summaryTrendStats: ['avg', 'min', 'med', 'max', 'p(95)', 'p(99)', 'p(99.99)', 'count'],

  thresholds: {
    http_req_duration: ['avg<100', 'p(95)<200'],
    'http_req_connecting{cdnAsset:true}': ['p(95)<100'],
  },

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

