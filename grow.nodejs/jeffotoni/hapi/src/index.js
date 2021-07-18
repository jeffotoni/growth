"use strict";

const Hapi = require("@hapi/hapi");
const port = process.env.PORT || 8080;
var map = new Map();

const init = async () => {
  const server = Hapi.server({
    port: port,
    host: "0.0.0.0",
  });

  await server.start();
  console.log("Server running on %s", server.info.uri);

  server.route({
    method: "GET",
    path: "/ping",
    handler: (req, res) => {
      const msg = { msg: "pong" };
      return res
        .response(msg)
        .code(202)
        .header("Content-Type", "application/json");
    },
  });

  server.route({
    config: {
      payload: {
        parse: true,
        maxBytes: 10485760,
      },
    },
    method: "POST",
    path: "/api/v1/growth",
    handler: (req, res) => {
      const payload = req.payload;
      var i = 0;
      payload.forEach(function (object) {
        var key =
          object.Country.toUpperCase() +
          object.Indicator.toUpperCase() +
          object.Year;
        var value = object.Value;
        map.set(key, value);
        i++;
      });
      const msg = { msg: "success", count: i };
      return res
        .response(msg)
        .code(202)
        .header("Content-Type", "application/json");
    },
  });

  server.route({
    config: {
      payload: {
        parse: true,
        maxBytes: 10485760,
      },
    },
    method: "PUT",
    path: "/api/v1/growth/{country}/{indicator}/{year}",
    handler: (req, res) => {
      const newVal = req.payload;
      if (newVal) {
        if (newVal.value) {
          var country = req.params.country;
          var indicator = req.params.indicator;
          var year = req.params.year;
          var key = country + indicator + year;
          var ok = map.set(key.toUpperCase(), newVal.value);
          if (ok) {
            const msg = { msg: "success" };
            return res
              .response(msg)
              .code(200)
              .header("Content-Type", "application/json");
          } else {
            const msg = { msg: "not found value" };
            return res
              .response(msg)
              .code(200)
              .header("Content-Type", "application/json");
          }
        }
      } else {
        const msg = { msg: "not found value body" };
        return res
          .response(msg)
          .code(200)
          .header("Content-Type", "application/json");
      }
    },
  });

  server.route({
    method: "GET",
    path: "/api/v1/growth/{country}/{indicator}/{year}",
    handler: (req, res) => {
      var country = req.params.country;
      var indicator = req.params.indicator;
      var year = req.params.year;
      var key = country + indicator + year;
      var value = map.get(key.toUpperCase());
      if (value) {
        var growth = {
          Country: country,
          Indicator: indicator,
          Year: year,
          Value: value,
        };
        return res
          .response(growth)
          .code(200)
          .header("Content-Type", "application/json");
      } else {
        const msg = { msg: "not found value" };
        return res
          .response(msg)
          .code(200)
          .header("Content-Type", "application/json");
      }
    },
  });
};

process.on("unhandledRejection", (err) => {
  console.log("error process:", +err);
  process.exit(1);
});

init();
