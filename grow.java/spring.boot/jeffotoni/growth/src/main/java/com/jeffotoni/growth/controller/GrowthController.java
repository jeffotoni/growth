package com.jeffotoni.growth.controller;

import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;

// import com.jeffotoni.growth.add.Add;
import com.jeffotoni.growth.model.Censo;

import java.net.URI;
import java.util.Arrays;
import java.util.List;

@RestController
@RequestMapping(path = "/v1/growth",
produces = MediaType.APPLICATION_JSON_VALUE,
consumes = MediaType.APPLICATION_JSON_VALUE)
public class GrowthController {
	@GetMapping
	@ResponseBody
	@ResponseStatus(HttpStatus.OK)
	public List <Censo> lista() {
		
		Censo censo = new Censo("BRZ", "NGDP_R",183.26, 2002);
		return Arrays.asList(censo, censo, censo);
		// return CensoDto.convert(Arrays.asList(censo, censo, censo));
	}

	@PostMapping
	@ResponseStatus(HttpStatus.CREATED)
	public ResponseEntity<Censo> add(@RequestBody Censo censo) {
		return ResponseEntity
            .created(URI.create(String.format("/v1/growth")))
            .body(censo);
	}
}