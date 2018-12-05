
# 	dhtxx - Activity
This activity provides your Flogo app the ability to read temperature and humidity data from a DHTXX sensor 
## Installation

```bash
flogo install github.com/pradyuz3rocool/flogo-workspace/activity/dhtxx
```
Link for flogo web:
```
https://github.com/pradyuz3rocool/flogo-workspace/activity/dhtxx
```

## Schema
Inputs and Outputs:

```json
"inputs":[
    {
      "name": "Pin Number",
      "type": "any",
      "required": true
    },
    {
      "name": "Sensor Type",
      "type": "any"
    }
  ]
```
## Inputs
| Input                          | Description    |
|:-------------------------------|:---------------|
| Sensor Type                    | DHT11/ DHT22   |
| GPIO Pin number                | GPIO PIN Connected to sensor data  |


## Ouputs
| Output       | Description                                            |
|:-------------|:-------------------------------------------------------|
| Temperature  | In Celsius                                             |
| status       | Percentage of Humidity                                 |