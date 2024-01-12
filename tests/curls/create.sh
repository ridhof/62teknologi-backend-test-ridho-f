curl "http://localhost:8080/business" \
	--include \
	--header "Content-Type: application/json" \
	--request "POST" \
	--data '{
		"alias": "test-6",
		"name": "Test 5",
		"image_url": "http://test5.com",
		"latitude": 64.23185121231231,
		"longitude": 114.123183124123,
		"location_address_1": "Jl Test 5",
		"city": "Kota Test",
		"zip_code": "92012",
		"country": "IND",
		"state": "EJV",
		"phone": "8125849213",
		"display_phone": "(+62) 852-1234-4213",
		"price": "$$",
		"transactions": [
			"delivery",
			"pickup"
		],
		"display_address": [
			"Jl Pahlawan No 32V"
		]
	}'
