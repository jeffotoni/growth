const express = require("express");
const app = express();
const bodyParser = require("body-parser");
const { json } = require("body-parser");
const jsonParser = bodyParser.json({ limit: "10mb" });
const port = process.env.PORT || 8080;
var map = new Map();

app.get("/ping", (req, res) => {
  res.setHeader("Content-Type", "application/json");
  res.status(200).json({ name: "pong" });
});

app.post("/api/v1/growth", jsonParser, (req, res) => {
  res.setHeader("Content-Type", "application/json");
  var i = 0;
  var array = req.body;
  array.forEach(function (object) {
    var key =
      object.Country.toUpperCase() +
      object.Indicator.toUpperCase() +
      object.Year;
    var value = object.Value;
    map.set(key, value);
    i++;
  });
  res.status(202).json({ msg: "success", count: i });
});

app.get("/api/v1/growth/:country/:indicator/:year", (req, res) => {
  res.setHeader("Content-Type", "application/json");
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
    res.status(200).json(growth);
  } else {
    res.status(200).json({ msg: "not found value" });
  }
});

app.get("/api/v1/growth/size", (req, res) => {
  res.setHeader("Content-Type", "application/json");
  var size = map.size;
  if (size == 0) {
    res.status(200).json({ msg: "not finished" });
  } else {
    res.status(200).json({ msg: "complete", count: size });
  }
});

app.get("/api/v1/growth/status", (req, res) => {
  res.setHeader("Content-Type", "application/json");
  var val = map.get("BRZNGDP_R2002");
  if (val) {
    var growth = {
      Country: "brz",
      Indicator: "ngdp_r",
      Year: "2002",
      Value: val,
    };
    res.status(200).json(growth);
  } else {
    res.status(200).json({ msg: "not found value" });
  }
});

app.get("/api/v1/growth/clean", (req, res) => {
  res.setHeader("Content-Type", "application/json");
  map.clear();
  res.status(200).json({ msg: "complete", count: map.size });
});

app.delete("/api/v1/growth/:country/:indicator/:year", (req, res) => {
  res.setHeader("Content-Type", "application/json");
  var country = req.params.country;
  var indicator = req.params.indicator;
  var year = req.params.year;
  var key = country + indicator + year;
  var ok = map.delete(key.toUpperCase());
  if (ok) {
    res.status(202).json({ msg: "success" });
  } else {
    res.status(200).json({ msg: "not found value" });
  }
});

app.put("/api/v1/growth/:country/:indicator/:year", jsonParser, (req, res) => {
  res.setHeader("Content-Type", "application/json");
  var newVal = req.body;
  if (newVal) {
    if (newVal.value) {
      var country = req.params.country;
      var indicator = req.params.indicator;
      var year = req.params.year;
      var key = country + indicator + year;
      var ok = map.set(key.toUpperCase(), newVal.value);
      if (ok) {
        res.status(200).json({ msg: "success" });
      } else {
        res.status(200).json({ msg: "not found value" });
      }
    }
  } else {
    res.status(200).json({ msg: "not found value" });
  }
});

app.listen(port);
console.log("Run Server " + port);
