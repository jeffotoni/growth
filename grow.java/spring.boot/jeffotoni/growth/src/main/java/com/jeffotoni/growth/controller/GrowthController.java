package com.jeffotoni.growth.controller;

import java.net.URI;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ThreadLocalRandom;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;

// import com.jeffotoni.growth.add.Add;
import com.jeffotoni.growth.model.Censo;

@RestController
@RequestMapping(path = "/v1/growth")
public class GrowthController {

	static final Map<String, Censo> mapCenso = new HashMap<>();

	@GetMapping
	@ResponseBody
	@ResponseStatus(HttpStatus.OK)

	// consumes = MediaType.APPLICATION_JSON_VALUE)
	public List<Censo> lista() {

		Censo censo = new Censo("00001", "BRZ", "NGDP_R", 183.26, 2002);
		return Arrays.asList(censo, censo, censo);
		// return CensoDto.convert(Arrays.asList(censo, censo, censo));
	}

	@ResponseStatus(HttpStatus.OK)
	@GetMapping(path = "/{key}")
	public List<Censo> getCenso(@PathVariable String key) {

		Censo censo = mapCenso.get(key);
		// System.out.println("here censo:"+censo.getContry());
		System.out.println("here censo:" + key);
		// System.out.println("here cabeca de pudim:" + censo.getContry());
		// return "ola nada ainda";

		if (censo == null) {
			Censo censo2 = new Censo("0", "", "", 0.0, 0);
			return Arrays.asList(censo2);
		}

		return Arrays.asList(censo);
	}

	@PostMapping
	@ResponseStatus(HttpStatus.CREATED)
	public ResponseEntity<Censo> add(@RequestBody Censo censo) {
		int randomNum = ThreadLocalRandom.current().nextInt(1, 100000000 + 1);
		String key;
		key = String.valueOf(randomNum);
		censo.setId(key);
		mapCenso.put(key, censo);

		// Censo sandrao = mapCenso.get(key);
		// System.out.println("censo:"+sandrao.getContry());
		// System.out.println("key:"+key);

		return ResponseEntity.created(URI.create(String.format("/v1/growth"))).header("Engine", "Spring Boot")
				.header("Country", censo.getContry()).header("Indicator", censo.getIndicator()).header("key", key)
				.body(censo);
	}
}