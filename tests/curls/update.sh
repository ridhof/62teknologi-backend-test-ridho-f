curl 'http://localhost:8080/business' \
	--include \
	--header "Content-Type: application/json" \
	--request "PUT" \
	--data '{
		"id": 18,
		"alias": "test-4",
		"price": "$",
		"transactions": [
			"pickup",
			"delivery"
		],
		"display_address": [
			"Test Street 200"
		],
		"name": "Test 9",
		"image_url": "https://google.com/image",
		"location_address_1": "1st test street number 9",
		"city": "9th Test",
		"zip_code": "14045",
		"country": "IND",
		"state": "EIN",
		"phone": "87123921311",
		"display_phone": "(+87) 123-921-311",
		"latitude": 56.12318912,
		"longitude": 112.12310312
	}'
