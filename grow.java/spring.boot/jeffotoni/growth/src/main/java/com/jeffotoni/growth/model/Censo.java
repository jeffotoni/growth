package com.jeffotoni.growth.model;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class Censo {

	private String Id;
	private String Country;
	private String Indicator;
	private Double Value;
	private int Year;

	public Censo(String Id, String Country, String Indicator, Double Value, int Year) {
		this.Id = Id;
		this.Country = Country;
		this.Indicator = Indicator;
		this.Value = Value;
		this.Year = Year;
	}

	public void setId(String id) {
		this.Id = id;
	}

	public String getId() {
		return Id;
	}

	public String getContry() {
		return Country;
	}

	public String getIndicator() {
		return Indicator;
	}

	public Double getValue() {
		return Value;
	}

	public int getYear() {
		return Year;
	}

	// public CensoDto(Censo censo) {
	// this.Country = censo.getCountry();
	// this.Indicator = censo.getIndicator();
	// this.Value = censo.getValue();
	// this.Year = censo.getYear();
	// }

	// public static List<CensoDto> convert(List<Censo> censos) {
	// return censos.stream().map(CensoDto::new).collect(Collectors.toList());
	// }
}