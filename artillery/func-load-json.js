"use strict";
module.exports = {
  generatePayload,
};
var contents = require("fs").readFileSync("./data/500kb-growth_json.json");

function generatePayload(userContext, events, done) {
  var payload = {
    data: "data",
  };
  payload = JSON.parse(contents);
  userContext.vars.payload = payload;
  return done();
}
