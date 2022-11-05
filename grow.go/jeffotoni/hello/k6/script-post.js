import http from 'k6/http';
import { SharedArray } from "k6/data";
import { sleep } from 'k6';

/*
var payload = new SharedArray("growth", function () {
  var f = JSON.parse(open("./data/500kb-growth_json.json"));
  //console.log(JSON.stringify(f));
  return f;
});*/

/*
export let options = {
  vus: 10,
  duration: '1m',
  summaryTrendStats: ['avg', 'min', 'med', 'max', 'p(95)', 'p(99)', 'p(99.99)', 'count'],

  thresholds: {
    http_req_duration: ['avg<100', 'p(95)<200'],
    'http_req_connecting{cdnAsset:true}': ['p(95)<100'],
  },

}*/

export default function () {
  var url = `http://localhost:8080/v1/growth`;
  var params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  http.post(url,  '{"Country":"BRZ","Indicator":"NGDP_R","Value":183.26,"Year":2002}', params);
}

