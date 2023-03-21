## Libraries used

1. Gorilla Mux - github.com/gorilla/mux
2. GORM - github.com/jinzhu/gorm
3. github.com/jinzhu/gorm/dialects/postgres


## Implementation done

Below APIs are implemented

1. GET API to fetch data from timeScaleDB based on UUID passed

    curl :- curl --location --request GET 'localhost:8089/rollerSpeed' \
            --header 'uuid: 7622ac5c-017a-44c9-812b-be5033767321'

2. POST API to create entry in the timeScaleDB table?

    curl :- curl --location --request POST 'localhost:8089/rollerSpeed/create' \
            --header 'Content-Type: application/json' \
            --data-raw '{
                "event_timestamp": 1679299917525,
                "roller_timestamp": "1679299917525",
                "roller_1_speed": 10,
                "roller_2_speed": 13,
                "roller_3_speed": 14,
                "roller_4_speed": 14,
                "roller_5_speed": 14,
                "roller_6_speed": 14,
                "roller_7_speed": 14,
                "roller_8_speed": 1.4,
                "roller_9_speed":	14,
                "roller_10_speed" :	14,
                "roller_11_speed" :	14,
                "roller_12_speed" :	14,
                "roller_13_speed" :	13,
                "roller_14_speed" :	10,
                "recipe_id" :	"0001",
                "equipment_id" :	"01CL010001",
                "mother_roll_id" :	"01AMR1BC177"
            }'

