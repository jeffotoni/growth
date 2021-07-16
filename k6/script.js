import http from 'k6/http';
import { SharedArray } from "k6/data";
import { sleep } from 'k6';

var payload = new SharedArray("growth", function () {
  var f = JSON.parse(open("./growth.json"));
  console.log(f);
  return f;
});

export let options = {
  vus: 30,
  duration: '1m',
};

// export default function () {
//   var url = 'http://192.168.0.70:8080/api/v1/growth';
//   var params = {
//     headers: {
//       'Content-Type': 'application/json',
//     },
//   };

//   http.post(url, payload, params);
//   sleep(1);
// }

//export default function () {
//	#  http.get('http://192.168.0.70:8080/ping');
//	#http.get('http://192.168.0.70:8080/api/v1/growth/brz/ngdp_r/2002');
//	#}
